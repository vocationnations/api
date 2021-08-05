package user

import (
	"fmt"
	"github.com/vocationnations/api/helper"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	_, err := fmt.Fprintf(w, "Taz is wrong!!!")
	if err != nil {
		return err
	}
	return nil;
}

func GetUser(w http.ResponseWriter, r *http.Request, ctx helper.AppContext) error {
	fmt.Fprintf(w, "IMPLEMENT ME!!")
	return nil;
}
