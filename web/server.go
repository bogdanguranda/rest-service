package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bogdanguranda/rest-service/service"
	"github.com/bogdanguranda/rest-service/util/logging"
	"github.com/gorilla/mux"
)

func StartServer(port string, searcher service.Searcher, logger logging.Logger) {
	reqHandler := NewGetValueHandler(searcher, logger)

	router := mux.NewRouter()
	router.HandleFunc("/value/{value}", reqHandler.Handle).Methods("GET")

	serverAddress := fmt.Sprintf(":%s", port)
	logger.Log(logging.LogLevelInfo, fmt.Sprintf("Server is running on: http://localhost%s/", serverAddress))

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
