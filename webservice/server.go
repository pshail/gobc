package webservice

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/heckdevice/gobc/core"
	"github.com/heckdevice/gobc/utils"
)

/**** Server Routes & Port ******/
var (
	geturl  = "/gobc"
	posturl = "/gobc/add"
)

// FetchGetURL - GET blockchain api relative url
func FetchGetURL() string {
	return geturl
}

// FetchPostURL - POST data to blockchain relative url
func FetchPostURL() string {
	return posturl
}

// GetPort - get Blockchain webservice port
func GetPort() string {
	return os.Getenv("PORT")
}

/****** Muxes and handlers *******/
func registerMuxes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(FetchGetURL(), handleGetBlockchain).Methods("GET")
	router.HandleFunc(FetchPostURL(), handleWriteBlock).Methods("POST")
	return router
}

// Run - start the webservice
func Run() error {
	mux := registerMuxes()
	log.Println("GoBC Server Listening on port :", GetPort())
	s := &http.Server{
		Addr:           ":" + GetPort(),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

// fetches the full blockchain as json data
// TODO we need to optimize this with pagination
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bcJSON, err := core.GetChain()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, *bcJSON)
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	_, newBlock, err := core.Add(data)
	if err != nil || newBlock == nil {
		respondWithJSON(w, r, http.StatusInternalServerError, r.Body)
		return
	}
	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := utils.InterfaceToJSONString(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	io.WriteString(w, *response)
}
