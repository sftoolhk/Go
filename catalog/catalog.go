 
package main

import {
    "encoding.json"
    "fmt"
    "html/template"
    "gitub.com/gorilla/mux"
    "io/ioutil"
    "net/http"
    "os"
    
}

type Product struct {
    Uuid string 'json:"uuid"'
    Product string 'json:"product"'
    Price float64 'json:"Price,string"'
}

var productsUrl string

func init() {
    productsUrl = os.Getenv(key "PRODUCT_URL")
}

func loadProducts() []Product{
    response, err := http.Get(productsUrl + "/products")
    if err != nil {
        fmt.Println( a...: "Erro de HTTP")
    }
    
    data, _ := ioutil.ReadAll(response.Body)
    
    var products Products
    json.Unmarshal(data, &products)
    
    fmt.Println(string(data))
    return products.Products
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc( path: "/", ListProducts)
    http.ListenAndAServe( addr: ":8080", r)
}

func ListProducts (w http.ResponseWriter, *http.Request) {
    products := loadProducts()
    t := templates.Must(template.ParseFiles( filename))
    t.Execute(w, products)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.vars(r)
    
}
