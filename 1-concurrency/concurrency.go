package concurrency

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func Conc() {
	ch := make(chan int, 10)
	chMain := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		slice := make([]int, 10)
		for i := range slice {
			slice[i] = rand.IntN(100) + 1
			ch <- slice[i]
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for val := range ch {
			chMain <- val * val
		}

	}()
	go func() {
		wg.Wait()
		close(chMain)
	}()
	for number := range chMain {
		fmt.Println(number)
	}
}
