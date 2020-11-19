package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// "Person type" (tipo um objeto)

type Doc struct {
    Type      string `json:"type,omitempty"`
    Value     string `json:"value,omitempty"`
}

type Person struct {
    ID        *Doc     `json:"id,omitempty"`
    Name string        `json:"firstname,omitempty"`
}

type Card struct {
    ID        string   `json:"id,omitempty"`
    Owner     string   `json:"owner,omitempty"`
    Validate  [2]int   `json:"validate,omitempty"`
    Number    string   `json:"number,omitempty"`
    CVV          int   `json:"cvv,omitempty"`
}

type BuyValue struct {

}

type BuyRequest struct {
    Token     string    `json:"token,omitempty"`
    Content   *BuyValue `json:"content,omitempty"`
    Seller    string    `json:"seller_id,omitempty"`
    Acquire   string    `json:"acquire_id,omitempty"`
}

type AcquireBuyRequest struct {
    Request   *BuyRequest `json:"request,omitempty"`
    Card      *Card       `json:"card,omitempty"`
}

var people []Person

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID.Value == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID.Value == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

// função principal para executar a api
func main() {
    router := mux.NewRouter()
    people = append(people, Person{
        ID: &Doc{
            Type:"id",
            Value: "1",
        },
        Name: "Joao Da Silva",
    })

    router.HandleFunc("/user", GetPeople).Methods("GET")
    router.HandleFunc("/user/", CreatePerson).Methods("POST")
    router.HandleFunc("/user/{id}", DeletePerson).Methods("DELETE")
    router.HandleFunc("/user/{id}", GetPerson).Methods("GET")
    log.Output(2, "Routing localhost:8000/user")
    log.Fatal(http.ListenAndServe(":8000", router))
}
