package skill

import (
	"fmt"
	"github.com/vocationnations/api/helper"
	"net/http"
)

func GetSkills(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	_, err := fmt.Fprintf(w, "Tayab is wrong!!!")
	if err != nil {
		return err
	}
	return nil;
}

func GetSkill(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	fmt.Fprintf(w, "IMPLEMENT ME!!")
	return nil;
}