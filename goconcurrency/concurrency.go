package goconcurrency

import (
	"fmt"
	"sync"
)

func GoRoutine(num int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() // once the wa

	waitGroup.Add(1) // it keeps the track of all the go routiunes that have been created
	fmt.Println(num)

}

func GoRoutine2(num int) {
	//there are different isnstructions in here that we want to execute
	// independently
	fmt.Println(num)

}
