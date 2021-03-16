package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Test struct{}

type HelloArgs struct {
	Name string
}

func (test *Test) Hello(args *HelloArgs, result *string) error {
	*result = "Hello " + args.Name
	return nil
}

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *HttpConn) Write(d []byte) (n int, err error) {
	return c.out.Write(d)
}
func (c *HttpConn) Close() error {
	return nil
}

func main() {
	server := rpc.NewServer()
	server.Register(&Test{})

	handler := func(w http.ResponseWriter, r *http.Request) {
		serverCodec := jsonrpc.NewServerCodec(&HttpConn{r.Body, w})
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		if err := server.ServeRequest(serverCodec); err != nil {
			log.Printf("Error while serving JSON request: %v", err)
			http.Error(w, "Error while serving JSON request, details have been logged.", 500)
			return
		}
	}

	http.HandleFunc("/rpc", handler)
	http.ListenAndServe(":10101", nil)
	//curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "Test.Hello", "params":[{"Name":"Human"}]}' http://localhost:10101/rpc
}
