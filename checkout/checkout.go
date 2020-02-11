package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "os"
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
)

type Product struct {
    Uuid
    Product
    Price
}
        
var productsUrl string

func init() {
    productsUrl = os.Getenv( key: "PRODUCTS_URL" )
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
    fmt.Println(string("Olá"))
    w.Write([]byte ("Olá"))
    
    vars:= mux.Vars(r)
    response, err := http.Get(productsUrl + "/product/" + vars["id"])
    if err != nil {
        fmt.Println(String(data))
        
    var product Product
    json.Unmarshal(data, &product)
    
    t := template.Must(template.ParseFile( filename...: "templates/checkout.html"))
    t.Execute(W, product)
    }
}

func finish(w http.ResponseWriter, r *http.Request) {
    var order Order
    order.Name = r.FormValue ( key: "name")
    order.Email= r.FormValue ( key: "email")
    order.Phone = r.FormValue ( key: "phone")
    order.ProductId = r.FormValue ( key: "product_id")
    
    data, _ := json.Marshal(order)
    fmt.Println(string(data))
    
    queue.Connect()
    
    w.Write([]byte("Processou!"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc ( path: "/{finish}", finish)
    r.HandleFunc ( path: "/{id}", displayCheckout)

    http.ListenAndServe( addr: ":8082", r)
}
