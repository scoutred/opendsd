package opendsd_test

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"testing"

	"github.com/scoutred/opendsd"
)

var remedyTypesTestData = `
10,1,Case Opened,5,102,Inspection Prep
10,1,Case Opened,20,82,Converted From M20 Legacy Application
10,1,Case Opened,30,103,Inspection
2010,3,Administrative Citation Warning,70,7,Issue
2010,3,Administrative Citation Warning,80,8,Compliance Period
2010,3,Administrative Citation Warning,90,9,Inspection
2020,4,Administrative Citation,110,11,Issue
2020,4,Administrative Citation,120,12,Compliance Period
2020,4,Administrative Citation,130,13,Inspection
2030,5,Administrative Citation Appeal Hearing,140,14,Packet Preparation
2030,5,Administrative Citation Appeal Hearing,150,15,Hearing Preparation
2030,5,Administrative Citation Appeal Hearing,160,16,Hearing
2030,5,Administrative Citation Appeal Hearing,170,17,Inspection
4050,12,Civil Penalty Notice And Order,470,98,Issued
4050,12,Civil Penalty Notice And Order,480,48,Compliance Period
4050,12,Civil Penalty Notice And Order,490,49,Inspection
4060,13,Civil Penalty Hearing,500,50,Packet Preparation
4060,13,Civil Penalty Hearing,510,51,Hearing Preparation
4060,13,Civil Penalty Hearing,520,52,Hearing
4060,13,Civil Penalty Hearing,530,53,Compliance Period
4060,13,Civil Penalty Hearing,540,54,Inspection
4070,14,Civil Penalty Hearing (2nd),550,55,Packet Preparation
4070,14,Civil Penalty Hearing (2nd),560,56,Hearing Preparation
4070,14,Civil Penalty Hearing (2nd),570,57,Hearing
4070,14,Civil Penalty Hearing (2nd),580,58,Compliance Period
4070,14,Civil Penalty Hearing (2nd),590,59,Inspection
8000,9,Administrative Abatement,300,99,Notice Issued
8000,9,Administrative Abatement,305,83,Record Notice
8000,9,Administrative Abatement,310,31,Appeal Period
8000,9,Administrative Abatement,320,32,Inspection
8000,9,Administrative Abatement,330,33,Packet Preparation
8000,9,Administrative Abatement,340,34,Hearing Preparation
8000,9,Administrative Abatement,350,35,Hearing
8000,9,Administrative Abatement,360,36,Abatement
8000,9,Administrative Abatement,370,37,Cost Confirmation Packet Prep
8000,9,Administrative Abatement,375,84,Cost Confirmation Hearing Preparation
8000,9,Administrative Abatement,380,38,Cost Confirmation Hearing
8000,10,Summary Abatement,390,39,Abatement
8000,10,Summary Abatement,400,40,Cost Confirmation Packet Prep
8000,10,Summary Abatement,410,41,Cost Confirmation Hearing Preparation
8000,10,Summary Abatement,420,42,Cost Confirmation Hearing
8100,6,Notice of Violation,180,18,Drafted
8100,6,Notice of Violation,190,19,Prepare Final
8100,6,Notice of Violation,200,20,Sent
8800,24,Vacant Property Monitoring,1000,94,Inspection
8810,7,Intent to Record Notice of Violation,210,21,Preparation
8810,7,Intent to Record Notice of Violation,220,22,Notice Appeal Period
8810,7,Intent to Record Notice of Violation,240,24,Record NOV
8810,7,Intent to Record Notice of Violation,870,87,NOV Recorded
8820,8,Intent to Record Notice of Violation Hearing,250,25,Packet Preparation
8820,8,Intent to Record Notice of Violation Hearing,260,26,Hearing Preparation
8820,8,Intent to Record Notice of Violation Hearing,270,27,Hearing
8820,8,Intent to Record Notice of Violation Hearing,280,28,Record NOV
8820,8,Intent to Record Notice of Violation Hearing,285,88,NOV Recorded
9000,15,City Attorney Referral,600,60,Packet Preparation
9000,15,City Attorney Referral,610,61,Final CA Case File
9000,15,City Attorney Referral,615,90,Case File Prepared
9000,15,City Attorney Referral,1100,92,Referred
9000,15,City Attorney Referral,1200,93,Inspection
9500,23,DSD Construction Permit,100,91,Compliance Period
10000,22,Administrative Fine/Costs,850,85,Invoice Request
10000,22,Administrative Fine/Costs,860,86,Invoiced
10100,11,Notice of Pending Enforcement Action,430,43,Preparation
10100,11,Notice of Pending Enforcement Action,440,44,Notice
10100,11,Notice of Pending Enforcement Action,450,45,Record
10200,21,Notice of Pending Enforcement Action Release (Administrative Action),780,78,Demand Received
10200,21,Notice of Pending Enforcement Action Release (Administrative Action),790,79,Inspection
10200,21,Notice of Pending Enforcement Action Release (Administrative Action),800,80,Preparation
10200,21,Notice of Pending Enforcement Action Release (Administrative Action),810,81,Record
10300,20,Notice of Compliance Recordation (Administrative Action),740,74,Demand Received
10300,20,Notice of Compliance Recordation (Administrative Action),750,75,Inspection
10300,20,Notice of Compliance Recordation (Administrative Action),760,76,Preparation
10300,20,Notice of Compliance Recordation (Administrative Action),770,77,Record
10400,16,Code Enforcement Lien (Administrative Action),620,62,Preparation
10400,16,Code Enforcement Lien (Administrative Action),630,63,Notice
10400,16,Code Enforcement Lien (Administrative Action),640,64,Record
10500,17,Code Enforcement Lien Release (Administrative Action),660,66,Demand Received
10500,17,Code Enforcement Lien Release (Administrative Action),670,67,Record
10600,18,Nuisance Abatement Lien (Administrative Action),680,68,Preparation
10600,18,Nuisance Abatement Lien (Administrative Action),690,69,Notice
10600,18,Nuisance Abatement Lien (Administrative Action),700,70,Record
10700,19,Nuisance Abatement Lien Release (Administrative Action),195,100,Compliance Period
10700,19,Nuisance Abatement Lien Release (Administrative Action),200,101,Inspection
10700,19,Nuisance Abatement Lien Release (Administrative Action),710,71,Demand Received
10700,19,Nuisance Abatement Lien Release (Administrative Action),720,72,Preparation
10700,19,Nuisance Abatement Lien Release (Administrative Action),730,73,Record
20000,25,Grafitti,100,95,Refer to Another Agency/Dept
20000,25,Grafitti,150,97,Refer to Urban Corp
20000,25,Grafitti,200,96,Abatement`

func TestDecodeRemedyTypes(t *testing.T) {

	buf := bytes.NewBufferString(remedyTypesTestData)

	r := csv.NewReader(buf)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
		}

		rec := opendsd.RemedyType{
			Sequence:       record[0],
			ID:             record[1],
			Remedy:         record[2],
			ActionSequence: record[3],
			ActionTypeID:   record[4],
			Action:         record[5],
		}

		fmt.Printf("%+v\n", rec)
	}
}
