package windows

import (
	"syscall"

	"github.com/siketyan/locksync/internal/pkg/os/windows/win32"
)

type (
	Instance win32.HINSTANCE
)

func GetInstanceHandle() (instance Instance, err *Win32Error) {
	instance = Instance(win32.GetModuleHandleW(nil))

	if instance == 0 {
		err = CreateWin32Error(win32.GetLastError())
	}

	return
}

func (instance Instance) CreateWindow(window Window) (handle WindowHandle, err *Win32Error) {
	lpClassName, _ := syscall.UTF16PtrFromString(window.ClassName)
	lpWindowName, _ := syscall.UTF16PtrFromString(window.Name)

	handle = WindowHandle(
		win32.CreateWindowExW(
			0,
			lpClassName,
			lpWindowName,
			win32.WS_POPUP,
			0, 0, 0, 0,
			win32.NULL,
			win32.NULL,
			win32.NULL,
			win32.NULL,
		),
	)

	if handle == 0 {
		err = CreateWin32Error(win32.GetLastError())
	}

	return
}

func (instance Instance) GetWindowClass(name string) (class WindowClass, err *Win32Error) {
	lpszClass, _ := syscall.UTF16PtrFromString(name)

	var wcx win32.WNDCLASSEXW
	if !win32.GetClassInfoExW(
		win32.HINSTANCE(instance),
		lpszClass,
		&wcx,
	) {
		err = CreateWin32Error(win32.GetLastError())
	} else {
		class = WindowClassFromWin32(wcx)
	}

	return
}

func (instance Instance) RegisterWindowClass(class WindowClass) (err *Win32Error) {
	wcx := class.ToWin32(instance)

	if win32.RegisterClassExW(&wcx) == 0 {
		err = CreateWin32Error(win32.GetLastError())
	}

	return
}
