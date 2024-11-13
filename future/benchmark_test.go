package future

import (
	"sync"
	"testing"
)

func BenchmarkFuture1Worker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f := Go(func() (string, error) {
			return "bar", nil
		})
		f.Wait()
	}
	b.StopTimer()
}

func BenchmarkWaitGroup1Worker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var swg sync.WaitGroup
		swg.Add(1)
		var str string
		var err error
		go func() {
			defer func() {
				if e := recover(); e != nil {
					b.Log(e)
				}
			}()
			defer swg.Done()
			str = "bar"
			err = nil
		}()
		swg.Wait()
		func(_ string, _ error) {
		}(str, err)
	}
	b.StopTimer()
}
