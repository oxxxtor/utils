package utils

import (
	"os"
	"os/signal"
)

var GWI WaiterInterrupter = newWaiterInterrupter()

type osSignalsProcessor struct {
	signals_channel chan os.Signal
}

func newWaiterInterrupter() WaiterInterrupter {
	waiter := new(osSignalsProcessor)
	waiter.signals_channel = make(chan os.Signal, 1)
	signal.Notify(waiter.signals_channel, os.Interrupt)
	return waiter
}

func (osp *osSignalsProcessor) Wait() {
	<-osp.signals_channel
}

func (osp *osSignalsProcessor) Interrupt() {
	osp.signals_channel <- os.Interrupt
}
