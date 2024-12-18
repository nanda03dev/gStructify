package common

// Operation defines the possible filter operations.
type Operation string

const (
	OperationEQ  Operation = "EQ"
	OperationIN  Operation = "IN"
	OperationGT  Operation = "GT"
	OperationLT  Operation = "LT"
	OperationGTE Operation = "GTE"
	OperationLTE Operation = "LTE"
	OperationNQ  Operation = "NQ"
)

type OrderType string

const (
	OrderASC  OrderType = "ASC"
	OrderDESC OrderType = "DESC"
)

// Order struct to hold sorting information
type Order struct {
	Key  string    `json:"Key"`
	Type OrderType `json:"Type"`
}

// Filter defines the filter criteria.
type Filter struct {
	Key       string    `json:"Key"`
	Value     string    `json:"Value"`
	Operation Operation `json:"Operation"`
}
