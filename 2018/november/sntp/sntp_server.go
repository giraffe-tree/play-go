package sntp

import (
	"github.com/btfak/sntp/netapp"
	"github.com/btfak/sntp/netevent"
)

func main() {
	var handler = netapp.GetHandler()
	netevent.Reactor.ListenUdp(123, handler)
	netevent.Reactor.Run()
}