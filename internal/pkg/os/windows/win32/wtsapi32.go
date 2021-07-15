package win32

import "syscall"

// noinspection GoSnakeCaseUsage
const (
	NOTIFY_FOR_THIS_SESSION = 0

	WM_WTSSESSION_CHANGE = 0x02B1

	WTS_SESSION_LOCK   = 0x7
	WTS_SESSION_UNLOCK = 0x8
)

var (
	wtsapi32                        = syscall.NewLazyDLL("wtsapi32.dll")
	pWTSRegisterSessionNotification = wtsapi32.NewProc("WTSRegisterSessionNotification")
)

func WTSRegisterSessionNotification(hWnd HWND, dwFlags DWORD) BOOL {
	ret, _, _ := pWTSRegisterSessionNotification.Call(
		uintptr(hWnd),
		uintptr(dwFlags),
	)

	return ret != 0
}
