package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	var xi = []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

func ExampleCenteredAvg() {
	var xi = []int{1, 2, 3, 4, 5}

	fmt.Println(CenteredAvg(xi))
	// Output:
	// 3
}

func TestCenteredAvg(t *testing.T) {
	xs := []struct {
		values []int
		expect float64
	}{
		{
			values: []int{1, 2, 3, 4, 5},
			expect: 3,
		},
		{
			values: []int{-1, -2, -3, -4, -5},
			expect: -3,
		},
		{
			values: []int{-2, -1, 1, 2, 3},
			expect: 0.6666666666666666,
		},
		{
			values: []int{-2, -1, 0, 2, 3},
			expect: 0.3333333333333333,
		},
		{
			values: []int{1, 2},
			expect: 0,
		},
		{
			values: []int{},
			expect: 0,
		},
		{
			values: nil,
			expect: 0,
		},
	}

	for _, v := range xs {
		got := CenteredAvg(v.values)
		if got != v.expect {
			t.Errorf("\nFor: %#v\nExpected: %v\nGot: %v", v.values, v.expect, got)
		}
	}

}
