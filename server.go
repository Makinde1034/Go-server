package main 

import (
	// "fmt"
	"net/http"
)



func getText( w http.ResponseWriter, r *http.Request){  
	w.Write([]byte("Hello world"))
}

func main(){
	// delare multiplexer that machtes incoming requests to handlers
	mux := http.NewServeMux()

	mux.HandleFunc("/",getText)
	

	http.ListenAndServe(":8000",mux)
}