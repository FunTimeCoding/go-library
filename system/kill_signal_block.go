package system

import (
	"os"
	"os/signal"
	"syscall"
)

func KillSignalBlock() {
	// kill    sends syscall.SIGTERM
	// kill -2 sends syscall.SIGINT
	// kill -9 sends syscall.SIGKILL, but cannot be caught
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
