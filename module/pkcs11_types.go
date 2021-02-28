package module

import "unsafe"

type CK_BYTE byte
type CK_ULONG uint
type CK_VOID_PTR unsafe.Pointer

const CK_FALSE = 0
const CK_TRUE = 1

var FalseAttr = []byte{CK_FALSE}
var TrueAttr = []byte{CK_TRUE}

type CK_RSA_PKCS_OAEP_PARAMS struct {
	HashAlg         CK_MECHANISM_TYPE
	Mgf             CK_RSA_PKCS_MGF_TYPE
	Source          CK_RSA_PKCS_OAEP_SOURCE_TYPE
	PSourceData     CK_VOID_PTR
	ULSourceDataLen CK_ULONG
}

// created with
// sed -nE 's/typedef +(CK_[^ ]+) +(CK_[^ ]+);/type \2 \1/p' core/pkcs11t.h

type CK_CHAR CK_BYTE
type CK_UTF8CHAR CK_BYTE
type CK_BBOOL CK_BYTE
type CK_FLAGS CK_ULONG
type CK_NOTIFICATION CK_ULONG
type CK_SLOT_ID CK_ULONG
type CK_SESSION_HANDLE CK_ULONG
type CK_USER_TYPE CK_ULONG
type CK_STATE CK_ULONG
type CK_OBJECT_HANDLE CK_ULONG
type CK_OBJECT_CLASS CK_ULONG
type CK_HW_FEATURE_TYPE CK_ULONG
type CK_KEY_TYPE CK_ULONG
type CK_CERTIFICATE_TYPE CK_ULONG
type CK_ATTRIBUTE_TYPE CK_ULONG
type CK_MECHANISM_TYPE CK_ULONG
type CK_RV CK_ULONG
type CK_RSA_PKCS_MGF_TYPE CK_ULONG
type CK_RSA_PKCS_OAEP_SOURCE_TYPE CK_ULONG
type CK_EC_KDF_TYPE CK_ULONG
type CK_X9_42_DH_KDF_TYPE CK_ULONG
type CK_RC2_PARAMS CK_ULONG
type CK_MAC_GENERAL_PARAMS CK_ULONG
type CK_EXTRACT_PARAMS CK_ULONG
type CK_PKCS5_PBKD2_PSEUDO_RANDOM_FUNCTION_TYPE CK_ULONG
type CK_PKCS5_PBKDF2_SALT_SOURCE_TYPE CK_ULONG
type CK_OTP_PARAM_TYPE CK_ULONG
type CK_PARAM_TYPE CK_OTP_PARAM_TYPE /* backward compatibility */
type CK_JAVA_MIDP_SECURITY_DOMAIN CK_ULONG
type CK_CERTIFICATE_CATEGORY CK_ULONG

// type CK_RV CK_ULONG
// type CK_OBJECT_HANDLE CK_ULONG
// type CK_SESSION_HANDLE CK_ULONG
// type CK_SLOT_ID CK_ULONG
// type CK_USER_TYPE CK_ULONG
// type CK_ATTRIBUTE_TYPE CK_ULONG
// type CK_MECHANISM_TYPE CK_ULONG
// type CK_FLAGS CK_ULONG
// type CK_STATE CK_ULONG

// type CK_CHAR C.CK_CHAR
// type CK_UTF8CHAR C.CK_UTF8CHAR

// type CK_LONG C.CK_LONG
// type CK_BYTE_PTR C.CK_BYTE_PTR
// type CK_CHAR_PTR C.CK_CHAR_PTR
// type CK_UTF8CHAR_PTR C.CK_UTF8CHAR_PTR
// type CK_ULONG_PTR C.CK_ULONG_PTR
// type CK_VOID_PTR C.CK_VOID_PTR
// type CK_VOID_PTR_PTR C.CK_VOID_PTR_PTR
// type CK_VERSION C.CK_VERSION
// type CK_VERSION_PTR C.CK_VERSION_PTR
// type CK_INFO C.CK_INFO
// type CK_INFO_PTR C.CK_INFO_PTR
// type CK_NOTIFICATION C.CK_NOTIFICATION
// type CK_SLOT_ID_PTR C.CK_SLOT_ID_PTR
// type CK_SLOT_INFO C.CK_SLOT_INFO
// type CK_SLOT_INFO_PTR C.CK_SLOT_INFO_PTR
// type CK_TOKEN_INFO C.CK_TOKEN_INFO
// type CK_TOKEN_INFO_PTR C.CK_TOKEN_INFO_PTR
// type CK_SESSION_HANDLE_PTR C.CK_SESSION_HANDLE_PTR
// type CK_SESSION_INFO C.CK_SESSION_INFO
// type CK_SESSION_INFO_PTR C.CK_SESSION_INFO_PTR

