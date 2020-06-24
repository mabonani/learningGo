package kart

import (
	"database/sql"
	"kartProject/db"
	"kartProject/product"
)

func FindByUserId(userId int32) Kart {
	dbConnection := db.DbConnection()

	result, err := queryUserKart(dbConnection, userId)
	var kartId int32
	items := fillKartItems(result, err, kartId)

	defer dbConnection.Close()
	return Kart{
		kartId,
		userId,
		items,
	}
}

func queryUserKart(dbConnection *sql.DB, userId int32) (*sql.Rows, error) {
	result, err := dbConnection.Query("SELECT k.id, k.userId, k.status, ki.id, ki.quantity, p.name, p.price "+
		"FROM kart k, kartitem ki, product p "+
		"WHERE k.userId = ? "+
		"AND k.itemId = ki.id "+
		"AND ki.productId = p.id "+
		"ORDER BY p.name", userId)
	if err != nil {
		panic(err.Error())
	}
	return result, err
}

func fillKartItems(result *sql.Rows, err error, kartId int32) []Item {
	var items []Item
	for result.Next() {
		var itemId int32
		var status, productName string
		var quantity int8
		var price float32

		err = result.Scan(&kartId, &status, &itemId, &quantity, &productName, &price)
		if err != nil {
			panic(err.Error())
		}

		items = append(items, Item{
			itemId, product.Product{
				productName,
				price,
			},
			quantity,
		})
	}
	return items
}
