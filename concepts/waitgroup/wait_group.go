package waitgroup

import (
	"fmt"
	"sync"
)

func worker(msg string) {
	fmt.Println("Worker", msg)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		worker("w1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		worker("w2")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		worker("w3")
	}()

	fmt.Println("waiting")
	wg.Wait()
	fmt.Println("exiting")
}
