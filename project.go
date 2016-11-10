package opendsd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Project struct {
	Customers []struct {
		ProjectID  int    `json:"ProjectId"`
		CustomerID int    `json:"CustomerId"`
		Role       string `json:"Role"`
		FirmName   string `json:"FirmName"`
		Name       string `json:"Name"`
	} `json:"Customers"`
	ReviewCycles []struct {
		ReviewCycleID  int         `json:"ReviewCycleId"`
		CycleNum       int         `json:"CycleNum"`
		Method         string      `json:"Method"`
		Status         string      `json:"Status"`
		StatusSequence int         `json:"StatusSequence"`
		SubmitDate     interface{} `json:"SubmitDate"`
		DueDate        interface{} `json:"DueDate"`
		CloseDate      interface{} `json:"CloseDate"`
		Performance    interface{} `json:"Performance"`
		Reviews        []struct {
			ReviewCycleID int         `json:"ReviewCycleId"`
			ReviewID      int         `json:"ReviewId"`
			Discipline    string      `json:"Discipline"`
			Status        string      `json:"Status"`
			DueDate       interface{} `json:"DueDate"`
			CompletedDate interface{} `json:"CompletedDate"`
			Performance   string      `json:"Performance"`
			Name          string      `json:"Name"`
			Phone         string      `json:"Phone"`
			Email         string      `json:"Email"`
			IsActive      bool        `json:"IsActive"`
		} `json:"Reviews"`
	} `json:"ReviewCycles"`
	Jobs []struct {
		SignOffs []struct {
			DisciplineID          int    `json:"DisciplineId"`
			DisciplineDescription string `json:"DisciplineDescription"`
			SignedDate            string `json:"SignedDate"`
		} `json:"SignOffs"`
		ApprovalInfo []interface{} `json:"ApprovalInfo"`
		Approvals    []struct {
			JobID               int         `json:"JobId"`
			ApprovalID          int         `json:"ApprovalId"`
			Type                string      `json:"Type"`
			Status              string      `json:"Status"`
			Scope               string      `json:"Scope"`
			Depiction           string      `json:"Depiction"`
			IssuedBy            string      `json:"IssuedBy"`
			IssueDate           Timestamp   `json:"IssueDate"`
			FirstInspectionDate Timestamp   `json:"FirstInspectionDate"`
			CompleteCancelDate  Timestamp   `json:"CompleteCancelDate"`
			PermitHolder        string      `json:"PermitHolder"`
			NetChangeDU         string      `json:"NetChangeDU"`
			Valuation           string      `json:"Valuation"`
			SquareFootage       interface{} `json:"SquareFootage"`
		} `json:"Approvals"`
		ProjectID             int         `json:"ProjectId"`
		JobID                 int         `json:"JobId"`
		Description           string      `json:"Description"`
		APN                   string      `json:"APN"`
		StreetAddress         string      `json:"StreetAddress"`
		MapReference          string      `json:"MapReference"`
		SortableStreetAddress string      `json:"SortableStreetAddress"`
		Latitude              float64     `json:"Latitude"`
		Longitude             float64     `json:"Longitude"`
		NAD83Northing         interface{} `json:"NAD83Northing"`
		NAD83Easting          interface{} `json:"NAD83Easting"`
		JobFeesSubTotal       interface{} `json:"JobFeesSubTotal"`
	} `json:"Jobs"`
	Fees []struct {
		FeeID            int    `json:"FeeId"`
		Description      string `json:"Description"`
		Category         string `json:"Category"`
		Unit             string `json:"Unit"`
		QuantityRequired int    `json:"QuantityRequired"`
		QuantityPaid     int    `json:"QuantityPaid"`
		InvoiceID        int    `json:"InvoiceId"`
		InvoiceStatus    string `json:"InvoiceStatus"`
		ProjectID        int    `json:"ProjectId"`
	} `json:"Fees"`
	Invoices []struct {
		InvoiceID        int    `json:"InvoiceId"`
		InvoiceIssueDate string `json:"InvoiceIssueDate"`
		InvoiceStatus    string `json:"InvoiceStatus"`
		ProjectID        int    `json:"ProjectId"`
	} `json:"Invoices"`
	ProjectID             int         `json:"ProjectId"`
	Title                 string      `json:"Title"`
	Scope                 string      `json:"Scope"`
	ApplicationExpiration Timestamp   `json:"ApplicationExpiration"`
	ApplicationExpired    bool        `json:"ApplicationExpired"`
	AdminHold             bool        `json:"AdminHold"`
	DevelopmentID         int         `json:"DevelopmentId"`
	DevelopmentTitle      string      `json:"DevelopmentTitle"`
	ApplicationDate       Timestamp   `json:"ApplicationDate"`
	AccountNum            string      `json:"AccountNum"`
	JobOrderNum           interface{} `json:"JobOrderNum"`
	Header                []struct {
		Jurisdiction  string                 `json:"Jurisdiction"`
		Agency        string                 `json:"Agency"`
		AgencyAddress string                 `json:"AgencyAddress"`
		AgencyWebsite string                 `json:"AgencyWebsite"`
		ExtractSystem string                 `json:"ExtractSystem"`
		ExtractDate   HeaderExtractTimestamp `json:"ExtractDate"`
		ExtractQuery  string                 `json:"ExtractQuery"`
	} `json:"Header"`
	ProjectManagerID int `json:"ProjectManagerId"`
	ProjectManager   struct {
		ProjectManagerID int    `json:"ProjectManagerId"`
		Name             string `json:"Name"`
		PhoneNum         string `json:"PhoneNum"`
		EmailAddress     string `json:"EmailAddress"`
		ActiveIndicator  bool   `json:"ActiveIndicator"`
	} `json:"ProjectManager"`
}

func DecodeProject(r io.Reader) (*Project, error) {
	var err error
	var project Project

	if err = json.NewDecoder(r).Decode(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

func ProjectByID(id int) (*Project, error) {
	//	build our request url
	url := fmt.Sprintf("http://opendsd.sandiego.gov/api/project/%v", id)
	//	build our request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	//	the API returns a 302 (Temporary Redirect) status code but does
	//	not provide a Location header for the redirect. Because of this
	//	we need to use a DefaultTransport rather than client.Do()
	//	make our request
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return DecodeProject(resp.Body)
}
