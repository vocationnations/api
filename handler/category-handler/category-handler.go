package category_handler

import (
	"fmt"
	"github.com/vocationnations/api/helper"
	"net/http"
)

func GetCultureCategories(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	_, err := fmt.Fprintf(w, "Tayab is wrong!!!")
	if err != nil {
		return err
	}
	return nil
}

func GetCultureCategory(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	fmt.Fprintf(w, "IMPLEMENT ME!!")
	return nil
}
