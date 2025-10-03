// Ларионова Арина 363
package main

import (
	"errors"
	"fmt"
)

type Customers struct {
	Name  string
	Email string
}

type OrderItem struct {
	ProductName string
	Quantity    int
	Price       float64
}

type Order struct {
	ID       string
	Costumer Customers
	Item     []OrderItem
	Status   string
}

func (o *Order) AddItem(item OrderItem) {
	o.Item = append(o.Item, item)
}

func (o *Order) RemoveItem(index int) error {
	if index < 0 || index >= len(o.Item) {
		return errors.New("недопустимый индекс для удаления")
	} else {
		o.Item = append(o.Item[:index], o.Item[index+1:]...)
	}
	return nil
}

func (o *Order) GetTotalAmount() float64 {
	total := 0.0
	for _, item := range o.Item {
		total += float64(item.Quantity) * item.Price
	}
	return total
}

func (o *Order) UpdateStatus(status string) {
	o.Status = status
}

func main() {
	customer := Customers{
		Name:  "Arnir",
		Email: "enen@mail.com",
	}

	order := Order{
		ID:       "1",
		Costumer: customer,
		Status:   "В процессе",
	}
	order.AddItem(OrderItem{
		ProductName: "Ябло",
		Quantity:    5,
		Price:       10.0,
	})
	Amount := order.GetTotalAmount()
	fmt.Printf("Общая сумма: %.2f\n", Amount)

	order.UpdateStatus("Выполнено")
	fmt.Printf("Статус заказа: %s\n", order.Status)

	err := order.RemoveItem(1)
	if err != nil {
		fmt.Println("Ошибка при удалении товара:", err)
	} else {
		fmt.Println("Товар успешно удален.")
	}
}
