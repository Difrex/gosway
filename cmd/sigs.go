package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func sigWait() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	go watchSignals(sigs)
}

func watchSignals(sigs chan os.Signal) {
	for {
		sig := <-sigs
		fmt.Println(sig)
		if sig == syscall.SIGINT || sig == syscall.SIGQUIT || sig == syscall.SIGTERM {
			err := cleanUpSocket()
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			os.Exit(0)
		}
	}
}
