package opendsd

import (
	"encoding/json"
	"fmt"
	"io"
)

type CodeEnforcement struct {
	Complaints []struct {
		ComplaintType string `json:"ComplaintType"`
		CaseID        int    `json:"CaseId"`
	} `json:"Complaints"`
	CaseID                int       `json:"CaseId"`
	CaseSource            string    `json:"CaseSource"`
	Description           string    `json:"Description"`
	OpenDate              Timestamp `json:"OpenDate"`
	CloseDate             Timestamp `json:"CloseDate"`
	StreetAddress         string    `json:"StreetAddress"`
	SortableStreetAddress string    `json:"SortableStreetAddress"`
	CloseNote             string    `json:"CloseNote"`
	CloseReason           string    `json:"CloseReason"`
	APN                   string    `json:"APN"`
	MapReference          string    `json:"MapReference"`
	Longitude             float64   `json:"Longitude"`
	Latitude              float64   `json:"Latitude"`
	NAD83Northing         string    `json:"NAD83Northing"`
	NAD83Easting          string    `json:"NAD83Easting"`
	Workgroup             string    `json:"Workgroup"`
	InvestigatorName      string    `json:"InvestigatorName"`
	InvestiagtorPhone     string    `json:"InvestiagtorPhone"`
	InvestigatorActive    string    `json:"InvestigatorActive"`
	InvestigatorEmail     string    `json:"InvestigatorEmail"`
	LastAction            string    `json:"LastAction"`
	LastActionDueDate     string    `json:"LastActionDueDate"`
	RemedyMsg             string    `json:"RemedyMsg"`
	Header                []struct {
		Jurisdiction  string                 `json:"Jurisdiction"`
		Agency        string                 `json:"Agency"`
		AgencyAddress string                 `json:"AgencyAddress"`
		AgencyWebsite string                 `json:"AgencyWebsite"`
		ExtractSystem string                 `json:"ExtractSystem"`
		ExtractDate   HeaderExtractTimestamp `json:"ExtractDate"`
		ExtractQuery  string                 `json:"ExtractQuery"`
	} `json:"Header"`
}

func DecodeCodeEnforcement(r io.Reader) (*CodeEnforcement, error) {
	var err error
	var ce CodeEnforcement

	if err = json.NewDecoder(r).Decode(&ce); err != nil {
		return nil, err
	}

	return &ce, nil
}

func (c *Client) CodeEnforcementByID(id int) (*CodeEnforcement, error) {
	var err error
	var ce CodeEnforcement

	uri := fmt.Sprintf("/codeenforcement/%v", id)
	if err = c.get(uri, &ce); err != nil {
		return nil, err
	}

	//	this is necessary since the API does not report a 404 on not found responses
	if ce.CaseID != id {
		return nil, APIError{
			ErrorMessage: fmt.Sprintf("CodeEnforcement with ID: %v could not be found.", id),
		}
	}

	return &ce, nil
}
