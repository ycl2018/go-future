package future

import (
	"errors"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	f := Go(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "foo", nil
	})
	_, err := f.WaitTimeout(50 * time.Millisecond)
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("expect timeout err, but got:%v", err)
	}
}

func TestTimeout2(t *testing.T) {
	f := Go2(func() (string, string, error) {
		time.Sleep(100 * time.Millisecond)
		return "foo", "bar", nil
	})
	_, _, err := f.WaitTimeout(50 * time.Millisecond)
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("expect timeout err, but got:%v", err)
	}
}

func TestJoinTimeout(t *testing.T) {
	f1 := Go(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "foo", nil
	})
	f2 := Go(func() (string, error) {
		return "var", nil
	})
	f3 := f2.JoinTimeout(f1, 50*time.Millisecond).Then(func(f1Ret string, f2Ret any) (any, error) {
		return T2[string, any]{f1Ret, f2Ret}, nil
	})
	ret, err := f3.Wait()
	if ret != nil {
		t.Fatalf("want nil ret, but got:%v", ret)
	}
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("expect timeout err, but got:%v", err)
	}
}

func TestJoinTimeout2(t *testing.T) {
	f1 := Go(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "foo", nil
	})
	f2 := Go2(func() (string, string, error) {
		time.Sleep(100 * time.Millisecond)
		return "foo", "bar", nil
	})
	f3 := f2.JoinTimeout(f1, 50*time.Millisecond).Then(func(f1Ret1, f1Ret2 string, f2Ret any) (any, error) {
		return T3[string, string, any]{f1Ret1, f1Ret2, f2Ret}, nil
	})
	ret, err := f3.Wait()
	if ret != nil {
		t.Fatalf("want nil ret, but got:%v", ret)
	}
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("expect timeout err, but got:%v", err)
	}
}
