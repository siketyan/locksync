package win32

type (
	BOOL  bool
	INT   int32
	LONG  int32
	UINT  uint32
	CHAR  uint8
	WCHAR uint16
	ATOM  uint16
	DWORD uint32

	PVOID   uintptr
	LPVOID  uintptr
	LPCSTR  *uint8
	LPCWSTR *uint16

	WNDPROC uintptr
	WPARAM  uintptr
	LPARAM  uintptr
	LRESULT uintptr

	HANDLE    PVOID
	HBRUSH    HANDLE
	HCURSOR   HANDLE
	HICON     HANDLE
	HINSTANCE HANDLE
	HMENU     HANDLE
	HMODULE   HANDLE
	HWND      HANDLE
)

const (
	NULL = 0
)
