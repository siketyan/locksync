package win32

import (
	"syscall"
)

type (
	BOOL  bool
	INT   int32
	LONG  int32
	UINT  uint32
	CHAR  uint8
	WCHAR uint16
	ATOM  uint16
	DWORD uint32

	LPVOID  uintptr
	LPCSTR  *uint8
	LPCWSTR *uint16

	WNDPROC uintptr
	WPARAM  uintptr
	LPARAM  uintptr
	LRESULT uintptr

	HBRUSH    syscall.Handle
	HCURSOR   syscall.Handle
	HICON     syscall.Handle
	HINSTANCE syscall.Handle
	HMENU     syscall.Handle
	HMODULE   syscall.Handle
	HWND      syscall.Handle
)

const (
	NULL = 0
)
