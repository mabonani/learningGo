package kart

import (
	"fmt"
	"sync"
	"testing"
)

func TestKart(t *testing.T) {
	kart := Kart{}

	item1 := Item{
		Id: 1,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Executando remocao")
			kart.RemoveItem(item1)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Executando adicao")
			kart.AddItem(item1)
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println()
	fmt.Println("===============")
	if len(kart.Items) != 0 {
		fmt.Println(len(kart.Items))
		t.Fail()
	}
}
