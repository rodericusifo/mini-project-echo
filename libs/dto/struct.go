package dto

type DTO[P, Q, B any] struct {
	Param P
	Query Q
	Body  B
}
