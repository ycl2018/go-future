# Go Future

Golang Future异步模型，用于异步获取执行结果，可以使用Go启动一个goroutine，它会返回一个Future来包装结果，在需要获取结果的地方通过Wait来获取结果。

## Feature
- [x] 支持范型，根据任务类型返回对应类型的Future，无需类型转换
- [x] 支持多种类型：从单返回值到至多5个返回值
- [x] 支持多次从future中获取结果，并发安全
- [x] 支持combine多个任务
- [ ] 支持设置超时时间
- [ ] 支持Then链式处理

## Install

```shell
go get github.com/ycl2018/go-future
```

## Example

### Example1: 基本用法

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

### Example2:合并多个并发的结果

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
	v1, v2, v3, err := Combine3(f1, f2, f3)
	log.Println(v1, v2, v3, err)
}

```