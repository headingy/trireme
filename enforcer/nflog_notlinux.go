//go:build darwin || !linux
// +build darwin !linux

package enforcer

import "github.com/headingy/trireme/collector"

// nfLog TODO
type nfLog struct {
}

func newNFLogger(ipv4groupSource, ipv4groupDest uint16, getPUInfo puInfoFunc, collector collector.EventCollector) nfLogger {
	return &nfLog{}
}

func (n *nfLog) start() {}
func (n *nfLog) stop()  {}
