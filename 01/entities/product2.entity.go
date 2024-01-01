package entities

import (
	"fmt"
)

type Product2 struct {
	Id       string
	Name     string
	Price    float64
	Quantity int
	Status   bool
}

func NewProduct(id string, name string, price float64, quantity int, status bool) Product2 {
	product := Product2{id, name, price, quantity, status}
	return product
}

func (product Product2) ToString() string {
	return fmt.Sprintf("id: %s\nname: %s\nprice: %0.2f\nquantity: %d\nstatus: %t", product.Id, product.Name, product.Price, product.Quantity, product.Status)
}

func (product Product2) Total() float64 {
	return product.Price * float64(product.Quantity)
}

func (product Product2) ChangeValue1(name string) {
	product.Name = name
}

func (product *Product2) ChangeValue2(name string) {
	product.Name = name
}
