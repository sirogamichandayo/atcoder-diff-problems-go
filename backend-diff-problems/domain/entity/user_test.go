package entity

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func Test(t *testing.T) {
	uuidObj := uuid.New()
	fmt.Println(uuidObj.MarshalBinary())
}
