package models

//go:generate ffjson $GOFILE
type Client struct {
	Id string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name string `json:"name" example:"João da Silva"`
	Tel string `json:"tel" example:"19970675070"`
	Zipcode string `json:"zipcode" example:"13024430"`
	Address string `json:"address" example:"R. Dr. Sampaio Ferraz, 216"`
}

type ClientCreateQuery struct {
	Name string `json:"name" example:"João da Silva"`
	Tel string `json:"tel" example:"19970675070"`
	Zipcode string `json:"zipcode" example:"13024430"`
	Address string `json:"address" example:"R. Dr. Sampaio Ferraz, 216"`
}

type ClientUpdateQuery ClientCreateQuery