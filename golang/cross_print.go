package golang

import "sync"

func crossPrint() {
	number, letter := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i < 27; i++ {
			<-letter
			println(i)
			if i == 26 {
				wg.Done()
				return
			}
			number <- true
		}
		wg.Done()
		return
	}()
	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			<-number
			println(string(i))
			letter <- true
		}
		wg.Done()
		return
	}()
	number <- true
	wg.Wait()
}
