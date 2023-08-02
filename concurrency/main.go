package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	msgch := make(chan string)
	signalch := make(chan struct{}, 10)

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go foo(&wg1, msgch, signalch, i)
	}

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go boo(&wg2, msgch, signalch)
	}

	fmt.Println("waiting for all go routines")
	wg1.Wait()
	close(msgch)
	fmt.Println("wg1 wait done")
	wg2.Wait()
	fmt.Println("wg2 wait done")

}

func foo(wg *sync.WaitGroup, msgch chan string, signal chan struct{}, poolno int) {
	defer wg.Done()
mloop:
	for i := 0; i < 2; i++ {
		select {
		case <-signal:
			fmt.Println("close signal", poolno)
			break mloop
		default:
			msgch <- "from foo poolno :" + strconv.Itoa(poolno) + " loop no :" + strconv.Itoa(i)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("exiting foo :", poolno)
}

func boo(wg *sync.WaitGroup, msgch chan string, signal chan struct{}) {
	defer wg.Done()
	i := 0
	for e := range msgch {
		fmt.Println("msg >>>", e)
		fmt.Println("about to send close signal", i)
		signal <- struct{}{}
	}
}
