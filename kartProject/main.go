package main

import (
	"kartProject/kart"
	"kartProject/product"
)

func main() {
	var myKart = kart.Kart{}

	bolacha := product.Product{
		Name:  "bolacha",
		Price: 3.99,
	}

	caneta := product.Product{
		Name:  "caneta",
		Price: 1.20,
	}

	lapis := product.Product{
		Name:  "lapis",
		Price: 1.00,
	}

	itemBolacha := kart.Item{
		Id:       nil,
		Product:  bolacha,
		Quantity: 10,
	}

	itemCaneta := kart.Item{
		Id:       nil,
		Product:  caneta,
		Quantity: 12,
	}

	itemLapis := kart.Item{
		Id:       4,
		Product:  lapis,
		Quantity: 6,
	}

	myKart.AddItem(itemBolacha)
	myKart.AddItem(itemCaneta)
	myKart.AddItem(itemLapis)

	myKart.ShowItems()
	myKart.CalculateKartPrice()
	myKart.RemoveItem(itemCaneta)

	myKart.ShowItems()
	myKart.CalculateKartPrice()

}
