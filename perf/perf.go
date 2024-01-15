package perf

import (
	"github.com/hardcore-os/plato/common/sdk"
	"net"
)

var (
	TcpConnNum int32
)

func RunMain() {
	for i := 0; i < int(TcpConnNum); i++ {
		sdk.NewChat(net.ParseIP("127.0.0.1"), 8900, "logic", "1223", "123")
	}
}
