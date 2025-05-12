package incident

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	WaroomURL = "https://api.app.waroom.com/api/v0/incidents"
)

type IncidentResponce struct {
	Incidents        []Incident       `json:"incidents"`
	ResponceMetadata ResponseMetadata `json:"response_metadata"`
}

type Incident struct {
	Uuid         string       `json:"uuid"`
	Title        string       `json:"title"`
	Severity     string       `json:"severity"`
	Status       string       `json:"status"`
	RootCause    string       `json:"root_cause"`
	Metrics      Metrics      `json:"metrics"`
	Experimental bool         `json:"experimental"`
	Service      Service      `json:"service"`
	Labels       []Label      `json:"labels"`
	CreatedAt    string       `json:"created_at"`
	Postmortem   []Postmortem `json:"postmortems"`
}

type Metrics struct {
	Ttd int `json:"ttd"`
	Tta int `json:"tta"`
	Tti int `json:"tti"`
	Ttf int `json:"ttf"`
	Ttr int `json:"ttr"`
}

type Service struct {
	Name string `json:"name"`
}

type Label struct {
	Name string `json:"name"`
}
type Postmortem struct {
	Uuid   string `json:"uuid"`
	Title  string `json:"title"`
	Blob   string `json:"blob"`
	Status string `json:"status"`
}

type ResponseMetadata struct {
	CurrentPage *int `json:"current_page"`
	NextPage    *int `json:"next_page"`
	PrevPage    *int `json:"prev_page"`
	TotalPages  *int `json:"total_pages"`
	TotalCount  *int `json:"total_count"`
}

type WaroomRequest struct {
	*HttpRequest
	apiKey string
}

func (wr *WaroomRequest) ExecHttpReq(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", wr.apiKey))
	req.Header.Set("Accept", "application/json")
	resp, err := wr.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		return nil, fmt.Errorf("HTTP status code: %d, error: %s", resp.StatusCode, body)
	}
	return body, nil
}

var _ = Requester(&WaroomRequest{})

type Collection struct {
	request Requester
}

type CollectionOption func(*Collection)

func CollectionWithRequest(request Requester) CollectionOption {
	return func(i *Collection) {
		i.request = request
	}
}

func NewCollection(apiKey string, options ...CollectionOption) *Collection {
	collection := &Collection{
		request: &WaroomRequest{
			HttpRequest: NewHttpRequest(),
			apiKey:      apiKey,
		},
	}
	for _, option := range options {
		option(collection)
	}
	return collection
}

func (c *Collection) GetIncidents() ([]Incident, error) {
	var incidents []Incident
	for i := 1; ; i++ {
		req, err := http.NewRequest("GET", WaroomURL, nil)
		if err != nil {
			return nil, err
		}
		q := req.URL.Query()
		q.Add("per_page", "50")
		q.Add("page", fmt.Sprintf("%d", i))
		req.URL.RawQuery = q.Encode()

		body, err := c.request.ExecHttpReq(req)
		if err != nil {
			return nil, err
		}

		var incidentResponce IncidentResponce
		err = json.Unmarshal(body, &incidentResponce)
		if err != nil {
			return nil, err
		}
		incidents = append(incidents, incidentResponce.Incidents...)

		if incidentResponce.ResponceMetadata.NextPage == nil {
			break
		}
	}
	return incidents, nil
}
