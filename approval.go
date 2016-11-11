package opendsd

import (
	"encoding/json"
	"fmt"
	"io"
)

type Approval struct {
	Header []struct {
		Jurisdiction  string                 `json:"Jurisdiction"`
		Agency        string                 `json:"Agency"`
		AgencyAddress string                 `json:"AgencyAddress"`
		AgencyWebsite string                 `json:"AgencyWebsite"`
		ExtractSystem string                 `json:"ExtractSystem"`
		ExtractDate   HeaderExtractTimestamp `json:"ExtractDate"`
		ExtractQuery  string                 `json:"ExtractQuery"`
	} `json:"Header"`
	ApprovalID int `json:"ApprovalId"`
	Project    struct {
		Customers             interface{} `json:"Customers"`
		ReviewCycles          interface{} `json:"ReviewCycles"`
		Jobs                  interface{} `json:"Jobs"`
		Fees                  interface{} `json:"Fees"`
		Invoices              interface{} `json:"Invoices"`
		ProjectID             int         `json:"ProjectId"`
		Title                 string      `json:"Title"`
		Scope                 string      `json:"Scope"`
		ApplicationExpiration string      `json:"ApplicationExpiration"`
		ApplicationExpired    bool        `json:"ApplicationExpired"`
		AdminHold             bool        `json:"AdminHold"`
		DevelopmentID         int         `json:"DevelopmentId"`
		DevelopmentTitle      string      `json:"DevelopmentTitle"`
		ApplicationDate       Timestamp   `json:"ApplicationDate"`
		AccountNum            string      `json:"AccountNum"`
		JobOrderNum           interface{} `json:"JobOrderNum"`
		Header                interface{} `json:"Header"`
		ProjectManagerID      int         `json:"ProjectManagerId"`
		ProjectManager        struct {
			ProjectManagerID int    `json:"ProjectManagerId"`
			Name             string `json:"Name"`
			PhoneNum         string `json:"PhoneNum"`
			EmailAddress     string `json:"EmailAddress"`
			ActiveIndicator  bool   `json:"ActiveIndicator"`
		} `json:"ProjectManager"`
	} `json:"Project"`
	Job struct {
		SignOffs              []interface{} `json:"SignOffs"`
		ApprovalInfo          []interface{} `json:"ApprovalInfo"`
		Approvals             []interface{} `json:"Approvals"`
		ProjectID             int           `json:"ProjectId"`
		JobID                 int           `json:"JobId"`
		Description           string        `json:"Description"`
		APN                   string        `json:"APN"`
		StreetAddress         string        `json:"StreetAddress"`
		MapReference          string        `json:"MapReference"`
		SortableStreetAddress string        `json:"SortableStreetAddress"`
		Latitude              float64       `json:"Latitude"`
		Longitude             float64       `json:"Longitude"`
		NAD83Northing         interface{}   `json:"NAD83Northing"`
		NAD83Easting          interface{}   `json:"NAD83Easting"`
		JobFeesSubTotal       interface{}   `json:"JobFeesSubTotal"`
	} `json:"Job"`
	BCCodes []struct {
		JobID    int    `json:"JobId"`
		BCCodeID int    `json:"BCCodeId"`
		BCCode   string `json:"BCCode"`
	} `json:"BCCodes"`
	Approval struct {
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
	} `json:"Approval"`
	Inspections []struct {
		ApprovalID            int           `json:"ApprovalId"`
		InspectionID          int           `json:"InspectionId"`
		InspectionTier        string        `json:"InspectionTier"`
		InspectionType        string        `json:"InspectionType"`
		InspectionDiscipline  string        `json:"InspectionDiscipline"`
		InspectionStatus      string        `json:"InspectionStatus"`
		InspectionStatusSeq   string        `json:"InspectionStatusSeq"`
		InspectorID           interface{}   `json:"InspectorId"`
		InspectorName         string        `json:"InspectorName"`
		InspectorPhone        string        `json:"InspectorPhone"`
		InspectorActive       string        `json:"InspectorActive"`
		SchedulingInstruction string        `json:"SchedulingInstruction"`
		CustRequestable       string        `json:"CustRequestable"`
		ScheduledDt           interface{}   `json:"ScheduledDt"`
		PerformedDt           interface{}   `json:"PerformedDt"`
		InspectionResult      string        `json:"InspectionResult"`
		InspectorEmail        string        `json:"InspectorEmail"`
		InspectionDetails     []interface{} `json:"InspectionDetails"`
	} `json:"Inspections"`
	InspectionIssues []struct {
		InspectionIssueID int         `json:"InspectionIssueId"`
		ApprovalID        int         `json:"ApprovalId"`
		InspectionTier    string      `json:"InspectionTier"`
		Issue             string      `json:"Issue"`
		Class             string      `json:"Class"`
		Visibility        string      `json:"Visibility"`
		CreatedBy         string      `json:"CreatedBy"`
		CreatedDt         string      `json:"CreatedDt"`
		ClearedBy         string      `json:"ClearedBy"`
		ClearedDt         interface{} `json:"ClearedDt"`
		ClearedNote       string      `json:"ClearedNote"`
	} `json:"InspectionIssues"`
	ApprovalFees []struct {
		FeeTypeID           int    `json:"FeeTypeId"`
		ProjectID           int    `json:"ProjectId"`
		FeeType             string `json:"FeeType"`
		FeeCategory         string `json:"FeeCategory"`
		FeeTypeUnit         string `json:"FeeTypeUnit"`
		FeeQuantityRequired string `json:"FeeQuantityRequired"`
		FeeQuantityPaid     string `json:"FeeQuantityPaid"`
		InvoiceID           int    `json:"InvoiceId"`
		InvoiceStatus       string `json:"InvoiceStatus"`
	} `json:"ApprovalFees"`
	Exceptions []struct {
		ApprovalID        int         `json:"ApprovalId"`
		ApprovalStatusSeq interface{} `json:"ApprovalStatusSeq"`
		Exception         string      `json:"Exception"`
		ApprovalStatus    string      `json:"ApprovalStatus"`
	} `json:"Exceptions"`
	DependentPackages  []interface{} `json:"DependentPackages"`
	DependantApprovals []struct {
		ApprovalID        int    `json:"ApprovalId"`
		InspectionTier    string `json:"InspectionTier"`
		DepApprApprovalID int    `json:"DepApprApprovalId"`
		DepApprType       string `json:"DepApprType"`
		DepApprStatus     string `json:"DepApprStatus"`
		StatusReq         string `json:"StatusReq"`
		RequirementMet    string `json:"RequirementMet"`
		AddedBy           string `json:"AddedBy"`
		AddedDt           string `json:"AddedDt"`
	} `json:"DependantApprovals"`
}

func DecodeApproval(r io.Reader) (*Approval, error) {
	var err error
	var approval Approval

	if err = json.NewDecoder(r).Decode(&approval); err != nil {
		return nil, err
	}

	return &approval, nil
}

func (c *Client) ApprovalByID(id int) (*Approval, error) {
	var err error
	var a Approval

	uri := fmt.Sprintf("/approval/%v", id)
	if err = c.get(uri, &a); err != nil {
		return nil, err
	}

	return &a, nil
}
