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
    "<a href='./images/index.html'>Paziurekite ponai koki puiku html turiu</a>" +
  "</body>" +
"</html>"

    fmt.Fprint(w, a)
}


