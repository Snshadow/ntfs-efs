package w32api

import (
	"golang.org/x/sys/windows"
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
