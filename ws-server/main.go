package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var cVal int
var a int
var b int

type wsRequest struct {
	Value    int    `json:"value"`
	Endpoint string `json:"endpoint"`
}

type wsResponse struct {
	Value int `json:"value"`
}

type wsServer struct {
	logf     func(format string, v ...interface{})
	serveMux http.ServeMux
}

func (ws *wsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.serveMux.ServeHTTP(w, r)
}

func (ws *wsServer) AEndpoint(w http.ResponseWriter, r *http.Request) {
	// Store the incoming value in the global variable "a" or "b" depending on the endpoint that's passed in the request
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		ws.logf("failed to accept websocket connection: %v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	var wr wsRequest
	wsjson.Read(r.Context(), c, &wr)
	if err != nil {
		ws.logf("failed to read message: %v", err)
		return
	}

	if wr.Endpoint == "a" {
		a = wr.Value
	}
	if wr.Endpoint == "b" {
		b = wr.Value
	}
	log.Printf("Posted value %d to %s", wr.Value, wr.Endpoint)
}

func (ws *wsServer) CEndpoint(w http.ResponseWriter, r *http.Request) {
	// Return the value of the global variable "c"
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		ws.logf("failed to accept websocket connection: %v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	cVal = a + b
	log.Printf("Returning value %d for c", cVal)
	resp := wsResponse{Value: cVal}
	wsjson.Write(r.Context(), c, resp)
}

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cs := &wsServer{
		logf: log.Printf,
	}
	cs.serveMux.HandleFunc("/send", cs.AEndpoint)
	cs.serveMux.HandleFunc("/sum", cs.CEndpoint)
	s := &http.Server{
		Handler:      cs,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Printf("listening on %s", l.Addr())
	s.Serve(l)
}
