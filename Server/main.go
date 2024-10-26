package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Method responsable for starting the server.
func Start_Server(ad string, handl http.Handler,r_timeout int,w_timeout int, MHB int){
	s := &http.Server{
		Addr:           ad,
		Handler:        handl,
		ReadTimeout:    time.Duration(r_timeout),
		WriteTimeout:   time.Duration(w_timeout),
		MaxHeaderBytes: MHB,
	}
	log.Fatal(s.ListenAndServe())
}


/*
Method responsible for handing a HTTP request.

@param request a string representing the request that the server has gotten through the http channel.
*/
func (s *http.Server) Handle_Request(request string){}

func main(){
	fmt.Print("Buld is complete")

}