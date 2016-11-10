package opendsd_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/scoutred/opendsd"
)

var approvalTestData = `
{
	"Header": [{
		"Jurisdiction": "City of San Diego",
		"Agency": "Development Services Department",
		"AgencyAddress": "1222 First Ave, San Diego, CA 92101",
		"AgencyWebsite": "www.sandiego.gov",
		"ExtractSystem": "Extract System",
		"ExtractDate": "10/30/2016 6:00:28 PM",
		"ExtractQuery": "Project Status XML"
	}],
	"ApprovalId": 1117208,
	"Project": {
		"Customers": null,
		"ReviewCycles": null,
		"Jobs": null,
		"Fees": null,
		"Invoices": null,
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
		"Header": null,
		"ProjectManagerId": 3329,
		"ProjectManager": {
			"ProjectManagerId": 3329,
			"Name": "Lynch, Pete",
			"PhoneNum": "",
			"EmailAddress": "dsdprojectinfo@sandiego.gov",
			"ActiveIndicator": true
		}
	},
	"Job": {
		"SignOffs": [],
		"ApprovalInfo": [],
		"Approvals": [],
		"ProjectId": 319781,
		"JobId": 635029,
		"Description": "37444-B Street Dedication on por lot 1of map 4674 ",
		"APN": "369-121-14-00",
		"StreetAddress": "5201 RUFFIN RD ",
		"MapReference": "XXXX-XX",
		"SortableStreetAddress": "RUFFIN RD 0000005201 ",
		"Latitude": 32.830113,
		"Longitude": -117.126156,
		"NAD83Northing": null,
		"NAD83Easting": null,
		"JobFeesSubTotal": null
	},
	"BCCodes": [{
		"JobId": 833119,
		"BCCodeId": 103,
		"BCCode": "Three or Four Family Apt"
	}],
	"Approval": {
		"JobId": 635029,
		"ApprovalId": 1117208,
		"Type": "Public Right of Way Dedication",
		"Status": "Issued",
		"Scope": "Street dedication over portion of lot 1 of map 4674. Recorded April 4, 2014. Document No. 2014-0133239 and Document No. 2014-0133247.",
		"Depiction": "2",
		"IssuedBy": "Lynch, Pete",
		"IssueDate": "2016-10-20T15:39:58",
		"FirstInspectionDate": null,
		"CompleteCancelDate": null,
		"PermitHolder": "Steven Doshay",
		"NetChangeDU": "",
		"Valuation": "",
		"SquareFootage": null
	},
	"Inspections": [{
		"ApprovalId": 1501285,
		"InspectionId": 2682138,
		"InspectionTier": "2",
		"InspectionType": "Final(2)",
		"InspectionDiscipline": "Mechanical",
		"InspectionStatus": "Available",
		"InspectionStatusSeq": "20",
		"InspectorId": null,
		"InspectorName": "",
		"InspectorPhone": "",
		"InspectorActive": "",
		"SchedulingInstruction": "Call (858) 581-7111 to schedule this inspection",
		"CustRequestable": "Y",
		"ScheduledDt": null,
		"PerformedDt": null,
		"InspectionResult": "",
		"InspectorEmail": "",
		"InspectionDetails": []
	}],
	"InspectionIssues": [{
		"InspectionIssueId": 196533,
		"ApprovalId": 1501283,
		"InspectionTier": "3",
		"Issue": "4,000 psi Concrete dtl E/S8",
		"Class": "Special Inspection: Concrete",
		"Visibility": "Regular",
		"CreatedBy": "Anderson, John",
		"CreatedDt": "2015-07-26T13:17:40",
		"ClearedBy": "",
		"ClearedDt": null,
		"ClearedNote": ""
	}],
	"ApprovalFees": [{
		"FeeTypeId": 3785,
		"ProjectId": 401996,
		"FeeType": "Recordation Fee-County",
		"FeeCategory": "Inspection Fees",
		"FeeTypeUnit": "Pages",
		"FeeQuantityRequired": "3",
		"FeeQuantityPaid": "3",
		"InvoiceId": 729995,
		"InvoiceStatus": "Paid on "
	}],
	"Exceptions": [{
		"ApprovalId": 1472397,
		"ApprovalStatusSeq": null,
		"Exception": "Inspection: You have 8 inspections remaining in the plan.",
		"ApprovalStatus": ""
	}],
	"DependentPackages": [],
	"DependantApprovals": [{
		"ApprovalId": 1501283,
		"InspectionTier": "3",
		"DepApprApprovalId": 1550442,
		"DepApprType": "Right Of Way Permit-Const Plan",
		"DepApprStatus": "Created",
		"StatusReq": "Completed",
		"RequirementMet": "N",
		"AddedBy": "Fridman, Bella",
		"AddedDt": "2016-04-21T14:21:20"
	}]
}`

func TestDecodApproval(t *testing.T) {
	var err error

	buf := bytes.NewBufferString(approvalTestData)

	approval, err := opendsd.DecodeApproval(buf)
	if err != nil {
		t.Errorf("error parsing: %v", err)
	}

	log.Printf("approval %+v", approval)
}

func TestApprovalByID(t *testing.T) {
	var err error

	approval, err := opendsd.ApprovalByID(1117208)
	if err != nil {
		t.Errorf("ApprovalByID error: %v", err)
	}

	log.Printf("approval %+v", approval)
}
