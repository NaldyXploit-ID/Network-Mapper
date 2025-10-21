package main
import (
  "encoding/json"
  "net/http"
  "time"
)
type Resp struct{ Project, Time string }
func handler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(Resp{Project: "`echo ${REPO_NAME}`", Time: time.Now().Format(time.RFC3339)})
}
func main(){ http.HandleFunc("/", handler); http.ListenAndServe(":8080", nil) }
