package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "fmt"

    "github.com/plankiton/PagarMeChallenge/util"
    "github.com/plankiton/PagarMeChallenge/user"
    "github.com/plankiton/PagarMeChallenge/card"
)

// função principal para executar a api
func main() {
    router := mux.NewRouter()

    user.AppendPerson(user.Person{
        Document: &user.Doc{
            Type: "cpf",
            Value: "123456789",
        },
        Name: "Joao Da Silva",
    })

    card.AppendCard(card.Card{
        ID: fmt.Sprintf("$pagarme$%s$%s$%s$%s",
           util.Hash("123456789"),
           util.Hash("Joao Da Silva"),
           util.Hash("987"),
           util.Hash("10/32")),
        Owner: "1",
        Validate: "10/32",
        CVV: "987",
    });

    router.HandleFunc("/card", card.GetCards).Methods("GET")
    router.HandleFunc("/card", card.CreateCard).Methods("POST")
    router.HandleFunc("/card/{id}", card.DeleteCard).Methods("DELETE")
    router.HandleFunc("/card/{id}", card.GetCard).Methods("GET")
    log.Output(2, "Routing /card - Card operations")

    router.HandleFunc("/user", user.GetPeople).Methods("GET")
    router.HandleFunc("/user", user.CreatePerson).Methods("POST")
    router.HandleFunc("/user/{id}", user.DeletePerson).Methods("DELETE")
    router.HandleFunc("/user/{id}", user.GetPerson).Methods("GET")
    log.Output(2, "Routing /user - User operations")

    log.Fatal(http.ListenAndServe(":8000", router))
}
