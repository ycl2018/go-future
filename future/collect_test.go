package future

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestCollect(t *testing.T) {
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
			v1, v2, err := Collect(
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

func TestCollect8(t *testing.T) {
	var mockErr1 = errors.New("error 1")
	var mockErr2 = errors.New("error 2")
	var mockErr3 = errors.New("error 3")
	var mockErr4 = errors.New("error 4")
	var mockErr5 = errors.New("error 5")
	var mockErr6 = errors.New("error 6")
	var mockErr7 = errors.New("error 7")
	var mockErr8 = errors.New("error 8")

	var cases = []struct {
		name                                           string
		v1, v2, v3, v4, v5, v6, v7, v8                 string
		err1, err2, err3, err4, err5, err6, err7, err8 error
		wantError                                      error
	}{
		{
			"no error",
			"1", "2", "3", "4", "5", "6", "7", "8",
			nil, nil, nil, nil, nil, nil, nil, nil,
			nil,
		},
		{
			"no error", "1", "2", "3", "4", "5", "6", "7", "8",
			mockErr1, nil, nil, nil, nil, nil, nil, nil,
			errors.Join(mockErr1),
		},
		{
			"no error", "1", "2", "3", "4", "5", "6", "7", "8",
			mockErr1, mockErr2, mockErr3, mockErr4, mockErr5, mockErr6, mockErr7,
			mockErr8,
			errors.Join(mockErr1, mockErr2, mockErr3, mockErr4, mockErr5, mockErr6, mockErr7, mockErr8),
		},
	}
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			v1, v2, v3, v4, v5, v6, v7, v8, err := Collect8(
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
			if s.wantError != nil && s.wantError.Error() != err.Error() {
				t.Fatalf("want err:%v but get:%v", s.wantError, err)
			}
		})
	}
}

func TestCollectSlice(t *testing.T) {
	var err1 = errors.New("ops 1")
	var err2 = errors.New("ops 2")
	var tests = []struct {
		caseName string
		errs     []error
		outputs  []string
	}{
		{
			caseName: "0 err",
			errs:     []error{nil, nil, nil},
			outputs:  []string{"1", "2", "3"},
		},
		{
			caseName: "1 err",
			errs:     []error{nil, err1, nil},
			outputs:  []string{"1", "2", "3"},
		},
		{
			caseName: "2 err",
			errs:     []error{nil, err1, err2},
			outputs:  []string{"1", "2", "3"},
		},
	}

	for _, test := range tests {
		t.Run(test.caseName, func(t *testing.T) {
			var ff []*Future[string]
			for i, output := range test.outputs {
				i, output := i, output
				ff = append(ff, Go(func() (string, error) {
					return output, test.errs[i]
				}))
			}
			ss, err := CollectSlice(ff...)
			for _, e := range test.errs {
				if e != nil {
					if !errors.Is(err, e) {
						t.Fatalf("want errs:%v,but got err:%v", e, err)
					}
				}
			}
			for i, want := range test.outputs {
				if ss[i] != want {
					t.Fatalf("want ss[%d]:%s but got:%s", i, want, ss[i])
				}
			}
		})
	}
}

func TestCollectAll(t *testing.T) {
	var err1 = errors.New("ops 1")
	var err2 = errors.New("ops 2")
	var tests = []struct {
		caseName string
		errs     []error
		outputs  []string
	}{
		{
			caseName: "0 err",
			errs:     []error{nil, nil, nil},
			outputs:  []string{"1", "2", "3"},
		},
		{
			caseName: "1 err",
			errs:     []error{nil, err1, nil},
			outputs:  []string{"1", "2", "3"},
		},
		{
			caseName: "2 err",
			errs:     []error{nil, err1, err2},
			outputs:  []string{"1", "2", "3"},
		},
	}

	for _, test := range tests {
		t.Run(test.caseName, func(t *testing.T) {
			var ff []futureI
			for i, output := range test.outputs {
				i, output := i, output
				ff = append(ff, Go(func() (string, error) {
					return output, test.errs[i]
				}))
			}
			ss, err := CollectAll(ff...)
			for _, e := range test.errs {
				if e != nil {
					if !errors.Is(err, e) {
						t.Fatalf("want errs:%v,but got err:%v", e, err)
					}
				}
			}
			for i, want := range test.outputs {
				if ss[i] != want {
					t.Fatalf("want ss[%d]:%s but got:%s", i, want, ss[i])
				}
			}
		})
	}
}

