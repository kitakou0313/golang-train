package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func getInstanceBySyncOnce() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instace now!")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created")
	}

	return singleInstance
}
