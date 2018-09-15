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
	"github.com/olesiapoz/microservices-demo"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"golang.org/x/net/context"
)
func main() {
    //program receives one argument: the JSON object as a string
    arg := os.Args[1]

    // unmarshal the string to a JSON object
    var obj map[string]interface{}
	json.Unmarshal([]byte(arg), &requestammount)
	
	var (
		declineAmount = flag.Float64("decline", 105, "Decline payments over certain amount")
	)

   // Mechanical stuff.
	errc := make(chan error)
	ctx := context.Background()

	handler, logger := payment.WireUp(ctx, float32(*declineAmount), tracer, ServiceName)

    // last line of stdout is the result JSON object as a string

}
