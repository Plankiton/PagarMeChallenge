package user

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"

    "github.com/plankiton/PagarMeChallenge/util"
)

type Doc struct {
    Type      string `json:"type,omitempty"`
    Value     string `json:"value,omitempty"`
}

type Person struct {
    ID         string  `json:"id,omitempty"`
    Document   *Doc    `json:"document,omitempty"`
    Name       string  `json:"firstname,omitempty"`
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
        if item.ID == params["id"] || item.Document.Value == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    json.NewDecoder(r.Body).Decode(&person)

    for _, item := range people {
        if item.Document.Value == person.Document.Value {
            w.WriteHeader(403)
            json.NewEncoder(w).Encode(util.ErrorTemplate{
                Message: "The user already exists!",
                Property: "slashtag",
                Code: "AlreadyExists",
            })
            return
        }
    }

    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] || item.Document.Value == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

func AppendPerson(p Person) {
    people = append(people, p)
}

func People() []Person {
    return people
}
