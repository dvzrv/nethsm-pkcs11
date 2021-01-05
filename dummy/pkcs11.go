package main

/*
#include <stdlib.h>
#include <string.h>
#include "pkcs11go.h"
extern CK_FUNCTION_LIST functionList;
*/
import "C"
import (
	"strings"
	"unsafe"
)

type logT struct{}

var log logT

func (v *logT) Panicf(...interface{}) {}
func (v *logT) Printf(...interface{}) {}

const (
	libManufacturerID = "Nitrokey GmbH"
	libDescription    = "NetHSM PKCS#11 module"
	libVersionMajor   = 0
	libVersionMinor   = 1
)

func str2Buf(s string, buffer interface{}) {
	var length int
	var pBuffer unsafe.Pointer
	if b, ok := buffer.(*[16]C.uchar); ok {
		length = 16
		pBuffer = unsafe.Pointer(&b[0])
	} else if b, ok := buffer.(*[32]C.uchar); ok {
		length = 32
		pBuffer = unsafe.Pointer(&b[0])
	} else if b, ok := buffer.(*[64]C.uchar); ok {
		length = 64
		pBuffer = unsafe.Pointer(&b[0])
	} else {
		log.Panicf("strBuf: buffer type %T not supported", buffer)
	}
	if len(s) > length {
		s = s[:length]
	}
	s += strings.Repeat(" ", length-len(s))
	cBytes := C.CBytes([]byte(s))
	defer C.free(unsafe.Pointer(cBytes))
	C.memcpy(pBuffer, cBytes, (C.size_t)(length))
}

//export C_Initialize
func C_Initialize(pInitArgs C.CK_VOID_PTR) C.CK_RV {
	log.Printf("Called: C_Initialize")
	return C.CKR_OK
}

//export C_Finalize
func C_Finalize(pReserved C.CK_VOID_PTR) C.CK_RV {
	log.Printf("Called: C_Finalize")
	return C.CKR_OK
}

