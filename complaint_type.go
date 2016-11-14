package opendsd

type ComplaintType struct {
	ID          int
	Description string
}

var ComplaintTypes = map[int]ComplaintType{
	850: {
		ID:          850,
		Description: "Building-Boarded Structure",
	},
	139: {
		ID:          139,
		Description: "Building-Comm/Ind Construction-Accessory Structure",
	},
	130: {
		ID:          130,
		Description: "Building-Comm/Ind Construction---Active",
	},
	143: {
		ID:          143,
		Description: "Building-Comm/Ind Construction-Disabled Violation",
	},
	132: {
		ID:          132,
		Description: "Building-Comm/Ind Construction--Electrical",
	},
	152: {
		ID:          152,
		Description: "Building-Comm/Ind Construction-Electrical Deficiency",
	},
	141: {
		ID:          141,
		Description: "Building-Comm/Ind Construction-Exiting Problem",
	},
	137: {
		ID:          137,
		Description: "Building-Comm/Ind Construction-Fire Damage Repair",
	},
	142: {
		ID:          142,
		Description: "Building-Comm/Ind Construction-Firewall Rating",
	},
	134: {
		ID:          134,
		Description: "Building-Comm/Ind Construction--Mechanical",
	},
	154: {
		ID:          154,
		Description: "Building-Comm/Ind Construction-Mechanical Deficiency",
	},
	150: {
		ID:          150,
		Description: "Building-Comm/Ind Construction-Misc Storage",
	},
	136: {
		ID:          136,
		Description: "Building-Comm/Ind Construction-Mixed Use",
	},
	135: {
		ID:          135,
		Description: "Building-Comm/Ind Construction-Occupancy Change",
	},
	133: {
		ID:          133,
		Description: "Building-Comm/Ind Construction--Plumbing",
	},
	153: {
		ID:          153,
		Description: "Building-Comm/Ind Construction-Plumbing Deficiency",
	},
	131: {
		ID:          131,
		Description: "Building-Comm/Ind Construction--Structural",
	},
	151: {
		ID:          151,
		Description: "Building-Comm/Ind Construction-Structural Deficiency",
	},
	140: {
		ID:          140,
		Description: "Building-Comm/Ind Construction-Tenant Improvement",
	},
	191: {
		ID:          191,
		Description: "Building-Encroachment-Barricade",
	},
	192: {
		ID:          192,
		Description: "Building-Encroachment-Driveway Cuts",
	},
	194: {
		ID:          194,
		Description: "Building-Encroachment-Newsracks",
	},
	199: {
		ID:          199,
		Description: "Building-Encroachment-Other",
	},
	193: {
		ID:          193,
		Description: "Building-Encroachment-Public Right Of Way",
	},
	190: {
		ID:          190,
		Description: "Building-Encroachment-Sidewalk",
	},
	171: {
		ID:          171,
		Description: "Building-Hazard-Electrical",
	},
	173: {
		ID:          173,
		Description: "Building-Hazard-Mechanical",
	},
	175: {
		ID:          175,
		Description: "Building-Hazard-Pedestrian",
	},
	172: {
		ID:          172,
		Description: "Building-Hazard-Plumbing",
	},
	170: {
		ID:          170,
		Description: "Building-Hazard-Structural",
	},
	185: {
		ID:          185,
		Description: "Building-LandDev-Dumping",
	},
	184: {
		ID:          184,
		Description: "Building-LandDev-Grading",
	},
	183: {
		ID:          183,
		Description: "Building-LandDev-Landslide",
	},
	181: {
		ID:          181,
		Description: "Building-LandDev-Stormwater",
	},
	109: {
		ID:          109,
		Description: "Building-Res Construction-Accessory Structure ( > 120 sqft)",
	},
	100: {
		ID:          100,
		Description: "Building-Res Construction---Active",
	},
	105: {
		ID:          105,
		Description: "Building-Res Construction--Combo",
	},
	114: {
		ID:          114,
		Description: "Building-Res Construction-Demolition",
	},
	118: {
		ID:          118,
		Description: "Building-Res Construction-Disabled Violation",
	},
	102: {
		ID:          102,
		Description: "Building-Res Construction--Electrical",
	},
	116: {
		ID:          116,
		Description: "Building-Res Construction-Exiting Problem",
	},
	107: {
		ID:          107,
		Description: "Building-Res Construction-Fire Damage Repair",
	},
	111: {
		ID:          111,
		Description: "Building-Res Construction-Fireplace/Woodstove",
	},
	117: {
		ID:          117,
		Description: "Building-Res Construction-Firewall Rating",
	},
	104: {
		ID:          104,
		Description: "Building-Res Construction--Mechanical",
	},
	112: {
		ID:          112,
		Description: "Building-Res Construction-Occupancy Change",
	},
	103: {
		ID:          103,
		Description: "Building-Res Construction--Plumbing",
	},
	108: {
		ID:          108,
		Description: "Building-Res Construction-Pool or Spa",
	},
	106: {
		ID:          106,
		Description: "Building-Res Construction-Retaining Wall or Fence",
	},
	101: {
		ID:          101,
		Description: "Building-Res Construction--Structural",
	},
	1001: {
		ID:          1001,
		Description: "Graffiti-Commercial/Industrial",
	},
	1002: {
		ID:          1002,
		Description: "Grafitti - Work Crew",
	},
	202: {
		ID:          202,
		Description: "Housing-Electrical Deficiency",
	},
	206: {
		ID:          206,
		Description: "Housing-Fire Prevention Deficiency",
	},
	203: {
		ID:          203,
		Description: "Housing-Mechanical Deficiency",
	},
	299: {
		ID:          299,
		Description: "Housing-Miscellaneous",
	},
	212: {
		ID:          212,
		Description: "Housing-Mobile Home Park",
	},
	207: {
		ID:          207,
		Description: "Housing-No Fence w/ Pool/Spa",
	},
	204: {
		ID:          204,
		Description: "Housing-No Manager (>12 units)",
	},
	201: {
		ID:          201,
		Description: "Housing-Plumbing Deficiency",
	},
	200: {
		ID:          200,
		Description: "Housing-Structural Deficiency",
	},
	210: {
		ID:          210,
		Description: "Housing-Structure Not Intended for Habitation",
	},
	211: {
		ID:          211,
		Description: "Housing-Transient Occupation",
	},
	205: {
		ID:          205,
		Description: "Housing-Weeds and Debris",
	},
	900: {
		ID:          900,
		Description: "Noise-Barking Dogs",
	},
	907: {
		ID:          907,
		Description: "Noise-Construction",
	},
	901: {
		ID:          901,
		Description: "Noise-Farm Animals",
	},
	914: {
		ID:          914,
		Description: "Noise-Other Offensive Noise",
	},
	521: {
		ID:          521,
		Description: "Sign-Banner",
	},
	526: {
		ID:          526,
		Description: "Sign-Billboard-Alcohol",
	},
	525: {
		ID:          525,
		Description: "Sign-Billboard-General",
	},
	527: {
		ID:          527,
		Description: "Sign-Billboard-Maintenance",
	},
	503: {
		ID:          503,
		Description: "Sign-Disrepair",
	},
	518: {
		ID:          518,
		Description: "Sign-Flash/Auto",
	},
	505: {
		ID:          505,
		Description: "Sign-Portable",
	},
	523: {
		ID:          523,
		Description: "Sign-Real Estate/Open House",
	},
	500: {
		ID:          500,
		Description: "Sign-Sign in Public Right Of Way",
	},
	519: {
		ID:          519,
		Description: "Sign-Temporary Construction",
	},
	520: {
		ID:          520,
		Description: "Sign-Temporary Real Estate",
	},
	524: {
		ID:          524,
		Description: "Sign-Temporary Subdivision",
	},
	522: {
		ID:          522,
		Description: "Sign-Temporary Vehicle Ad",
	},
	501: {
		ID:          501,
		Description: "Sign-Unpermitted Sign",
	},
	540: {
		ID:          540,
		Description: "Sign-Window Sign",
	},
	1005: {
		ID:          1005,
		Description: "Zoning - Lighting",
	},
	1003: {
		ID:          1003,
		Description: "Zoning - Miscellaneous",
	},
	1004: {
		ID:          1004,
		Description: "Zoning - PVPO",
	},
	1006: {
		ID:          1006,
		Description: "Zoning - Sidewalk Caf‚ Annual",
	},
	350: {
		ID:          350,
		Description: "Zoning-Accessory Structure-Not Visible",
	},
	340: {
		ID:          340,
		Description: "Zoning-Accessory Structure-Visible",
	},
	302: {
		ID:          302,
		Description: "Zoning-Adult Entertainment",
	},
	345: {
		ID:          345,
		Description: "Zoning-Agriculture Zone",
	},
	357: {
		ID:          357,
		Description: "Zoning-Animals-Large (farm, ranch, zoological)",
	},
	356: {
		ID:          356,
		Description: "Zoning-Animals-Small (>25 fowl/rabbits, >100 pigeons)",
	},
	348: {
		ID:          348,
		Description: "Zoning-Auto Repair in Commercial Zone",
	},
	317: {
		ID:          317,
		Description: "Zoning-Auto Repair in Residential Zone",
	},
	339: {
		ID:          339,
		Description: "Zoning-Barbed Wire",
	},
	327: {
		ID:          327,
		Description: "Zoning-Bed & Breakfast",
	},
	323: {
		ID:          323,
		Description: "Zoning-Business in Residential Zone",
	},
	305: {
		ID:          305,
		Description: "Zoning-Business Not OK - Commercial",
	},
	320: {
		ID:          320,
		Description: "Zoning-Child Care",
	},
	330: {
		ID:          330,
		Description: "Zoning-Church Use",
	},
	304: {
		ID:          304,
		Description: "Zoning-Coastal Development",
	},
	358: {
		ID:          358,
		Description: "Zoning-Condominium Conversion",
	},
	300: {
		ID:          300,
		Description: "Zoning-Danger to the Public",
	},
	313: {
		ID:          313,
		Description: "Zoning-Discretionary Permit Required",
	},
	355: {
		ID:          355,
		Description: "Zoning-Fence-Disrepair",
	},
	347: {
		ID:          347,
		Description: "Zoning-Fence-Front/Street Side Yard",
	},
	346: {
		ID:          346,
		Description: "Zoning-Fence-Side/Rear",
	},
	1007: {
		ID:          1007,
		Description: "Zoning-Food Trucks",
	},
	314: {
		ID:          314,
		Description: "Zoning-Garage Conversion",
	},
	329: {
		ID:          329,
		Description: "Zoning-Garage Sales",
	},
	324: {
		ID:          324,
		Description: "Zoning-Garage Storage-MDU",
	},
	303: {
		ID:          303,
		Description: "Zoning-Grading/Environmentally Sensitive Lands",
	},
	308: {
		ID:          308,
		Description: "Zoning-Guest Quarters",
	},
	336: {
		ID:          336,
		Description: "Zoning-Historic Site",
	},
	332: {
		ID:          332,
		Description: "Zoning-Hotel (>6 Rooms)",
	},
	307: {
		ID:          307,
		Description: "Zoning-Illegal Housing Units",
	},
	316: {
		ID:          316,
		Description: "Zoning-Impound Storage Yard",
	},
	335: {
		ID:          335,
		Description: "Zoning-Inoperable Vehicles",
	},
	309: {
		ID:          309,
		Description: "Zoning-Junk Yards",
	},
	341: {
		ID:          341,
		Description: "Zoning-Landscaping",
	},
	301: {
		ID:          301,
		Description: "Zoning-Medical Marijuana",
	},
	318: {
		ID:          318,
		Description: "Zoning-Outdoor Display",
	},
	311: {
		ID:          311,
		Description: "Zoning-Oversize Structure - Floor Area Ratio",
	},
	333: {
		ID:          333,
		Description: "Zoning-Parking in Front Yard",
	},
	334: {
		ID:          334,
		Description: "Zoning-Parking Lots",
	},
	312: {
		ID:          312,
		Description: "Zoning-Previous Conforming Use",
	},
	326: {
		ID:          326,
		Description: "Zoning-Pushcarts",
	},
	338: {
		ID:          338,
		Description: "Zoning-Recycling Facility",
	},
	321: {
		ID:          321,
		Description: "Zoning-Residential Care-Rehab",
	},
	337: {
		ID:          337,
		Description: "Zoning-Residential High Occupancy",
	},
	352: {
		ID:          352,
		Description: "Zoning-Rooming House",
	},
	342: {
		ID:          342,
		Description: "Zoning-RV/Motorhomes-MDU",
	},
	354: {
		ID:          354,
		Description: "Zoning-RV/Motorhomes-SDU",
	},
	325: {
		ID:          325,
		Description: "Zoning-Sidewalk Caf‚",
	},
	319: {
		ID:          319,
		Description: "Zoning-Storage/Hoarding",
	},
	328: {
		ID:          328,
		Description: "Zoning-Swap Meet",
	},
	322: {
		ID:          322,
		Description: "Zoning-Work Furlough",
	},
}
