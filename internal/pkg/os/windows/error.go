package windows

import (
	"fmt"

	"locksync/internal/pkg/os/windows/win32"
)

type (
	Win32Error win32.DWORD
)

func CreateWin32Error(hResult win32.DWORD) *Win32Error {
	return (*Win32Error)(&hResult)
}

func (err *Win32Error) Error() string {
	return fmt.Sprintf("Win32 Error %d", *err)
}
