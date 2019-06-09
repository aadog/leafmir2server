package login

import (
	"leafmir2server/login/internal"
)

var (
	Module     = new(internal.Module)
	ChanRPC    = internal.ChanRPC
	GetTcpInfo = internal.GetTcpInfo
	SetTcpInfo = internal.SetTcpInfo
)
