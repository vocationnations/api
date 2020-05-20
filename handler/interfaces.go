package handler

import (
	"github.com/vocationnations/api/helpter"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, helpter.AppContext) error
