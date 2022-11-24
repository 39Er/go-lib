package call_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"go-lib/pkg/call"

	"github.com/stretchr/testify/assert"
)

func TestTaskRunner(t *testing.T) {
	runner := call.NewTaskRunner(3)
	var wg sync.WaitGroup
	var count int32
	times := 10
	m := make(map[int]call.PlaceHolderType, times)
	var lock sync.Mutex
	for i := 0; i < times; i++ {
		j := i //参数拷贝，防止循环遍历协程参数传递问题
		wg.Add(1)
		runner.Schedule(func() {
			lock.Lock()
			m[j] = call.PlaceHolder
			lock.Unlock()
			atomic.AddInt32(&count, 1)
			wg.Done()
		})
	}
	wg.Wait()
	assert.Equal(t, times, int(count))
	assert.Equal(t, times, len(m), "map length should equal times") //校验参数传递
}
