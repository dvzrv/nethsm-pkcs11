package core

/*
#include "pkcs11go.h"
*/
import "C"
import (
	"encoding/binary"
)

// SignContext represents a structure which groups parameters that allow to sign
// a document.
type SignContext interface {
	Init(metaBytes []byte) error
	SignatureLength() int
	Update(data []byte) error
	Final() ([]byte, error)
	Initialized() bool
}

// VerifyContext represents a structure which groups parameters that allow to verify a signature of
// a document.
type VerifyContext interface {
	Init(metaBytes []byte) error
	Length() int
	Update(data []byte) error
	Final([]byte) error
	Initialized() bool
}

func NewSignContext(session *Session, mechanism *Mechanism, hKey C.CK_OBJECT_HANDLE) (context SignContext, err error) {
	keyObject, err := session.GetObject(hKey)
	if err != nil {
		return nil, err
	}
	keyIDAttr := keyObject.FindAttribute(C.CKA_ID)
	if keyIDAttr == nil {
		return nil, NewError("NewSignContext", "object does not contain a key ID", C.CKR_ARGUMENTS_BAD)
	}
	// keyMetaAttr := keyObject.FindAttribute(AttrTypeKeyMeta)
	// if keyMetaAttr == nil {
	// 	return nil, NewError(" NewSignContext", "object handle does not contain any key metainfo attribute", C.CKR_ARGUMENTS_BAD)
	// }
	//log.Printf("Type: %v", mechanism.Type)
	switch mechanism.Type {
	case C.CKM_RSA_PKCS, C.CKM_MD5_RSA_PKCS, C.CKM_SHA1_RSA_PKCS, C.CKM_SHA256_RSA_PKCS, C.CKM_SHA384_RSA_PKCS, C.CKM_SHA512_RSA_PKCS, C.CKM_RSA_PKCS_PSS, C.CKM_SHA1_RSA_PKCS_PSS, C.CKM_SHA256_RSA_PKCS_PSS, C.CKM_SHA384_RSA_PKCS_PSS, C.CKM_SHA512_RSA_PKCS_PSS:
		c := &SignContextRSA{
			// randSrc:   session.randSrc,
			apiCtx:    session.apiCtx,
			keyID:     string(keyIDAttr.Value),
			mechanism: mechanism,
			data:      make([]byte, 0),
		}
		if err := c.Init(nil); err != nil {
			return nil, err
		}
		context = c
	case C.CKM_ECDSA, C.CKM_ECDSA_SHA1, C.CKM_ECDSA_SHA256, C.CKM_ECDSA_SHA384, C.CKM_ECDSA_SHA512:
		// Get PK
		pkAttr := keyObject.FindAttribute(C.CKA_EC_POINT)
		if pkAttr == nil {
			return nil, NewError("NewSignContext", "object handle does not contain any ec public key attribute", C.CKR_ARGUMENTS_BAD)
		}
		c := &ECDSASignContext{
			// randSrc:   session.randSrc,
			keyID:     string(keyIDAttr.Value),
			mechanism: mechanism,
			data:      make([]byte, 0),
		}
		if err := c.Init(nil); err != nil {
			return nil, err
		}
		// pk, err := utils.ASN1BytesToPubKey(c.pubKey.Curve, pkAttr.Value)
		// if err != nil {
		// 	return nil, NewError("NewSignContext", fmt.Sprintf("%s", err), C.CKR_ARGUMENTS_BAD)
		// }
		// c.pubKey = pk
		context = c
	default:
		err = NewError("NewSignContext", "sign mechanism invalid", C.CKR_MECHANISM_INVALID)
		return nil, err
	}
	return context, nil
}

func NewVerifyContext(session *Session, mechanism *Mechanism, hKey C.CK_OBJECT_HANDLE) (context VerifyContext, err error) {
	keyObject, err := session.GetObject(hKey)
	if err != nil {
		return nil, err
	}
	keyIDAttr := keyObject.FindAttribute(C.CKA_ID)
	if keyIDAttr == nil {
		return nil, NewError("NewSignContext", "object does not contain a key ID", C.CKR_ARGUMENTS_BAD)
	}
	// keyMetaAttr := keyObject.FindAttribute(AttrTypeKeyMeta)
	// if keyMetaAttr == nil {
	// 	return nil, NewError(" NewSignContext", "object handle does not contain any key metainfo attribute", C.CKR_ARGUMENTS_BAD)
	// }

	switch mechanism.Type {
	case C.CKM_RSA_PKCS, C.CKM_MD5_RSA_PKCS, C.CKM_SHA1_RSA_PKCS, C.CKM_SHA256_RSA_PKCS, C.CKM_SHA384_RSA_PKCS, C.CKM_SHA512_RSA_PKCS, C.CKM_SHA1_RSA_PKCS_PSS, C.CKM_SHA256_RSA_PKCS_PSS, C.CKM_SHA384_RSA_PKCS_PSS, C.CKM_SHA512_RSA_PKCS_PSS:
		c := &VerifyContextRSA{
			randSrc:   session.randSrc,
			keyID:     string(keyIDAttr.Value),
			mechanism: mechanism,
			data:      make([]byte, 0),
		}
		if err := c.Init(nil); err != nil {
			return nil, err
		}
		context = c
	case C.CKM_ECDSA, C.CKM_ECDSA_SHA1, C.CKM_ECDSA_SHA256, C.CKM_ECDSA_SHA384, C.CKM_ECDSA_SHA512:
		// Get PK
		pkAttr := keyObject.FindAttribute(C.CKA_EC_POINT)
		if pkAttr == nil {
			return nil, NewError("NewVerifyContext", "object handle does not contain any ec public key attribute", C.CKR_ARGUMENTS_BAD)
		}
		c := &ECDSAVerifyContext{
			// randSrc:   session.randSrc,
			keyID:     string(keyIDAttr.Value),
			mechanism: mechanism,
			data:      make([]byte, 0),
		}
		if err := c.Init(nil); err != nil {
			return nil, err
		}
		// pk, err := utils.ASN1BytesToPubKey(c.pubKey.Curve, pkAttr.Value)
		// if err != nil {
		// 	return nil, NewError("NewVerifyContext", fmt.Sprintf("%s", err), C.CKR_ARGUMENTS_BAD)
		// }
		// c.pubKey = pk
		context = c
	default:
		err = NewError("NewVerifyContext", "sign mechanism invalid", C.CKR_MECHANISM_INVALID)
		return nil, err
	}
	return context, nil
}

func ulongToArr(n C.ulong) []byte {
	arr := make([]byte, 8)
	binary.LittleEndian.PutUint64(arr, uint64(n))
	return arr
}

func boolToArr(n C.CK_BBOOL) []byte {
	return []byte{byte(n)}
}
