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

var invoiceTestData = `
{
	"InvoiceId": 483626,
	"Status": "Paid",
	"IssueDate": "2013-04-10T15:02:18",
	"IssuedBy": "Huff, Sandra",
	"CustomerName": "Cobian, Neva ",
	"CustomerFirmName": "RBF Consulting",
	"TotalAmount": 19207.0,
	"PaidDate": "2013-04-10T15:58:57",
	"ReverseDate": null,
	"InvoiceTotalAmount": null,
	"Header": [{
		"Jurisdiction": "City of San Diego",
		"Agency": "Development Services Department",
		"AgencyAddress": "1222 First Ave, San Diego, CA 92101",
		"AgencyWebsite": "www.sandiego.gov",
		"ExtractSystem": "Extract System",
		"ExtractDate": "11/7/2016 7:40:25 PM",
		"ExtractQuery": "Invoice XML"
	}],
	"InvoiceNotes": [{
		"InvoiceId": 483626,
		"NoteId": 1,
		"Note": "Upon payment of any Development  Impact Fees (DIF), Regional Transportation Congestion Improvement Program (RTCIP), or Facilities Benefit Assessment (FBA) fees, the 90-day protest period in which you may protest these fees under Government Code section 66020 will begin.  A written protest must be filed with the City Clerk pursuant to Government Code section 66020.  The protest procedures under section 66020 are additional to other procedures authorized or required under the San Diego Municipal Code. Please contact Facilities Financing at 619-533-3670 to request additional information."
	}],
	"ProjectDetails": [{
		"Jobs": [],
		"ProjectFees": [{
			"InvoiceDetialId": 3694253,
			"InvoiceId": 483626,
			"ProjectId": 319781,
			"JobId": null,
			"ApprovalId": null,
			"FeeTypeId": 4440,
			"FeeDescription": "Close Out -Engineering/Mapping",
			"FeeQuantityRequired": "1",
			"FeeTypeUnit": "Each",
			"FeeAmount": 262.0,
			"PreviousCreditInv": "",
			"FeeAuthority": "Municipal Code",
			"InvoiceCalcRules": [{
				"InvoiceDetailId": 3694253,
				"FeeThresholdId": 30076,
				"BaseQuantity": "0",
				"BaseIncrement": "1",
				"RuleRate": "262",
				"RuleRateAmt": 0.0,
				"FeeRule": "$262 each = $262.00",
				"PreSurchargeAmt": 262.0,
				"PreviousCreditAmt": "",
				"FeeSurchargeRules": []
			}]
		}, {
			"InvoiceDetialId": 3694254,
			"InvoiceId": 483626,
			"ProjectId": 319781,
			"JobId": null,
			"ApprovalId": null,
			"FeeTypeId": 1656,
			"FeeDescription": "Deposit Account",
			"FeeQuantityRequired": "18600",
			"FeeTypeUnit": "Dollars",
			"FeeAmount": 18600.0,
			"PreviousCreditInv": "",
			"FeeAuthority": "Municipal Code",
			"InvoiceCalcRules": [{
				"InvoiceDetailId": 3694254,
				"FeeThresholdId": 11213,
				"BaseQuantity": "0",
				"BaseIncrement": "0.01",
				"RuleRate": "0.01",
				"RuleRateAmt": 0.0,
				"FeeRule": "$.01 for each .01 \"Dollars\" over 0 or portion thereof = $18,600.00",
				"PreSurchargeAmt": 18600.0,
				"PreviousCreditAmt": "",
				"FeeSurchargeRules": []
			}]
		}, {
			"InvoiceDetialId": 3694255,
			"InvoiceId": 483626,
			"ProjectId": 319781,
			"JobId": null,
			"ApprovalId": null,
			"FeeTypeId": 2347,
			"FeeDescription": "General Plan Maintenance",
			"FeeQuantityRequired": "1",
			"FeeTypeUnit": "Each",
			"FeeAmount": 275.0,
			"PreviousCreditInv": "",
			"FeeAuthority": "Municipal Code",
			"InvoiceCalcRules": [{
				"InvoiceDetailId": 3694255,
				"FeeThresholdId": 29736,
				"BaseQuantity": "0",
				"BaseIncrement": "1",
				"RuleRate": "275",
				"RuleRateAmt": 0.0,
				"FeeRule": "$275 each = $275.00",
				"PreSurchargeAmt": 275.0,
				"PreviousCreditAmt": "",
				"FeeSurchargeRules": []
			}]
		}, {
			"InvoiceDetialId": 3694256,
			"InvoiceId": 483626,
			"ProjectId": 319781,
			"JobId": null,
			"ApprovalId": null,
			"FeeTypeId": 2385,
			"FeeDescription": "Mapping",
			"FeeQuantityRequired": "1",
			"FeeTypeUnit": "Each",
			"FeeAmount": 10.0,
			"PreviousCreditInv": "",
			"FeeAuthority": "Municipal Code",
			"InvoiceCalcRules": [{
				"InvoiceDetailId": 3694256,
				"FeeThresholdId": 1925,
				"BaseQuantity": "0",
				"BaseIncrement": "1",
				"RuleRate": "10",
				"RuleRateAmt": 0.0,
				"FeeRule": "$10 each = $10.00",
				"PreSurchargeAmt": 10.0,
				"PreviousCreditAmt": "",
				"FeeSurchargeRules": []
			}]
		}, {
			"InvoiceDetialId": 3694257,
			"InvoiceId": 483626,
			"ProjectId": 319781,
			"JobId": null,
			"ApprovalId": null,
			"FeeTypeId": 3784,
			"FeeDescription": "Records-Prelim(MF)",
			"FeeQuantityRequired": "1",
			"FeeTypeUnit": "Each",
			"FeeAmount": 60.0,
			"PreviousCreditInv": "",
			"FeeAuthority": "CC Ordinance R-305326 - Bulltn:",
			"InvoiceCalcRules": [{
				"InvoiceDetailId": 3694257,
				"FeeThresholdId": 27839,
				"BaseQuantity": "0",
				"BaseIncrement": "1",
				"RuleRate": "60",
				"RuleRateAmt": 0.0,
				"FeeRule": "$60 each = $60.00",
				"PreSurchargeAmt": 60.0,
				"PreviousCreditAmt": "",
				"FeeSurchargeRules": []
			}]
		}],
		"ProjectId": 319781,
		"Title": "Kaiser SD PI/Grading/Ded",
		"Scope": "KEARNY MESA Grading and public improvement, and public right of way dedication and Street Name for the new Kaiser Hospital. ",
		"ApplicationExpiration": "2014-04-11T17:00:00",
		"ApplicationExpired": true,
		"AdminHold": false,
		"DevelopmentId": 176644,
		"DevelopmentTitle": "Kaiser Hospital",
		"ApplicationDate": "2013-04-11T17:00:00",
		"AccountNum": "24003691",
		"JobOrderNum": null,
		"ProjectFeesTotal": 19207.0,
		"ProjectFeesSubTotal": 19207.0,
		"ProjectManagerId": 3329,
		"ProjectManager": {
			"ProjectManagerId": 3329,
			"Name": "Lynch, Pete",
			"PhoneNum": "",
			"EmailAddress": "dsdprojectinfo@sandiego.gov",
			"ActiveIndicator": true
		}
	}],
	"InvoiceRevenue": [{
		"InvoiceId": null,
		"FundNum": "100000",
		"GLNum": "416158",
		"CostObjectNum": "1619110011",
		"Fund": "GENERAL FUND",
		"GL": "DSD APPL FEE XFER-PLANNG",
		"Amount": 275.0
	}, {
		"InvoiceId": null,
		"FundNum": "200459",
		"GLNum": "422020",
		"CostObjectNum": "1611170012",
		"Fund": "DEPOSITS",
		"GL": "PLANNING SUBDIVISION DPST",
		"Amount": 18600.0
	}, {
		"InvoiceId": null,
		"FundNum": "700036",
		"GLNum": "416090",
		"CostObjectNum": "1611000013",
		"Fund": "PLANNING & DEVEL REV ENTR FUND",
		"GL": "STREET ADDRESS CHANGE FEE",
		"Amount": 10.0
	}, {
		"InvoiceId": null,
		"FundNum": "700036",
		"GLNum": "422043",
		"CostObjectNum": "1611120011",
		"Fund": "PLANNING & DEVEL REV ENTR FUND",
		"GL": "OTHER FEES",
		"Amount": 262.0
	}, {
		"InvoiceId": null,
		"FundNum": "700036",
		"GLNum": "423011",
		"CostObjectNum": "1611000015",
		"Fund": "PLANNING & DEVEL REV ENTR FUND",
		"GL": "Sale of Publications",
		"Amount": 60.0
	}]
}`

func TestDecodeInvoice(t *testing.T) {
	var err error

	buf := bytes.NewBufferString(invoiceTestData)

	invoice, err := opendsd.DecodeInvoice(buf)
	if err != nil {
		t.Errorf("error parsing: %v", err)
	}

	log.Printf("invoice %+v", invoice)
}

func TestInvoiceByID(t *testing.T) {
	var err error

	//	setup test muxer
	mux := http.NewServeMux()
	//	register our test route
	mux.HandleFunc("/invoice/483626",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, invoiceTestData)
		},
	)

	//	setup test server
	server := httptest.NewServer(mux)
	//	setup test client
	client := opendsd.NewClient()
	//	use the dynamically generated testing url
	client.APIRoot = server.URL

	invoice, err := client.InvoiceByID(483626)
	if err != nil {
		t.Errorf("InvoiceByID error: %v", err)
	}

	log.Printf("invoice %+v", invoice)
}
