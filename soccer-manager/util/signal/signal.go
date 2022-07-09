package signal

import (
	"os"
	"os/signal"
	"syscall"
)

var gracefulStop = make(chan os.Signal)

func SetupSignals() {
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
}

func CleanupOnSignal(cleanup func()) {
	go func() {
		_ = <-gracefulStop
		cleanup()
		os.Exit(0)
	}()
}
