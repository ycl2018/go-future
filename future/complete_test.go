package future

import (
	"errors"
	"testing"
)

func TestWhenComplete(t *testing.T) {
	ret, err := Go(func() (string, error) {
		return "foo", nil
	}).WhenComplete(func(ret string, err error) (string, error) {
		if ret == "foo" {
			return "bar", nil
		}
		return "", errors.New("ops")
	}).Wait()
	if ret != "bar" {
		t.Fatalf("want %s, but got:%s", "bar", ret)
	}
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
}

func TestWhenComplete2(t *testing.T) {
	ret, ret2, err := Go2(func() (string, string, error) {
		return "foo", "foo", nil
	}).WhenComplete(func(ret string, ret2 string, err error) (string, string, error) {
		if ret == "foo" {
			return "bar", "bar", nil
		}
		return "", "", errors.New("ops")
	}).Wait()
	if ret != "bar" {
		t.Fatalf("want %s, but got:%s", "bar", ret)
	}
	if ret2 != "bar" {
		t.Fatalf("want %s, but got:%s", "bar", ret2)
	}
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
}