func TestCollectAny(t *testing.T) {
	f1 := Go(func() (value string, err error) {
		time.Sleep(20 * time.Millisecond)
		return "1", nil
	})
	f2 := Go(func() (value string, err error) {
		time.Sleep(40 * time.Millisecond)
		return "2", nil
	})
	f3 := Go(func() (value string, err error) {
		time.Sleep(60 * time.Millisecond)
		return "3", nil
	})
	ret, err := CollectAny(f1, f2, f3)
	if err != nil {
		t.Fatalf("got err:%v", err)
	}
	if ret != "1" {
		t.Fatalf("want ret:%s,but got:%s", "1", ret)
	}
}

func TestCollectAnyTimeout(t *testing.T) {
	f1 := Go(func() (value string, err error) {
		time.Sleep(20 * time.Millisecond)
		return "1", nil
	})
	f2 := Go(func() (value string, err error) {
		time.Sleep(40 * time.Millisecond)
		return "2", nil
	})
	f3 := Go(func() (value string, err error) {
		time.Sleep(60 * time.Millisecond)
		return "3", nil
	})
	ret, err := CollectAnyTimeout(5*time.Millisecond, f1, f2, f3)
	if ret != "" {
		t.Fatalf("want empty string, but got:%s", ret)
	}
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("expect timeout err, but got:%v", err)
	}
}

func TestCollectAllTimeout(t *testing.T) {
	f1 := Go(func() (value string, err error) {
		time.Sleep(20 * time.Millisecond)
		return "1", nil
	})
	f2 := Go(func() (value string, err error) {
		time.Sleep(40 * time.Millisecond)
		return "2", nil
	})
	f3 := Go(func() (value string, err error) {
		time.Sleep(60 * time.Millisecond)
		return "3", nil
	})
	_, err := CollectAllTimeout(5*time.Millisecond, f1, f2, f3)
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("expect timeout err, but got:%v", err)
	}
}

func TestCollect1x2(t *testing.T) {
	type args[T any, V any, M any] struct {
		f1 *Future[T]
		f2 *Future2[V, M]
	}
	type testCase[T any, V any, M any] struct {
		name    string
		args    args[T, V, M]
		want    T
		want1   V
		want2   M
		wantErr bool
	}
	var ret = "bar"
	tests := []testCase[string, *string, int]{
		{
			name: "case 1",
			args: args[string, *string, int]{
				f1: Go(func() (string, error) {
					return "foo", nil
				}),
				f2: Go2(func() (*string, int, error) {
					return &ret, 1, nil
				}),
			},
			want:    "foo",
			want1:   &ret,
			want2:   1,
			wantErr: false,
		},
		{
			name: "case 2",
			args: args[string, *string, int]{
				f1: Go(func() (string, error) {
					return "", errors.New("ops")
				}),
				f2: Go2(func() (*string, int, error) {
					return nil, 0, errors.New("bar")
				}),
			},
			want:    "",
			want1:   nil,
			want2:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := Collect1u2(tt.args.f1, tt.args.f2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collect1x2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect1x2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Collect1x2() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Collect1x2() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCollect2x2u3(t *testing.T) {
	f1 := Go2(func() (string, string, error) {
		return "foo", "bar", nil
	})
	f2 := Go2(func() (string, int, error) {
		return "foo2", 1, nil
	})
	f3 := Go3(func() (string, string, int, error) {
		return "foo3", "bar3", 2, nil
	})
	v1, v2, v3, v4, v5, v6, v7, err := Collect2x2u3(f1, f2, f3)
	t.Log(v1, v2, v3, v4, v5, v6, v7, err)
}
