package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "routes"
)

func main() {
  r := mux.NewRouter().StrictSlash(true)

  /*-------------------------------------
   *         `/api` router
   *------------------------------------*/
  api := r.PathPrefix("/api").Subrouter()

  /*-------------------------------------
   *     `/api/isalive` test route
   *------------------------------------*/
  t := api.Path("/isalive").Subrouter()
  t.Methods("GET").HandlerFunc(routes.IsAlive)

  /*-------------------------------------
   *      `/api/videos` sub-route
   *------------------------------------*/
  v := api.PathPrefix("/videos").Subrouter()

  v.Methods("POST").HandlerFunc(routes.CreateVideo)
  v.Methods("OPTIONS").Path("/sign").HandlerFunc(routes.AllowAccess)
  v.Methods("POST").Path("/sign").HandlerFunc(routes.SignVideo)
  v.Methods("GET").Path("/{url}").HandlerFunc(routes.GetVideo)
  v.Methods("POST").Path("/process").HandlerFunc(routes.ProcessVideo)
  v.Methods("POST").HandlerFunc(routes.CreateVideo)


  /*-------------------------------------
   *      `/` static file server
   *------------------------------------*/
  r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./doc/"))))

  // Start up server and error log
  log.Println("Listening at port 3001")
  log.Fatal(http.ListenAndServe(":3001", r))
}
