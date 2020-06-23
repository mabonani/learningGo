package kart

import "kartProject/product"

type Item struct {
	Id string
	Product  product.Product
	Quantity int8
}
