package avgcounter

import (
	"fmt"
	"time"
)

func Example() {
	counter := New(time.Minute)
	counter.Add(42)

	// We use special function in the test to emulate time passing. You don't
	// need to do that, just let time fly.
	passTime(time.Minute)

	// The result is 42 * exp(-1)
	fmt.Println(counter.Value())
	// Output: 15.450936529200579
}

func Example_multi_add() {
	counter := New(time.Minute)

	for i := 0; i < 600; i++ {
		counter.Add(1)
		passTime(time.Second)
	}

	// The result about 60 (60 adds/minute)
	fmt.Println(counter.Value())
	// Output: 59.49868752358285
}

func ExampleNew() {
	counter := New(time.Minute)

	fmt.Println(counter.Value())
	// Output: 0
}

func ExampleCounter_Add() {
	counter := New(time.Minute)
	counter.Add(42)

	fmt.Println(counter.Value())
	// Output: 42
}

func ExampleCounter_Add_more() {
	counter := New(time.Minute)
	counter.Add(42)

	passTime(2 * time.Minute)
	counter.Add(42)

	// The result is 42 + 42 * exp(-2)
	fmt.Println(counter.Value())
	// Output: 47.68408189593774
}

func ExampleCounter_ValuePer() {
	counter := New(time.Minute)
	counter.Add(42)

	// 42 per minute is the 0.7 per second
	fmt.Println(counter.ValuePer(time.Second))
	// Output: 0.7
}

func passTime(d time.Duration) {
	nowTime = nowTime.Add(d)
}

func init() {
	// Stop the clock
	nowTime = time.Now()
}
