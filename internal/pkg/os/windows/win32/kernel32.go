package win32

import (
	"syscall"
	"unsafe"
)

var (
	kernel32          = syscall.NewLazyDLL("kernel32.dll")
	pGetLastError     = kernel32.NewProc("GetLastError")
	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

func GetLastError() DWORD {
	ret, _, _ := pGetLastError.Call()

	return DWORD(ret)
}

func GetModuleHandleW(lpModuleName LPCWSTR) HMODULE {
	ret, _, _ := pGetModuleHandleW.Call(
		uintptr(unsafe.Pointer(lpModuleName)),
	)

	return HMODULE(ret)
}
