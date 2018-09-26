package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	//      "os"
	//      "strconv"
	//      "github.com/olesiapoz/microservices-demo"
)

func decodeAuthoriseRequest(arg string) float32 {
	var req map[string]interface{}
	json.Unmarshal([]byte(arg), &req)

	// Decode auth request

	//json.Unmarshal(arg, &req)
	am, ok := req["amount"].(float64)
	if !ok {
		return 0
	}
	return float32(am)
}

type UnmarshalKeyError struct {
	Key  string
	JSON string
}

func (e *UnmarshalKeyError) Error() string {
	return fmt.Sprintf("Cannot unmarshal object key %q from JSON: %s", e.Key, e.JSON)
}

var ErrInvalidJson = errors.New("Invalid json")

// AuthoriseRequest represents a request for payment authorisation.
// The Amount is the total amount of the transaction

type Service interface {
	Authorise(total float32) (Authorisation, error) // GET /paymentAuth
	//      Health() []Health                               // GET /health
}

type Authorisation struct {
	Authorised bool   `json:"authorised"`
	Message    string `json:"message"`
}

func NewAuthorisationService(declineOverAmount float32) Service {
	return &service{
		declineOverAmount: declineOverAmount,
	}
}

//ammount of money that shoudl be declined
type service struct {
	declineOverAmount float32
}

func (s *service) Authorise(amount float32) (Authorisation, error) {
	if amount == 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	if amount < 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	authorised := false
	message := "Payment declined"
	if amount <= s.declineOverAmount {
		authorised = true
		message = "Payment authorised"
	} else {
		message = fmt.Sprintf("Payment declined: amount exceeds %.2f", s.declineOverAmount)
	}
	return Authorisation{
		Authorised: authorised,
		Message:    message,
	}, nil
}

//error
var ErrInvalidPaymentAmount = errors.New("Invalid payment amount")

func main() {
	//program receives one argument: the JSON object as a string
	var declineAmount = flag.Float64("decline", 105, "Decline payments over certain amount")
	arg := os.Args[1]

	service := NewAuthorisationService(float32(*declineAmount))
	// unmarshal the string to a JSON object

	amount := decodeAuthoriseRequest(arg)
	auth, _ := service.Authorise(amount)

	json, _ := json.Marshal(auth)
	fmt.Println(string(json))
}
