package w32api

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// EFS file and directory handling functions

//sys	fileEncryptionStatusW(lpFileName *uint16, lpStatus *uint32) (err error) = Advapi32.FileEncryptionStatusW
//sys	encryptionDisable(dirPath *uint16, disable bool) (err error) = Advapi32.EncryptionDisable
//sys	encryptFileW(lpFileName *uint16) (err error) = Advapi32.EncryptFileW
//sys	decryptFileW(lpFileName *uint16, dwReserved uint32) (err error) = Advapi32.DecryptFileW

// EFS files and user key functions

//sys	addUsersToEncryptedFileW(lpFileName *uint16, pEncryptionCertificates *ENCRYPTION_CERTIFICATE_LIST) (ret error) = Advapi32.AddUsersToEncryptedFileW
//sys	duplicateEncryptionInfoFile(srcFileName *uint16, dstFileName *uint16, creationDistribution uint32, attributes uint32, lpSecurityAttributes *windows.SecurityAttributes) (ret error) = Advapi32.DuplicateEncryptionInfoFile
//sys	FreeEncryptionCertificateHashList(pUsers *ENCRYPTION_CERTIFICATE_HASH_LIST) = Advapi32.FreeEncryptionCertificateHashList
//sys	queryRecoveryAgentsOnEncryptedFile(lpFileName *uint16, pRecoveryAgents *ENCRYPTION_CERTIFICATE_HASH_LIST) (ret error) = Advapi32.QueryRecoveryAgentsOnEncryptedFile
//sys	queryUsersOnEncryptedFile(lpFileName *uint16, pUsers *ENCRYPTION_CERTIFICATE_HASH_LIST) (ret error) = Advapi32.QueryUsersOnEncryptedFile
//sys	removeUsersFromEncryptedFile(lpFileName *uint16, pHashes *ENCRYPTION_CERTIFICATE_HASH_LIST) (ret error) = Advapi32.RemoveUsersFromEncryptedFile
//sys	SetUserFileEncryptionKey(pEncryptionCertificate *ENCRYPTION_CERTIFICATE) (ret error) = Advapi32.SetUserFileEncryptionKey

// EFS file backup and restore functions

//sys	openEncryptedFileRawW(lpFileName *uint16, ulFlags uint32, pvContext *unsafe.Pointer) (ret error) = Advapi32.OpenEncryptedFileRawW
//sys	CloseEncryptedFileRaw(pvContext unsafe.Pointer) = Advapi32.CloseEncryptedFileRaw
//sys	ReadEncryptedFileRaw(pfExportCallback uintptr, pvCallbackContext unsafe.Pointer, pvContext unsafe.Pointer) (ret error) = Advapi32.ReadEncryptedFileRaw
//sys	WriteEncryptedFileRaw(pfImportCallback uintptr, pvCallbackContext unsafe.Pointer, pvContext unsafe.Pointer) (ret error) = Advapi32.WriteEncryptedFileRaw


func FileEncryptionStatus(fileName string) (status uint32, err error) {
	u16ptr, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return
	}

	err = fileEncryptionStatusW(u16ptr, &status)

	return
}

func EncryptionDisable(dirPath string, disable bool) error {
	u16ptr, err := windows.UTF16PtrFromString(dirPath)
	if err != nil {
		return err
	}

	err = encryptionDisable(u16ptr, disable)

	return err
}

func EncryptFile(fileName string) error {
	u16ptr, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return err
	}

	return encryptFileW(u16ptr)
}

// reserved should be 0
func DecryptFile(fileName string, reserved uint32) error {
	u16ptr, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return err
	}

	return decryptFileW(u16ptr, reserved)
}

func AddUsersToEncryptedFile(fileName string, encryptionCertificates []ENCRYPTION_CERTIFICATE) error {
	u16ptr, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return err
	}

	var list ENCRYPTION_CERTIFICATE_LIST
	list.NumUser = uint32(len(encryptionCertificates))
	list.Users = (*ENCRYPTION_CERTIFICATE)(unsafe.Pointer(&encryptionCertificates[0]))

	return addUsersToEncryptedFileW(u16ptr, &list)
}

func DuplicateEncryptionInfoFile(srcFileName, dstFilename string, creationDistribution, attributes uint32, securityAttributes *windows.SecurityAttributes) error {
	u16src, err := windows.UTF16PtrFromString(srcFileName)
	if err != nil {
		return err
	}
	u16dst, err := windows.UTF16PtrFromString(dstFilename)
	if err != nil {
		return err
	}

	return duplicateEncryptionInfoFile(u16src, u16dst, creationDistribution, attributes, securityAttributes)
}

func QueryRecoveryAgentsOnEncryptedFile(fileName string) (*ENCRYPTION_CERTIFICATE_HASH_LIST, error) {
	u16str, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return nil, err
	}

	list := &ENCRYPTION_CERTIFICATE_HASH_LIST{}

	err = queryRecoveryAgentsOnEncryptedFile(u16str, list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func QueryUsersOnEncryptedFile(fileName string) (*ENCRYPTION_CERTIFICATE_HASH_LIST, error) {
	u16str, err := windows.UTF16PtrFromString(fileName)
	if err != nil{
		return nil, err
	}

	list := &ENCRYPTION_CERTIFICATE_HASH_LIST{}

	err = queryUsersOnEncryptedFile(u16str, list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func RemoveUsersFromEncryptedFile(fileName string, hashes []ENCRYPTION_CERTIFICATE_HASH) error {
	u16str, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return err
	}

	var list ENCRYPTION_CERTIFICATE_HASH_LIST
	list.NumCertHash = uint32(len(hashes))
	list.Users = (*ENCRYPTION_CERTIFICATE_HASH)(unsafe.Pointer(&hashes[0]))

	return removeUsersFromEncryptedFile(u16str, &list)
}

func OpenEncryptedFileRaw(fileName string, flags uint32) (context unsafe.Pointer, err error) {
	u16str, err := windows.UTF16PtrFromString(fileName)
	if err != nil {
		return
	}

	err = openEncryptedFileRawW(u16str, flags, &context)

	return
}
