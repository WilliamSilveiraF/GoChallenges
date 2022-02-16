package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Address struct {
    Cep	  string `json: "cep"`
    Address string `json: "logradouro"`
    Complement string `json: "complemento"`
    District string `json: "bairro"`
    City string `json: "localidade"`
    State string `json: "uf"`
    IbgeId string `json: "ibge"`
    Gia string `json: "gia"`
    DDD string `json: "ddd"`
    Siafi string `json: "siafi"`
}


func main() {
    foo1 := new(Address)

    var myClient = &http.Client{}
    r, err := myClient.Get("https://viacep.com.br/ws/88036002/json/")
    if err != nil {
        return
    }
    defer r.Body.Close()

    json.NewDecoder(r.Body).Decode(foo1)

    println(foo1.Address)
    println(foo1.Complement)
}