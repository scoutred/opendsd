package opendsd_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/scoutred/opendsd"
)

var codeEnforcementTestCaseData = `
<?xml version="1.0" encoding="UTF-8"?>
<extract_results>
  <!--City of San Diego Development Services Department-->
  <!--Please see dtd for extract explanation-->
  <metadata>
    <jurisdiction>City of San Diego</jurisdiction>
    <agency>Development Services Department</agency>
    <agency_address>1222 First AV, San Diego, CA 92101</agency_address>
    <agency_website>http://www.sandiego.gov/development-services/</agency_website>
    <data_extract>
      <extract_system>PTSxml Report Id: 203</extract_system>
      <query>VWXML_CECASES</query>
      <extract_date>11/06/2016 22:35</extract_date>
      <report_title>Code Enforcement Report</report_title>
      <request_id/>
    </data_extract>
  </metadata>
  <cases>
    <case case_id="3351">
      <case_source>Other City Department</case_source>
      <description>GARAGE CONVERSN -   [Addr:  2034 JULIAN  AV   APN: 538-330-28   Owner: CORDERO, MARIE/NORMA ] </description>
      <open_date>1993-03-30</open_date>
      <APN>5383302800</APN>
      <street_address>2034 JULIAN AV </street_address>
      <sortable_street_address>JULIAN AV 0000002034 </sortable_street_address>
      <map_reference>XXXX-XX</map_reference>
      <longitude>-117.140165</longitude>
      <latitude>32.70176</latitude>
      <nad83_northing>1836282</nad83_northing>
      <nad83_easting>6287823</nad83_easting>
      <workgroup>Zoning</workgroup>
      <investigator_name>Vasquez, Edward</investigator_name>
      <investigator_phone_number>(619)533-6275</investigator_phone_number>
      <investigator_email_address>EAVasquez@sandiego.gov</investigator_email_address>
      <investigator_active>Y</investigator_active>
      <last_action>City Attorney Referral - Referred</last_action>
      <last_action_due_date>2015-06-06</last_action_due_date>
      <remedy_msg>New plans submitted &amp; approved  with the CED Stamp for submittal to DSD Reveiw Section per VMN.</remedy_msg>
      <complaints>
        <complaint complaint_type_id="314">
          <complaint_type>Zoning-Garage Conversion</complaint_type>
        </complaint>
      </complaints>
    </case>
  </cases>
</extract_results>`

func TestDecodCodeEnforcementCases(t *testing.T) {
	var err error

	buf := bytes.NewBufferString(codeEnforcementTestCaseData)
	cases, err := opendsd.DecodeCodeEnforcementCases(buf)
	if err != nil {
		t.Errorf("error parsing: %v", err)
	}

	log.Printf("%+v", cases)
}
