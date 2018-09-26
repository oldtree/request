package main

import (
	"os"
	"os/signal"
	"request"
	"syscall"
)

func main() {
	_ = request.Request{}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}
