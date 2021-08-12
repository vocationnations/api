package handler

import (
	"github.com/vocationnations/api/helper"
	"net/http"
)

type Func func(http.ResponseWriter, *http.Request, helper.AppContext) error
