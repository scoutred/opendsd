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

//  curl 'http://opendsd.sandiego.gov/api/approval/1117208' -H 'Pragma: no-cache' -H 'Accept-Encoding: gzip, deflate, sdch' -H 'Accept-Language: en-US,en;q=0.8' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36' -H 'Accept: application/json' -H 'Cache-Control: no-cache' -H 'Cookie: __cfduid=da5f8a189ad7f88ad881e454b387d611b1477864804; vxv=2015041001; vxu=EO_47RRkFu7I0JF2z9edzg; vxr=87.81; vxp=12; vxl=1477865568; _ga=GA1.2.445451898.1477864807' -H 'Connection: keep-alive' --compressed
