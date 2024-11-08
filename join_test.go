package cc

import (
	"log"
	"testing"
	"time"
)

func TestJoin(t *testing.T) {
	log.Println("program starting...")
	f1 := Go(func() (string, error) {
		log.Println("f1 starting...")
		consumeTime(time.Second)
		return "1", nil
	})
	f2 := Go(func() (string, error) {
		log.Println("f2 starting...")
		consumeTime(time.Second)
		return "2", nil
	})
	f3 := f2.Join(f1).Then(func(v1 string, v2 any) (any, error) {
		log.Println("f3 starting...")
		consumeTime(time.Second)
		str2 := v2.(string)
		return v1 + str2, nil
	})
	log.Printf("do something else")
	val, err := f3.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if val != "21" {
		t.Fatalf("got:%s,want:%s", val, "21")
	}
	log.Println("got result:" + val.(string))
}

func TestJoin2(t *testing.T) {
	log.Println("program starting...")
	f1 := Go2(func() (string, string, error) {
		log.Println("f1 starting...")
		consumeTime(time.Second)
		return "1", "1", nil
	})
	f2 := Go(func() (string, error) {
		log.Println("f2 starting...")
		consumeTime(time.Second)
		return "2", nil
	})
	f3 := f1.Join(f2).Then(func(v1, v2 string, v3 any) (any, error) {
		log.Println("f3 starting...")
		consumeTime(time.Second)
		str3 := v3.(string)
		return v1 + v2 + str3, nil
	})
	log.Printf("do something else")
	val, err := f3.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if val != "112" {
		t.Fatalf("got:%s,want:%s", val, "112")
	}
	log.Println("got result:" + val.(string))
}

func TestJoin6(t *testing.T) {
	log.Println("program starting...")
	f1 := Go5(func() (string, string, string, string, string, error) {
		log.Println("f1 starting...")
		consumeTime(time.Second)
		return "1", "1", "1", "1", "1", nil
	})
	f2 := Go(func() (string, error) {
		log.Println("f2 starting...")
		consumeTime(time.Second)
		return "2", nil
	})
	f3 := f1.Join(f2).Then(func(vv [6]any) (any, error) {
		log.Println("f3 starting...")
		consumeTime(time.Second)
		var ret string
		for _, a := range vv {
			ret += a.(string)
		}
		return ret, nil
	})
	log.Printf("do something else")
	val, err := f3.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if val != "111112" {
		t.Fatalf("got:%s,want:%s", val, "111112")
	}
	log.Println("got result:" + val.(string))
}
