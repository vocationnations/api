package jobfamily_handler

import (
	"encoding/json"
	"fmt"
	"github.com/vocationnations/api/helper"
	"io/ioutil"
	"net/http"
)

func GetJobFamilies(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	client := &http.Client{}
	r, _ = http.NewRequest(http.MethodGet, "https://services.onetcenter.org/ws/online/job_families/", nil) // URL-encoded payload
	r.Header.Add("Authorization", "Basic Y2FyZWVyY3VwaWQ6MzUzN3RuYg==")
	r.Header.Add("Content-Type", "application/vnd.org.onetcenter.online.job_families+json")
	r.Header.Add("Accept", "application/json")

	resp, _ := client.Do(r)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the body, err: %v", err)
	}

	// TODO: Encode the body so that the JSON can be added to the Encoder.
	if err := json.NewEncoder(w).Encode(body); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}
