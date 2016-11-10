package opendsd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Invoice struct {
	InvoiceID          int         `json:"InvoiceId"`
	Status             string      `json:"Status"`
	IssueDate          Timestamp   `json:"IssueDate"`
	IssuedBy           string      `json:"IssuedBy"`
	CustomerName       string      `json:"CustomerName"`
	CustomerFirmName   string      `json:"CustomerFirmName"`
	TotalAmount        float64     `json:"TotalAmount"`
	PaidDate           string      `json:"PaidDate"`
	ReverseDate        Timestamp   `json:"ReverseDate"`
	InvoiceTotalAmount interface{} `json:"InvoiceTotalAmount"`
	Header             []struct {
		Jurisdiction  string                 `json:"Jurisdiction"`
		Agency        string                 `json:"Agency"`
		AgencyAddress string                 `json:"AgencyAddress"`
		AgencyWebsite string                 `json:"AgencyWebsite"`
		ExtractSystem string                 `json:"ExtractSystem"`
		ExtractDate   HeaderExtractTimestamp `json:"ExtractDate"`
		ExtractQuery  string                 `json:"ExtractQuery"`
	} `json:"Header"`
	InvoiceNotes []struct {
		InvoiceID int    `json:"InvoiceId"`
		NoteID    int    `json:"NoteId"`
		Note      string `json:"Note"`
	} `json:"InvoiceNotes"`
	ProjectDetails []struct {
		Jobs        []interface{} `json:"Jobs"`
		ProjectFees []struct {
			InvoiceDetialID     int         `json:"InvoiceDetialId"`
			InvoiceID           int         `json:"InvoiceId"`
			ProjectID           int         `json:"ProjectId"`
			JobID               interface{} `json:"JobId"`
			ApprovalID          interface{} `json:"ApprovalId"`
			FeeTypeID           int         `json:"FeeTypeId"`
			FeeDescription      string      `json:"FeeDescription"`
			FeeQuantityRequired string      `json:"FeeQuantityRequired"`
			FeeTypeUnit         string      `json:"FeeTypeUnit"`
			FeeAmount           float64     `json:"FeeAmount"`
			PreviousCreditInv   string      `json:"PreviousCreditInv"`
			FeeAuthority        string      `json:"FeeAuthority"`
			InvoiceCalcRules    []struct {
				InvoiceDetailID   int           `json:"InvoiceDetailId"`
				FeeThresholdID    int           `json:"FeeThresholdId"`
				BaseQuantity      string        `json:"BaseQuantity"`
				BaseIncrement     string        `json:"BaseIncrement"`
				RuleRate          string        `json:"RuleRate"`
				RuleRateAmt       float64       `json:"RuleRateAmt"`
				FeeRule           string        `json:"FeeRule"`
				PreSurchargeAmt   float64       `json:"PreSurchargeAmt"`
				PreviousCreditAmt string        `json:"PreviousCreditAmt"`
				FeeSurchargeRules []interface{} `json:"FeeSurchargeRules"`
			} `json:"InvoiceCalcRules"`
		} `json:"ProjectFees"`
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
		ProjectFeesTotal      float64     `json:"ProjectFeesTotal"`
		ProjectFeesSubTotal   float64     `json:"ProjectFeesSubTotal"`
		ProjectManagerID      int         `json:"ProjectManagerId"`
		ProjectManager        struct {
			ProjectManagerID int    `json:"ProjectManagerId"`
			Name             string `json:"Name"`
			PhoneNum         string `json:"PhoneNum"`
			EmailAddress     string `json:"EmailAddress"`
			ActiveIndicator  bool   `json:"ActiveIndicator"`
		} `json:"ProjectManager"`
	} `json:"ProjectDetails"`
	InvoiceRevenue []struct {
		InvoiceID     int     `json:"InvoiceId"`
		FundNum       string  `json:"FundNum"`
		GLNum         string  `json:"GLNum"`
		CostObjectNum string  `json:"CostObjectNum"`
		Fund          string  `json:"Fund"`
		GL            string  `json:"GL"`
		Amount        float64 `json:"Amount"`
	} `json:"InvoiceRevenue"`
}

func DecodeInvoice(r io.Reader) (*Invoice, error) {
	var err error
	var invoice Invoice

	if err = json.NewDecoder(r).Decode(&invoice); err != nil {
		return nil, err
	}

	return &invoice, nil
}

func InvoiceByID(id int) (*Invoice, error) {
	//	build our request url
	url := fmt.Sprintf("http://opendsd.sandiego.gov/api/invoice/%v", id)
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

	return DecodeInvoice(resp.Body)
}
