package utils

type Waiter interface {
	Wait()
}

type Interrupter interface {
	Interrupt()
}

type WaiterInterrupter interface {
	Waiter
	Interrupter
}
