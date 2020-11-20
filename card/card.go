package card

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"

    "github.com/plankiton/PagarMeChallenge/util"
)

type Card struct {
    ID        string   `json:"id,omitempty"`
    Owner     string   `json:"owner,omitempty"`
    Validate  string   `json:"validate,omitempty"`
    Number    string   `json:"number,omitempty"`
    CVV       string   `json:"cvv,omitempty"`
}
var cards []Card

// GetCards mostra todos os contatos da variável cards
func GetCards(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(cards)
}

// GetCard mostra apenas um contato
func GetCard(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range cards {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Card{})
}

// CreateCard cria um novo contato
func CreateCard(w http.ResponseWriter, r *http.Request) {
    var card Card
    _ = json.NewDecoder(r.Body).Decode(&card)

    for _, item := range cards {
        if item.ID == card.ID {
            w.WriteHeader(403)
            json.NewEncoder(w).Encode(util.ErrorTemplate{
                Message: "The card already exists!",
                Property: "slashtag",
                Code: "AlreadyExists",
            })
            return
        }
    }

    cards = append(cards, card)
    json.NewEncoder(w).Encode(cards)
}

// DeleteCard deleta um contato
func DeleteCard(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range cards {
        if item.ID == params["id"] {
            cards = append(cards[:index], cards[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(cards)
    }
}

func AppendCard(c Card) {
    cards = append(cards, c)
}

func Cards() []Card {
    return cards
}
