package future

import (
	"errors"
	"testing"
)

func TestCombine(t *testing.T) {
	var mockErr1 = errors.New("error 1")
	var mockErr2 = errors.New("error 2")

	var cases = []struct {
		name       string
		v1, v2     string
		err1, err2 error
		wantError  error
	}{
		{
			"no error", "1", "2", nil, nil, nil,
		},
		{
			"v1 error", "1", "2", mockErr1, nil, errors.Join(mockErr1, nil),
		},
		{
			"v2 error", "1", "2", nil, mockErr2, errors.Join(mockErr2, nil),
		},
		{
			"3 error", "1", "2", mockErr1, mockErr2, errors.Join(mockErr1, mockErr2),
		},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			v1, v2, err := Combine(
				Go(func() (value string, err error) {
					return s.v1, s.err1
				}),
				Go(func() (value string, err error) {
					return s.v2, s.err2
				}))
			if v1 != s.v1 {
				t.Fatalf("want v1:%s but get:%s", s.v1, v1)
			}
			if v2 != s.v2 {
				t.Fatalf("want v2:%s but get:%s", s.v2, v2)
			}
			if s.wantError != nil && s.wantError.Error() != err.Error() {
				t.Fatalf("want err:%v but get:%v", s.wantError, err)
			}
		})
	}
}

func TestCombine10(t *testing.T) {
	var mockErr1 = errors.New("error 1")
	var mockErr2 = errors.New("error 2")
	var mockErr3 = errors.New("error 3")
	var mockErr4 = errors.New("error 4")
	var mockErr5 = errors.New("error 5")
	var mockErr6 = errors.New("error 6")
	var mockErr7 = errors.New("error 7")
	var mockErr8 = errors.New("error 8")
	var mockErr9 = errors.New("error 9")
	var mockErr10 = errors.New("error 10")

	var cases = []struct {
		name                                                        string
		v1, v2, v3, v4, v5, v6, v7, v8, v9, v10                     string
		err1, err2, err3, err4, err5, err6, err7, err8, err9, err10 error
		wantError                                                   error
	}{
		{
			"no error", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			nil,
		},
		{
			"no error", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
			mockErr1, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			errors.Join(mockErr1),
		},
		{
			"no error", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
			mockErr1, mockErr2, mockErr3, mockErr4, mockErr5, mockErr6, mockErr7,
			mockErr8, mockErr9, mockErr10,
			errors.Join(mockErr1, mockErr2, mockErr3, mockErr4, mockErr5, mockErr6, mockErr7, mockErr8,
				mockErr9, mockErr10),
		},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, err := Combine10(
				Go(func() (value string, err error) {
					return s.v1, s.err1
				}),
				Go(func() (value string, err error) {
					return s.v2, s.err2
				}),
				Go(func() (value string, err error) {
					return s.v3, s.err3
				}),
				Go(func() (value string, err error) {
					return s.v4, s.err4
				}),
				Go(func() (value string, err error) {
					return s.v5, s.err5
				}),
				Go(func() (value string, err error) {
					return s.v6, s.err6
				}),
				Go(func() (value string, err error) {
					return s.v7, s.err7
				}),
				Go(func() (value string, err error) {
					return s.v8, s.err8
				}),
				Go(func() (value string, err error) {
					return s.v9, s.err9
				}),
				Go(func() (value string, err error) {
					return s.v10, s.err10
				}),
			)
			if v1 != s.v1 {
				t.Fatalf("want v1:%s but get:%s", s.v1, v1)
			}
			if v2 != s.v2 {
				t.Fatalf("want v2:%s but get:%s", s.v2, v2)
			}
			if v3 != s.v3 {
				t.Fatalf("want v3:%s but get:%s", s.v3, v3)
			}
			if v4 != s.v4 {
				t.Fatalf("want v4:%s but get:%s", s.v4, v4)
			}
			if v5 != s.v5 {
				t.Fatalf("want v5:%s but get:%s", s.v5, v5)
			}
			if v6 != s.v6 {
				t.Fatalf("want v6:%s but get:%s", s.v6, v6)
			}
			if v7 != s.v7 {
				t.Fatalf("want v7:%s but get:%s", s.v7, v7)
			}
			if v8 != s.v8 {
				t.Fatalf("want v8:%s but get:%s", s.v8, v8)
			}
			if v9 != s.v9 {
				t.Fatalf("want v9:%s but get:%s", s.v9, v9)
			}
			if v10 != s.v10 {
				t.Fatalf("want v10:%s but get:%s", s.v10, v10)
			}
			if s.wantError != nil && s.wantError.Error() != err.Error() {
				t.Fatalf("want err:%v but get:%v", s.wantError, err)
			}
		})
	}
}

func TestCombineN(t *testing.T) {
	vv, err := CombineN(
		Go(func() (value string, err error) {
			return "v1", nil
		}),
		Go(func() (value string, err error) {
			return "v2", nil
		}),
		Go(func() (value string, err error) {
			return "v3", nil
		}),
	)
	if err != nil {
		t.Fatalf("get err:%v", err)
	}
	var want = []string{"v1", "v2", "v3"}
	for i, s := range vv {
		if want[i] != s {
			t.Fatalf("num:%d want:%s but get:%s", i, want[i], s)
		}
	}
}

func TestCombineAll(t *testing.T) {
	all, err := CombineAny(
		Go(func() (value string, err error) {
			return "v1", nil
		}),
		Go2(func() (value int, value2 string, err error) {
			return 2, "2", nil
		}),
		Go(func() (value bool, err error) {
			return true, nil
		}),
	)
	if err != nil {
		t.Fatalf("get err:%v", err)
	}
	var want = []any{"v1", 2, "2", true}
	for i, s := range all {
		if want[i] != s {
			t.Fatalf("num:%d want:%s but get:%s", i, want[i], s)
		}
	}
}
