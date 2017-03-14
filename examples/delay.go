package main

import (
	"fmt"
	"grqueue"
	"time"
)

func main() {
	t1 := time.Now()
	queue := grqueue.NewGoroutineQueue(6, 20)
	for i := 0; i < 20; i++ {
		index := i
		task := func() interface{} {
			return delay(1, index)
		}
		queue.AddTask(task)
	}
	queue.SetTaskEndCallback(tfCallback)
	queue.SetFinishCallback(endCallback)
	queue.Start()
	diff := time.Now().Sub(t1)
	fmt.Println("diff:", diff)
}

func delay(sec uint, index int) int {
	time.Sleep(time.Duration(sec) * time.Second)
	return index
}
func tfCallback(result interface{}) {
	index := result.(int)
	fmt.Printf("filished task index:%d\n", index)
}
func endCallback() {
	fmt.Println("all tasks finished")
}
