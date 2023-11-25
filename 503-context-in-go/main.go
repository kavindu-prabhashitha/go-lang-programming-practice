package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCool(ctx context.Context) {
	rId := ctx.Value("request-id")
	rId02 := ctx.Value("request-id2")
	fmt.Println(rId)
	fmt.Println(rId02)
}

func doSomethingCooler(ctx context.Context) {
	rId := ctx.Value("request-id")
	fmt.Println("doSomethingCooler() ==> ", rId)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("do something cooler")

		}
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	fmt.Println("Go Context Example")
	// ctx := context.Background()
	// ctx = enrichContext(ctx)
	// doSomethingCool(ctx)

	ctxNew, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctxNew = enrichContext(ctxNew)
	go doSomethingCooler(ctxNew)
	select {
	case <-ctxNew.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}

/*
We can use context to pass information between the layers of your application
use this only for things that truly need to be propagated
Shouldn't use context as a bucket for all of your information
*/

/*

Go Context Example
doSomethingCooler() ==>  12345
do something cooler
do something cooler
do something cooler
do something cooler
oh no, I've exceeded the deadline
timed out

*/
