package statement

import (
	"fmt"
	"github.com/vocationnations/api/helper"
	"net/http"
)

func GetCultureStatements(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	_, err := fmt.Fprintf(w, "Tayab is wrong!!!")
	if err != nil {
		return err
	}
	return nil;
}

func GetCultureStatement(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	fmt.Fprintf(w, "IMPLEMENT ME!!")
	return nil;
}
