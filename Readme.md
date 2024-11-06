# Go Future

Golang Future异步模型，用于异步获取执行结果，可以使用Go启动一个goroutine，它会返回一个Future来包装结果，在需要获取结果的地方通过Wait来获取结果。

## Feture
- [ ] 支持范型，根据任务类型返回对应类型的Future，无需类型转换
- [ ] 支持多种类型：从单返回值到至多5个返回值
- [ ] 支持多次从future中获取结果，并发安全

## Install

```shell
go get github.com/ycl2018/go-future
```

## Example

```go
package main

import (
	"log"
	"time"

	gf "github.com/ycl2018/go-future"
)

func main() {
	// do something in main
	...
	// do something take long time async
	future := gf.Go(func() (value string, err error) {
		// take a long time to complete
		time.Sleep(1 * time.Second) 
		return "complete", nil
	})
	
	// do something else...
	...
	// wait when you need result
	str, err := future.Wait()
	if err != nil {
		log.Printf("something went wrong:%v",err)
	} else {
		log.Printf("get future:%s",str)
	}
}
```