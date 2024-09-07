package ntfs_efs

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// application-defined callback functions for ReadEncryptedFileRaw and WriteEncryptedFileRaw
type PfeExportFunc func(data *byte, callbackContext unsafe.Pointer, length *uint32) uintptr
type PfeImportFunc func(data *byte, callbackContext unsafe.Pointer, length uint32) uintptr

// NewCallback creates callback function pointer to be used by ReadEncryptedFileRaw or WriteEncryptedFileRaw.
// Note that only a limited number of callbacks may be created in a single Go process.
func NewCallback[callback PfeExportFunc | PfeImportFunc](cb callback) uintptr {
	return windows.NewCallback(cb)
}

type EfsClient struct {
	
}
