// Ларионова Арина 363
package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	products map[string]*Product
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]*Product),
	}
}

func (inv *Inventory) AddProduct(product Product) {
	if existingProduct, found := inv.products[product.ID]; found {
		existingProduct.Quantity += product.Quantity
	}
	inv.products[product.ID] = &product
}

func (inv *Inventory) WriteOff(productID string, quantity int) error {
	if product, exists := inv.products[productID]; exists {
		if product.Quantity < quantity {
			return errors.New("недостаточно количества для списания")
		}
		product.Quantity -= quantity
		return nil
	}
	return errors.New("Такого продукта")
}

func (inv *Inventory) RemoveProduct(productID string) error {
	if _, exists := inv.products[productID]; exists {
		delete(inv.products, productID)
		return nil
	} else {
		return errors.New("такого продукта нет")
	}
}

func (inv *Inventory) GetTotalValue() float64 {
	totalValue := 0.0
	for _, product := range inv.products {
		totalValue += product.Price * float64(product.Quantity) 
	}
	return totalValue
}

func main() {
	inventory := NewInventory()
	product := Product{
	ID : "1",
	Name: "Блабла",
	Price: 100.0,
	Quantity: 10,}
	inventory.AddProduct(product)
	fmt.Printf("Общая стоимость: %.2f\n", inventory.GetTotalValue())
}

