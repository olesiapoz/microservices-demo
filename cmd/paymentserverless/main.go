package main

import "encoding/json"
import "fmt"
import "os"
import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	
	"github.com/go-kit/kit/log"
	"github.com/olesiapoz/microservices-demo/service"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"golang.org/x/net/context"
)
func main() {
    //program receives one argument: the JSON object as a string
    arg := os.Args[1]

    // unmarshal the string to a JSON object
    var req map[string]interface{}
	json.Unmarshal([]byte(arg), &req)

	ammount, ok := req["ammont"].(string)
	
	// last line of stdout is the result JSON object as a string
	msg := map[string]string{"ammount": (ammount)}
	res, _ := json.Marshal(msg)
	
	/*var (
		declineAmount = flag.Float64("decline", 105, "Decline payments over certain amount")
	)

	func decodeAuthoriseRequest(r *arg) (interface{}, error) {
		// Read the content
		var bodyBytes []byte
		if r != nil {
			var err error
			bodyBytes, err = ioutil.ReadAll(r)
			if err != nil {
				return nil, err
			}
		}
		// Save the content
		bodyString := string(bodyBytes)
	
		// Decode auth request
		var request AuthoriseRequest
		if err := json.Unmarshal(bodyBytes, &arg); err != nil {
			return nil, err
		}
	
		// If amount isn't present, error
		if arg.Amount == 0.0 {
			return nil, &UnmarshalKeyError{
				Key:  "amount",
				JSON: bodyString,
			}
		}
		return request, nil
	}
	
	type UnmarshalKeyError struct {
		Key  string
		JSON string
	}
	
	func (e *UnmarshalKeyError) Error() string {
		return fmt.Sprintf("Cannot unmarshal object key %q from JSON: %s", e.Key, e.JSON)
	}
	
	var ErrInvalidJson = errors.New("Invalid json")
	
	func encodeAuthoriseResponse(response interface{}) error {
		resp := response.(AuthoriseResponse)
		if resp.Err != nil {
			encodeError(resp.Err)
			return nil
		}
		return encodeResponse(resp.Authorisation)
	}
	
	func encodeResponse(response interface{}) error {
		// All of our response objects are JSON serializable, so we just do that.
		return json.Encode(response)

	}

// AuthoriseRequest represents a request for payment authorisation.
// The Amount is the total amount of the transaction
type AuthoriseRequest struct {
	Amount float32 `json:"amount"`
}

// AuthoriseResponse returns a response of type Authorisation and an error, Err.
type AuthoriseResponse struct {
	Authorisation Authorisation
	Err           error
}

    // last line of stdout is the result JSON object as a string

}
*/