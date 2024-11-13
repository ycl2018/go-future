package future

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	log.Printf("start program...")
	f := Go(func() (string, error) {
		log.Printf("start task 1...")
		consumeTime(time.Second)
		return "1", nil
	}).Handle(func(str string, _ error) (any, error) {
		log.Printf("start task 2...")
		consumeTime(time.Second)
		return str + str, nil
	}).Handle(func(str any, _ error) (any, error) {
		log.Printf("start task 3...")
		consumeTime(time.Second)
		return str.(string) + str.(string), nil
	})
	log.Printf("do something else...")
	wait, err := f.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if s := wait.(string); s != "1111" {
		t.Fatalf("got ret:%s want:%s", s, "1111")
	}
	log.Printf(wait.(string))
}

func TestHandle2(t *testing.T) {
	log.Printf("start program...")
	f := Go2(func() (string, string, error) {
		log.Printf("start task 1...")
		consumeTime(time.Second)
		return "1", "2", nil
	}).Handle(func(str1, str2 string, err error) (any, error) {
		if err != nil {
			return nil, err
		}
		log.Printf("start task 2...")
		consumeTime(time.Second)
		return str1 + str2, nil
	}).Handle(func(str any, err error) (any, error) {
		if err != nil {
			return nil, err
		}
		log.Printf("start task 3...")
		consumeTime(time.Second)
		return str.(string) + str.(string), nil
	})
	log.Printf("do something else...")
	wait, err := f.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if s := wait.(string); s != "1212" {
		t.Fatalf("got ret:%s want:%s", s, "1212")
	}
	log.Printf(wait.(string))
}

func TestHandleErr(t *testing.T) {
	var mockErr = errors.New("mock err")
	log.Printf("start program...")
	f := Go(func() (string, error) {
		log.Printf("start task 1...")
		consumeTime(time.Second)
		return "1", mockErr
	}).Handle(func(str string, err error) (any, error) {
		if err != nil {
			return nil, err
		}
		log.Printf("start task 2...")
		consumeTime(time.Second)
		return str + str, nil
	}).Handle(func(str any, err error) (any, error) {
		if err != nil {
			return nil, err
		}
		log.Printf("start task 3...")
		consumeTime(time.Second)
		return str.(string) + str.(string), nil
	})
	log.Printf("do something else...")
	wait, err := f.Wait()
	if wait != nil {
		t.Fatalf("want nil,but got %v", wait)
	}
	if !errors.Is(err, mockErr) {
		t.Fatalf("wang err:%v,but got:%v", mockErr, err)
	}
}
