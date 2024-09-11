use std::{
    collections::BTreeMap,
    sync::{
        atomic::{AtomicUsize, Ordering::Relaxed},
        mpsc::{self, RecvTimeoutError},
        Arc, Condvar, LazyLock, Mutex, RwLock,
    },
    thread,
    time::{Duration, Instant},
};

use nethsm_sdk_rs::apis::{configuration::Configuration, default_api::health_ready_get};

use crate::{backend::db::Db, data::THREADS_ALLOWED};

use super::config_file::{RetryConfig, UserConfig};

static RETRY_THREAD: LazyLock<mpsc::Sender<(Duration, InstanceData)>> = LazyLock::new(|| {
    let (tx, rx) = mpsc::channel();
    let (tx_instance, rx_instance) = mpsc::channel();
    thread::spawn(background_thread(rx_instance));
    thread::spawn(background_timer(rx, tx_instance));
    tx
});

fn background_timer(
    rx: mpsc::Receiver<(Duration, InstanceData)>,
    tx_instance: mpsc::Sender<InstanceData>,
) -> impl FnOnce() {
    let mut jobs: BTreeMap<Instant, InstanceData> = BTreeMap::new();
    move || loop {
        let next_job = jobs.pop_first();
        let Some((next_job_deadline, next_job_instance)) = next_job else {
            // No jobs in the queue, we can just run the next
            let Ok((new_job_duration, new_state)) = rx.recv() else {
                return;
            };

            jobs.insert(Instant::now() + new_job_duration, new_state);
            continue;
        };

        let now = Instant::now();

        if now >= next_job_deadline {
            tx_instance.send(next_job_instance).unwrap();
            continue;
        }
        jobs.insert(next_job_deadline, next_job_instance);

        let timeout = next_job_deadline.duration_since(now);
        match rx.recv_timeout(timeout) {
            Ok((run_in, new_instance)) => {
                jobs.insert(now + run_in, new_instance);
                continue;
            }
            Err(RecvTimeoutError::Timeout) => continue,
            Err(RecvTimeoutError::Disconnected) => break,
        }
    }
}

fn background_thread(rx: mpsc::Receiver<InstanceData>) -> impl FnOnce() {
    move || loop {
        while let Ok(instance) = rx.recv() {
            match health_ready_get(&instance.config) {
                Ok(_) => instance.clear_failed(),
                Err(_) => instance.bump_failed(),
            }
        }
    }
}

// stores the global configuration of the module
#[derive(Debug, Clone)]
pub struct Device {
    pub slots: Vec<Arc<Slot>>,
    pub enable_set_attribute_value: bool,
}

#[derive(Debug, Clone, Default, PartialEq, Eq)]
pub enum InstanceState {
    #[default]
    Working,
    Failed {
        retry_count: u8,
        last_retry_at: Instant,
    },
}

impl InstanceState {
    pub fn new_failed() -> InstanceState {
        InstanceState::Failed {
            retry_count: 0,
            last_retry_at: Instant::now(),
        }
    }
}

#[derive(Debug, Clone)]
pub struct InstanceData {
    pub config: Configuration,
    pub state: Arc<RwLock<InstanceState>>,
}

pub enum InstanceAttempt {
    /// The instance is in the failed state and should not be used
    Failed,
    /// The instance is in the failed  state but a connection should be attempted
    Retry,
    /// The instance is in the working state
    Working,
}

impl InstanceData {
    pub fn should_try(&self) -> InstanceAttempt {
        let this = self.state.read().unwrap();
        match *this {
            InstanceState::Working => InstanceAttempt::Working,
            InstanceState::Failed {
                retry_count,
                last_retry_at,
            } => {
                if last_retry_at.elapsed() < retry_duration_from_count(retry_count) {
                    InstanceAttempt::Failed
                } else {
                    InstanceAttempt::Retry
                }
            }
        }
    }

    pub fn clear_failed(&self) {
        *self.state.write().unwrap() = InstanceState::Working;
    }

    pub fn bump_failed(&self) {
        let mut write = self.state.write().unwrap();
        let retry_count = match *write {
            InstanceState::Working => {
                *write = InstanceState::new_failed();
                0
            }
            InstanceState::Failed {
                retry_count: prev_retry_count,
                last_retry_at,
            } => {
                // We only bump if it's a "real" retry. This is to avoid race conditions where
                // the same instance stops working when multiple threads are simultaneously connecting
                // to it
                if last_retry_at.elapsed() >= retry_duration_from_count(prev_retry_count) {
                    let retry_count = prev_retry_count.saturating_add(1);
                    *write = InstanceState::Failed {
                        retry_count,
                        last_retry_at: Instant::now(),
                    };
                    retry_count
                } else {
                    prev_retry_count
                }
            }
        };
        drop(write);
        if THREADS_ALLOWED.load(Relaxed) {
            RETRY_THREAD
                .send((retry_duration_from_count(retry_count), self.clone()))
                .ok();
        }
    }
}

fn retry_duration_from_count(retry_count: u8) -> Duration {
    let secs = match retry_count {
        0 | 1 => 1,
        2 => 2,
        3 => 5,
        4 => 10,
        5 => 60,
        6.. => 60 * 5,
    };

    Duration::from_secs(secs)
}

#[derive(Debug, Clone)]
pub struct Slot {
    pub label: String,
    pub retries: Option<RetryConfig>,
    pub _description: Option<String>,
    pub instances: Vec<InstanceData>,
    pub operator: Option<UserConfig>,
    pub administrator: Option<UserConfig>,
    pub db: Arc<(Mutex<Db>, Condvar)>,
    pub instance_balancer: Arc<AtomicUsize>,
}

impl Slot {
    // the user is connected if the basic auth is filled with an username and a password, otherwise the user will have to login
    pub fn is_connected(&self) -> bool {
        let Some(instance_data) = self.instances.first() else {
            return false;
        };
        let Some(auth) = &instance_data.config.basic_auth else {
            return false;
        };

        let Some(pwd) = &auth.1 else { return false };

        !pwd.is_empty()
    }
}
