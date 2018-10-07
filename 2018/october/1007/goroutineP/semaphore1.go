package main

import (
	"fmt"
	"time"
)

func main() {
	s := new(semaphore)
	fmt.Println("begin 0")
	go func() {
		s.Wait(1)
		defer s.Signal()
		fmt.Println("begin 1")
		time.Sleep(1e9)

		return
	}()

	go func() {
		s.Wait(1)
		defer s.Signal()
		fmt.Println("begin 2")
		time.Sleep(1e9)
	}()
	fmt.Println("end 0")
	time.Sleep(3e9)
}

type Empty interface{}
type semaphore chan Empty

// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
