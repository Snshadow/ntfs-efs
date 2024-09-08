package w32api

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// application-defined callback functions for ReadEncryptedFileRaw and WriteEncryptedFileRaw

/*
DWORD PfeExportFunc(
  [in]           PBYTE pbData,
  [in, optional] PVOID pvCallbackContext,
  [in]           ULONG ulLength
)
{...}
*/
type PRE_EXPORT_FUNC func(data *byte, callbackContext unsafe.Pointer, length *uint32) uintptr

/*
DWORD PfeImportFunc(
  [in]           PBYTE pbData,
  [in, optional] PVOID pvCallbackContext,
  [in, out]      PULONG ulLength
)
{...}
*/
type PRE_IMPORT_FUNC func(data *byte, callbackContext unsafe.Pointer, length uint32) uintptr

// flag for OpenEncryptedFileRaw
const (
	CREATE_FOR_IMPORT = 1
	CREATE_FOR_DIR    = 2
	OVERWRITE_HIDDEN  = 4
)

type EFS_CERTIFICATE_BLOB struct {
	CertEncodingType uint32
	CbData           uint32
	PbData           *byte
}

type ENCRYPTION_CERTIFICATE struct {
	TotalLength uint32
	UserSid     *windows.SID
	CertBlob    *EFS_CERTIFICATE_BLOB
}

type ENCRYPTION_CERTIFICATE_LIST struct {
	NumUser uint32
	Users   *ENCRYPTION_CERTIFICATE // []ENCRYPTION_CERTIFICATE
}

type EFS_HASH_BLOB struct {
	CbData uint32
	PbData *byte
}

type ENCRYPTION_CERTIFICATE_HASH struct {
	TotalLength        uint32
	UserSid            *windows.SID
	Hash               *EFS_HASH_BLOB
	DisplayInformation *uint16
}

type ENCRYPTION_CERTIFICATE_HASH_LIST struct {
	NumCertHash uint32
	Users       *ENCRYPTION_CERTIFICATE_HASH // []ENCRYPTION_CERTIFICATE_HASH
}
