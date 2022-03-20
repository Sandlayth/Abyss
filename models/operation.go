package model

import (
    "time"
)

type Operation struct {
    Id            int       `json: id`
    Date          time.Time `json: "date"`
    EffectiveDate time.Time `json: "effectiveDate"`
    Label         string    `json: "id"`
    Balance       float     `json: "balance"`
}
