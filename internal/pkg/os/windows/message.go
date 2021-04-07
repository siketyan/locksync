package windows

import (
	"github.com/siketyan/locksync/internal/pkg/os/windows/win32"
)

type (
	Message                 win32.MSG
	MessageType             uint
	WtsSessionChangeMessage Message
	WtsSessionChangeReason  uint
)

const (
	// MessageType
	MsgTypeUnknown = iota
	MsgTypeWtsSessionChange
)

const (
	// WtsSessionChangeReason
	WtsReasonUnknown = iota
	WtsReasonSessionLock
	WtsReasonSessionUnlock
)

func (message *Message) GetType() MessageType {
	switch message.Message {
	case win32.WM_WTSSESSION_CHANGE:
		return MsgTypeWtsSessionChange

	default:
		return MsgTypeUnknown
	}
}

func (message *Message) Translate() (err *Win32Error) {
	if !win32.TranslateMessage((*win32.MSG)(message)) {
		err = CreateWin32Error(win32.GetLastError())
	}

	return
}

func (message *Message) Dispatch() (err *Win32Error) {
	if !win32.DispatchMessage((*win32.MSG)(message)) {
		err = CreateWin32Error(win32.GetLastError())
	}

	return
}

func (message *WtsSessionChangeMessage) GetReason() WtsSessionChangeReason {
	switch message.WParam {
	case win32.WTS_SESSION_LOCK:
		return WtsReasonSessionLock

	case win32.WTS_SESSION_UNLOCK:
		return WtsReasonSessionUnlock

	default:
		return WtsReasonUnknown
	}
}