//export C_InitToken
func C_InitToken(slotID C.CK_SLOT_ID, pPin C.CK_UTF8CHAR_PTR, ulPinLen C.CK_ULONG, pLabel C.CK_UTF8CHAR_PTR) C.CK_RV {
	log.Printf("Called: C_InitToken")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_InitPIN
func C_InitPIN(hSession C.CK_SESSION_HANDLE, pPin C.CK_UTF8CHAR_PTR, ulPinLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_InitPIN")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SetPIN
func C_SetPIN(hSession C.CK_SESSION_HANDLE, pOldPin C.CK_UTF8CHAR_PTR, ulOldPinLen C.CK_ULONG, pNewPin C.CK_UTF8CHAR_PTR, ulNewPinLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_SetPIN")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetInfo
func C_GetInfo(pInfo C.CK_INFO_PTR) C.CK_RV {
	log.Printf("Called: C_GetInfo\n")
	if pInfo == nil {
		return C.CKR_ARGUMENTS_BAD
	}
	info := (*C.CK_INFO)(unsafe.Pointer(pInfo))

	log.Printf("%v", &info.manufacturerID[0])
	str2Buf(libManufacturerID, &info.manufacturerID)
	str2Buf(libDescription, &info.libraryDescription)

	info.flags = 0
	info.cryptokiVersion.major = 2
	info.cryptokiVersion.minor = 40
	info.libraryVersion.major = libVersionMajor
	info.libraryVersion.minor = libVersionMinor
	return C.CKR_OK
}

//export C_GetFunctionList
func C_GetFunctionList(ppFunctionList C.CK_FUNCTION_LIST_PTR_PTR) C.CK_RV {
	log.Printf("Called: C_GetFunctionList\n")
	if ppFunctionList == nil {
		return C.CKR_ARGUMENTS_BAD
	}
	*ppFunctionList = &C.functionList
	return C.CKR_OK
}

//export C_GetSlotList
func C_GetSlotList(tokenPresent C.CK_BBOOL, pSlotList C.CK_SLOT_ID_PTR, pulCount C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_GetSlotList")
	return C.CKR_OK
}

//export C_GetSlotInfo
func C_GetSlotInfo(slotId C.CK_SLOT_ID, pInfo C.CK_SLOT_INFO_PTR) C.CK_RV {
	log.Printf("Called: C_GetSlotInfo")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetTokenInfo
func C_GetTokenInfo(slotId C.CK_SLOT_ID, pInfo C.CK_TOKEN_INFO_PTR) C.CK_RV {
	log.Printf("Called: C_GetTokenInfo")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_OpenSession
func C_OpenSession(slotId C.CK_SLOT_ID, flags C.CK_FLAGS, pApplication C.CK_VOID_PTR, notify C.CK_NOTIFY, phSession C.CK_SESSION_HANDLE_PTR) C.CK_RV {
	log.Printf("Called: C_OpenSession")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_CloseSession
func C_CloseSession(hSession C.CK_SESSION_HANDLE) C.CK_RV {
	log.Printf("Called: C_CloseSession")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_CloseAllSessions
func C_CloseAllSessions(slotId C.CK_SLOT_ID) C.CK_RV {
	log.Printf("Called: C_CloseAllSessions")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetSessionInfo
func C_GetSessionInfo(hSession C.CK_SESSION_HANDLE, pInfo C.CK_SESSION_INFO_PTR) C.CK_RV {
	log.Printf("Called: C_GetSessionInfo")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_Login
func C_Login(hSession C.CK_SESSION_HANDLE, userType C.CK_USER_TYPE, pPin C.CK_UTF8CHAR_PTR, ulPinLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_Login")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_Logout
func C_Logout(hSession C.CK_SESSION_HANDLE) C.CK_RV {
	log.Printf("Called: C_Logout")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_CreateObject
func C_CreateObject(hSession C.CK_SESSION_HANDLE, pTemplate C.CK_ATTRIBUTE_PTR, ulCount C.CK_ULONG, phObject C.CK_OBJECT_HANDLE_PTR) C.CK_RV {
	log.Printf("Called: C_CreateObject")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DestroyObject
func C_DestroyObject(hSession C.CK_SESSION_HANDLE, hObject C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_DestroyObject")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_FindObjectsInit
func C_FindObjectsInit(hSession C.CK_SESSION_HANDLE, pTemplate C.CK_ATTRIBUTE_PTR, ulCount C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_FindObjectsInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_FindObjects
func C_FindObjects(hSession C.CK_SESSION_HANDLE, phObject C.CK_OBJECT_HANDLE_PTR, ulMaxObjectCount C.CK_ULONG, pulObjectCount C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_FindObjects")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_FindObjectsFinal
func C_FindObjectsFinal(hSession C.CK_SESSION_HANDLE) C.CK_RV {
	log.Printf("Called: C_FindObjectsFinal")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SetAttributeValue
func C_SetAttributeValue(hSession C.CK_SESSION_HANDLE, hObject C.CK_OBJECT_HANDLE, pTemplate C.CK_ATTRIBUTE_PTR,
	ulCount C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_SetAttributeValue")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetAttributeValue
func C_GetAttributeValue(hSession C.CK_SESSION_HANDLE, hObject C.CK_OBJECT_HANDLE, pTemplate C.CK_ATTRIBUTE_PTR, ulCount C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_GetAttributeValue")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GenerateKeyPair
func C_GenerateKeyPair(hSession C.CK_SESSION_HANDLE, pMechanism C.CK_MECHANISM_PTR, pPublicKeyTemplate C.CK_ATTRIBUTE_PTR, ulPublicKeyAttributeCount C.CK_ULONG, pPrivateKeyTemplate C.CK_ATTRIBUTE_PTR, ulPrivateKeyAttributeCount C.CK_ULONG, phPublicKey C.CK_OBJECT_HANDLE_PTR, phPrivateKey C.CK_OBJECT_HANDLE_PTR) C.CK_RV {
	log.Printf("Called: C_GenerateKeyPair")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SignInit
func C_SignInit(hSession C.CK_SESSION_HANDLE, pMechanism C.CK_MECHANISM_PTR, hKey C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_SignInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SignUpdate
func C_SignUpdate(hSession C.CK_SESSION_HANDLE, pPart C.CK_BYTE_PTR, ulPartLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_SignUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SignFinal
func C_SignFinal(hSession C.CK_SESSION_HANDLE, pSignature C.CK_BYTE_PTR, pulSignatureLen C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_SignFinal")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_Sign
func C_Sign(hSession C.CK_SESSION_HANDLE, pData C.CK_BYTE_PTR, ulDataLen C.CK_ULONG, pSignature C.CK_BYTE_PTR, pulSignatureLen C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_Sign")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_VerifyInit
func C_VerifyInit(hSession C.CK_SESSION_HANDLE, pMechanism C.CK_MECHANISM_PTR, hKey C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_VerifyInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_Verify
func C_Verify(hSession C.CK_SESSION_HANDLE, pData C.CK_BYTE_PTR, ulDataLen C.CK_ULONG, pSignature C.CK_BYTE_PTR, ulSignatureLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_Verify")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_VerifyUpdate
func C_VerifyUpdate(hSession C.CK_SESSION_HANDLE, pPart C.CK_BYTE_PTR, ulPartLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_VerifyUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_VerifyFinal
func C_VerifyFinal(hSession C.CK_SESSION_HANDLE, pSignature C.CK_BYTE_PTR, ulSignatureLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_VerifyFinal")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DigestInit
func C_DigestInit(hSession C.CK_SESSION_HANDLE, pMechanism C.CK_MECHANISM_PTR) C.CK_RV {
	log.Printf("Called: C_DigestInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_Digest
func C_Digest(hSession C.CK_SESSION_HANDLE, pData C.CK_BYTE_PTR, ulDataLen C.CK_ULONG, pDigest C.CK_BYTE_PTR, pulDigestLen C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_Digest")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SeedRandom
func C_SeedRandom(hSession C.CK_SESSION_HANDLE, pSeed C.CK_BYTE_PTR, ulSeedLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_SeedRandom")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GenerateRandom
func C_GenerateRandom(hSession C.CK_SESSION_HANDLE, pRandomData C.CK_BYTE_PTR, ulRandomLen C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_GenerateRandom")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

// NOTE: Not implemented functions...

//export C_GetMechanismList
func C_GetMechanismList(C.CK_SLOT_ID, C.CK_MECHANISM_TYPE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_GetMechanismList")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetMechanismInfo
func C_GetMechanismInfo(C.CK_SLOT_ID, C.CK_MECHANISM_TYPE, C.CK_MECHANISM_INFO_PTR) C.CK_RV {
	log.Printf("Called: C_GetMechanismInfo")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetOperationState
func C_GetOperationState(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_GetOperationState")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SetOperationState
func C_SetOperationState(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_OBJECT_HANDLE, C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_SetOperationState")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_CopyObject
func C_CopyObject(C.CK_SESSION_HANDLE, C.CK_OBJECT_HANDLE, C.CK_ATTRIBUTE_PTR, C.CK_ULONG, C.CK_OBJECT_HANDLE_PTR) C.CK_RV {
	log.Printf("Called: C_CopyObject")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetObjectSize
func C_GetObjectSize(C.CK_SESSION_HANDLE, C.CK_OBJECT_HANDLE, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_GetObjectSize")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_EncryptInit
func C_EncryptInit(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_EncryptInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_Encrypt
func C_Encrypt(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG,
	C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_EncryptUpdate
func C_EncryptUpdate(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_EncryptUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_EncryptFinal
func C_EncryptFinal(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_EncryptFinal")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DecryptInit
func C_DecryptInit(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_DecryptInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_Decrypt
func C_Decrypt(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG,
	C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DecryptUpdate
func C_DecryptUpdate(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_DecryptUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DecryptFinal
func C_DecryptFinal(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_DecryptFinal")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DigestUpdate
func C_DigestUpdate(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG) C.CK_RV {
	log.Printf("Called: C_DigestUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DigestKey
func C_DigestKey(C.CK_SESSION_HANDLE, C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_DigestKey")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DigestFinal
func C_DigestFinal(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_DigestFinal")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SignRecoverInit
func C_SignRecoverInit(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_SignRecoverInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SignRecover
func C_SignRecover(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_SignRecover")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_VerifyRecoverInit
func C_VerifyRecoverInit(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_OBJECT_HANDLE) C.CK_RV {
	log.Printf("Called: C_VerifyRecoverInit")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_VerifyRecover
func C_VerifyRecover(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_VerifyRecover")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DigestEncryptUpdate
func C_DigestEncryptUpdate(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_DigestEncryptUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DecryptDigestUpdate
func C_DecryptDigestUpdate(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_DecryptDigestUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_SignEncryptUpdate
func C_SignEncryptUpdate(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_SignEncryptUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DecryptVerifyUpdate
func C_DecryptVerifyUpdate(C.CK_SESSION_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG, C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	log.Printf("Called: C_DecryptVerifyUpdate")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GenerateKey
func C_GenerateKey(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_ATTRIBUTE_PTR, C.CK_ULONG, C.CK_OBJECT_HANDLE_PTR) C.CK_RV {
	log.Printf("Called: C_GenerateKey")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_WrapKey
func C_WrapKey(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_OBJECT_HANDLE, C.CK_OBJECT_HANDLE,
	C.CK_BYTE_PTR, C.CK_ULONG_PTR) C.CK_RV {
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_UnwrapKey
func C_UnwrapKey(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_OBJECT_HANDLE, C.CK_BYTE_PTR, C.CK_ULONG,
	C.CK_ATTRIBUTE_PTR, C.CK_ULONG, C.CK_OBJECT_HANDLE_PTR) C.CK_RV {
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_DeriveKey
func C_DeriveKey(C.CK_SESSION_HANDLE, C.CK_MECHANISM_PTR, C.CK_OBJECT_HANDLE, C.CK_ATTRIBUTE_PTR,
	C.CK_ULONG, C.CK_OBJECT_HANDLE_PTR) C.CK_RV {
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_GetFunctionStatus
func C_GetFunctionStatus(C.CK_SESSION_HANDLE) C.CK_RV {
	log.Printf("Called: C_GetFunctionStatus")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_CancelFunction
func C_CancelFunction(C.CK_SESSION_HANDLE) C.CK_RV {
	log.Printf("Called: C_CancelFunction")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

//export C_WaitForSlotEvent
func C_WaitForSlotEvent(C.CK_FLAGS, C.CK_SLOT_ID_PTR, C.CK_VOID_PTR) C.CK_RV {
	log.Printf("Called: C_WaitForSlotEvent")
	return C.CKR_FUNCTION_NOT_SUPPORTED
}

func main() {}
