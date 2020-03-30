package point

import (
	"api/model"
	"context"
	"github.com/go-resty/resty/v2"
)

type Service interface {
	UserPoints(ctx context.Context, month, year int) ([]model.UserPoint, error)
}

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

type issueFields struct {
	Summary  string         `json:"summary"`
	Assignee *issueAssignee `json:"assignee"`
	Point    *float64       `json:"customfield_10106"`
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

func (s *RestService) UserPoints(ctx context.Context, month, year int) ([]model.UserPoint, error) {
	resp, err := s.restyClient.R().
		SetResult(&searchResponse{}).
		SetQueryParam("maxResults", "1000").
		SetQueryParam("fields", "assignee,project,customfield_10106,summary").
		SetQueryParam("jql", "status = Done AND resolved >= 2020-03-01 AND resolved <= 2020-03-31").
		Get("/rest/api/2/search")
	if err != nil {
		return nil, err
	}

	res := resp.Result().(*searchResponse)
	m := make(map[string][]model.Issue)
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
		mp[as.Key] += p
		m[as.Key] = append(m[as.Key], model.Issue{
			ID:      is.ID,
			Key:     is.Key,
			Summary: is.Fields.Summary,
			Point:   p,
		})
	}

	ups := make([]model.UserPoint, 0)
	for uk := range m {
		ups = append(ups, model.UserPoint{
			Name:        ma[uk].Name,
			DisplayName: ma[uk].DisplayName,
			Issues:      m[uk],
			PointTotal:  mp[uk],
		})
	}

	return ups, nil
}

func getFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}
