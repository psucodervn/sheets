package model

type Issue struct {
	ID      string  `json:"id"`
	Key     string  `json:"key"`
	Summary string  `json:"summary"`
	Point   float64 `json:"point"`
}

type UserPoint struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	Issues      []Issue `json:"issues"`
	PointTotal  float64 `json:"pointTotal"`
}
