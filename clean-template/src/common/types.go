package common

type ConditionOperation string

type Condition struct {
	Key      string             `json:"key"`
	Value    string             `json:"value"`
	Operator ConditionOperation `json:"operator"`
}

type SortDirection string

type Sort struct {
	Key  string        `json:"key"`
	Type SortDirection `json:"type"`
}

type FilterQuery struct {
	Conditions []Condition `json:"conditions"`
	Sorts      []Sort      `json:"sorts"`
	MaxResults uint        `json:"maxResults"`
	Offset     uint        `json:"offset"`
	Logic      string      `json:"logic"`
}

type EventType string

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
