package opendsd_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/scoutred/opendsd"
)

var codeEnforcementTestData = `
{
  "Complaints": [
    {
      "ComplaintType": "Zoning-Garage Conversion",
      "CaseId": 3351
    }
  ],
  "CaseId": 3351,
  "CaseSource": "Other City Department",
  "Description": "GARAGE CONVERSN -   [Addr:  2034 JULIAN  AV   APN: 538-330-28   Owner: CORDERO, MARIE/NORMA ] ",
  "OpenDate": "1993-03-30T00:00:00",
  "CloseDate": null,
  "StreetAddress": "2034 JULIAN AV ",
  "SortableStreetAddress": "JULIAN AV 0000002034 ",
  "CloseNote": "",
  "CloseReason": "",
  "APN": "538-330-28-00",
  "MapReference": "XXXX-XX",
  "Longitude": -117.140165,
  "Latitude": 32.70176,
  "NAD83Northing": "1836282",
  "NAD83Easting": "6287823",
  "Workgroup": "Zoning",
  "InvestigatorName": "Vasquez, Edward",
  "InvestiagtorPhone": "(619)533-6275",
  "InvestigatorActive": "Y",
  "InvestigatorEmail": "EAVasquez@sandiego.gov",
  "LastAction": "City Attorney Referral-Referred",
  "LastActionDueDate": "2015-06-06T00:00:00",
  "RemedyMsg": "New plans submitted & approved  with the CED Stamp for submittal to DSD Reveiw Section per VMN.",
  "Header": [
    {
      "Jurisdiction": "City of San Diego",
      "Agency": "Development Services Department",
      "AgencyAddress": "1222 First Ave, San Diego, CA 92101",
      "AgencyWebsite": "www.sandiego.gov",
      "ExtractSystem": "Extract System",
      "ExtractDate": "11/13/2016 7:28:54 PM",
      "ExtractQuery": "Code Enforcement Case XML"
    }
  ]
}`

func TestDecodCodeEnforcement(t *testing.T) {
	var err error

	buf := bytes.NewBufferString(codeEnforcementTestData)

	codeEnforcement, err := opendsd.DecodeCodeEnforcement(buf)
	if err != nil {
		t.Errorf("error parsing: %v", err)
	}

	log.Printf("code enforcement %+v", codeEnforcement)
}

func TestCodeEnforcementByID(t *testing.T) {
	//	setup test muxer
	mux := http.NewServeMux()
	//	register our test route
	mux.HandleFunc("/codeenforcement/3351", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, codeEnforcementTestData)
	})

	//	setup test server
	server := httptest.NewServer(mux)
	//	setup test client
	client := opendsd.NewClient()
	//	use the generated testing url
	client.APIRoot = server.URL

	codeEnforcement, err := client.CodeEnforcementByID(3351)
	if err != nil {
		t.Errorf("CodeEnforcementByID error: %v", err)
	}

	log.Printf("code enforcement %+v", codeEnforcement)
}
