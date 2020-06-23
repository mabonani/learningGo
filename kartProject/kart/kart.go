package kart

import "fmt"

type Kart struct {
	Items []Item
}

func (kart *Kart) AddItem(item Item) {
	kart.Items = append(kart.Items, item)
	fmt.Printf("\nitem: %v added to the kart", item)

}

func (kart *Kart) RemoveItem(item Item) {
	itemIndex := -1
	for i, currentItem := range kart.Items {
		if item.Id == currentItem.Id {
			itemIndex = i
			break
		}
	}
	if itemIndex >= 0 {
		kart.Items = append(kart.Items[:itemIndex], kart.Items[itemIndex+1:]...)
		fmt.Printf("\nitem: %v removed from kart", item)
	}

}

func (kart *Kart) ShowItems() {
	fmt.Printf("\nKart Items: %v", kart.Items)
}
