package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Years(5)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = YearsTwo(5)
	}
}

func ExampleYears() {
	age := 2
	years := Years(age)
	fmt.Println("Dog's Age:", years)
	// Output:
	// Dog's Age: 14

}
func ExampleYearsTwo() {
	age := 2
	years := YearsTwo(age)
	fmt.Println("Dog's Age:", years)
	// Output:
	// Dog's Age: 14
}

func TestYears(t *testing.T) {
	ages := []struct {
		value  int
		expect int
	}{
		{
			value:  1,
			expect: 7,
		},
		{
			value:  2,
			expect: 14,
		},
		{
			value:  3,
			expect: 21,
		},
		{
			value:  4,
			expect: 28,
		},
	}

	for _, age := range ages {
		got := Years(age.value)
		if got != age.expect {
			t.Errorf("\nFor: %#v\nExpected: %v\nGot: %v", age.value, age.expect, got)
		}
	}
}
func TestYearsTwo(t *testing.T) {
	ages := []struct {
		value  int
		expect int
	}{
		{
			value:  1,
			expect: 7,
		},
		{
			value:  2,
			expect: 14,
		},
		{
			value:  3,
			expect: 21,
		},
		{
			value:  4,
			expect: 28,
		},
	}

	for _, age := range ages {
		got := YearsTwo(age.value)
		if got != age.expect {
			t.Errorf("\nFor: %#v\nExpected: %v\nGot: %v", age.value, age.expect, got)
		}
	}
}
