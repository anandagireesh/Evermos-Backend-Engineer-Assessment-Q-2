package models

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-2/database"
)

type Product struct {
	ProductName     string `json:"productname"`
	ProductQuantity int    `json:"quantity"`
	ProductPrice    int    `json:"price"`
}

type StatusMessage struct {
	Message string `json:"message"`
}
type Response struct {
	Code   int         `json:"code"`   // 200 , 400
	Status string      `json:"status"` // "Ok" "Error"
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func AddProduct(data Product) string {

	database.DbConnection()
	var message string

	row := database.Db.QueryRow("SELECT productname,quantity FROM products where productname = '" + data.ProductName + "'")
	product := Product{}
	err := row.Scan(&product.ProductName, &product.ProductQuantity)
	if err != nil {
		if err == sql.ErrNoRows {

			insert, err := database.Db.Prepare("INSERT INTO products(productname,quantity,price) VALUES ( ?,?,? )")

			//if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			insert.Exec(data.ProductName, data.ProductQuantity, data.ProductPrice)

			message = "products successfully registered"

		} else {
			panic(err)
		}
	} else {

		log.Println(product.ProductQuantity)

		qty := product.ProductQuantity + data.ProductQuantity

		log.Println("new quantity")
		log.Println(qty)

		insForm, err := database.Db.Prepare("UPDATE products SET quantity=?, price=? WHERE productname=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(qty, data.ProductPrice, data.ProductName)

		message = "products successfully updated"
		//message = "products already registered"
	}

	return message

}

func ButProduct(data Product) string {

	database.DbConnection()
	var message string

	row := database.Db.QueryRow("SELECT productname,quantity FROM products where productname = '" + data.ProductName + "'")
	product := Product{}
	err := row.Scan(&product.ProductName, &product.ProductQuantity)
	if err != nil {
		if err == sql.ErrNoRows {

			message = "No products Available"

		} else {
			panic(err)
		}
	} else {

		if product.ProductQuantity < data.ProductQuantity {

			availableStock := strconv.Itoa(product.ProductQuantity)

			message = "Not much stocks available. Available stocks : " + availableStock

		} else {

			log.Println(product.ProductQuantity)

			qty := product.ProductQuantity - data.ProductQuantity

			log.Println("new quantity")
			log.Println(qty)

			insForm, err := database.Db.Prepare("UPDATE products SET quantity=? WHERE productname=?")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(qty, data.ProductName)

			message = "purchase successfully done"

		}

		//message = "products already registered"
	}

	return message

}

func CheckProductQuantity() string {

	//Either we can mail or message or notify the products less than 5 quantity in stock. I am just printing these products here

	database.DbConnection()
	var message string

	row, err := database.Db.Query("SELECT productname,quantity,price FROM products where quantity < 5")
	if err != nil {
		defer panic(err)
	}
	product := Product{}
	prod := []Product{}

	for row.Next() {
		err = row.Scan(&product.ProductName, &product.ProductQuantity, &product.ProductPrice)

		if err != nil {
			defer panic(err)
		}

		prod = append(prod, product)

		log.Println(prod)
	}

	return message

}
