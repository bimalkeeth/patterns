package main

import (
	"fmt"
	"sync"
	"time"
)

type ObjectNext struct {
	Action *sync.Cond
}

func attachListner(cd *sync.Cond, fn func()) {
	cd.L.Lock()
	defer cd.L.Unlock()
	cd.Wait()
	fn()
	go attachListner(cd, fn)
}

func main() {
	obj := ObjectNext{Action: sync.NewCond(&sync.Mutex{})}

	go attachListner(obj.Action, func() {
		fmt.Println("Now i feel like a javascript thing:First")
	})
	go attachListner(obj.Action, func() {
		fmt.Println("Now i feel like a javascript thing:Two")
	})
	go attachListner(obj.Action, func() {
		fmt.Println("Now i feel like a javascript thing:Three")
	})

	for range time.Tick(time.Second * 2) {
		obj.Action.Broadcast()
	}

}
