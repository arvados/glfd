package main

import "fmt"
import "io"
import "net/http"
import "io/ioutil"

func (glfd *GLFD) Interactive(w http.ResponseWriter, req *http.Request) {

  str,e := ioutil.ReadFile("html/index.html")
  if e!=nil { io.WriteString(w, "error") ; return }

  io.WriteString(w, string(str))
}


func (glfd *GLFD) Hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "hello, friend")
}

func (glfd *GLFD) Process(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  fmt.Printf("got>>>\n%s\n\n", body)

  io.WriteString(w, `{"value":"ok"}`)
}

func (glfd *GLFD) WebRun(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  fmt.Printf("webrun got>>>\n%s\n\n", body)

  io.WriteString(w, `{"value":"ok"}`)
}

func (glfd *GLFD) StartSrv() {
  http.HandleFunc("/", glfd.Hello)
  http.HandleFunc("/t", glfd.Process)
  http.HandleFunc("/js", glfd.WebRun)
  http.HandleFunc("/i", glfd.Interactive)
  http.ListenAndServe(":8081", nil)
}
