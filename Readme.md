# Go Future

![Build Status](https://github.com/ycl2018/go-future/actions/workflows/test.yml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/ycl2018/go-future)](https://goreportcard.com/report/github.com/ycl2018/go-future)

Golang Future异步模型，用于异步获取执行结果，使用Go启动一个goroutine，它会返回一个Future来包装结果，在需要获取结果的地方通过Wait来获取结果。

## 核心概念和使用方法

- 使用`Go`创建一个Future任务
- 使用`Wait`等待Future返回结果
- 使用`Then`链接Future后置处理流程
- 使用`Collect`收集多个Future返回值和错误

## 核心Feature

- [x] 支持范型，根据任务类型返回对应类型的Future，无需类型转换
- [x] 支持多返回值类型任务：从单返回值到至多5个返回值+error
- [x] 支持重复从future中`Wait`获取结果，并发高效、安全
- [x] 支持`Collect`多个Future任务，等待完成并收集结果和错误，8个值内提供范型支持
- [x] 支持`Then`在Future任务完成且无错误返回时执行后置任务
- [x] 支持`Handle`在Future任务完成时检查错误并执行后置处理
- [x] 支持`Check`在链路节点完成时进行错误处理/结果检查
- [x] 支持链式`Join`其他Future任务
- [x] 支持设置超时时间
- [x] 支持使用 `Group` `AnyGroup` `ErrGroup` 等类WaitGroup用法

## BenchMark

```text
func BenchmarkFuture(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f := Go(func() (string, error) {
			return "bar", nil
		})
		f.Wait()
	}
	b.StopTimer()
}
==================================================
goos: darwin
goarch: arm64
pkg: github.com/ycl2018/go-future/future
cpu: Apple M3 Pro
BenchmarkFuture-12       2929560               396.0 ns/op           152 B/op          4 allocs/op
```

## Install

```shell
go get github.com/ycl2018/go-future
```

## Examples

### Example1: 创建一个Future并等待执行结果

使用`Go`创建一个Future任务，并在需要的地方获取其返回值。

```go
package main

import (
	"log"
	"time"

	. "github.com/ycl2018/go-future"
)

func main() {
	// 启动异步任务
	future := Go(func() (value string, err error) {
		// take a long time to complete
		time.Sleep(1 * time.Second)
		return "complete", nil
	})

	// 做其他事情...

	// 在需要的地方获取结果
	str, err := future.Wait()
	if err != nil {
		log.Printf("something went wrong:%v", err)
	} else {
		log.Printf("get future:%s", str)
	}
}

```

### Example2: 收集多个Future的结果

使用`Collect`收集多个Future的结果和错误，一并处理

```go
package main

import (
	"log"

	. "github.com/ycl2018/go-future/future"
)

func main() {
	// 启动三个异步任务
	var f1 = Go(func() (value string, err error) {
		return "v1", nil
	})
	var f2 = Go(func() (value int, err error) {
		return 2, nil
	})
	var f3 = Go(func() (value int, err error) {
		return 3, nil
	})

	// 做其他事情...
	
	// 在需要的地方获取结果，范型实现，返回值无需转换类型可直接使用
	v1, v2, v3, err := Collect3(f1, f2, f3)
	log.Println(v1, v2, v3, err)
}
```

### Example3: 使用Then链式处理任务

创建Future任务后，可以使用`Then`链接后置处理任务，可以链接多个，它会返回一个新的Future任务，在需要的地方使用`Wait`获取整个链式任务的处理结果。

```go
package future

import (
	"log"
	"testing"
	"time"

	. "github.com/ycl2018/go-future/future"
)

func TestThen(t *testing.T) {
	
	// 链式组装一个Future任务
	log.Printf("start program...")
	f := Go(func() (string, error) {
		log.Printf("start task 1...")
		time.Sleep(time.Second)
		return "1", nil
	}).Then(func(str string) (any, error) {
		log.Printf("start task 2...")
		time.Sleep(time.Second)
		return str + str, nil
	}).Then(func(str any) (any, error) {
		log.Printf("start task 3...")
		time.Sleep(time.Second)
		return str.(string) + str.(string), nil
	})
	// 做其他事情
	log.Printf("do something else...")
	
	// 等待所有链式任务执行完毕
	wait, err := f.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if s := wait.(string); s != "1111" {
		t.Fatalf("got ret:%s want:%s", s, "1111")
	}
	log.Printf(wait.(string))
}
```

### Example4: 综合场景：通过URL下载图片

这个示例说明了使用go-future方便地批量下载url二进制图片信息的方式

```go
func TestFuture(t *testing.T) {
	var urls = []string{
		"https://www.test.com/pic1",
		"https://www.test.com/pic2",
		"https://www.test.com/pic3",
		"https://www.test.com/pic4",
	}
	// 声明`[]byte`类型Group
	var g Group[[]byte]
	for _, url := range urls {
		// 启动下载任务
		g.Run(func() ([]byte, error) {
			resp, err := http.DefaultClient.Get(url)
			if err != nil {
				return nil, err
			}
			bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return bytes, nil
		})
	}
	// 收集Futures获取结果
	ret, err := g.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	t.Log(ret)
}
```