package controller

import (
    "fmt"
    "os"
    "bufio"
    "net/http"
)

const root_url = "/"
const home_url = "./index.html"
const test_url = "./view/test.html"
const page_404 = "./view/404.html"

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/test/test.html", handler_test)
    http.HandleFunc("/test/404.html", handler_404)
}

/**
 * This func loads home_page.
 */
func handler(w http.ResponseWriter, r *http.Request) {

  //home_page := read_file_content("./view/index.html")
  home_page_html := read_file_content(home_url)
  fmt.Fprint(w, home_page_html)
}

/**
 * This func load test page.
 */
func handler_test(w http.ResponseWriter, r *http.Request) {
  test_page_html := read_file_content(test_url)
  fmt.Fprint(w, test_page_html)
}

/**
 * This func loads 404 page.
 */
func handler_404(w http.ResponseWriter, r *http.Request) {
  page_404_html := read_file_content(page_404)
  fmt.Fprint(w, page_404_html)
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

