package util

import (
    "crypto/sha1"
    "fmt"
)

type ErrorTemplate struct {
    Message   string `json:"message,omitempty"`
    Code      string `json:"code,omitempty"`
    Property  string `json:"property,omitempty"`
}

func Hash(s string) string{
    h := sha1.New()
    return fmt.Sprintf("%x",h.Sum([]byte(s)))
}
