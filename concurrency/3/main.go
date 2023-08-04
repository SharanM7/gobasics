package main

import (
	"context"
	"log"
	"time"
)

type Result struct {
	S   string
	err error
}

func main() {
	startTime := time.Now()
	res, err := foo()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("res : %v | time took : %v\n", res, time.Since(startTime))
}

func foo() (string, error) {
	resultch := make(chan Result, 2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	go func() {
		s, err := boo()
		resultch <- Result{
			S:   s,
			err: err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.S, res.err
	}
}

func boo() (string, error) {
	time.Sleep(time.Millisecond * 50)
	return "a", nil
}
