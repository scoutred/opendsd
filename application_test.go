package opendsd_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/scoutred/opendsd"
)

var applicationTestData = `
<?xml version="1.0" encoding="UTF-8"?>
<extract_results>
  <!--City of San Diego Development Services Department-->
  <!--Please see dtd for extract explanation-->
  <metadata>
    <jurisdiction>City of San Diego</jurisdiction>
    <agency>Development Services Department</agency>
    <agency_address>APPLICATION_PARM</agency_address>
    <agency_website>APPLICATION_PARM</agency_website>
    <data_extract>
      <extract_system>PTSxml Report Id:  47</extract_system>
      <query>VWXML_WEEKLYPERMITSAPPLRECV</query>
      <extract_date>10/23/2016 21:23</extract_date>
      <start_date>10/17/2016</start_date>
      <end_date>10/23/2016</end_date>
      <report_title>Weekly_PermitsApplRecv</report_title>
      <request_id/>
    </data_extract>
  </metadata>
  <approvals>
    <approval approval_id="1795317">
      <approval_type>Map Waiver</approval_type>
      <approval_type_id>244</approval_type_id>
      <project_id>447621</project_id>
      <map_reference>XXXX-XX</map_reference>
      <latitude>32.746437</latitude>
      <longitude>-117.139728</longitude>
      <application_date>10/21/2016 11:05</application_date>
    </approval>
  </approvals>
</extract_results>
`

func TestDecodApplications(t *testing.T) {
	var err error

	buf := bytes.NewBufferString(applicationTestData)

	applications, err := opendsd.DecodeApplication(buf)
	if err != nil {
		t.Errorf("error parsing: %v", err)
	}

	log.Printf("%+v", applications)
}
