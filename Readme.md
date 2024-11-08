# Go Future

Golang Future异步模型，用于异步获取执行结果，使用Go启动一个goroutine，它会返回一个Future来包装结果，在需要获取结果的地方通过Wait来获取结果。

## 核心概念和使用方法

- 使用`Go`创建一个Future任务
- 使用`Wait`等待Future返回结果
- 使用`Then`链接Future后置处理流程
- 使用`Combine`合并多个Future返回值和错误

## 核心Feature

- [x] 支持范型，根据任务类型返回对应类型的Future，无需类型转换
- [x] 支持多返回值类型任务：从单返回值到至多5个返回值
- [x] 支持重复从future中获取结果，并发安全
- [x] 支持Combine多个Future任务，等待完成并合并结果和错误
- [x] 支持Then链接其他Future任务
- [x] 支持链式Join其他Future任务
- [ ] 支持在链路节点发生错误时处理
- [x] 支持设置超时时间

## Install

```shell
go get github.com/ycl2018/go-future
```

## Example

### Example1: 创建一个Future并等待执行结果

使用`Go`创建一个Future任务，并在需要的地方获取其返回值。

```go
package main

import (
	"log"
	"time"

	gf "github.com/ycl2018/go-future"
)

func main() {
	// 启动异步任务
	future := gf.Go(func() (value string, err error) {
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

### Example2: 合并多个Future的结果

使用`Combine`合并多个Future的结果和错误，一并处理

```go
package main

import (
	"log"

	. "github.com/ycl2018/go-future"
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
	
	// 在需要的地方获取结果
	v1, v2, v3, err := Combine3(f1, f2, f3)
	log.Println(v1, v2, v3, err)
}
```

## Example3: 使用Then链式处理任务

创建Future任务后，可以使用`Then`链接后置处理任务，可以链接多个，它会返回一个新的Future任务，在需要的地方使用`Wait`获取整个链式任务的处理结果。

```go
package future

import (
	"log"
	"testing"
	"time"

	. "github.com/ycl2018/go-future"
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

## Example4: 综合场景：通过URL下载图片

这个示例说明了使用go-future方便地批量下载url二进制图片信息的方式

```go
func TestFuture(t *testing.T) {
	var urls = []string{
		"https://www.test.com/pic1",
		"https://www.test.com/pic2",
		"https://www.test.com/pic3",
		"https://www.test.com/pic3",
	}
	var futures []*Future[[]byte]
	for _, url := range urls {
		// 启动下载任务
		f := Go(func() ([]byte, error) {
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
		// 收集Futures
		futures = append(futures, f)
	}
	// 合并Futures获取结果
	ret, err := CombineSame(futures...)
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	t.Log(ret)
}
```