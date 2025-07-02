package future

import (
	"testing"
)

func TestGoIfTrue(t *testing.T) {
	f1 := GoIf(true, func() (string, error) {
		return "foo", nil
	})
	wait, err := f1.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if wait != "foo" {
		t.Fatalf("got:%s,want:%s", wait, "foo")
	}
}

func TestGoIfFalse(t *testing.T) {
	f1 := GoIf(false, func() (string, error) {
		return "foo", nil
	})
	wait, err := f1.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if wait != "" {
		t.Fatalf("got:%s,want:%s", wait, "")
	}
}

func TestGo2IfFalse(t *testing.T) {
	f1 := Go2If(false, func() (string, string, error) {
		return "foo", "bar", nil
	})
	foo, bar, err := f1.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if foo != "" || bar != "" {
		t.Fatalf("got foo:%s bar:%s want foo:%s bar:%s", foo, bar, "", "")
	}
}

func TestGo2IfTrue(t *testing.T) {
	f1 := Go2If(true, func() (string, string, error) {
		return "foo", "bar", nil
	})
	foo, bar, err := f1.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if foo != "foo" || bar != "bar" {
		t.Fatalf("got foo:%s bar:%s want foo:%s bar:%s", foo, bar, "foo", "bar")
	}
}

func TestGo3IfFalse(t *testing.T) {
	f1 := Go3If(false, func() (string, string, int, error) {
		return "foo", "bar", 1, nil
	})
	foo, bar, num, err := f1.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if foo != "" || bar != "" || num != 0 {
		t.Fatalf("got foo:%s bar:%s num:%d want zero values", foo, bar, num)
	}
}

func TestGo4IfFalse(t *testing.T) {
	f1 := Go4If(false, func() (string, string, int, int, error) {
		return "foo", "bar", 1, 2, nil
	})
	foo, bar, num, num2, err := f1.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if foo != "" || bar != "" || num != 0 || num2 != 0 {
		t.Fatalf("got foo:%s bar:%s num:%d num2:%d want zero values", foo, bar, num, num2)
	}
}

func TestGo5IfFalse(t *testing.T) {
	f1 := Go5If(false, func() (string, string, int, int, int, error) {
		return "foo", "bar", 1, 2, 3, nil
	})
	foo, bar, num, num2, num3, err := f1.Wait()
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if foo != "" || bar != "" || num != 0 || num2 != 0 || num3 != 0 {
		t.Fatalf("got foo:%s bar:%s num:%d num2:%d num3:%d want zero values", foo, bar, num, num2, num3)
	}
}
