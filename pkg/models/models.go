package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     *string   `json:"phone,omitempty"`
	Email     *string   `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Route struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	OriginName      string    `json:"origin_name"`
	DestinationName string    `json:"destination_name"`
	BasePriceCents  int       `json:"base_price_cents"`
	Currency        string    `json:"currency"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Stop struct {
	ID        string    `json:"id"`
	RouteID   string    `json:"route_id"`
	Name      string    `json:"name"`
	StopOrder int       `json:"stop_order"`
	Latitude  *float64  `json:"latitude,omitempty"`
	Longitude *float64  `json:"longitude,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Trip struct {
	ID            string    `json:"id"`
	PassengerID   string    `json:"passenger_id"`
	RouteID       string    `json:"route_id"`
	PickupStopID  *string   `json:"pickup_stop_id,omitempty"`
	DropoffStopID *string   `json:"dropoff_stop_id,omitempty"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	FareCents     *int      `json:"fare_cents,omitempty"`
	Notes         *string   `json:"notes,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RouteWithStops struct {
	Route Route  `json:"route"`
	Stops []Stop `json:"stops"`
}

type CreateTripRequest struct {
	RouteID       string  `json:"route_id"`
	PickupStopID  *string `json:"pickup_stop_id,omitempty"`
	DropoffStopID *string `json:"dropoff_stop_id,omitempty"`
	PaymentMethod string  `json:"payment_method"`
	Notes         *string `json:"notes,omitempty"`
}
