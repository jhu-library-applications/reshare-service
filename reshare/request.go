package reshare

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const apiBase = "https://borrowdirect.reshare.indexdata.com/api/v1/search"
const recordBaseURL = "https://borrowdirect.reshare.indexdata.com/Record/"

type Request struct {
	Isn        string
	Title      string
	Author     string
	requestURL string
}

func (r Request) GenTitleAuthorRequestUrl() string {
	fields := []string{"id", "authors", "lendingStatus"}

	requestURL, err := url.Parse(apiBase)
	if err != nil {
		log.Fatal(err)
	} else {
		q := requestURL.Query()
		q.Add("join", "AND")
		q.Add("lookfor0[]", r.Title)
		q.Add("type0[]", "Title")
		q.Add("lookfor1[]", r.Author)
		q.Add("type1[]", "Authors")

		for _, f := range fields {
			q.Add("field[]", f)
		}
		requestURL.RawQuery = q.Encode()
	}
	log.Println(requestURL.String())
	return requestURL.String()
}

func (r Request) GenIsnRequestUrl() string {
	fields := []string{"id", "authors", "lendingStatus"}

	requestURL, err := url.Parse(apiBase)
	if err != nil {
		log.Fatal(err)
	} else {
		q := requestURL.Query()
		q.Add("type", "ISN")
		q.Add("lookfor", r.Isn)
		for _, f := range fields {
			q.Add("field[]", f)
		}
		requestURL.RawQuery = q.Encode()
	}
	return requestURL.String()
}

func (r Request) ItemRequest() string {
	var records Records

	// Make an HTTP Request to ReShare's VuFind API.

	// Generate a request URL.
	if r.Isn == "" || len(r.Isn) > 1 {
		r.requestURL = r.GenTitleAuthorRequestUrl()
	} else {
		r.requestURL = r.GenIsnRequestUrl()
	}

	log.Println(r.requestURL)
	resp, err := http.Get(r.requestURL)

	if err != nil {
		log.Fatal(err)
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		// Unmarshal the JSON into Records
		json.Unmarshal(body, &records)

		// Iterate over the records that are returned
		for i := range records.InnerRecord {
			// Iterate over the lending statuses.
			for _, l := range records.InnerRecord[i].LendingStatus {
				if l == "LOANABLE" {
					var loanableRecordUrl string = fmt.Sprintf("%s%s", recordBaseURL, records.InnerRecord[i].Id)
					log.Printf("Found loanable record: %s \n", loanableRecordUrl)
					return loanableRecordUrl
				}
			}
		}
	}

	// If you tried with ISN and got nothing try again with title/author
	if len(r.Isn) < 1 && len(r.Title) < 1 {
		r.Isn = ""
		r.ItemRequest()
	} else {
		return "This item is not available for loan via BorrowDirect."
	}

	return "This item is not available for loan via BorrowDirect."
}
