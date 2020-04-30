package wakatime

import (
	"time"
)

type SummaryData struct {
	Categories []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"categories"`
	Dependencies []interface{} `json:"dependencies"`
	Editors      []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"editors"`
	GrandTotal struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"grand_total"`
	Languages []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"languages"`
	Machines []struct {
		Digital       string  `json:"digital"`
		Hours         int     `json:"hours"`
		MachineNameID string  `json:"machine_name_id"`
		Minutes       int     `json:"minutes"`
		Name          string  `json:"name"`
		Percent       float64 `json:"percent"`
		Seconds       int     `json:"seconds"`
		Text          string  `json:"text"`
		TotalSeconds  float64 `json:"total_seconds"`
	} `json:"machines"`
	OperatingSystems []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"operating_systems"`
	Projects []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"projects"`
	Range struct {
		Date     string    `json:"date"`
		End      time.Time `json:"end"`
		Start    time.Time `json:"start"`
		Text     string    `json:"text"`
		Timezone string    `json:"timezone"`
	} `json:"range"`
}

type LeaderboardLeaders struct {
	CurrentUser struct {
		Rank         int `json:"rank"`
		RunningTotal struct {
			DailyAverage              int    `json:"daily_average"`
			HumanReadableDailyAverage string `json:"human_readable_daily_average"`
			HumanReadableTotal        string `json:"human_readable_total"`
			Languages                 []struct {
				Name         string  `json:"name"`
				TotalSeconds float64 `json:"total_seconds"`
			} `json:"languages"`
			ModifiedAt   time.Time `json:"modified_at"`
			TotalSeconds float64   `json:"total_seconds"`
		} `json:"running_total"`
		User struct {
			FullName             string `json:"full_name"`
			HumanReadableWebsite string `json:"human_readable_website"`
			ID                   string `json:"id"`
			Location             string `json:"location"`
			Photo                string `json:"photo"`
			Username             string `json:"username"`
			Website              string `json:"website"`
		} `json:"user"`
	} `json:"current_user"`
	Data []struct {
		Rank         int `json:"rank"`
		RunningTotal struct {
			DailyAverage              int    `json:"daily_average"`
			HumanReadableDailyAverage string `json:"human_readable_daily_average"`
			HumanReadableTotal        string `json:"human_readable_total"`
			IsUpToDate                bool   `json:"is_up_to_date"`
			Languages                 []struct {
				Name         string  `json:"name"`
				TotalSeconds float64 `json:"total_seconds"`
			} `json:"languages"`
			ModifiedAt   time.Time `json:"modified_at"`
			TotalSeconds float64   `json:"total_seconds"`
		} `json:"running_total"`
		User struct {
			DisplayName          string      `json:"display_name"`
			Email                interface{} `json:"email"`
			EmailPublic          bool        `json:"email_public"`
			FullName             string      `json:"full_name"`
			HumanReadableWebsite string      `json:"human_readable_website"`
			ID                   string      `json:"id"`
			Location             string      `json:"location"`
			Photo                string      `json:"photo"`
			PhotoPublic          bool        `json:"photo_public"`
			Username             string      `json:"username"`
			Website              string      `json:"website"`
		} `json:"user"`
	} `json:"data"`
	Language   interface{} `json:"language"`
	ModifiedAt time.Time   `json:"modified_at"`
	Page       int         `json:"page"`
	Range      struct {
		EndDate   string `json:"end_date"`
		EndText   string `json:"end_text"`
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		StartText string `json:"start_text"`
		Text      string `json:"text"`
	} `json:"range"`
	Timeout    int  `json:"timeout"`
	TotalPages int  `json:"total_pages"`
	WritesOnly bool `json:"writes_only"`
}

type UserID string
type UserTime struct {
	Username    string  `json:"username"`
	DisplayName string  `json:"displayName"`
	Human       string  `json:"human"`
	Seconds     float64 `json:"seconds"`
}
type Users map[UserID]UserTime
