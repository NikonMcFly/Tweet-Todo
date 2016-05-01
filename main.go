package main

import (
	"fmt"
	"net/http"
  "net/url"
	"time"
	"encoding/json"
  "github.com/ChimeraCoder/anaconda"
)


type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Todos []Todo


func helloServer(w http.ResponseWriter, res *http.Request){
	fmt.Fprintln(w, "hello World")
}

func todoApi(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
        Todo{Name: "what is going on"},
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(todos)
}

func TwitterApi(w http.ResponseWriter, r *http.Request){
const (
  ConsumerKey    = "Key"
  ConsumerSecret = "Secret"
)
anaconda.SetConsumerKey(ConsumerKey)
anaconda.SetConsumerSecret(ConsumerSecret)
api := anaconda.NewTwitterApi("", "")

v := url.Values{}
v.Set("count", "1")  
  searchResult, _ := api.GetRetweetsOfMe(v)

  for _ , tweet := range searchResult {
    w.Header().Set("Content-Type", "Application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
     
    json.NewEncoder(w).Encode(tweet)
  }
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/todo", todoApi)
  http.HandleFunc("/tweet", TwitterApi)
	http.ListenAndServe(":8080", nil)
}
