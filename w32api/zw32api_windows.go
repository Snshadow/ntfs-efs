// Code generated by 'go generate'; DO NOT EDIT.

package w32api

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modAdvapi32 = windows.NewLazySystemDLL("Advapi32.dll")

	procAddUsersToEncryptedFileW           = modAdvapi32.NewProc("AddUsersToEncryptedFileW")
	procCloseEncryptedFileRaw              = modAdvapi32.NewProc("CloseEncryptedFileRaw")
	procDecryptFileW                       = modAdvapi32.NewProc("DecryptFileW")
	procDuplicateEncryptionInfoFile        = modAdvapi32.NewProc("DuplicateEncryptionInfoFile")
	procEncryptFileW                       = modAdvapi32.NewProc("EncryptFileW")
	procEncryptionDisable                  = modAdvapi32.NewProc("EncryptionDisable")
	procFileEncryptionStatusW              = modAdvapi32.NewProc("FileEncryptionStatusW")
	procFreeEncryptionCertificateHashList  = modAdvapi32.NewProc("FreeEncryptionCertificateHashList")
	procOpenEncryptedFileRawW              = modAdvapi32.NewProc("OpenEncryptedFileRawW")
	procQueryRecoveryAgentsOnEncryptedFile = modAdvapi32.NewProc("QueryRecoveryAgentsOnEncryptedFile")
	procQueryUsersOnEncryptedFile          = modAdvapi32.NewProc("QueryUsersOnEncryptedFile")
	procReadEncryptedFileRaw               = modAdvapi32.NewProc("ReadEncryptedFileRaw")
	procRemoveUsersFromEncryptedFile       = modAdvapi32.NewProc("RemoveUsersFromEncryptedFile")
	procSetUserFileEncryptionKey           = modAdvapi32.NewProc("SetUserFileEncryptionKey")
	procWriteEncryptedFileRaw              = modAdvapi32.NewProc("WriteEncryptedFileRaw")
)

func addUsersToEncryptedFileW(lpFileName *uint16, pEncryptionCertificates *ENCRYPTION_CERTIFICATE_LIST) (ret error) {
	r0, _, _ := syscall.Syscall(procAddUsersToEncryptedFileW.Addr(), 2, uintptr(unsafe.Pointer(lpFileName)), uintptr(unsafe.Pointer(pEncryptionCertificates)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func CloseEncryptedFileRaw(pvContext unsafe.Pointer) {
	syscall.Syscall(procCloseEncryptedFileRaw.Addr(), 1, uintptr(pvContext), 0, 0)
	return
}

func decryptFileW(lpFileName *uint16, dwReserved uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procDecryptFileW.Addr(), 2, uintptr(unsafe.Pointer(lpFileName)), uintptr(dwReserved), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func duplicateEncryptionInfoFile(srcFileName *uint16, dstFileName *uint16, creationDistribution uint32, attributes uint32, lpSecurityAttributes *windows.SecurityAttributes) (ret error) {
	r0, _, _ := syscall.Syscall6(procDuplicateEncryptionInfoFile.Addr(), 5, uintptr(unsafe.Pointer(srcFileName)), uintptr(unsafe.Pointer(dstFileName)), uintptr(creationDistribution), uintptr(attributes), uintptr(unsafe.Pointer(lpSecurityAttributes)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func encryptFileW(lpFileName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procEncryptFileW.Addr(), 1, uintptr(unsafe.Pointer(lpFileName)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func encryptionDisable(dirPath *uint16, disable bool) (err error) {
	var _p0 uint32
	if disable {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procEncryptionDisable.Addr(), 2, uintptr(unsafe.Pointer(dirPath)), uintptr(_p0), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func fileEncryptionStatusW(lpFileName *uint16, lpStatus *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procFileEncryptionStatusW.Addr(), 2, uintptr(unsafe.Pointer(lpFileName)), uintptr(unsafe.Pointer(lpStatus)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func FreeEncryptionCertificateHashList(pUsers *ENCRYPTION_CERTIFICATE_HASH_LIST) {
	syscall.Syscall(procFreeEncryptionCertificateHashList.Addr(), 1, uintptr(unsafe.Pointer(pUsers)), 0, 0)
	return
}

func openEncryptedFileRawW(lpFileName *uint16, ulFlags uint32, pvContext *unsafe.Pointer) (ret error) {
	r0, _, _ := syscall.Syscall(procOpenEncryptedFileRawW.Addr(), 3, uintptr(unsafe.Pointer(lpFileName)), uintptr(ulFlags), uintptr(unsafe.Pointer(pvContext)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func queryRecoveryAgentsOnEncryptedFile(lpFileName *uint16, pRecoveryAgents *ENCRYPTION_CERTIFICATE_HASH_LIST) (ret error) {
	r0, _, _ := syscall.Syscall(procQueryRecoveryAgentsOnEncryptedFile.Addr(), 2, uintptr(unsafe.Pointer(lpFileName)), uintptr(unsafe.Pointer(pRecoveryAgents)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func queryUsersOnEncryptedFile(lpFileName *uint16, pUsers *ENCRYPTION_CERTIFICATE_HASH_LIST) (ret error) {
	r0, _, _ := syscall.Syscall(procQueryUsersOnEncryptedFile.Addr(), 2, uintptr(unsafe.Pointer(lpFileName)), uintptr(unsafe.Pointer(pUsers)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func ReadEncryptedFileRaw(pfExportCallback uintptr, pvCallbackContext unsafe.Pointer, pvContext unsafe.Pointer) (ret error) {
	r0, _, _ := syscall.Syscall(procReadEncryptedFileRaw.Addr(), 3, uintptr(pfExportCallback), uintptr(pvCallbackContext), uintptr(pvContext))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func removeUsersFromEncryptedFile(lpFileName *uint16, pHashes *ENCRYPTION_CERTIFICATE_HASH_LIST) (ret error) {
	r0, _, _ := syscall.Syscall(procRemoveUsersFromEncryptedFile.Addr(), 2, uintptr(unsafe.Pointer(lpFileName)), uintptr(unsafe.Pointer(pHashes)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func SetUserFileEncryptionKey(pEncryptionCertificate *ENCRYPTION_CERTIFICATE) (ret error) {
	r0, _, _ := syscall.Syscall(procSetUserFileEncryptionKey.Addr(), 1, uintptr(unsafe.Pointer(pEncryptionCertificate)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func WriteEncryptedFileRaw(pfImportCallback uintptr, pvCallbackContext unsafe.Pointer, pvContext unsafe.Pointer) (ret error) {
	r0, _, _ := syscall.Syscall(procWriteEncryptedFileRaw.Addr(), 3, uintptr(pfImportCallback), uintptr(pvCallbackContext), uintptr(pvContext))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}
