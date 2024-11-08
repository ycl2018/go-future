package future

import (
	"errors"
	"testing"
	"time"
)

var debugMode = false

func consumeTime(duration time.Duration) {
	if debugMode {
		time.Sleep(duration)
	} else {
		return
	}
}

func TestGo(t *testing.T) {
	var cases = []struct {
		name string
		val  int
		err  error
	}{
		{"no err", 1, nil},
		{"has err", 0, errors.New("error")},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			go1 := Go(func() (t int, err error) {
				consumeTime(100 * time.Millisecond)
				return s.val, s.err
			})
			for i := 0; i < 10; i++ {
				val, err := go1.Wait()
				if val != s.val {
					t.Fatalf("want val:%d but get:%d", s.val, val)
				}
				if s.err != nil && !errors.Is(err, s.err) {
					t.Fatalf("want err:%v but get:%v", s.err, err)
				}
			}
		})
	}
}

func TestGo2(t *testing.T) {
	var cases = []struct {
		name  string
		retV1 int
		retV2 string
		err   error
	}{
		{"no err", 1, "1", nil},
		{"has err", 0, "2", errors.New("error")},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			go2 := Go2(func() (t int, t2 string, err error) {
				consumeTime(100 * time.Millisecond)
				return s.retV1, s.retV2, s.err
			})
			for i := 0; i < 10; i++ {
				val1, val2, err := go2.Wait()
				if val1 != s.retV1 {
					t.Fatalf("want val:%d but get:%d", s.retV1, val1)
				}
				if val2 != s.retV2 {
					t.Fatalf("want val:%s but get:%s", s.retV2, val2)
				}
				if s.err != nil && !errors.Is(err, s.err) {
					t.Fatalf("want err:%v but get:%v", s.err, err)
				}
			}
		})
	}
}

func ptrOf[T any](t T) *T {
	return &t
}

func TestGo3(t *testing.T) {
	var cases = []struct {
		name  string
		retV1 int
		retV2 string
		retV3 string
		err   error
	}{
		{"no err", 1, "1", "1", nil},
		{"has err", 0, "2", "", errors.New("error")},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			go3 := Go3(func() (int, string, string, error) {
				consumeTime(100 * time.Millisecond)
				return s.retV1, s.retV2, s.retV3, s.err
			})
			for i := 0; i < 10; i++ {
				val1, val2, val3, err := go3.Wait()
				if val1 != s.retV1 {
					t.Fatalf("want val:%d but get:%d", s.retV1, val1)
				}
				if val2 != s.retV2 {
					t.Fatalf("want val:%s but get:%s", s.retV2, val2)
				}
				if val3 != s.retV3 {
					t.Fatalf("want val:%s but get:%s", s.retV3, val3)
				}
				if s.err != nil && !errors.Is(err, s.err) {
					t.Fatalf("want err:%v but get:%v", s.err, err)
				}
			}
		})
	}
}

func TestGo4(t *testing.T) {
	var cases = []struct {
		name  string
		retV1 int
		retV2 string
		retV3 string
		retV4 *string
		err   error
	}{
		{"no err", 1, "1", "1", ptrOf("str"), nil},
		{"has err", 0, "2", "", nil, errors.New("error")},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			go4 := Go4(func() (int, string, string, *string, error) {
				consumeTime(100 * time.Millisecond)
				return s.retV1, s.retV2, s.retV3, s.retV4, s.err
			})
			for i := 0; i < 10; i++ {
				val1, val2, val3, val4, err := go4.Wait()
				if val1 != s.retV1 {
					t.Fatalf("want val:%d but get:%d", s.retV1, val1)
				}
				if val2 != s.retV2 {
					t.Fatalf("want val:%s but get:%s", s.retV2, val2)
				}
				if val3 != s.retV3 {
					t.Fatalf("want val:%s but get:%s", s.retV3, val3)
				}
				if val4 != s.retV4 {
					t.Fatalf("want val:%p but get:%p", s.retV4, val4)
				}
				if s.err != nil && !errors.Is(err, s.err) {
					t.Fatalf("want err:%v but get:%v", s.err, err)
				}
			}
		})
	}
}

func TestGo5(t *testing.T) {
	var cases = []struct {
		name  string
		retV1 int
		retV2 string
		retV3 string
		retV4 *string
		retV5 *string
		err   error
	}{
		{"no err", 1, "1", "1", ptrOf("str"), ptrOf("str2"), nil},
		{"has err", 0, "2", "", nil, nil, errors.New("error")},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			go5 := Go5(func() (int, string, string, *string, *string, error) {
				consumeTime(100 * time.Millisecond)
				return s.retV1, s.retV2, s.retV3, s.retV4, s.retV5, s.err
			})
			for i := 0; i < 10; i++ {
				val1, val2, val3, val4, val5, err := go5.Wait()
				if val1 != s.retV1 {
					t.Fatalf("want val:%d but get:%d", s.retV1, val1)
				}
				if val2 != s.retV2 {
					t.Fatalf("want val:%s but get:%s", s.retV2, val2)
				}
				if val3 != s.retV3 {
					t.Fatalf("want val:%s but get:%s", s.retV3, val3)
				}
				if val4 != s.retV4 {
					t.Fatalf("want val:%p but get:%p", s.retV4, val4)
				}
				if val5 != s.retV5 {
					t.Fatalf("want val:%p but get:%p", s.retV5, val5)
				}
				if s.err != nil && !errors.Is(err, s.err) {
					t.Fatalf("want err:%v but get:%v", s.err, err)
				}
			}
		})
	}
}
