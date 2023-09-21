package dto

type FileDto struct {
	Indicador1 int `json:"Indicador1" validate:"required"`
	Indicador2 int `json:"Indicador2" validate:"required"`
	Indicador3 int `json:"Indicador3" validate:"required"`
}
