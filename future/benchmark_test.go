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

func BenchmarkFuture3Worker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var f []*Future[string]
		for j := 0; j < 3; j++ {
			f = append(f, Go(func() (string, error) {
				return "bar", nil
			}))
		}
		CollectSlice(f...)
	}
	b.StopTimer()
}

func BenchmarkWaitGroup3Worker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var swg sync.WaitGroup
		swg.Add(3)
		var strs [3]string
		var errs [3]error
		for j := 0; j < 3; j++ {
			j := j
			go func() {
				defer func() {
					if e := recover(); e != nil {
						b.Log(e)
					}
				}()
				defer swg.Done()
				strs[j] = "bar"
				errs[j] = nil
			}()
		}
		swg.Wait()
	}
	b.StopTimer()
}

func BenchmarkFuture10Worker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var f []*Future[string]
		for j := 0; j < 10; j++ {
			f = append(f, Go(func() (string, error) {
				return "bar", nil
			}))
		}
		CollectSlice(f...)
	}
	b.StopTimer()
}

func BenchmarkWaitGroup10Worker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var swg sync.WaitGroup
		swg.Add(10)
		var strs [10]string
		var errs [10]error
		for j := 0; j < 10; j++ {
			j := j
			go func() {
				defer func() {
					if e := recover(); e != nil {
						b.Log(e)
					}
				}()
				defer swg.Done()
				strs[j] = "bar"
				errs[j] = nil
			}()
		}
		swg.Wait()
	}
	b.StopTimer()
}
