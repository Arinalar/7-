// Ларионова Арина 363
package main

import (
	"fmt"
	"time"
)

type Room struct {
	ID            string  
	Type          string  
	PricePerNight float64 
	Available     bool    
}

type Reservation struct {
	RoomID    string    
	StartDate time.Time 
	EndDate   time.Time 
}

type Hotel struct {
	Rooms []Room 
}

func NewHotel(rooms []Room) *Hotel {
	return &Hotel{
		Rooms: rooms,
	}
}

func (h *Hotel) CheckAvailability(roomID string, startDate, endDate time.Time) (bool, error) {
	for _, room := range h.Rooms {
		if room.ID == roomID {
			if !room.Available {
				return false, fmt.Errorf("номер %s недоступен", roomID)
			}
			return true, nil
		}
	}
	return false, fmt.Errorf("номер %s не найден", roomID)
}

func (h *Hotel) CreateReservation(roomID string, startDate, endDate time.Time) (*Reservation, error) {
	available, err := h.CheckAvailability(roomID, startDate, endDate)
	if !available {
		return nil, err
	}

	reservation := &Reservation{
		RoomID:    roomID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	for i, room := range h.Rooms {
		if room.ID == roomID {
			room.Available = false
			h.Rooms[i] = room 
			break
		}
	}

	return reservation, nil
}

func (r *Reservation) CalculateCost(rooms []Room) (float64, error) {
	var roomPrice float64
	var nights int

	nights = int(r.EndDate.Sub(r.StartDate).Hours() / 24)

	for _, room := range rooms {
		if room.ID == r.RoomID {
			roomPrice = room.PricePerNight
			break
		}
	}

	if roomPrice == 0 {
		return 0, fmt.Errorf("номер с ID %s не найден", r.RoomID)
	}

	return roomPrice * float64(nights), nil
}

func main() {

	rooms := []Room{
		{ID: "101", Type: "Стандарт", PricePerNight: 50, Available: true},
		{ID: "102", Type: "Люкс", PricePerNight: 100, Available: true},
		{ID: "103", Type: "Стандарт", PricePerNight: 50, Available: true},
	}

	hotel := NewHotel(rooms) 

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 3) 

	available, err := hotel.CheckAvailability("101", startDate, endDate)
	if err != nil {
		fmt.Println("Ошибка проверки доступности:", err)
	} else if available {

		reservation, err := hotel.CreateReservation("101", startDate, endDate)
		if err != nil {
			fmt.Println("Ошибка при создании брони:", err)
		} else {

			fmt.Printf("Бронь успешно создана! Номер: %s, Начало: %s, Конец: %s\n", reservation.RoomID, reservation.StartDate.Format("2006-01-02"), reservation.EndDate.Format("2006-01-02"))

			cost, err := reservation.CalculateCost(hotel.Rooms)
			if err != nil {
				fmt.Println("Ошибка при расчете стоимости:", err)
			} else {
				fmt.Printf("Стоимость бронирования: %.2f\n", cost)
			}
		}
	}
}