// type CK_OBJECT_HANDLE_PTR C.CK_OBJECT_HANDLE_PTR
// type CK_OBJECT_CLASS C.CK_OBJECT_CLASS
// type CK_OBJECT_CLASS_PTR C.CK_OBJECT_CLASS_PTR
// type CK_HW_FEATURE_TYPE C.CK_HW_FEATURE_TYPE
// type CK_KEY_TYPE C.CK_KEY_TYPE
// type CK_CERTIFICATE_TYPE C.CK_CERTIFICATE_TYPE
// type CK_ATTRIBUTE C.CK_ATTRIBUTE
// type CK_ATTRIBUTE_PTR C.CK_ATTRIBUTE_PTR
// type CK_DATE C.CK_DATE
// type CK_MECHANISM_TYPE_PTR C.CK_MECHANISM_TYPE_PTR
// type CK_MECHANISM C.CK_MECHANISM
// type CK_MECHANISM_PTR C.CK_MECHANISM_PTR
// type CK_MECHANISM_INFO C.CK_MECHANISM_INFO
// type CK_MECHANISM_INFO_PTR C.CK_MECHANISM_INFO_PTR
// type CK_FUNCTION_LIST C.CK_FUNCTION_LIST
// type CK_FUNCTION_LIST_PTR C.CK_FUNCTION_LIST_PTR
// type CK_FUNCTION_LIST_PTR_PTR C.CK_FUNCTION_LIST_PTR_PTR
// type CK_C_INITIALIZE_ARGS C.CK_C_INITIALIZE_ARGS
// type CK_C_INITIALIZE_ARGS_PTR C.CK_C_INITIALIZE_ARGS_PTR
// type CK_RSA_PKCS_MGF_TYPE C.CK_RSA_PKCS_MGF_TYPE
// type CK_RSA_PKCS_MGF_TYPE_PTR C.CK_RSA_PKCS_MGF_TYPE_PTR
// type CK_RSA_PKCS_OAEP_SOURCE_TYPE C.CK_RSA_PKCS_OAEP_SOURCE_TYPE
// type CK_RSA_PKCS_OAEP_SOURCE_TYPE_PTR C.CK_RSA_PKCS_OAEP_SOURCE_TYPE_PTR
// type CK_RSA_PKCS_OAEP_PARAMS C.CK_RSA_PKCS_OAEP_PARAMS
// type CK_RSA_PKCS_OAEP_PARAMS_PTR C.CK_RSA_PKCS_OAEP_PARAMS_PTR
// type CK_RSA_PKCS_PSS_PARAMS C.CK_RSA_PKCS_PSS_PARAMS
// type CK_RSA_PKCS_PSS_PARAMS_PTR C.CK_RSA_PKCS_PSS_PARAMS_PTR
// type CK_EC_KDF_TYPE C.CK_EC_KDF_TYPE
// type CK_ECDH1_DERIVE_PARAMS C.CK_ECDH1_DERIVE_PARAMS
// type CK_ECDH1_DERIVE_PARAMS_PTR C.CK_ECDH1_DERIVE_PARAMS_PTR
// type CK_ECDH2_DERIVE_PARAMS C.CK_ECDH2_DERIVE_PARAMS
// type CK_ECDH2_DERIVE_PARAMS_PTR C.CK_ECDH2_DERIVE_PARAMS_PTR
// type CK_ECMQV_DERIVE_PARAMS C.CK_ECMQV_DERIVE_PARAMS
// type CK_ECMQV_DERIVE_PARAMS_PTR C.CK_ECMQV_DERIVE_PARAMS_PTR
// type CK_X9_42_DH_KDF_TYPE C.CK_X9_42_DH_KDF_TYPE
// type CK_X9_42_DH_KDF_TYPE_PTR C.CK_X9_42_DH_KDF_TYPE_PTR
// type CK_X9_42_DH1_DERIVE_PARAMS C.CK_X9_42_DH1_DERIVE_PARAMS
// type CK_X9_42_DH1_DERIVE_PARAMS_PTR C.CK_X9_42_DH1_DERIVE_PARAMS_PTR
// type CK_X9_42_DH2_DERIVE_PARAMS C.CK_X9_42_DH2_DERIVE_PARAMS
// type CK_X9_42_DH2_DERIVE_PARAMS_PTR C.CK_X9_42_DH2_DERIVE_PARAMS_PTR
// type CK_X9_42_MQV_DERIVE_PARAMS C.CK_X9_42_MQV_DERIVE_PARAMS
// type CK_X9_42_MQV_DERIVE_PARAMS_PTR C.CK_X9_42_MQV_DERIVE_PARAMS_PTR
// type CK_KEA_DERIVE_PARAMS C.CK_KEA_DERIVE_PARAMS
// type CK_KEA_DERIVE_PARAMS_PTR C.CK_KEA_DERIVE_PARAMS_PTR
// type CK_RC2_PARAMS C.CK_RC2_PARAMS
// type CK_RC2_PARAMS_PTR C.CK_RC2_PARAMS_PTR
// type CK_RC2_CBC_PARAMS C.CK_RC2_CBC_PARAMS
// type CK_RC2_CBC_PARAMS_PTR C.CK_RC2_CBC_PARAMS_PTR
// type CK_RC2_MAC_GENERAL_PARAMS C.CK_RC2_MAC_GENERAL_PARAMS
// type CK_RC2_MAC_GENERAL_PARAMS_PTR C.CK_RC2_MAC_GENERAL_PARAMS_PTR
// type CK_RC5_PARAMS C.CK_RC5_PARAMS
// type CK_RC5_PARAMS_PTR C.CK_RC5_PARAMS_PTR
// type CK_RC5_CBC_PARAMS C.CK_RC5_CBC_PARAMS
// type CK_RC5_CBC_PARAMS_PTR C.CK_RC5_CBC_PARAMS_PTR
// type CK_RC5_MAC_GENERAL_PARAMS C.CK_RC5_MAC_GENERAL_PARAMS
// type CK_RC5_MAC_GENERAL_PARAMS_PTR C.CK_RC5_MAC_GENERAL_PARAMS_PTR
// type CK_MAC_GENERAL_PARAMS C.CK_MAC_GENERAL_PARAMS
// type CK_MAC_GENERAL_PARAMS_PTR C.CK_MAC_GENERAL_PARAMS_PTR
// type CK_DES_CBC_ENCRYPT_DATA_PARAMS C.CK_DES_CBC_ENCRYPT_DATA_PARAMS
// type CK_DES_CBC_ENCRYPT_DATA_PARAMS_PTR C.CK_DES_CBC_ENCRYPT_DATA_PARAMS_PTR
// type CK_AES_CBC_ENCRYPT_DATA_PARAMS C.CK_AES_CBC_ENCRYPT_DATA_PARAMS
// type CK_AES_CBC_ENCRYPT_DATA_PARAMS_PTR C.CK_AES_CBC_ENCRYPT_DATA_PARAMS_PTR
// type CK_SKIPJACK_PRIVATE_WRAP_PARAMS C.CK_SKIPJACK_PRIVATE_WRAP_PARAMS
// type CK_SKIPJACK_PRIVATE_WRAP_PARAMS_PTR C.CK_SKIPJACK_PRIVATE_WRAP_PARAMS_PTR
// type CK_SKIPJACK_RELAYX_PARAMS C.CK_SKIPJACK_RELAYX_PARAMS
// type CK_SKIPJACK_RELAYX_PARAMS_PTR C.CK_SKIPJACK_RELAYX_PARAMS_PTR
// type CK_PBE_PARAMS C.CK_PBE_PARAMS
// type CK_PBE_PARAMS_PTR C.CK_PBE_PARAMS_PTR
// type CK_KEY_WRAP_SET_OAEP_PARAMS C.CK_KEY_WRAP_SET_OAEP_PARAMS
// type CK_KEY_WRAP_SET_OAEP_PARAMS_PTR C.CK_KEY_WRAP_SET_OAEP_PARAMS_PTR
// type CK_SSL3_RANDOM_DATA C.CK_SSL3_RANDOM_DATA
// type CK_SSL3_MASTER_KEY_DERIVE_PARAMS C.CK_SSL3_MASTER_KEY_DERIVE_PARAMS
// type CK_SSL3_MASTER_KEY_DERIVE_PARAMS_PTR C.CK_SSL3_MASTER_KEY_DERIVE_PARAMS_PTR
// type CK_SSL3_KEY_MAT_OUT C.CK_SSL3_KEY_MAT_OUT
// type CK_SSL3_KEY_MAT_OUT_PTR C.CK_SSL3_KEY_MAT_OUT_PTR
// type CK_SSL3_KEY_MAT_PARAMS C.CK_SSL3_KEY_MAT_PARAMS
// type CK_SSL3_KEY_MAT_PARAMS_PTR C.CK_SSL3_KEY_MAT_PARAMS_PTR
// type CK_TLS_PRF_PARAMS C.CK_TLS_PRF_PARAMS
// type CK_TLS_PRF_PARAMS_PTR C.CK_TLS_PRF_PARAMS_PTR
// type CK_WTLS_RANDOM_DATA C.CK_WTLS_RANDOM_DATA
// type CK_WTLS_RANDOM_DATA_PTR C.CK_WTLS_RANDOM_DATA_PTR
// type CK_WTLS_MASTER_KEY_DERIVE_PARAMS C.CK_WTLS_MASTER_KEY_DERIVE_PARAMS
// type CK_WTLS_MASTER_KEY_DERIVE_PARAMS_PTR C.CK_WTLS_MASTER_KEY_DERIVE_PARAMS_PTR
// type CK_WTLS_PRF_PARAMS C.CK_WTLS_PRF_PARAMS
// type CK_WTLS_PRF_PARAMS_PTR C.CK_WTLS_PRF_PARAMS_PTR
// type CK_WTLS_KEY_MAT_OUT C.CK_WTLS_KEY_MAT_OUT
// type CK_WTLS_KEY_MAT_OUT_PTR C.CK_WTLS_KEY_MAT_OUT_PTR
// type CK_WTLS_KEY_MAT_PARAMS C.CK_WTLS_KEY_MAT_PARAMS
// type CK_WTLS_KEY_MAT_PARAMS_PTR C.CK_WTLS_KEY_MAT_PARAMS_PTR
// type CK_CMS_SIG_PARAMS C.CK_CMS_SIG_PARAMS
// type CK_CMS_SIG_PARAMS_PTR C.CK_CMS_SIG_PARAMS_PTR
// type CK_KEY_DERIVATION_STRING_DATA C.CK_KEY_DERIVATION_STRING_DATA
// type CK_KEY_DERIVATION_STRING_DATA_PTR C.CK_KEY_DERIVATION_STRING_DATA_PTR
// type CK_EXTRACT_PARAMS C.CK_EXTRACT_PARAMS
// type CK_EXTRACT_PARAMS_PTR C.CK_EXTRACT_PARAMS_PTR
// type CK_PKCS5_PBKD2_PSEUDO_RANDOM_FUNCTION_TYPE C.CK_PKCS5_PBKD2_PSEUDO_RANDOM_FUNCTION_TYPE
// type CK_PKCS5_PBKD2_PSEUDO_RANDOM_FUNCTION_TYPE_PTR C.CK_PKCS5_PBKD2_PSEUDO_RANDOM_FUNCTION_TYPE_PTR
// type CK_PKCS5_PBKDF2_SALT_SOURCE_TYPE C.CK_PKCS5_PBKDF2_SALT_SOURCE_TYPE
// type CK_PKCS5_PBKDF2_SALT_SOURCE_TYPE_PTR C.CK_PKCS5_PBKDF2_SALT_SOURCE_TYPE_PTR
// type CK_PKCS5_PBKD2_PARAMS C.CK_PKCS5_PBKD2_PARAMS
// type CK_PKCS5_PBKD2_PARAMS_PTR C.CK_PKCS5_PBKD2_PARAMS_PTR
// type CK_PKCS5_PBKD2_PARAMS2 C.CK_PKCS5_PBKD2_PARAMS2
// type CK_PKCS5_PBKD2_PARAMS2_PTR C.CK_PKCS5_PBKD2_PARAMS2_PTR
// type CK_OTP_PARAM_TYPE C.CK_OTP_PARAM_TYPE
// type CK_PARAM_TYPE C.CK_PARAM_TYPE
// type CK_OTP_PARAM C.CK_OTP_PARAM
// type CK_OTP_PARAM_PTR C.CK_OTP_PARAM_PTR
// type CK_OTP_PARAMS C.CK_OTP_PARAMS
// type CK_OTP_PARAMS_PTR C.CK_OTP_PARAMS_PTR
// type CK_OTP_SIGNATURE_INFO C.CK_OTP_SIGNATURE_INFO
// type CK_OTP_SIGNATURE_INFO_PTR C.CK_OTP_SIGNATURE_INFO_PTR
// type CK_KIP_PARAMS C.CK_KIP_PARAMS
// type CK_KIP_PARAMS_PTR C.CK_KIP_PARAMS_PTR
// type CK_AES_CTR_PARAMS C.CK_AES_CTR_PARAMS
// type CK_AES_CTR_PARAMS_PTR C.CK_AES_CTR_PARAMS_PTR
// type CK_GCM_PARAMS C.CK_GCM_PARAMS
// type CK_GCM_PARAMS_PTR C.CK_GCM_PARAMS_PTR
// type CK_CCM_PARAMS C.CK_CCM_PARAMS
// type CK_CCM_PARAMS_PTR C.CK_CCM_PARAMS_PTR
// type CK_AES_GCM_PARAMS C.CK_AES_GCM_PARAMS
// type CK_AES_GCM_PARAMS_PTR C.CK_AES_GCM_PARAMS_PTR
// type CK_AES_CCM_PARAMS C.CK_AES_CCM_PARAMS
// type CK_AES_CCM_PARAMS_PTR C.CK_AES_CCM_PARAMS_PTR
// type CK_CAMELLIA_CTR_PARAMS C.CK_CAMELLIA_CTR_PARAMS
// type CK_CAMELLIA_CTR_PARAMS_PTR C.CK_CAMELLIA_CTR_PARAMS_PTR
// type CK_CAMELLIA_CBC_ENCRYPT_DATA_PARAMS C.CK_CAMELLIA_CBC_ENCRYPT_DATA_PARAMS
// type CK_CAMELLIA_CBC_ENCRYPT_DATA_PARAMS_PTR C.CK_CAMELLIA_CBC_ENCRYPT_DATA_PARAMS_PTR
// type CK_ARIA_CBC_ENCRYPT_DATA_PARAMS C.CK_ARIA_CBC_ENCRYPT_DATA_PARAMS
// type CK_ARIA_CBC_ENCRYPT_DATA_PARAMS_PTR C.CK_ARIA_CBC_ENCRYPT_DATA_PARAMS_PTR
// type CK_DSA_PARAMETER_GEN_PARAM C.CK_DSA_PARAMETER_GEN_PARAM
// type CK_DSA_PARAMETER_GEN_PARAM_PTR C.CK_DSA_PARAMETER_GEN_PARAM_PTR
// type CK_ECDH_AES_KEY_WRAP_PARAMS C.CK_ECDH_AES_KEY_WRAP_PARAMS
// type CK_ECDH_AES_KEY_WRAP_PARAMS_PTR C.CK_ECDH_AES_KEY_WRAP_PARAMS_PTR
// type CK_JAVA_MIDP_SECURITY_DOMAIN C.CK_JAVA_MIDP_SECURITY_DOMAIN
// type CK_CERTIFICATE_CATEGORY C.CK_CERTIFICATE_CATEGORY
// type CK_RSA_AES_KEY_WRAP_PARAMS C.CK_RSA_AES_KEY_WRAP_PARAMS
// type CK_RSA_AES_KEY_WRAP_PARAMS_PTR C.CK_RSA_AES_KEY_WRAP_PARAMS_PTR
// type CK_TLS12_MASTER_KEY_DERIVE_PARAMS C.CK_TLS12_MASTER_KEY_DERIVE_PARAMS
// type CK_TLS12_MASTER_KEY_DERIVE_PARAMS_PTR C.CK_TLS12_MASTER_KEY_DERIVE_PARAMS_PTR
// type CK_TLS12_KEY_MAT_PARAMS C.CK_TLS12_KEY_MAT_PARAMS
// type CK_TLS12_KEY_MAT_PARAMS_PTR C.CK_TLS12_KEY_MAT_PARAMS_PTR
// type CK_TLS_KDF_PARAMS C.CK_TLS_KDF_PARAMS
// type CK_TLS_KDF_PARAMS_PTR C.CK_TLS_KDF_PARAMS_PTR
// type CK_TLS_MAC_PARAMS C.CK_TLS_MAC_PARAMS
// type CK_TLS_MAC_PARAMS_PTR C.CK_TLS_MAC_PARAMS_PTR
// type CK_GOSTR3410_DERIVE_PARAMS C.CK_GOSTR3410_DERIVE_PARAMS
// type CK_GOSTR3410_DERIVE_PARAMS_PTR C.CK_GOSTR3410_DERIVE_PARAMS_PTR
// type CK_GOSTR3410_KEY_WRAP_PARAMS C.CK_GOSTR3410_KEY_WRAP_PARAMS
// type CK_GOSTR3410_KEY_WRAP_PARAMS_PTR C.CK_GOSTR3410_KEY_WRAP_PARAMS_PTR
// type CK_SEED_CBC_ENCRYPT_DATA_PARAMS C.CK_SEED_CBC_ENCRYPT_DATA_PARAMS
// type CK_SEED_CBC_ENCRYPT_DATA_PARAMS_PTR C.CK_SEED_CBC_ENCRYPT_DATA_PARAMS_PTR
