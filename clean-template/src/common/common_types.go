package common

// Operation defines the possible filter operations.
type ConditionOperation string

const (
	CONDITION_EQ  ConditionOperation = "EQ"
	CONDITION_IN  ConditionOperation = "IN"
	CONDITION_GT  ConditionOperation = "GT"
	CONDITION_LT  ConditionOperation = "LT"
	CONDITION_GTE ConditionOperation = "GTE"
	CONDITION_LTE ConditionOperation = "LTE"
	CONDITION_NQ  ConditionOperation = "NQ"
)

type SortDirection string

const (
	SORT_ASC  SortDirection = "ASC"
	SORT_DESC SortDirection = "DESC"
)

type Sort struct {
	Key  string        `json:"key"`
	Type SortDirection `json:"type"`
}

type Condition struct {
	Key      string             `json:"key"`
	Value    string             `json:"value"`
	Operator ConditionOperation `json:"operator"`
}

type EventType string

const (
	ENTITY_CREATED EventType = "ENTITY_CREATED"
	ENTITY_UPDATED EventType = "ENTITY_UPDATED"
	ENTITY_DELETED EventType = "ENTITY_DELETED"
)

type EntityName string

type EntityConfig struct {
	EventStore bool
}

type Event struct {
	ID         string
	EntityId   string
	EntityName EntityName
	Type       EventType
	Config     EntityConfig
}
