package web

import (
	"net/http"

	"github.com/bogdanguranda/rest-service/service"
	"github.com/bogdanguranda/rest-service/util/logging"
)

type RequestHandler interface {
	Handle(w http.ResponseWriter, r *http.Request, searcher service.Searcher, logger logging.Logger)
}
