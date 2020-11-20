package processor

import (
    "encoding/json"
    "net/http"
    // "fmt"

    "github.com/plankiton/PagarMeChallenge/user"
    "github.com/plankiton/PagarMeChallenge/util"
    "github.com/plankiton/PagarMeChallenge/card"
)

type BuyValue struct {
    Value     float64   `json:"value,omitempty"`
    Items     []string  `json:"items,omitempty"`
    Plots     int       `json:"plots,omitempty"`
}

type BuyRequest struct {
    Token     string    `json:"token,omitempty"`
    Content   *BuyValue `json:"content,omitempty"`
    Seller    string    `json:"seller_id,omitempty"`
    Acquire   string    `json:"acquire_id,omitempty"`
}

type AcquireBuyRequest struct {
    Request   *BuyRequest `json:"request,omitempty"`
    Card      *card.Card  `json:"card,omitempty"`
}

var acquires []string = []string{"cielo", "getnet", "rede", "stone"}

func Processor(w http.ResponseWriter, r *http.Request) {
    people := user.People()
    cards := card.Cards()

    var request BuyRequest
    json.NewDecoder(r.Body).Decode(&request)
    if request.Token == "" || request.Seller == "" {
        w.WriteHeader(400)
        json.NewEncoder(w).Encode(util.ErrorTemplate{
            Message: "Bad buy request!",
            Property: "slashtag",
            Code: "BadRequest",
        })
        return
    }


    {
        var acquire string = ""
        for _, item := range acquires {
            if item == request.Acquire {
                acquire = item
            }
        }

        if acquire == "" {
            w.WriteHeader(404)
            json.NewEncoder(w).Encode(util.ErrorTemplate{
                Message: "Acquire not found!",
                Property: "slashtag",
                Code: "AcquireNotFound",
            })
            return
        }
    }


    var seller user.Person
    for _, item := range people {
        if item.ID == request.Seller {
            seller = item
        }
    }

    if seller.ID == "" {
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(util.ErrorTemplate{
            Message: "Seller not found on user list!",
            Property: "slashtag",
            Code: "SellerNotFound",
        })
        return
    }

    var card card.Card
    for _, item := range cards {
        if item.ID == request.Token {
            card = item
        }
    }

    if card.ID == "" {
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(util.ErrorTemplate{
            Message: "Card not found!",
            Property: "slashtag",
            Code: "CardNotFound",
        })
        return
    }

    card.ID = ""
    request.Token = ""
    json.NewEncoder(w).Encode(AcquireBuyRequest{
        Request: &request,
        Card: &card,
    })
}
