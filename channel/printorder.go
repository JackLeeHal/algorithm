package channel

import (
	"fmt"
	"time"
)

type Token struct{}

func printNums(id int, ch, nextCh chan Token) {
	for {
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(time.Second)
		nextCh <- token
	}
}
