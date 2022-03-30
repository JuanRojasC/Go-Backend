package main

import (
	"context"
	. "fmt"
	"log"
	"os"
	"time"
)

// PANIC
/*
	Break program in time of execution, error index of, invoke a nil pointer, close channel are very commons situtions
*/

func goPanic() {
	_, err := os.Open("nonFile.txt")
	if err != nil {
		panic(err)
	}
}

// DEFER
/*
	Defer asure that certains functions will be executed before finish the program. Its main job in clean resources while program is executing
*/

func goDefer() {
	defer func() {
		Println("Esta funcion se ejecuta al final del programa")
	}()
	panic("Forzando panic")
}

// RECOVER
/*
	Recover allow us intercept a panic and avoid the break
*/

func goRecover() {
	divideByZero()
	Println("Panic Intercepted")
}

func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic:", err)
		}
	}()
	Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}

// PACKAGE CONTEXT
/*
	Work to define the context that can be pass trough code. (ctx context.Context) ctx like name and first arg always
	Allow us cancel the execution of function, that context can be pass to childs functions
	Interface
	type Context interface {
		Deadline() (deadline time.Time, ok bool)
		Done() <- chan struc{}
		Err() error
		Value(key interface{}) interface{}
	}

	ctx.Background() -> allow us create a empty context
	WithValue(ctx, "myKey") -> return a new context
	WithDeadline(ctx context.Context, d time.Time) -> return a new context and work to cancel manually the context
	WithTimeout(ctx context.Context, d time.Time) -> same that deadline but jut need the time that must to wait
*/

func goContext() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "saludo", "hola DH")
	contextWrapper(ctx)

	ctxDeadline := context.Background()
	deadline := time.Now().Add(time.Second * 5)

	ctxDeadline, cancelFunction := context.WithDeadline(ctxDeadline, deadline)
	// ctxDeadline, cancelFunction := context.WithTimeout(ctxDeadline, time.Second * 5)

	<-ctxDeadline.Done()
	cancelFunction()
	Println(ctxDeadline.Err().Error())
}

func contextWrapper(ctx context.Context) {
	Println(ctx.Value("saludo"))
}

func main() {

	// goPanic()
	// goDefer()
	// goRecover()
	goContext()
}
