package main

// https://golang.org/pkg/net/url/#URL

//  type URL struct {
//          Scheme   string
//          Opaque   string    // encoded opaque data
//          User     *Userinfo // username and password information
//          Host     string    // host or host:port
//          Path     string
//          RawPath  string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
//          RawQuery string // encoded query values, without '?'
//          Fragment string // fragment for references, without '#'
//  }
//
//  A URL represents a parsed URL (technically, a URI reference). The general form represented is:
//
//    scheme://[userinfo@]host/path[?query][#fragment]
//
//  URLs that do not start with a slash after the scheme are interpreted as:
//
//    scheme:opaque[?query][#fragment]


import "fmt"
import "io"
import "net/http"
import "io/ioutil"

import "strconv"

func (glfd *GLFD) WebDefault(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  url := req.URL
  fmt.Printf("default:\n")
  fmt.Printf("  method: %s\n", req.Method)
  fmt.Printf("  proto:  %s\n", req.Proto)
  fmt.Printf("  scheme: %s\n", url.Scheme)
  fmt.Printf("  host:   %s\n", url.Host)
  fmt.Printf("  path:   %s\n", url.Path)
  fmt.Printf("  frag:   %s\n", url.Fragment)
  fmt.Printf("  body:   %s\n\n", body)

  io.WriteString(w, `{"value":"ok"}`)
}


func (glfd *GLFD) WebAbout(w http.ResponseWriter, req *http.Request) {
  //str,e := ioutil.ReadFile("html/about.html")
  str,e := ioutil.ReadFile( glfd.HTMLDir + "/about.html")
  if e!=nil { io.WriteString(w, "error") ; return }
  io.WriteString(w, string(str))
}

func (glfd *GLFD) WebInteractive(w http.ResponseWriter, req *http.Request) {
  //str,e := ioutil.ReadFile("html/index.html")
  str,e := ioutil.ReadFile( glfd.HTMLDir + "/index.html")
  if e!=nil { io.WriteString(w, "error") ; return }
  io.WriteString(w, string(str))
}

func (glfd *GLFD) WebExec(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  fmt.Printf("webexec got>>>\n%s\n\n", body)

  rstr,e := glfd.JSVMRun(string(body))
  if e!=nil {
    rerr := strconv.Quote(fmt.Sprintf("%v", e))
    io.WriteString(w, `{"value":"error","error":` + rerr + `}`)
    return
  }

  io.WriteString(w, rstr)
}

func (glfd *GLFD) StartSrv() error {
  http.HandleFunc("/", glfd.WebDefault)
  http.HandleFunc("/exec", glfd.WebExec)
  http.HandleFunc("/exec/", glfd.WebExec)
  http.HandleFunc("/about", glfd.WebAbout)
  http.HandleFunc("/about/", glfd.WebAbout)
  http.HandleFunc("/i", glfd.WebInteractive)
  http.HandleFunc("/i/", glfd.WebInteractive)

  //http.ListenAndServe(":8081", nil)
  port_str := fmt.Sprintf("%d", glfd.Port)
  return http.ListenAndServe(":" + port_str, nil)
}
