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
		Color: "preta",
	}

	canetaAzul := product.Product{
		Name:  "caneta",
		Price: 1.20,
		Color: "azul",
	}

	canetaVermelha:= product.Product{
		Name:  "caneta",
		Price: 1.20,
		Color: "vermelho",
	}

	lapisVermelho := product.Product{
		Name:  "lapis",
		Price: 1.00,
		Color: "vermelho",
	}

	lapisPreto := product.Product{
		Name:  "lapis",
		Price: 1.00,
		Color: "preto",
	}

	itemBolacha := kart.Item{
		"1",
		bolacha,
		10,
	}

	itemCanetaAzul := kart.Item{
		"2",
		canetaAzul,
		12,
	}

	itemCanetaVermelha := kart.Item{
		"3",
		canetaVermelha,
		12,
	}

	itemLapisVermelho := kart.Item{
		"4",
		lapisVermelho,
		6,
	}

	itemLapisPreto:= kart.Item{
		"5",
		lapisPreto,
		6,
	}

	myKart.AddItem(itemBolacha)
	myKart.AddItem(itemCanetaAzul)
	myKart.AddItem(itemCanetaVermelha)
	myKart.AddItem(itemLapisVermelho)

	myKart.ShowItems()

	myKart.RemoveItem(itemLapisPreto) 

	myKart.ShowItems()

	myKart.AddItem(itemLapisPreto)

	myKart.ShowItems()
}
