package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/xgolis/wac-tg-jv-backend/docs"
	"github.com/xgolis/wac-tg-jv-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Message string        `json:"message,omitempty"`
	Results []primitive.M `json:"results,omitempty"`
}

type requestID struct {
	ID string `json:"id"`
}

type recordReq struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	PatientName   string `json:"patientName,omitempty"`
	DateOfBirth   string `json:"dateOfBirth,omitempty"`
	Description   string `json:"description,omitempty"`
	RequirementID string `json:"requirementID,omitempty"`
}

func MakeHandlers() *http.ServeMux {
	mux := *http.NewServeMux()
	// mux.HandleFunc("/", sendHello)
	mux.HandleFunc("/records", enableCORS(getRecord))
	mux.HandleFunc("/docs/", enableCORS(swaggerHandler))
	mux.HandleFunc("/record", enableCORS(putRecord))
	mux.HandleFunc("/filter", enableCORS(filterRecord))
	mux.HandleFunc("/delete", enableCORS(deleteRecord))
	mux.HandleFunc("/update", enableCORS(updateRecord))

	return &mux
}

func swaggerHandler(res http.ResponseWriter, req *http.Request) {
	httpSwagger.WrapHandler(res, req)
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func sendError(w *http.ResponseWriter, err error, statusCode int) {
	status := Response{
		Message: err.Error(),
	}

	fmt.Println(status)
	statusJson, err := json.Marshal(&status)
	if err != nil {
		http.Error(*w, err.Error(), statusCode)
		panic(err)
	}

	(*w).WriteHeader(statusCode)
	(*w).Write(statusJson)
}

func sendHello(w http.ResponseWriter, req *http.Request) {
	resp := &Response{
		Message: "Hello",
	}

	byteResp, err := json.Marshal(resp)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	w.Write(byteResp)
}

// getRecord
//
//		@Summary		Get all records
//		@Description	Get all records from selected collection in database
//		@Produce		json
//	 	@Success		200 {object} string
//	 	@Param       collection query     string         true  "The collection name"  example(patients)
//		@Error       400        {string}  string "Bad Request"
//		@Error       500        {string}  string "Internal Server Error"
//		@Router			/records [GET]
func getRecord(w http.ResponseWriter, req *http.Request) {
	collection := req.URL.Query().Get("collection")

	results, err := db.GetAllColection(DB, collection)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&Response{Results: results})
}

// putRecord
//
//		@Summary		Insert record to DB
//		@Description	The endpoint inserts sent data to the database
//		@Accept 		json
//		@Produce		json
//	 	@Success		200 {object} string
//		@Param       collection query     string         true  "The collection name"  example(patients)
//		@Param       body       body      recordReq  true  "Insert Request Body"
//		@Error       400        {string}  string "Bad Request"
//		@Error       500        {string}  string "Internal Server Error"
//		@Router			/record [PUT]
func putRecord(w http.ResponseWriter, req *http.Request) {
	collection := req.URL.Query().Get("collection")

	var record bson.M
	if err := json.NewDecoder(req.Body).Decode(&record); err != nil {
		sendError(&w, fmt.Errorf("failed to decode JSON: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("Inserted record: %s \n", record)

	err := db.PutRecord(DB, collection, record)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ok"))
}

// deleteRecord
//
//		@Summary		Delete a record
//		@Description	Delete a record from Database
//		@Produce		json
//	 	@Success		200 {object} string
//	 	@Param       collection query     string         true  "The collection name"  example(patients)
//		@Param       body       body      requestID  true  "Delete Request Body"
//		@Error       400        {string}  string "Bad Request"
//		@Error       500        {string}  string "Internal Server Error"
//		@Router			/delete [DELETE]
func deleteRecord(w http.ResponseWriter, req *http.Request) {
	collection := req.URL.Query().Get("collection")

	var requestBody requestID

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = db.DeleteRecord(DB, collection, requestBody.ID)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ok"))
}

// updateRecord
//
//		@Summary		Update a record
//		@Description	Update a record from Database
//		@Produce		json
//	 	@Success		200 {object} string
//	 	@Param       collection query     string         true  "The collection name"  example(patients)
//		@Param       body       body      recordReq  true  "Delete Request Body"
//		@Error       400        {string}  string "Bad Request"
//		@Error       500        {string}  string "Internal Server Error"
//		@Router			/update [POST]
func updateRecord(w http.ResponseWriter, req *http.Request) {
	collection := req.URL.Query().Get("collection")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	var record bson.M
	err = json.Unmarshal(body, &record)
	if err != nil {
		sendError(&w, err, http.StatusBadRequest)
		return
	}

	delete(record, "id")

	var requestBody requestID
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		sendError(&w, err, http.StatusBadRequest)
		return
	}

	err = db.UpdateRecord(DB, collection, requestBody.ID, record)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ok"))
}

func filterRecord(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Got request %s\n", req.URL.RawQuery)

	collection := req.URL.Query().Get("collection")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	var requestBody bson.M

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		sendError(&w, err, http.StatusBadRequest)
		return
	}

	results, err := db.FilterCollection(DB, collection, requestBody)
	if err != nil {
		sendError(&w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&Response{Results: results})
}
