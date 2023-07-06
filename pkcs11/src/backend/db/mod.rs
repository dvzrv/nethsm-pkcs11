// Copyright 2020-2021 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// modified from https://github.com/aws/aws-nitro-enclaves-acm

pub mod attr;
pub mod object;
pub use object::{Object, ObjectHandle, ObjectKind};
use openapi::apis::default_api::{self, KeysGetError, KeysKeyIdGetError};

use crate::config::device::Slot;

// NOTE: for now, we use these *Info structs to construct key objects. The source PEM is
// preserved, so that a crypto::Pkey (an EVP_PKEY wrapper) can be constructed whenever
// it is needed (e.g. at operation context initialization).
// If the PEM to EVP_PKEY conversion turns out to impact performance, we could construct
// the crypto::Pkey object at DB creation time, and replace the *Info structs with it,
// provided we also implement a proper cloning mechanism for crypto::Pkey. This is needed
// in order to make sure that each session gets its own copy of each key, and maintain
// thread safety.
// Cloning could be done via RSAPrivateKey_dup() and EC_KEY_dup(), together with a TryClone
// trait, since these operations can fail.
#[derive(Clone)]
pub struct RsaKeyInfo {
    pub priv_pem: String,
    pub id: cryptoki_sys::CK_BYTE,
    pub label: String,
    pub num_bits: cryptoki_sys::CK_ULONG,
    pub modulus: Vec<u8>,
    pub public_exponent: Vec<u8>,
}

#[derive(Clone)]
pub struct EcKeyInfo {
    pub priv_pem: String,
    pub id: cryptoki_sys::CK_BYTE,
    pub label: String,
    pub params_x962: Vec<u8>,
    pub point_q_x962: Vec<u8>,
}
/// Certificate object type
#[derive(Clone)]
pub enum CertCategory {
    #[allow(dead_code)]
    /// Default (unverified)
    Unverified,
    /// Token certificate
    Token,
    /// CA certificate
    Authority,
    #[allow(dead_code)]
    /// Other
    Other,
}
#[derive(Clone)]
pub struct CertInfo {
    pub categ: CertCategory,
    pub id: cryptoki_sys::CK_BYTE,
    pub label: String,
    pub subject_der: Vec<u8>,
    pub issuer_der: Vec<u8>,
    pub serno_der: Vec<u8>,
    pub cert_der: Vec<u8>,
}

#[derive(Debug)]
pub enum Error {
    ListKeys(openapi::apis::Error<KeysGetError>),
    GetKey(openapi::apis::Error<KeysKeyIdGetError>),
}

#[derive(Debug, Clone)]
pub struct Db {
    objects: Vec<Object>,
}

impl Db {
    pub fn new(api_config: openapi::apis::configuration::Configuration) -> Result<Self, Error> {
        let mut objects = Vec::new();

        let keys = default_api::keys_get(&api_config, None).map_err(Error::ListKeys)?;

        for key in keys.iter() {
            let key_data =
                default_api::keys_key_id_get(&api_config, &key.key).map_err(Error::GetKey)?;
            objects.push(Object::from_key_data(key_data, key.key.clone()));
        }

        Ok(Self { objects })
    }

    pub fn enumerate(&self) -> impl Iterator<Item = (ObjectHandle, &Object)> {
        self.objects
            .iter()
            .enumerate()
            .map(|(i, o)| (ObjectHandle::from(i), o))
    }

    pub fn object(&self, handle: ObjectHandle) -> Option<&Object> {
        if self.objects.len() <= usize::from(handle) {
            return None;
        }
        Some(&self.objects[usize::from(handle)])
    }
}
