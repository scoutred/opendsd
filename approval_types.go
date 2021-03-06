package opendsd

type ApprovalType struct {
	ID            int
	ProcessCode   int
	Short         string
	Approval      string
	Category      string
	EffectiveFrom string
	EffectiveTo   string
	ApplDays      int
	Authority     string
}

//	map of approval types with the ID property is also used as the map key
var ApprovalTypes = map[int]ApprovalType{
	100: {
		ID:            100,
		ProcessCode:   1,
		Short:         "CE",
		Approval:      "Code Enforcement",
		Category:      "S",
		EffectiveFrom: "12/31/2019",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "City Council Section Code input here",
	},
	106: {
		ID:            106,
		ProcessCode:   1,
		Short:         "P",
		Approval:      "Plumbing Permit",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0402",
	},
	107: {
		ID:            107,
		ProcessCode:   1,
		Short:         "E",
		Approval:      "Electrical Permit",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0302",
	},
	108: {
		ID:            108,
		ProcessCode:   1,
		Short:         "M",
		Approval:      "Mechanical Permit",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0402",
	},
	113: {
		ID:            113,
		ProcessCode:   1,
		Short:         "DA",
		Approval:      "Development Agreement",
		Category:      "P",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 124.0101",
	},
	120: {
		ID:            120,
		ProcessCode:   1,
		Short:         "G",
		Approval:      "Grading Permit",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0602",
	},
	127: {
		ID:            127,
		ProcessCode:   1,
		Short:         "RD",
		Approval:      "Public Right of Way Dedication",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	129: {
		ID:            129,
		ProcessCode:   1,
		Short:         "CUP",
		Approval:      "Conditional Use Permit",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0303",
	},
	131: {
		ID:            131,
		ProcessCode:   1,
		Short:         "CDP",
		Approval:      "Coastal Development Permit",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0702",
	},
	132: {
		ID:            132,
		ProcessCode:   1,
		Short:         "SDP",
		Approval:      "Site Development Permit",
		Category:      "V",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0502",
	},
	133: {
		ID:            133,
		ProcessCode:   1,
		Short:         "B",
		Approval:      "Building Permit",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0202",
	},
	134: {
		ID:            134,
		ProcessCode:   1,
		Short:         "COC",
		Approval:      "Certificate of Compliance",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0210",
	},
	135: {
		ID:            135,
		ProcessCode:   1,
		Short:         "LLA",
		Approval:      "Lot-Line Adjustment",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0310",
	},
	136: {
		ID:            136,
		ProcessCode:   1,
		Short:         "TM",
		Approval:      "Tentative Map",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0410",
	},
	137: {
		ID:            137,
		ProcessCode:   1,
		Short:         "VTM",
		Approval:      "Vesting Tentative Map",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0410",
	},
	138: {
		ID:            138,
		ProcessCode:   1,
		Short:         "FM",
		Approval:      "Final Map",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0610",
	},
	139: {
		ID:            139,
		ProcessCode:   1,
		Short:         "EV",
		Approval:      "Easement Vacation",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.1010",
	},
	140: {
		ID:            140,
		ProcessCode:   1,
		Short:         "ED",
		Approval:      "Easement Dedication",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 62.",
	},
	141: {
		ID:            141,
		ProcessCode:   1,
		Short:         "SNC",
		Approval:      "Street Name Change",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 62.",
	},
	142: {
		ID:            142,
		ProcessCode:   1,
		Short:         "SA",
		Approval:      "Public Right of Way Vacation",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0910",
	},
	144: {
		ID:            144,
		ProcessCode:   1,
		Short:         "COR",
		Approval:      "Certificate of Correction",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0140",
	},
	145: {
		ID:            145,
		ProcessCode:   1,
		Short:         "AM",
		Approval:      "Amendment of Approved Map",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0140",
	},
	146: {
		ID:            146,
		ProcessCode:   1,
		Short:         "REZ",
		Approval:      "Rezone",
		Category:      "P",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 123.0102",
	},
	205: {
		ID:            205,
		ProcessCode:   1,
		Short:         "S",
		Approval:      "Sign Permit",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0802",
	},
	208: {
		ID:            208,
		ProcessCode:   1,
		Short:         "VAR",
		Approval:      "Variance",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0802",
	},
	223: {
		ID:            223,
		ProcessCode:   1,
		Short:         "ERA",
		Approval:      "Encroachment Agreement",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 62.",
	},
	227: {
		ID:            227,
		ProcessCode:   1,
		Short:         "PM",
		Approval:      "Parcel Map",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0510",
	},
	230: {
		ID:            230,
		ProcessCode:   1,
		Short:         "SCR",
		Approval:      "Substantial Conformance Review",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	240: {
		ID:            240,
		ProcessCode:   1,
		Short:         "ZUC",
		Approval:      "Zoning Use Certificate",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 123.0302",
	},
	244: {
		ID:            244,
		ProcessCode:   1,
		Short:         "MW",
		Approval:      "Map Waiver",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0120",
	},
	245: {
		ID:            245,
		ProcessCode:   1,
		Short:         "RM",
		Approval:      "Reversion to Acreage",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0810",
	},
	247: {
		ID:            247,
		ProcessCode:   1,
		Short:         "PDP",
		Approval:      "Planned Development Permit",
		Category:      "V",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0603",
	},
	248: {
		ID:            248,
		ProcessCode:   1,
		Short:         "HRD",
		Approval:      "Historic Resource Designation",
		Category:      "V",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 123.0202",
	},
	249: {
		ID:            249,
		ProcessCode:   1,
		Short:         "MP",
		Approval:      "Merger of Parcels",
		Category:      "S",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0710",
	},
	250: {
		ID:            250,
		ProcessCode:   1,
		Short:         "NDP",
		Approval:      "Neighborhood Developmnt Permit",
		Category:      "V",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0402",
	},
	251: {
		ID:            251,
		ProcessCode:   1,
		Short:         "NUP",
		Approval:      "Neighborhood Use Permit",
		Category:      "V",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0203",
	},
	252: {
		ID:            252,
		ProcessCode:   1,
		Short:         "R",
		Approval:      "Right Of Way Permit",
		Category:      "S",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0702",
	},
	254: {
		ID:            254,
		ProcessCode:   1,
		Short:         "C",
		Approval:      "Combination Building Permit",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Ca. Building Code",
	},
	255: {
		ID:            255,
		ProcessCode:   1,
		Short:         "SIA",
		Approval:      "Subdivision Improvement Agrmnt",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 144.0430",
	},
	256: {
		ID:            256,
		ProcessCode:   1,
		Short:         "DIA",
		Approval:      "Deferred Improvement Agrmnt",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	257: {
		ID:            257,
		ProcessCode:   1,
		Short:         "CRD",
		Approval:      "Cost Recovery District",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	258: {
		ID:            258,
		ProcessCode:   1,
		Short:         "LUP",
		Approval:      "Land Use Plan",
		Category:      "P",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 122.0102",
	},
	259: {
		ID:            259,
		ProcessCode:   1,
		Short:         "FP",
		Approval:      "Financing Plan",
		Category:      "P",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	260: {
		ID:            260,
		ProcessCode:   1,
		Short:         "SP",
		Approval:      "Special Permit",
		Category:      "B",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Stephen Haase",
	},
	261: {
		ID:            261,
		ProcessCode:   1,
		Short:         "LMA",
		Approval:      "Landscape Maintenance Agrmnt",
		Category:      "S",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	262: {
		ID:            262,
		ProcessCode:   1,
		Short:         "CRA",
		Approval:      "Cost Reimbursement Agreement",
		Category:      "V",
		EffectiveFrom: "03/16/1998",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	263: {
		ID:            263,
		ProcessCode:   1,
		Short:         "CCA",
		Approval:      "City Council Approval",
		Category:      "V",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      720,
		Authority:     "M.C.",
	},
	264: {
		ID:            264,
		ProcessCode:   1,
		Short:         "CMA",
		Approval:      "City Manager Approval",
		Category:      "S",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      720,
		Authority:     "M.C.",
	},
	265: {
		ID:            265,
		ProcessCode:   1,
		Short:         "ADD",
		Approval:      "Street Address Change",
		Category:      "B",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 62.",
	},
	267: {
		ID:            267,
		ProcessCode:   1,
		Short:         "A",
		Approval:      "Agreement",
		Category:      "S",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	268: {
		ID:            268,
		ProcessCode:   1,
		Short:         "D",
		Approval:      "Demolition Permit",
		Category:      "B",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0502",
	},
	269: {
		ID:            269,
		ProcessCode:   1,
		Short:         "TUP",
		Approval:      "Temporary Use Permit",
		Category:      "S",
		EffectiveFrom: "01/01/2000",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 123.0401",
	},
	271: {
		ID:            271,
		ProcessCode:   1,
		Short:         "CCB",
		Approval:      "Construction Change - Building",
		Category:      "B",
		EffectiveFrom: "03/20/2003",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "MC",
	},
	272: {
		ID:            272,
		ProcessCode:   1,
		Short:         "RTC",
		Approval:      "ROW Permit-Traffic Control",
		Category:      "B",
		EffectiveFrom: "06/29/2004",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0702",
	},
	273: {
		ID:            273,
		ProcessCode:   1,
		Short:         "CCE",
		Approval:      "Construction Change - Eng.",
		Category:      "S",
		EffectiveFrom: "03/21/2004",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C.",
	},
	275: {
		ID:            275,
		ProcessCode:   1,
		Short:         "GR",
		Approval:      "Grading + Right of Way Permit",
		Category:      "S",
		EffectiveFrom: "11/08/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0602",
	},
	277: {
		ID:            277,
		ProcessCode:   1,
		Short:         "SWA",
		Approval:      "Storm Water Maintenance Agrmnt",
		Category:      "S",
		EffectiveFrom: "11/08/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 43.0310",
	},
	278: {
		ID:            278,
		ProcessCode:   1,
		Short:         "DDR",
		Approval:      "Deferred Document Review",
		Category:      "B",
		EffectiveFrom: "04/27/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Strohminger",
	},
	279: {
		ID:            279,
		ProcessCode:   1,
		Short:         "N",
		Approval:      "Newsrack Permit",
		Category:      "S",
		EffectiveFrom: "04/27/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 62.1001",
	},
	280: {
		ID:            280,
		ProcessCode:   1,
		Short:         "T",
		Approval:      "Transportation Permit",
		Category:      "B",
		EffectiveFrom: "04/27/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 85.21",
	},
	281: {
		ID:            281,
		ProcessCode:   1,
		Short:         "ZCA",
		Approval:      "Zone Code Amendment",
		Category:      "S",
		EffectiveFrom: "04/27/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      720,
		Authority:     "M.C. 103.0106",
	},
	282: {
		ID:            282,
		ProcessCode:   1,
		Short:         "CNP",
		Approval:      "Construction Noise Permit",
		Category:      "S",
		EffectiveFrom: "07/29/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 59.5.0404",
	},
	283: {
		ID:            283,
		ProcessCode:   1,
		Short:         "DAB",
		Approval:      "Damage Appraisal - Bldg",
		Category:      "B",
		EffectiveFrom: "07/29/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Afsaneh Ahmadi",
	},
	284: {
		ID:            284,
		ProcessCode:   1,
		Short:         "DAC",
		Approval:      "Damage Appraisal - Combo",
		Category:      "B",
		EffectiveFrom: "07/29/2007",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Afsaneh Ahmadi",
	},
	288: {
		ID:            288,
		ProcessCode:   1,
		Short:         "RHO",
		Approval:      "Res. High Occupancy Permit",
		Category:      "B",
		EffectiveFrom: "04/20/2008",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 123.0502",
	},
	289: {
		ID:            289,
		ProcessCode:   1,
		Short:         "SC5",
		Approval:      "Sprinkler Certification-5 Year",
		Category:      "B",
		EffectiveFrom: "04/20/2008",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "CCR 904.4",
	},
	290: {
		ID:            290,
		ProcessCode:   1,
		Short:         "DFA",
		Approval:      "Deferred Fire Approval",
		Category:      "B",
		EffectiveFrom: "02/15/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Afsaneh Ahmadi",
	},
	291: {
		ID:            291,
		ProcessCode:   1,
		Short:         "RW",
		Approval:      "Reclaimed Water Permit",
		Category:      "B",
		EffectiveFrom: "02/15/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 64.0807(e)",
	},
	292: {
		ID:            292,
		ProcessCode:   1,
		Short:         "CCF",
		Approval:      "Construction Change - Fire",
		Category:      "B",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Afsaneh Ahmadi",
	},
	293: {
		ID:            293,
		ProcessCode:   1,
		Short:         "EPV",
		Approval:      "Electrical-Photovoltaic",
		Category:      "B",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0302",
	},
	294: {
		ID:            294,
		ProcessCode:   1,
		Short:         "EAE",
		Approval:      "Emergency Authorization-ESL",
		Category:      "S",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 143.0126",
	},
	295: {
		ID:            295,
		ProcessCode:   1,
		Short:         "EAH",
		Approval:      "Emergency Authorization-Hist",
		Category:      "S",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 143.0214",
	},
	296: {
		ID:            296,
		ProcessCode:   1,
		Short:         "ECD",
		Approval:      "Emergency Coastal Dev Permit",
		Category:      "S",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0718",
	},
	297: {
		ID:            297,
		ProcessCode:   1,
		Short:         "EOT",
		Approval:      "Extension of Time - Dev. Pmt",
		Category:      "V",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 126.0111",
	},
	298: {
		ID:            298,
		ProcessCode:   1,
		Short:         "EOT",
		Approval:      "Extension of Time - TM",
		Category:      "V",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0461",
	},
	299: {
		ID:            299,
		ProcessCode:   1,
		Short:         "NOT",
		Approval:      "Notice of Termination",
		Category:      "S",
		EffectiveFrom: "12/20/2009",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "Jeff Strohminger",
	},
	301: {
		ID:            301,
		ProcessCode:   1,
		Short:         "R",
		Approval:      "Right Of Way Permit-Const Plan",
		Category:      "S",
		EffectiveFrom: "03/13/2010",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0702",
	},
	303: {
		ID:            303,
		ProcessCode:   1,
		Short:         "EU",
		Approval:      "Elec.-UG Program-City only",
		Category:      "B",
		EffectiveFrom: "09/11/2011",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0302",
	},
	304: {
		ID:            304,
		ProcessCode:   1,
		Short:         "EU",
		Approval:      "Elec.-UG Program-Cust Upgrades",
		Category:      "B",
		EffectiveFrom: "09/11/2011",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 129.0302",
	},
	305: {
		ID:            305,
		ProcessCode:   1,
		Short:         "LTA",
		Approval:      "Lot Tie Agreement",
		Category:      "B",
		EffectiveFrom: "09/11/2011",
		EffectiveTo:   "12/31/2019",
		ApplDays:      360,
		Authority:     "M.C. 125.0760",
	},
	306: {
		ID:            306,
		ProcessCode:   1,
		Short:         "FA",
		Approval:      "Fire Permit - Alarm",
		Category:      "B",
		EffectiveFrom: "09/27/2012",
		EffectiveTo:   "12/31/2019",
		ApplDays:      720,
		Authority:     "M.C. 129.0102",
	},
	307: {
		ID:            307,
		ProcessCode:   1,
		Short:         "FKH",
		Approval:      "Fire Permit - Kitchen Hood",
		Category:      "B",
		EffectiveFrom: "09/27/2012",
		EffectiveTo:   "12/31/2019",
		ApplDays:      720,
		Authority:     "M.C. 129.0102",
	},
	308: {
		ID:            308,
		ProcessCode:   1,
		Short:         "FSP",
		Approval:      "Fire Permit - Suppression",
		Category:      "B",
		EffectiveFrom: "09/27/2012",
		EffectiveTo:   "12/31/2019",
		ApplDays:      720,
		Authority:     "M.C. 129.0102",
	},
}
