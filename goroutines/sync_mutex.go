package goroutines

import (
	"fmt"
	"sync"
)


var (
	mx sync.Mutex // mutex to synchronize access to task
	task []TaskDone // slice to store TaskDone
)

type TaskDone struct {
	Number int  `json:"number"`
	Status bool `json:"status"`
}



func RunJob() []TaskDone {
	var wg sync.WaitGroup // wait group to wait for all goroutines to finish
	for i := 0; i < 10; i++ {
		wg.Add(1) // increment the wait group counter
		go func(i int) {
			defer wg.Done() // decrement the wait group counter
			mx.Lock() // lock the mutex
			task = append(task, TaskDone{Number: i, Status: true}) // append to the slice
			fmt.Printf("Task %d is done\n", i)
			mx.Unlock() // unlock the mutex
		}(i)
	}
	wg.Wait() // wait for all goroutines to finish
	return task
}