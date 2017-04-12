package opendsd

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
	ReportTitle   string               `xml:"report_title"`
	RequestID     string               `xml:"request_id"`
}
