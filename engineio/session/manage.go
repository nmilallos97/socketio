package session

import "time"

type sessionCtxKey string

const (
	SessionTimeoutKey       sessionCtxKey = "timeout"
	SessionIntervalKey      sessionCtxKey = "interval"
	SessionExtendTimeoutKey sessionCtxKey = "timeout-extend"

	SessionCloseChannelKey    sessionCtxKey = "cancel-channel"
	SessionCloseFunctionKey   sessionCtxKey = "cancel-function"
	SessionCloseDisconnectKey sessionCtxKey = "close-disconnect"
)

type (
	TimeoutChannel    func() <-chan struct{}
	IntervalChannel   func() <-chan time.Time
	ExtendTimeoutFunc func()
)
