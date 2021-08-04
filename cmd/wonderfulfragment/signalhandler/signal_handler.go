package signalhandler

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		time.Sleep(5 * time.Second)
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

func SetupShutdownContext() context.Context {
	stopCh := SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-stopCh
		cancel()
	}()
	return ctx
}
