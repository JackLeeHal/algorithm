package channel

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doWork() {
	closing := make(chan struct{})
	closed := make(chan struct{})

	go func() {
		for {
			select {
			case <-closing:
				return
			default:
				fmt.Println("doing job")
			}
		}
	}()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	close(closing)
	go doClean(closed)

	select {
	case <-closed:
		fmt.Println("qinglijieshu")
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
}

func doClean(closed chan struct{}) {
	time.Sleep(time.Second)
	close(closed)
}
