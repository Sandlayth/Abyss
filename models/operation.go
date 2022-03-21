package models

import (
    "time"
    "errors"
    "strings"
    "strconv"
)

var IdCounter = 0

type Operation struct {
    Id            int       `json: id`
    Date          time.Time `json: "date"`
    EffectiveDate time.Time `json: "effectiveDate"`
    Label         string    `json: "id"`
    Amount        float64   `json: "amount"`
}

func (h Operation) parseOperation(s string) (*Operation, error) {
    parsedString := strings.Split(s, ";")[:]
    if len(parsedString) < 6 {
        return nil, errors.New("Failed to parse line")
    }
    IdCounter += 1
    date, err := time.Parse("02/01/2006", parsedString[0])
    effectiveDate, err := time.Parse("02/01/06", strings.Fields(parsedString[2])[1])
    label := strings.SplitAfterN(parsedString[2], " ", 3)[2]
    amount, err := strconv.ParseFloat(strings.Replace(parsedString[6], ",", ".", 1), 2)

    o := Operation {
        Id:            IdCounter,
        Date:          date,
        EffectiveDate: effectiveDate,
        Label:         label,
        Amount:        amount,
    }

    return &o, err
}
