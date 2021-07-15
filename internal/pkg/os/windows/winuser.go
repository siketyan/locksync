package windows

import (
	"github.com/siketyan/locksync/internal/pkg/os/windows/win32"
)

func LockWorkStation() (err *Win32Error) {
	if !win32.LockWorkStation() {
		err = CreateWin32Error(win32.GetLastError())
	}

	return
}
