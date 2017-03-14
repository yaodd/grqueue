package main

import (
	"fmt"
	"github.com/yaodd/grqueue"
	"time"
)

func main() {
	t1 := time.Now()
	queue := grqueue.NewGoroutineQueue(5, 20)
	for i := 0; i < 20; i++ {
		index := i
		task := func() interface{} {
			delay(1, index)
			return nil
		}
		queue.AddTask(task)
	}
	queue.SetFinishCallback(endCallback)
	queue.Start()
	diff := time.Now().Sub(t1)
	fmt.Println("diff:", diff)
}

func delay(sec uint, index int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("finished task index:%d\n", index)
}
func endCallback() {
	fmt.Println("all tasks finished")
}
