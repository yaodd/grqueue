package grqueue

import (
	"sync"
)

type GoroutineQueue struct {
	Number int //并发执行的任务个数
	Total  int //总任务数

	tasks             chan func() interface{}
	task_end_callback func(result interface{})
	finish_callback   func()
}

func NewGoroutineQueue(number int, total int) *GoroutineQueue {
	queue := &GoroutineQueue{
		tasks:  make(chan func() interface{}, total),
		Number: number,
		Total:  total}
	return queue
}

var wg sync.WaitGroup

//开始执行task
func (queue *GoroutineQueue) Start() {
	defer close(queue.tasks)
	//加锁，锁的数量是tasks的数量
	wg.Add(len(queue.tasks))
	for i := 0; i < queue.Number; i++ {

		//分number个routine执行work
		go queue.work()
	}

	//等待routine执行完毕
	wg.Wait()

	//所有task完毕，若finish回调函数存在则执行则回调
	if queue.finish_callback != nil {
		queue.finish_callback()
	}
}

func (queue *GoroutineQueue) work() {

	for {

		//不断取出task执行，直到chan关闭
		task, ok := <-queue.tasks
		if !ok {
			break
		}
		res := task()

		//完成一个task立即回调
		if queue.task_end_callback != nil {
			queue.task_end_callback(res)
		}

		//每执行完一个task，解锁一次
		wg.Done()
	}

}

func (queue *GoroutineQueue) AddTask(task func() interface{}) {
	queue.tasks <- task
}
func (queue *GoroutineQueue) SetFinishCallback(callback func()) {
	queue.finish_callback = callback
}
func (queue *GoroutineQueue) SetTaskEndCallback(callback func(result interface{})) {
	queue.task_end_callback = callback
}
