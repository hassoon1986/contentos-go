package main

import (
	"fmt"
	"github.com/coschain/contentos-go/prototype"
//	"github.com/coschain/contentos-go/p2p/message/msg_pack"
	"time"

	myp2p "github.com/coschain/contentos-go/p2p"
	"github.com/coschain/contentos-go/p2p/common"
	//	conn "github.com/coschain/contentos-go/p2p/depend/common"
	"github.com/coschain/contentos-go/p2p/depend/common/log"
)

var ch chan int

func init() {
	log.InitLog(log.DebugLog)
	fmt.Println("Start test the netserver...")

}
func main() {
	log.Init(log.Stdout)
	fmt.Println("Start test new p2pserver...")

	p2p := myp2p.NewServer()

	err := p2p.Start()
	if err != nil {
		fmt.Println("Start p2p error: ", err)
	}

	time.Sleep(28 * time.Second)

	for i:=0;i<1;i++ {
		// Broadcast signedTransaction
		trx := &prototype.Transaction{
			RefBlockNum:    1,
			RefBlockPrefix: 2,
		}

		sigtrx := new(prototype.SignedTransaction)
		sigtrx.Trx = trx
		p2p.Xmit(sigtrx)
	}

	if p2p.GetVersion() != common.PROTOCOL_VERSION {
		log.Error("TestNewP2PServer p2p version error", p2p.GetVersion())
	}

	if p2p.GetVersion() != common.PROTOCOL_VERSION {
		log.Error("TestNewP2PServer p2p version error")
	}
	sync, cons := p2p.GetPort()
	if sync != 20338 {
		log.Error("TestNewP2PServer sync port error")
	}

	if cons != 20339 {
		log.Error("TestNewP2PServer consensus port error")
	}

	<- ch
}
