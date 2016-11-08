package opendsd

import (
	"encoding/xml"
	"io"
	"time"
)

const (
	ApplicationTimestampFormat = "01/02/2006 15:04"
	ApplicationDatestampFormat = "01/02/2006"
)

type Applications struct {
	Metadata  Metadata              `xml:"metadata"`
	Approvals []ApplicationApproval `xml:"approvals>approval"`
}

type Metadata struct {
	Jurisdiction  string      `xml:"jurisdiction"`
	Agency        string      `xml:"agency"`
	AgencyAddress string      `xml:"agency_address"`
	AgencyWebsite string      `xml:"agency_website"`
	DataExtract   DataExtract `xml:"data_extract"`
}

type DataExtract struct {
	ExtractSystem string               `xml:"extract_system"`
	Query         string               `xml:"query"`
	ExtractDate   ApplicationTimestamp `xml:"extract_date"`
	StartDate     ApplicationDatestamp `xml:"start_date"`
	EndDate       ApplicationDatestamp `xml:"end_date"`
	ReportTitle   string               `xml:"report_title"`
	RequestID     string               `xml:"request_id"`
}

type ApplicationApproval struct {
	XMLName         xml.Name             `xml:"approval"`
	ID              string               `xml:"approval_id,attr"`
	Type            string               `xml:"approval_type"`
	TypeID          int                  `xml:"approval_type_id"`
	ProjectID       int                  `xml:"project_id"`
	MapReference    string               `xml:"map_reference"`
	Lat             float64              `xml:"latitude"`
	Lon             float64              `xml:"longitude"`
	ApplicationDate ApplicationTimestamp `xml:"application_date"`
}

func DecodeApplication(r io.Reader) (*Applications, error) {
	var err error
	var applications Applications

	if err = xml.NewDecoder(r).Decode(&applications); err != nil {
		return nil, err
	}

	return &applications, nil
}

type ApplicationTimestamp time.Time

func (at *ApplicationTimestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var val string
	d.DecodeElement(&val, &start)

	v, err := time.Parse(ApplicationTimestampFormat, val)
	if err != nil {
		return err
	}
	*at = ApplicationTimestamp(v)
	return nil
}

func (at ApplicationTimestamp) String() string {
	return time.Time(at).Format(time.RFC3339)
}

type ApplicationDatestamp time.Time

func (ad ApplicationDatestamp) String() string {
	return time.Time(ad).Format(time.RFC3339)
}

func (ad *ApplicationDatestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tstr string
	d.DecodeElement(&tstr, &start)

	//	empty time check
	if tstr == "null" || tstr == "" {
		return nil
	}

	v, err := time.Parse(ApplicationDatestampFormat, tstr)
	if err != nil {
		return err
	}
	*ad = ApplicationDatestamp(v)
	return nil
}
