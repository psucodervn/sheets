package point

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	dateLayout = "2006-01-02"
)

type Service interface {
	UserPointsInMonth(ctx context.Context, month, year int) ([]UserPoint, error)
	WorkingIssues(ctx context.Context, from time.Time, to time.Time) ([]UserPoint, error)
}

var _ Service = &RestService{}

type RestService struct {
	restyClient *resty.Client
}

func NewRestService(username, password, host string) *RestService {
	cli := resty.New().
		SetBasicAuth(username, password).
		SetHostURL(host).
		SetHeader("Accept", "application/json")
	return &RestService{restyClient: cli}
}

type issueAssignee struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type issueStatus struct {
	Name string `json:"name"`
}

type issueFields struct {
	Summary  string         `json:"summary"`
	Assignee *issueAssignee `json:"assignee"`
	Point    *float64       `json:"customfield_10106"`
	Created  string         `json:"created"`
	Updated  string         `json:"updated"`
	Resolved string         `json:"resolutiondate"`
	Status   *issueStatus   `json:"status"`
}

type issue struct {
	ID     string      `json:"id"`
	Key    string      `json:"key"`
	Self   string      `json:"self"`
	Fields issueFields `json:"fields"`
}

type searchResponse struct {
	Issues []issue `json:"issues"`
}

var fields = strings.Split("assignee,project,customfield_10106,summary,status,updated,created,resolutiondate", ",")

func (s *RestService) WorkingIssues(ctx context.Context, from time.Time, to time.Time) ([]UserPoint, error) {
	lower, upper := getTimeBound(from, to)
	jql := fmt.Sprintf(`(resolved >= %s AND resolved < %s) OR (status = 'In Progress')`, lower, upper)

	resp, err := s.restyClient.R().
		SetResult(&searchResponse{}).
		SetQueryParam("maxResults", "1000").
		SetQueryParam("fields", strings.Join(fields, ",")).
		SetQueryParam("jql", jql).
		SetContext(ctx).
		Get("/rest/api/2/search")
	if err != nil {
		return nil, err
	}

	res := resp.Result().(*searchResponse)
	ups := searchResponseToUserPoints(res)

	return ups, nil
}

func (s *RestService) UserPointsInMonth(ctx context.Context, month, year int) ([]UserPoint, error) {
	lower, upper := getTimeBoundByMonth(month, year)
	jql := fmt.Sprintf(`status = Done AND resolved >= %s AND resolved < %s`, lower, upper)

	resp, err := s.restyClient.R().
		SetResult(&searchResponse{}).
		SetQueryParam("maxResults", "1000").
		SetQueryParam("fields", strings.Join(fields, ",")).
		SetQueryParam("jql", jql).
		SetContext(ctx).
		Get("/rest/api/2/search")
	if err != nil {
		return nil, err
	}

	res := resp.Result().(*searchResponse)
	ups := searchResponseToUserPoints(res)

	return ups, nil
}

func searchResponseToUserPoints(res *searchResponse) []UserPoint {
	m := make(map[string][]Issue)
	ma := make(map[string]issueAssignee)
	mp := make(map[string]float64)
	for _, is := range res.Issues {
		if is.Fields.Assignee == nil {
			continue
		}
		as := is.Fields.Assignee
		if _, ok := ma[as.Key]; !ok {
			ma[as.Key] = *as
		}
		p := getFloat64(is.Fields.Point)
		m[as.Key] = append(m[as.Key], Issue{
			ID:      is.ID,
			Key:     is.Key,
			Summary: is.Fields.Summary,
			Point:   p,
			Status:  is.Fields.Status.Name,
			Created: parseUpdatedTime(is.Fields.Created),
			Updated: parseUpdatedTime(is.Fields.Updated),
		})
		if len(is.Fields.Resolved) > 0 {
			mp[as.Key] += p
			val := parseUpdatedTime(is.Fields.Resolved)
			m[as.Key][len(m[as.Key])-1].Resolved = &val
		}
	}

	ups := make([]UserPoint, 0)
	for uk := range m {
		ups = append(ups, UserPoint{
			Name:        ma[uk].Name,
			DisplayName: ma[uk].DisplayName,
			Issues:      m[uk],
			PointTotal:  mp[uk],
		})
	}
	return ups
}

func getTimeBoundByMonth(month int, year int) (string, string) {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	finish := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	//start = start.Add(-24 * time.Hour)
	return start.Format(dateLayout), finish.Format(dateLayout)
}

func getTimeBound(from time.Time, to time.Time) (string, string) {
	return from.Format(dateLayout), to.Add(24 * time.Hour).Format(dateLayout)
}

func getFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func parseUpdatedTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05.999+0000", s)
	return t
}
