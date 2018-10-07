// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to returned channel
func generate2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()
		for {
			prime := <-ch
			ch = filter2(ch, prime)
			out <- prime
		}
	}()
	return out
}

/**
这里还没有想通, 到底有多少个协程在工作?
*/
func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
