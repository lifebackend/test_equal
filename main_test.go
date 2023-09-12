package main

import "testing"

func Test_ci_Compare(t *testing.T) {
	var value0 T0
	var value1 T1

	ptr0a := (*T0)(&value0)
	ptr1a := (*T1)(ptr0a)

	ptr0b := (*T0)(&value1)
	ptr1b := (*T1)(ptr0b)

	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"int", args{a: 1, b: 1}, true},
		{"double", args{a: 0.11, b: 0.11}, true},
		{"v0", args{a: ptr0a, b: ptr1a}, true},
		{"v1", args{a: ptr0b, b: ptr1b}, true},
		{"v3", args{a: ptr0a, b: ptr0b}, false},
		{"v4", args{a: ptr1a, b: ptr1b}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ci{}
			if got := c.Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
