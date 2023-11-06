package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bogdanguranda/rest-service/service"
	"github.com/bogdanguranda/rest-service/util/logging"
	"github.com/gorilla/mux"
)

type GetValueHandler struct {
	searcher service.Searcher
	logger   logging.Logger
}

func NewGetValueHandler(searcher service.Searcher, logger logging.Logger) *GetValueHandler {
	return &GetValueHandler{searcher: searcher, logger: logger}
}

type Response struct {
	Index   int    `json:"index"`
	Value   int    `json:"value"`
	Message string `json:"error,omitempty"`
}

func (gvh GetValueHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var resp Response

	vars := mux.Vars(r)
	value, ok := vars["value"]

	gvh.logger.Log(logging.LogLevelDebug, fmt.Sprintf("Path params are: %v", vars))
	gvh.logger.Log(logging.LogLevelDebug, fmt.Sprintf("Path param value is: %v", value))

	if !ok || value == "" {
		gvh.logger.Log(logging.LogLevelError, "Missing or invalid 'value' path parameter.")
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "ERROR: Missing or invalid 'value' path parameter."
		gvh.writeResponse(w, resp)
		return
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		gvh.logger.Log(logging.LogLevelError, "Invalid 'value' path parameter, must be a number.")
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "ERROR: Invalid 'value' path parameter, must be a number."
		gvh.writeResponse(w, resp)
		return
	}

	index, actualValue := gvh.searcher.Search(valueInt)
	if index == -1 {
		gvh.logger.Log(logging.LogLevelInfo, fmt.Sprintf("Value %s not found", value))
		w.WriteHeader(http.StatusNotFound)
		resp.Message = fmt.Sprintf("ERROR: Value %s not found", value)
		gvh.writeResponse(w, resp)
		return
	}

	gvh.logger.Log(logging.LogLevelInfo, fmt.Sprintf("Found value: %d at index %d", actualValue, index))
	resp.Index = index
	resp.Value = actualValue
	gvh.writeResponse(w, resp)
}

func (gvh GetValueHandler) writeResponse(w http.ResponseWriter, resp Response) {
	jsonData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}
