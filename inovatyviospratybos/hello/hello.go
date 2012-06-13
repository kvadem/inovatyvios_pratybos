package hello

import (
    "fmt"
    "os"
    "bufio"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

  //home_page := read_file_content("./view/index.html")
  home_page := read_file_content("./index.html")
  fmt.Fprint(w, home_page)
}

/**
 * This function reads file content and return it as a string.
 * author: Dainius Jocas
 * input: file_name
 * output: content
 */
func read_file_content(file_name string) string{
  f, _ := os.Open(file_name)
  defer f.Close()
  r := bufio.NewReader(f)
  var content string = "";
  for {
    line, ok := r.ReadString('\n')
    if ok != nil { break }
    content = content + line
  }
  return content
}

