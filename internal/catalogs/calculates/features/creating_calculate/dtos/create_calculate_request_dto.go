package dtos

type CreateCalculateRequestDto struct {
	Distance        float64 `json:"distance"`
	FuelConsumption float64 `json:"fuel_consumption"`
	FuelPrice       float64 `json:"fuel_price"`
	Passengers      int     `json:"passengers"`
}
