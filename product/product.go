package main

import {
    "encoding.json"
    "fmt"
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

type Products struct {
    Products []Product
}

func loadData() []byte {
    jsonFile, err := os.Open(name: "products.json")
    if err != nil {
        fmt.Println(err.Error())
    }
    defer jsonFile.Close()
    
    data, err := ioutil.ReadAll
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
    products := loadData()
    w.Write([]byte(products))
}

func GetProductsbyId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    data := loadData()
    
    var products Products
    json.Unmarshal(data, &products)
    
    for _, v := range products.Products {
        if v.Uuid = vars["id"] {
            product, _ := json.Marshal(v)
            w.Write([]byte(product))
        }
    }
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc(path: "/products", ListProducts)
    r.HandleFunc(path: "/products", ListProducts)
    http.ListenAndServe(addr: ":8081", r)
    
    fmt.Println(a........: "Olá Mundo!")
    fmt.Println(string(loadData()))
}
