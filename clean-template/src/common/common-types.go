package common

// Operation defines the possible filter operations.
type ConditionOperation string

const (
	ConditionEQ  ConditionOperation = "EQ"
	ConditionIN  ConditionOperation = "IN"
	ConditionGT  ConditionOperation = "GT"
	ConditionLT  ConditionOperation = "LT"
	ConditionGTE ConditionOperation = "GTE"
	ConditionLTE ConditionOperation = "LTE"
	ConditionNQ  ConditionOperation = "NQ"
)

type SortDirection string

const (
	SortASC  SortDirection = "ASC"
	SortDESC SortDirection = "DESC"
)

// Order struct to hold sorting information
type Sort struct {
	Key  string        `json:"key"`
	Type SortDirection `json:"type"`
}

// Filter defines the filter criteria.
type Condition struct {
	Key      string             `json:"key"`
	Value    string             `json:"value"`
	Operator ConditionOperation `json:"operator"`
}
