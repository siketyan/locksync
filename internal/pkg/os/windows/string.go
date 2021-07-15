package windows

import "C"

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/siketyan/locksync/internal/pkg/os/windows/win32"
)

func WStrToString(str win32.LPCWSTR) string {
	size := C.wcslen(str)
	xbuf := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(str)),
		Len:  int(size + 1),
		Cap:  int(size + 1),
	}

	buf := (*[]uint16)(unsafe.Pointer(xbuf))

	return syscall.UTF16ToString(*buf)
}
