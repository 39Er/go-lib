package call

// A TaskRunner is used to control the concurrency of goroutines.
type TaskRunner struct {
	limitChan chan PlaceHolderType
}

// NewTaskRunner returns a TaskRunner.
// 可配合WaitGroup使用
func NewTaskRunner(concurrency int) *TaskRunner {
	return &TaskRunner{
		limitChan: make(chan PlaceHolderType, concurrency),
	}
}

// Schedule schedules a task to run under concurrency control.
// 使用时注意循环遍历时的协程参数传递问题
func (r *TaskRunner) Schedule(task func()) {
	r.limitChan <- PlaceHolder
	go func() {
		defer Recover(func() {
			<-r.limitChan
		})
		task()
	}()
}
