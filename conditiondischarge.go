package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	Action *sync.Cond
}

func main() {
	obj := Object{Action: sync.NewCond(&sync.Mutex{})}

	var attachListner func(cd *sync.Cond, fn func())

	attachListner = func(cd *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			cd.L.Lock()
			defer cd.L.Unlock()
			cd.Wait()
			fn()

		}()
		wg.Wait()
		go attachListner(cd, fn)
	}
	attachListner(obj.Action, func() {
		fmt.Println("Now i feel like a javascript thing:First")
	})
	attachListner(obj.Action, func() {
		fmt.Println("Now i feel like a javascript thing:Two")
	})
	attachListner(obj.Action, func() {
		fmt.Println("Now i feel like a javascript thing:Three")
	})

	for range time.Tick(time.Second * 2) {
		obj.Action.Broadcast()
	}

}
