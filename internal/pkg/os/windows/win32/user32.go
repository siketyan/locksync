package win32

import (
	"syscall"
	"unsafe"
)

type (
	WNDCLASSEXW struct {
		CbSize        UINT
		Style         UINT
		LpfnWndProc   WNDPROC
		CbClsExtra    INT
		CbWndExtra    INT
		HInstance     HINSTANCE
		HIcon         HICON
		HCursor       HCURSOR
		HprBackground HBRUSH
		LpszMenuName  LPCWSTR
		LpszClassName LPCWSTR
		HIconSm       HICON
	}

	POINT struct {
		X LONG
		Y LONG
	}

	MSG struct {
		HWnd     HWND
		Message  UINT
		WParam   WPARAM
		LParam   LPARAM
		Time     DWORD
		Pt       POINT
		LPrivate DWORD
	}

	LPWNDCLASSEXW *WNDCLASSEXW
	LPMSG         *MSG
)

// noinspection GoSnakeCaseUsage
const (
	COLOR_WINDOW = 5

	WS_POPUP = 0x80000000
)

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	pCreateWindowExW  = user32.NewProc("CreateWindowExW")
	pDispatchMessage  = user32.NewProc("DispatchMessage")
	pGetClassInfoExW  = user32.NewProc("GetClassInfoExW")
	pGetMessage       = user32.NewProc("GetMessage")
	pLockWorkStation  = user32.NewProc("LockWorkStation")
	pRegisterClassExW = user32.NewProc("RegisterClassExW")
	pTranslateMessage = user32.NewProc("TranslateMessage")
)

func CreateWindowExW(dwExStyle DWORD, lpClassName, lpWindowName LPCWSTR, dwStyle DWORD, x, y, nWidth, nHeight INT, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam LPVOID) HWND {
	ret, _, _ := pCreateWindowExW.Call(
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam),
	)

	return HWND(ret)
}

func DispatchMessage(lpMsg LPMSG) BOOL {
	ret, _, _ := pDispatchMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)

	return ret != 0
}

func GetClassInfoExW(hInstance HINSTANCE, lpszClass LPCWSTR, lpwcx LPWNDCLASSEXW) BOOL {
	ret, _, _ := pGetClassInfoExW.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpszClass)),
		uintptr(unsafe.Pointer(lpwcx)),
	)

	return ret != 0
}

func GetMessage(lpMsg LPMSG, hWnd HWND, wMsgFilterMin UINT, wMsgFilterMax UINT) BOOL {
	ret, _, _ := pGetMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)

	return ret != 0
}

func LockWorkStation() BOOL {
	ret, _, _ := pLockWorkStation.Call()

	return ret != 0
}

func RegisterClassExW(lpwcx LPWNDCLASSEXW) ATOM {
	ret, _, _ := pRegisterClassExW.Call(
		uintptr(unsafe.Pointer(lpwcx)),
	)

	return ATOM(ret)
}

func TranslateMessage(lpMsg LPMSG) BOOL {
	ret, _, _ := pTranslateMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
	)

	return ret != 0
}
