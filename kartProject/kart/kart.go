package kart

import (
	"fmt"
	"time"
)

type Kart struct {
	Id             int64
	UserId         int64
	Status         string
	CreateDateTime time.Time
	Items          []Item
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

func (kart *Kart) CalculateKartPrice() float32 {
	var kartPrice float32
	for _, item := range kart.Items {
		itemPrice := float32(item.Quantity) * item.Product.Price
		kartPrice += itemPrice
	}

	fmt.Printf("\nkart total price is %v", kartPrice)

	return kartPrice
}
