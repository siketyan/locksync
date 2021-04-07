package windows

import (
	"syscall"
	"unsafe"

	"locksync/internal/pkg/os/windows/win32"
)

type (
	WindowProc *func(hWnd win32.HWND, msg win32.UINT, wParam win32.WPARAM, lParam win32.LPARAM) win32.LRESULT

	WindowClass struct {
		Name       string
		WindowProc WindowProc
	}
)

func CreateWindowClass(name string, onMessage func(message Message)) WindowClass {
	proc := func(hWnd win32.HWND, msg win32.UINT, wParam win32.WPARAM, lParam win32.LPARAM) win32.LRESULT {
		onMessage(
			Message{
				HWnd:    hWnd,
				Message: msg,
				WParam:  wParam,
				LParam:  lParam,
			},
		)

		return 0
	}

	return WindowClass{
		Name:       name,
		WindowProc: &proc,
	}
}

func WindowClassFromWin32(wcx win32.WNDCLASSEXW) (class WindowClass) {
	class.Name = WStrToString(wcx.LpszClassName)
	class.WindowProc = WindowProc(unsafe.Pointer(wcx.LpfnWndProc))

	return
}

func (class WindowClass) ToWin32(instance Instance) (wcx win32.WNDCLASSEXW) {
	wcx.CbSize = win32.UINT(unsafe.Sizeof(wcx))
	wcx.HInstance = win32.HINSTANCE(instance)
	wcx.LpszClassName, _ = syscall.UTF16PtrFromString(class.Name)
	wcx.LpfnWndProc = win32.WNDPROC(syscall.NewCallback(class.WindowProc))
	wcx.HprBackground = win32.COLOR_WINDOW + 1

	return
}
