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
	url := ctx.VNConfiguration.ONET.Base + OnetService + keywordStr
	authorization := ctx.VNConfiguration.ONET.AuthType + " " + ctx.VNConfiguration.ONET.AuthToken
	contentType := ctx.VNConfiguration.ONET.ContentType
	acceptType := ctx.VNConfiguration.ONET.AcceptType
	fmt.Println(url)
	r, _ = http.NewRequest(http.MethodGet, url, nil) // URL-encoded payload
	r.Header.Add("Authorization", authorization)
	r.Header.Add("Content-Type", contentType)
	r.Header.Add("Accept", acceptType)

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
