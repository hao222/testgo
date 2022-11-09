package main

import "log"

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work) // start the goroutine for that work
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Work failed with %s in %v", err, work)
		}
	}()
	// do(work) 发生 panic()，错误会被记录且协程会退出并释放，而其他协程不受影响
	do(work)
}
