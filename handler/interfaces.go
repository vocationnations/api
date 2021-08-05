package handler

import (
	"github.com/vocationnations/api/helper"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, helper.AppContext) error
