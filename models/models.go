package models

import (
	"AlthCart/config"
	"fmt"
)

type Products struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price       int64  `json:"price" db:"price"`
	Image       string `json:"image" db:"image"`
}

func GetProducts() ([]Products, error) {
	db := config.OpenConn()
	defer db.Close()

	data := []Products{}
	err := db.Select(&data, "SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	return data, nil
}

func GetProductById(id int64) (Products, error) {
	db := config.OpenConn()
	defer db.Close()

	data := Products{}
	err := db.Get(&data, "SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func TestGetProducts() {
	data, err := GetProducts()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
}

func TestGetProductById() {
	data, err := GetProductById(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
}
