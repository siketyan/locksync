package windows

import (
	"locksync/internal/pkg/os/windows/win32"
)

type (
	WindowHandle win32.HWND
)

func (handle WindowHandle) GetMessage() (message Message, err *Win32Error) {
	var msg win32.MSG

	if !win32.GetMessage(&msg, win32.HWND(handle), 0, 0) {
		err = CreateWin32Error(win32.GetLastError())
	} else {
		message = Message(msg)
	}

	return
}

func (handle WindowHandle) RegisterSessionNotification() (err *Win32Error) {
	if !win32.WTSRegisterSessionNotification(
		win32.HWND(handle),
		win32.NOTIFY_FOR_THIS_SESSION,
	) {
		err = CreateWin32Error(win32.GetLastError())
	}

	return
}
