//go:build linux
// +build linux

package enforcer

// Go libraries
import (
	"fmt"
	"strconv"

	"github.com/headingy/trireme/enforcer/utils/packet"
	nfqueue "go.aporeto.io/netlink-go/nfqueue"
	"go.uber.org/zap"
)

func errorCallback(err error, data interface{}) {
	zap.L().Error("Error while processing packets on queue", zap.Error(err))
}
func networkCallback(packet *nfqueue.NFPacket, d interface{}) {
	d.(*Datapath).processNetworkPacketsFromNFQ(packet)
}

func appCallBack(packet *nfqueue.NFPacket, d interface{}) {
	d.(*Datapath).processApplicationPacketsFromNFQ(packet)
}

// startNetworkInterceptor will the process that processes  packets from the network
// Still has one more copy than needed. Can be improved.
func (d *Datapath) startNetworkInterceptor() {
	var err error
	d.netStop = make([]chan bool, d.filterQueue.GetNumNetworkQueues())
	for i := uint16(0); i < d.filterQueue.GetNumNetworkQueues(); i++ {
		d.netStop[i] = make(chan bool)
	}

	nfq := make([]nfqueue.Verdict, d.filterQueue.GetNumNetworkQueues())

	for i := uint16(0); i < d.filterQueue.GetNumNetworkQueues(); i++ {

		// Initialize all the queues
		nfq[i], err = nfqueue.CreateAndStartNfQueue(d.filterQueue.GetNetworkQueueStart()+i, d.filterQueue.GetNetworkQueueSize(), nfqueue.NfDefaultPacketSize, networkCallback, errorCallback, d)
		if err != nil {
			zap.L().Fatal("Unable to initialize netfilter queue", zap.Error(err))
		}
		go func(j uint16) {
			for range d.netStop[j] {
				return
			}
		}(i)

	}
}

// startApplicationInterceptor will create a interceptor that processes
// packets originated from a local application
func (d *Datapath) startApplicationInterceptor() {

	var err error
	d.appStop = make([]chan bool, d.filterQueue.GetNumApplicationQueues())
	for i := uint16(0); i < d.filterQueue.GetNumApplicationQueues(); i++ {
		d.appStop[i] = make(chan bool)
	}

	nfq := make([]nfqueue.Verdict, d.filterQueue.GetNumApplicationQueues())

	for i := uint16(0); i < d.filterQueue.GetNumApplicationQueues(); i++ {
		nfq[i], err = nfqueue.CreateAndStartNfQueue(d.filterQueue.GetApplicationQueueStart()+i, d.filterQueue.GetApplicationQueueSize(), nfqueue.NfDefaultPacketSize, appCallBack, errorCallback, d)

		if err != nil {
			zap.L().Fatal("Unable to initialize netfilter queue", zap.Error(err))
		}
		go func(j uint16) {
			for range d.appStop[j] {
				return
			}

		}(i)

	}
}

// processNetworkPacketsFromNFQ processes packets arriving from the network in an NF queue
func (d *Datapath) processNetworkPacketsFromNFQ(p *nfqueue.NFPacket) {

	// Parse the packet - drop if parsing fails
	netPacket, err := packet.New(packet.PacketTypeNetwork, p.Buffer, strconv.Itoa(int(p.Mark)))

	if err != nil {
		netPacket.Print(packet.PacketFailureCreate)
	} else if netPacket.IPProto == packet.IPProtocolTCP {
		err = d.processNetworkTCPPackets(netPacket)
	} else {
		err = fmt.Errorf("Invalid IP Protocol %d", netPacket.IPProto)
	}
	if err != nil {
		length := uint32(len(p.Buffer))
		buffer := p.Buffer
		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 0, uint32(p.Mark), length, uint32(p.ID), buffer)
		return
	}

	// // Accept the packet
	buffer := make([]byte, len(netPacket.Buffer)+netPacket.TCPOptionLength()+netPacket.TCPDataLength())
	copyIndex := copy(buffer, netPacket.Buffer)
	copyIndex += copy(buffer[copyIndex:], netPacket.GetTCPOptions())
	copyIndex += copy(buffer[copyIndex:], netPacket.GetTCPData())
	// buffer = append(buffer, netPacket.GetTCPOptions()...)
	// buffer = append(buffer, netPacket.GetTCPData()...)
	// length = uint32(len(buffer))
	p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 1, uint32(p.Mark), uint32(copyIndex), uint32(p.ID), buffer)

}

// processApplicationPackets processes packets arriving from an application and are destined to the network
func (d *Datapath) processApplicationPacketsFromNFQ(p *nfqueue.NFPacket) {

	// Being liberal on what we transmit - malformed TCP packets are let go
	// We are strict on what we accept on the other side, but we don't block
	// lots of things at the ingress to the network
	appPacket, err := packet.New(packet.PacketTypeApplication, p.Buffer, strconv.Itoa(int(p.Mark)))

	if err != nil {
		appPacket.Print(packet.PacketFailureCreate)
	} else if appPacket.IPProto == packet.IPProtocolTCP {
		err = d.processApplicationTCPPackets(appPacket)
	} else {
		err = fmt.Errorf("Invalid IP Protocol %d", appPacket.IPProto)
	}

	if err != nil {
		length := uint32(len(p.Buffer))
		buffer := p.Buffer
		p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 0, uint32(p.Mark), length, uint32(p.ID), buffer)
		return
	}

	// Accept the packet
	buffer := make([]byte, len(appPacket.Buffer)+appPacket.TCPOptionLength()+appPacket.TCPDataLength())
	copyIndex := copy(buffer, appPacket.Buffer)
	copyIndex += copy(buffer[copyIndex:], appPacket.GetTCPOptions())
	copyIndex += copy(buffer[copyIndex:], appPacket.GetTCPData())
	// buffer := appPacket.Buffer
	// buffer = append(buffer, appPacket.GetTCPOptions()...)
	// buffer = append(buffer, appPacket.GetTCPData()...)
	// length = uint32(len(buffer))

	p.QueueHandle.SetVerdict2(uint32(p.QueueHandle.QueueNum), 1, uint32(p.Mark), uint32(copyIndex), uint32(p.ID), buffer)

}
