package opendsd

import (
	"encoding/xml"
	"io"
)

type CodeEnforcementCases struct {
	Metadata Metadata              `xml:"metadata"`
	Cases    []CodeEnforcementCase `xml:"cases>case"`
}

type CodeEnforcementCase struct {
	XMLName                  xml.Name                       `xml:"case"`
	ID                       string                         `xml:"case_id,attr"`
	CaseSource               string                         `xml:"case_source"`
	Description              string                         `xml:"description"`
	OpenDate                 string                         `xml:"open_date"`
	CloseDate                string                         `xml:"close_date"`
	CloseReason              string                         `xml:"close_reason"`
	CloseNote                string                         `xml:"close_note"`
	APN                      string                         `xml:"APN"`
	StreetAddress            string                         `xml:"street_address"`
	SortableStreetAddress    string                         `xml:"sortable_street_address"`
	MapReference             string                         `xml:"map_reference"`
	Lat                      float64                        `xml:"latitude"`
	Lon                      float64                        `xml:"longitude"`
	NAD83Northing            string                         `xml:"nad83_northing"`
	NAD83Easting             string                         `xml:"nad83_easting"`
	Workgroup                string                         `xml:"workgroup"`
	InvestigatorName         string                         `xml:"investigator_name"`
	InvestigatorPhoneNumber  string                         `xml:"investigator_phone_number"`
	InvestigatorEmailAddress string                         `xml:"investigator_email_address"`
	InvestigatorActive       string                         `xml:"investigator_active"`
	LastAction               string                         `xml:"last_action"`
	LastActionDueDate        string                         `xml:"last_action_due_date"`
	RemedyMsg                string                         `xml:"remedy_msg"`
	Complaints               []CodeEnforcementCaseComplaint `xml:"complaints>complaint"`
}

type CodeEnforcementCaseComplaint struct {
	XMLName xml.Name `xml:"complaint"`
	TypeID  string   `xml:"complaint_type_id,attr"`
	Type    string   `xml:"complaint_type"`
}

func DecodeCodeEnforcementCases(r io.Reader) (*CodeEnforcementCases, error) {
	var err error
	var cec CodeEnforcementCases

	if err = xml.NewDecoder(r).Decode(&cec); err != nil {
		return nil, err
	}

	return &cec, nil
}
