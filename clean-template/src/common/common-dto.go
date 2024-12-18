package common

type FilterQueryDTO struct {
	Conditions []Condition `json:"conditions"` // Renamed from Filters
	Sorts      []Sort      `json:"sorts"`      // Renamed from Orders
	MaxResults uint        `json:"maxResults"` // Renamed from Limit
	Offset     uint        `json:"offset"`     // Renamed from OffSet
	Logic      string      `json:"logic"`      // Renamed from Operation (e.g., AND/OR)
}
