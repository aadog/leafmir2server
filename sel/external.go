package sel

import (
	"leafmir2server/sel/internal"
)

var (
	Module     = new(internal.Module)
	ChanRPC    = internal.ChanRPC
	GetTcpInfo = internal.GetTcpInfo
	SetTcpInfo = internal.SetTcpInfo
)
