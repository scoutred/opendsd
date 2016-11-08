package opendsd

type Approval struct {
	Header []struct {
		Jurisdiction  string `json:"Jurisdiction"`
		Agency        string `json:"Agency"`
		AgencyAddress string `json:"AgencyAddress"`
		AgencyWebsite string `json:"AgencyWebsite"`
		ExtractSystem string `json:"ExtractSystem"`
		ExtractDate   string `json:"ExtractDate"`
		ExtractQuery  string `json:"ExtractQuery"`
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
		ApplicationDate       string      `json:"ApplicationDate"`
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
	BCCodes  []interface{} `json:"BCCodes"`
	Approval struct {
		JobID               int         `json:"JobId"`
		ApprovalID          int         `json:"ApprovalId"`
		Type                string      `json:"Type"`
		Status              string      `json:"Status"`
		Scope               string      `json:"Scope"`
		Depiction           string      `json:"Depiction"`
		IssuedBy            string      `json:"IssuedBy"`
		IssueDate           string      `json:"IssueDate"`
		FirstInspectionDate interface{} `json:"FirstInspectionDate"`
		CompleteCancelDate  interface{} `json:"CompleteCancelDate"`
		PermitHolder        string      `json:"PermitHolder"`
		NetChangeDU         string      `json:"NetChangeDU"`
		Valuation           string      `json:"Valuation"`
		SquareFootage       interface{} `json:"SquareFootage"`
	} `json:"Approval"`
	Inspections        []interface{} `json:"Inspections"`
	InspectionIssues   []interface{} `json:"InspectionIssues"`
	ApprovalFees       []interface{} `json:"ApprovalFees"`
	Exceptions         []interface{} `json:"Exceptions"`
	DependentPackages  []interface{} `json:"DependentPackages"`
	DependantApprovals []interface{} `json:"DependantApprovals"`
}
