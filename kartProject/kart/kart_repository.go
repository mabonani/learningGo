package kart

import (
	"database/sql"
	"kartProject/db"
	"kartProject/product"
	"time"
)

func FindByUserId(userId int64) Kart {
	dbConnection := db.DbConnection()

	result, err := queryUserKart(dbConnection, userId)
	var kartId int64
	var status string
	var createDateTime time.Time
	items := getKartItems(result, err, kartId, status, createDateTime)

	defer dbConnection.Close()
	return Kart{
		kartId,
		userId,
		status,
		createDateTime,
		items,
	}
}

func CreateKart(kart Kart) Kart {
	dbConnection := db.DbConnection()
	kartQuery, err := dbConnection.Prepare("INSERT INTO kart(userId) values (?)")
	kartResult, err := kartQuery.Exec(kart.UserId)
	if err != nil {
		panic(err.Error())
	}
	kart.Id, _ = kartResult.LastInsertId()

	for _, item := range kart.Items {
		kartItemQuery, err := dbConnection.Prepare("INSERT INTO kartitem(kartId, quantity, productName) values ( ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		kartItemQuery.Exec(kart.Id, item.Quantity, item.Product.Name)
	}

	defer dbConnection.Close()

	return FindByUserId(kart.UserId)
}

func queryUserKart(dbConnection *sql.DB, userId int64) (*sql.Rows, error) {
	result, err := dbConnection.Query("SELECT k.id, k.status, k.createDateTime, ki.id, ki.quantity, p.name, p.price "+
		"FROM kart k, kartitem ki, product p "+
		"WHERE k.userId = ? "+
		"AND k.id in ( SELECT max(lk.id) from kart lk where userId= ? ) "+
		"AND k.id = ki.kartId "+
		"AND ki.productName = p.name "+
		"ORDER BY p.name", userId, userId)
	if err != nil {
		panic(err.Error())
	}
	return result, err
}

func getKartItems(result *sql.Rows, err error, kartId int64, status string, createDateTime time.Time) []Item {
	var items []Item
	for result.Next() {
		var itemId int32
		var productName string
		var quantity int8
		var price float32

		err = result.Scan(&kartId, &status, &createDateTime, &itemId, &quantity, &productName, &price)
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
