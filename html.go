package main
import "net/http"
import "fmt"
import "io/ioutil"


func init(){
  http.HandleFunc("/"+folder+"/"+index, h_index)
  http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))
}

func h_index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, html())
}


func run(){
  http.ListenAndServe(":"+port, nil)
}


func html() string{
      file, _ := ioutil.ReadFile("go.html")
    return string(file)
}
 