package main
import(
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)
func main(){
    r := mux.NewRouter()
    r.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
        fmt.Fprintf(w,"Welcome to this life-changing API,refreshi22,cccc")
    })
    fmt.Println("Server listening")
    http.ListenAndServe(":80",r)


}
