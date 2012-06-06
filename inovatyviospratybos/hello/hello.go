package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

    a :=
"<html>" +
  "<body bgcolor='silver'>" +
    "<h1>Prasideda inovatyviu pratybu programavimo darbai</h1>" +
  "</body>" +
"</html>"

    fmt.Fprint(w, a)
}


