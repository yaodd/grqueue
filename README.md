# grqueue
Grqueue is a simple goroutine synchronous queue.It can make the mutiple tasks what cost too many time and execute synchronously in your program run at different goroutines.
## Quickstart
```
go get github.com/yaodd/grqueue
```
Import in your package and use it.Let's look a simple example.
```
package main

import (
	"fmt"
	"github.com/yaodd/grqueue"
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
	fmt.Println("cost:", diff)
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
```
After you execute `go run` it will be display:
```
$ go run
filished task index:2
filished task index:3
filished task index:0
filished task index:4
filished task index:5
filished task index:1
filished task index:11
filished task index:9
filished task index:10
filished task index:7
filished task index:6
filished task index:8
filished task index:17
filished task index:14
filished task index:13
filished task index:12
filished task index:16
filished task index:15
filished task index:19
filished task index:18
all tasks finished
cost: 4.007254648s
```

That's all.