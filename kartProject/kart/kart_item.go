package kart

import "kartProject/product"

type Item struct {
	Id       int32
	Product  product.Product
	Quantity int8
}
