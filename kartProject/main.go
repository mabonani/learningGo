package main

import (
	"kartProject/kart"
	"kartProject/product"
)

func main() {
	var myKart = kart.Kart{}

	banana := product.Product{
		Name:  "banana",
		Price: 3.99,
	}

	/*caneta := product.Product{
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
	myKart.RemoveItem(itemCaneta)*/

	itemBanana := kart.Item{
		Product:  banana,
		Quantity: 6,
	}
	myKart.AddItem(itemBanana)
	myKart.ShowItems()
	myKart.CalculateKartPrice()

	createdKart := kart.CreateKart(myKart)

	createdKart.ShowItems()

}
