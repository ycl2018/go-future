package future

import (
	"errors"
	"testing"
)

func TestGroup(t *testing.T) {
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
			var g Group[string]
			for i, output := range test.outputs {
				var i, output = i, output
				g.Run(func() (string, error) {
					return output, test.errs[i]
				})
			}
			ss, err := g.Wait()
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

func TestGroup_Add(t *testing.T) {
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
			var g Group[string]
			for i, output := range test.outputs {
				var i, output = i, output
				g.Add(Go(func() (string, error) {
					return output, test.errs[i]
				}))
			}
			ss, err := g.Wait()
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

func TestAnyGroup(t *testing.T) {
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
			var g AnyGroup
			for i, output := range test.outputs {
				var i, output = i, output
				g.Run(func() (any, error) {
					return output, test.errs[i]
				})
			}
			ss, err := g.Wait()
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

func TestErrGroup(t *testing.T) {
	var err1 = errors.New("ops 1")
	var err2 = errors.New("ops 2")
	var tests = []struct {
		caseName string
		errs     []error
	}{
		{
			caseName: "0 err",
			errs:     []error{nil, nil, nil},
		},
		{
			caseName: "1 err",
			errs:     []error{nil, err1, nil},
		},
		{
			caseName: "2 err",
			errs:     []error{nil, err1, err2},
		},
	}

	for _, test := range tests {
		t.Run(test.caseName, func(t *testing.T) {
			var g ErrGroup
			for _, err := range test.errs {
				err := err
				g.Run(func() error {
					return err
				})
			}
			err := g.Wait()
			for _, e := range test.errs {
				if e != nil {
					if !errors.Is(err, e) {
						t.Fatalf("want errs:%v,but got err:%v", e, err)
					}
				}
			}
		})
	}
}
