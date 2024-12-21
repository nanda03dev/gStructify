package common

const (
	CONDITION_EQ  ConditionOperation = "EQ"
	CONDITION_IN  ConditionOperation = "IN"
	CONDITION_GT  ConditionOperation = "GT"
	CONDITION_LT  ConditionOperation = "LT"
	CONDITION_GTE ConditionOperation = "GTE"
	CONDITION_LTE ConditionOperation = "LTE"
	CONDITION_NQ  ConditionOperation = "NQ"
)

const (
	SORT_ASC  SortDirection = "ASC"
	SORT_DESC SortDirection = "DESC"
)

const (
	ENTITY_CREATED EventType = "ENTITY_CREATED"
	ENTITY_UPDATED EventType = "ENTITY_UPDATED"
	ENTITY_DELETED EventType = "ENTITY_DELETED"
)
