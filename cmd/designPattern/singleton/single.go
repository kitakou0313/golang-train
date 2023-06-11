package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getSingleInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating single instace now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance is already created")
		}

	} else {
		fmt.Println("Single instance is already created")
	}

	return singleInstance
}
