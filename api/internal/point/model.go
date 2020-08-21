package point

import (
	"time"
)

type GetReportResponse struct {
}

type Issue struct {
	ID       string     `json:"id"`
	Key      string     `json:"key"`
	Summary  string     `json:"summary"`
	Point    float64    `json:"point"`
	Status   string     `json:"status"`
	Created  time.Time  `json:"created"`
	Updated  time.Time  `json:"updated"`
	Resolved *time.Time `json:"resolved,omitempty"`
}

type UserPoint struct {
	Name            string  `json:"name"`
	DisplayName     string  `json:"displayName"`
	Issues          []Issue `json:"issues"`
	PointTotal      float64 `json:"pointTotal"`
	WakatimeHuman   string  `json:"wakatimeHuman,omitempty"`
	WakatimeSeconds float64 `json:"wakatimeSeconds"`
}
