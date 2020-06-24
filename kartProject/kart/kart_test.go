package kart

import (
	"fmt"
	"sync"
	"testing"
)

func TestKart(t *testing.T) {
	var wg sync.WaitGroup
	sum := 0
	wg.Add(1)
	go func() {
		for i := 0; i < 1000000; i++ {
			sum = sum + 1
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 1000000; i++ {
			sum = sum - 1
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println()
	fmt.Println("===============")

	fmt.Printf("Resultado: %d\n", sum)
	t.Fail()
}
