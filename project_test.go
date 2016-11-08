package opendsd_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/scoutred/opendsd"
)

var projectTestData = `
{
	"Customers": [{
		"ProjectId": 512321,
		"CustomerId": 151887,
		"Role": "Owner",
		"FirmName": "",
		"Name": "Robertson Whittemore"
	}],
	"ReviewCycles": [{
		"ReviewCycleId": 1673991,
		"CycleNum": 3,
		"Method": "City Engineer-Civil",
		"Status": "Opened",
		"StatusSequence": 100,
		"SubmitDate": null,
		"DueDate": null,
		"CloseDate": null,
		"Performance": null,
		"Reviews": [{
			"ReviewCycleId": 1673991,
			"ReviewId": 1421425,
			"Discipline": "City Engineer-Civil",
			"Status": "Assignment Pending",
			"DueDate": null,
			"CompletedDate": null,
			"Performance": "Completed On Time",
			"Name": "",
			"Phone": "()-",
			"Email": "",
			"IsActive": true
		}]
	}],
	"Jobs": [{
		"SignOffs": [{
			"DisciplineId": 104,
			"DisciplineDescription": "LDR-Environmental",
			"SignedDate": "Not Signed-Off"
		}, {
			"DisciplineId": 284,
			"DisciplineDescription": "BDR-Landscaping",
			"SignedDate": "09/21/2016"
		}],
		"ApprovalInfo": [],
		"Approvals": [{
			"JobId": 986815,
			"ApprovalId": 1794629,
			"Type": "Grading Permit",
			"Status": "Created",
			"Scope": "test boring on existing residential lot.",
			"Depiction": "1",
			"IssuedBy": "",
			"IssueDate": null,
			"FirstInspectionDate": null,
			"CompleteCancelDate": null,
			"PermitHolder": " ",
			"NetChangeDU": "",
			"Valuation": "",
			"SquareFootage": null
		}],
		"ProjectId": 512321,
		"JobId": 986815,
		"Description": "8470 EL PASEO GRANDE ",
		"APN": "346-050-01-00",
		"StreetAddress": "8470 EL PASEO GRANDE ",
		"MapReference": "XXXX-XX",
		"SortableStreetAddress": "EL PASEO GRANDE 0000008470 ",
		"Latitude": 32.860901,
		"Longitude": -117.255248,
		"NAD83Northing": null,
		"NAD83Easting": null,
		"JobFeesSubTotal": null
	}],
	"Fees": [{
		"FeeId": 4440,
		"Description": "Close Out -Engineering/Mapping",
		"Category": "Submittal Fees",
		"Unit": "Each",
		"QuantityRequired": 1,
		"QuantityPaid": 1,
		"InvoiceId": 724015,
		"InvoiceStatus": "Paid on ",
		"ProjectId": 512321
	}],
	"Invoices": [{
		"InvoiceId": 724015,
		"InvoiceIssueDate": "2016-09-15T13:47:22",
		"InvoiceStatus": "Paid",
		"ProjectId": 512321
	}],
	"ProjectId": 512321,
	"Title": "EPG3 Residence Test Boring",
	"Scope": "LA JOLLA: Grading Permit for test boring on existing residential lot.",
	"ApplicationExpiration": "2018-09-16T17:00:00",
	"ApplicationExpired": false,
	"AdminHold": false,
	"DevelopmentId": 322701,
	"DevelopmentTitle": "Devel Num 322701",
	"ApplicationDate": "2016-09-16T16:10:00",
	"AccountNum": "24006988",
	"JobOrderNum": null,
	"Header": [{
		"Jurisdiction": "City of San Diego",
		"Agency": "Development Services Department",
		"AgencyAddress": "1222 First Ave, San Diego, CA 92101",
		"AgencyWebsite": "www.sandiego.gov",
		"ExtractSystem": "Extract System",
		"ExtractDate": "9/27/2016 7:23:28 PM",
		"ExtractQuery": "Project Status XML"
	}],
	"ProjectManagerId": 3152,
	"ProjectManager": {
		"ProjectManagerId": 3152,
		"Name": "Gonzalez, Dolores",
		"PhoneNum": "",
		"EmailAddress": "dsdprojectinfo@sandiego.gov",
		"ActiveIndicator": true
	}
}`

func TestDecodeProject(t *testing.T) {
	var err error

	buf := bytes.NewBufferString(projectTestData)

	project, err := opendsd.DecodeProject(buf)
	if err != nil {
		t.Errorf("error parsing: %v", err)
	}

	log.Printf("project %+v", project)
}
