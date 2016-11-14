# opendsd

[![GoDoc](https://godoc.org/github.com/scoutred/opendsd?status.svg)](http://godoc.org/github.com/scoutred/opendsd)

Package for interfacing with the City of San Diego's [Open DSD API](https://www.sandiego.gov/development-services/opendsd/developers) and parsing permit activity reports.

Supports the following:

- Permit activity reports (XML) - weekly permit feeds published online [here](https://www.sandiego.gov/development-services/opendsd/permitactivity)
- Project API (JSON)
- Invoice API (JSON)
- Approval API (JSON)
- Code Enforcement API (JSON)
- Code Enforcement cases (XML) - published periodically [here](https://www.sandiego.gov/development-services/opendsd/codenforcement)

TODO 
- [ ] Support [complaint type codes](https://www.sandiego.gov/sites/default/files/legacy/development-services/opendsd/csv/complaintypes.csv)
- [ ] Support [code enforcement remedies](https://www.sandiego.gov/sites/default/files/legacy/development-services/opendsd/csv/codenforcementremedies.csv)

## Usage

```go

pacakge main

import (
	"log"

	"github.com/scoutred/opendsd"
)

func main(){
	//	instantiate a new API client
	client := opendsd.NewClient()

	//	fetch a project by id
	project, err := client.ProjectByID(319781)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("project %v", project)
}

```

## Notes

- The Invoice, Approval and Project data structures were cobbled together based on combining several API results from each endpoint. I was not able to find a comprehensive version of the data structure so some fields might be missing. I will happily accept pull requests to help flush out the entire data structure.
- The timestamps from the API come in multiple versions. This package parses the various time formats to time.Time structs.

## License

The MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.