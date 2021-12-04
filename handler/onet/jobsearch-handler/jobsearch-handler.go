package jobsearch_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"io/ioutil"
	"net/http"
)

const OnetService = "search?keyword="
// returns top 20 most relevant jobs based on the keyword given
func GetJobResults(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	params := mux.Vars(r)

	// get the id from the url
	keywordStr := params["keyword"]

	client := &http.Client{}
	url := ctx.OnetApiBase + OnetService + keywordStr
	fmt.Println(url)
	r, _ = http.NewRequest(http.MethodGet, url, nil) // URL-encoded payload
	r.Header.Add("Authorization", "Basic Y2FyZWVyY3VwaWQ6MzUzN3RuYg==")
	r.Header.Add("Content-Type", "application/vnd.org.onetcenter.online.search.occupations+json")
	r.Header.Add("Accept", "application/json")

	resp, _ := client.Do(r)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the body, err: %v", err)
	}

	keywordWrapper := &model.KeywordWrapper{}
	// casting it to a string removes all the weird spaces, escapes the quotes, etc.
	sb := string(body)
	fmt.Println(sb)
	if err := json.Unmarshal([]byte(sb), &keywordWrapper); err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(keywordWrapper); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}