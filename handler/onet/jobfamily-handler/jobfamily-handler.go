package jobfamily_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vocationnations/api/helper"
	"github.com/vocationnations/api/model"
	"io/ioutil"
	"net/http"
)

const OnetService = "job_families"

func GetJobFamilies(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	client := &http.Client{}
	url := ctx.VNConfiguration.ONET.Base + OnetService + "/"
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

	// casting it to a string removes all the weird spaces, escapes the quotes, etc.
	sb := string(body)
	jobFamilies := &model.JobFamilies{}
	if err := json.Unmarshal([]byte(sb), &jobFamilies); err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(jobFamilies); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}

func GetOccupationsFromFamily(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {

	params := mux.Vars(r)

	// get the id from the url
	idStr := params["id"]

	client := &http.Client{}
	url := ctx.VNConfiguration.ONET.Base + OnetService + "/" + idStr
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

	occupationWrapper := &model.OccupationWrapper{}
	// casting it to a string removes all the weird spaces, escapes the quotes, etc.
	sb := string(body)
	fmt.Println(sb)
	if err := json.Unmarshal([]byte(sb), &occupationWrapper); err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(occupationWrapper); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}

func GetTasksByOccupation(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	// get the id from the url
	idStr := params["id"]
	authorization := ctx.VNConfiguration.ONET.AuthType + " " + ctx.VNConfiguration.ONET.AuthToken
	acceptType := ctx.VNConfiguration.ONET.AcceptType
	client := &http.Client{}
	url := ctx.VNConfiguration.ONET.Base + "related_task/tasks/" + idStr
	fmt.Println(url)
	r, _ = http.NewRequest(http.MethodGet, url, nil) // URL-encoded payload
	r.Header.Add("Authorization", authorization)
	r.Header.Add("Content-Type", "Content-Type: application/vnd.org.onetcenter.online.related_task.tasks+json")
	r.Header.Add("Accept", acceptType)

	resp, _ := client.Do(r)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the body, err: %v", err)
	}

	occupationTasks := &model.OccupationTasks{}
	// casting it to a string removes all the weird spaces, escapes the quotes, etc.
	sb := string(body)
	fmt.Println(sb)
	if err := json.Unmarshal([]byte(sb), &occupationTasks); err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(occupationTasks); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}

func GetSkillsByOccupation(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	// get the id from the url
	idStr := params["id"]
	authorization := ctx.VNConfiguration.ONET.AuthType + " " + ctx.VNConfiguration.ONET.AuthToken
	acceptType := ctx.VNConfiguration.ONET.AcceptType
	client := &http.Client{}
	url := ctx.VNConfiguration.ONET.Base + "occupations/" + idStr + "/summary/skills"
	fmt.Println(url)
	r, _ = http.NewRequest(http.MethodGet, url, nil) // URL-encoded payload
	r.Header.Add("Authorization", authorization)
	r.Header.Add("Content-Type", "Content-Type: application/vnd.org.onetcenter.online.occupation.skills+json")
	r.Header.Add("Accept", acceptType)

	resp, _ := client.Do(r)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the body, err: %v", err)
	}

	occupationSkills := &model.OccupationSkills{}
	// casting it to a string removes all the weird spaces, escapes the quotes, etc.
	sb := string(body)
	fmt.Println(sb)
	if err := json.Unmarshal([]byte(sb), &occupationSkills); err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(occupationSkills); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}

func GetKnowledgeSummaryByOccupation(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	// get the id from the url
	idStr := params["id"]
	authorization := ctx.VNConfiguration.ONET.AuthType + " " + ctx.VNConfiguration.ONET.AuthToken
	acceptType := ctx.VNConfiguration.ONET.AcceptType
	client := &http.Client{}
	url := ctx.VNConfiguration.ONET.Base + "occupations/" + idStr + "/summary/knowledge"
	fmt.Println(url)
	r, _ = http.NewRequest(http.MethodGet, url, nil) // URL-encoded payload
	r.Header.Add("Authorization", authorization)
	r.Header.Add("Content-Type", "Content-Type: application/vnd.org.onetcenter.online.occupation.knowledge+json")
	r.Header.Add("Accept", acceptType)

	resp, _ := client.Do(r)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the body, err: %v", err)
	}

	// Knowledge summary has the exact same json model as skills summary, so I am using that model for now to avoid duplication.
	occupationSkills := &model.OccupationSkills{}
	// casting it to a string removes all the weird spaces, escapes the quotes, etc.
	sb := string(body)
	fmt.Println(sb)
	if err := json.Unmarshal([]byte(sb), &occupationSkills); err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(occupationSkills); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}

func GetEducationOccupation(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	params := mux.Vars(r)

	// get the id from the url
	idStr := params["id"]
	authorization := ctx.VNConfiguration.ONET.AuthType + " " + ctx.VNConfiguration.ONET.AuthToken
	acceptType := ctx.VNConfiguration.ONET.AcceptType
	client := &http.Client{}
	url := ctx.VNConfiguration.ONET.Base + "occupations/" + idStr + "/summary/education"
	fmt.Println(url)
	r, _ = http.NewRequest(http.MethodGet, url, nil) // URL-encoded payload
	r.Header.Add("Authorization", authorization)
	r.Header.Add("Content-Type", "Content-Type: application/vnd.org.onetcenter.online.occupation.education+json")
	r.Header.Add("Accept", acceptType)

	resp, _ := client.Do(r)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the body, err: %v", err)
	}

	occupationEducation := &model.OccupationEducation{}
	// casting it to a string removes all the weird spaces, escapes the quotes, etc.
	sb := string(body)
	fmt.Println(sb)
	if err := json.Unmarshal([]byte(sb), &occupationEducation); err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(occupationEducation); err != nil {
		return fmt.Errorf("cannot write the json, err: %v", err)
	}
	return nil
}
