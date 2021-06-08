package channel

import (
	"testing"
)

//func TestPrint(t *testing.T) {
//	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
//	// 创建4个worker
//	for i := 0; i < 4; i++ { go printNums(i, chs[i], chs[(i+1)%4]) }
//	//首先把令牌交给第一个worker
//	chs[0] <- struct{}{}
//	select{}
//}

func TestSig(t *testing.T) {
	doWork()
}
