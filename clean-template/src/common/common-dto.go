package common

type FilterQueryDTO struct {
	Filters []Filter `json:"filters"`
	Orders  []Order  `json:"orders"`
	Limit   uint     `json:"limit"`
	OffSet  uint     `json:"offSet"`
}
