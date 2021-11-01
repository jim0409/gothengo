// Connection Pooling
package main

import "fmt"

func getData() {
    fmt.Println("getDat")
}

func main() {
    // semaphore
    sema := Sema(10)
    for i := 0; i < 100; i++ {
        sema.run(getData)
    }
}

type Semaphore struct {
    c chan struct{}
}

func (s Semaphore) run(f func()) {
    s.c <- struct{}{}
    go func() {
        f()
        <-s.c
    }()
}

func Sema(len int) *Semaphore {
    return &Semaphore{c: make(chan struct{}, len)}
}

