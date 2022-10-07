package main

import (
	"atm/signals"
	"atm/starters"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BodyContent struct {
	HttpCode int
	Message  string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/withdraw/{amountValue}", Withdraw)
	r.HandleFunc("/sendCodePin/{workflowID}/{codePin}", SendCodePin)
	log.Fatal(http.ListenAndServe(":3310", r))
}

// Return body content as JSON content type.
func returnJsonContent(writer http.ResponseWriter, bodyContent BodyContent, httpCode int) {
	// Setup return as JSON content
	writer.Header().Set("Content-Type", "application/json")

	// Encode JSON
	bytesBodyContent, _ := json.Marshal(bodyContent)

	// Set http return code
	writer.WriteHeader(httpCode)

	// Return json content
	writer.Write(bytesBodyContent)

}

// Withdraw amount from account
func Withdraw(writer http.ResponseWriter, request *http.Request) {
	// Create unique uuid
	workflowUUID, err := uuid.NewUUID()
	if err != nil {
		bodyContent := BodyContent{http.StatusInternalServerError, err.Error()}
		returnJsonContent(writer, bodyContent, http.StatusInternalServerError)
	}

	// Get amount value from HTTP request
	vars := mux.Vars(request)
	amountValue, err := strconv.Atoi(vars["amountValue"])
	if err != nil {
		bodyContent := BodyContent{http.StatusInternalServerError, err.Error()}
		returnJsonContent(writer, bodyContent, http.StatusInternalServerError)
	}

	// Run Withdraw workflow
	starters.StartWithdrawWorkflowFunc(workflowUUID.String(), amountValue)

	bodyContent := BodyContent{http.StatusOK, fmt.Sprintf("workflowid: %s", workflowUUID)}
	returnJsonContent(writer, bodyContent, http.StatusOK)
}

func SendCodePin(writer http.ResponseWriter, request *http.Request) {
	// Get Workflow ID and CodePin parameters from HTTP request
	vars := mux.Vars(request)
	workflowId := vars["workflowID"]
	codePin, err := strconv.Atoi(vars["codePin"])
	if err != nil {
		bodyContent := BodyContent{http.StatusInternalServerError, err.Error()}
		returnJsonContent(writer, bodyContent, http.StatusInternalServerError)
	}

	// Check if Code Pin is valid
	if err := signals.VerifyCreditCardCodePin(workflowId, codePin); err != nil {
		bodyContent := BodyContent{http.StatusInternalServerError, err.Error()}
		returnJsonContent(writer, bodyContent, http.StatusInternalServerError)
	}

	bodyContent := BodyContent{http.StatusOK, "Code PIN is valid."}
	returnJsonContent(writer, bodyContent, http.StatusOK)
}
