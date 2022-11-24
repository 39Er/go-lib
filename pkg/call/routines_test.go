package call_test

import (
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"testing"
	"time"

	"go-lib/pkg/call"
	"go-lib/pkg/runtimex/debugx"

	"golang.org/x/sync/errgroup"
)

func TestRunSafeError(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			debug.PrintStack()
			fmt.Println(debugx.LocatePanic())
		}
	}()
	var wg errgroup.Group
	var a, b int
	wg.Go(call.RunSafeWrapper(func() error {
		time.Sleep(time.Second)
		a = 1
		var tmp []int
		_ = tmp[11]
		return errors.New("1")
	}))

	wg.Go(call.RunSafeWrapper(func() error {
		time.Sleep(2 * time.Second)
		b = 2
		return nil
	}))

	err := wg.Wait()
	t.Log(err, a, b)
	time.Sleep(5 * time.Second)
}

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	semi := make(chan struct{}, 3) //控制50并发数
	for i := 0; i < 10; i++ {
		fmt.Println("--------->", i)
		wg.Add(1)
		semi <- struct{}{}
		go func(j int) {
			defer func() {
				<-semi
				wg.Done()
			}()
			fmt.Println("++++++", j)
			//time.Sleep(time.Millisecond*10 + time.Duration(rand.Intn(10))*5)
		}(i)
	}

	wg.Wait()
}
