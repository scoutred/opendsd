package opendsd_test

approvalTestData = `
{
  "Header": [
    {
      "Jurisdiction": "City of San Diego",
      "Agency": "Development Services Department",
      "AgencyAddress": "1222 First Ave, San Diego, CA 92101",
      "AgencyWebsite": "www.sandiego.gov",
      "ExtractSystem": "Extract System",
      "ExtractDate": "10/30/2016 6:00:28 PM",
      "ExtractQuery": "Project Status XML"
    }
  ],
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
  "BCCodes": [],
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
  "Inspections": [],
  "InspectionIssues": [],
  "ApprovalFees": [],
  "Exceptions": [],
  "DependentPackages": [],
  "DependantApprovals": []
}
`