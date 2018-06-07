package astatine

import (
	"github.com/satori/go.uuid"
	"encoding/json"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func toString(i interface{}) string {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(j)
}
