package helper

import (
	"fmt"
	"github.com/vocationnations/api/database"
	"net/http"
)

const ENVLocal = "local"

type AppContext struct {
	VNConfiguration VNConfiguration
	DB              database.Database
}

// MakeHandler makeHandler allows us to pass the API context. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc.
func MakeHandler(ctx AppContext, fn func(http.ResponseWriter, *http.Request, AppContext) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r, ctx)
		if err != nil {
			_, _ = fmt.Fprintf(w, "ERROR: %s", err)
			return
		}
	}
}
