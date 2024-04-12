package utils

import (
	"fmt"
	"testing"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/dongri/phonenumber"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

func TestRandomAccount(t *testing.T) {
	// assert := assert.New(t)

	r := RandomAccountData(entity.Account{})
	fmt.Printf("%+v", r)

	// assert equality
	// assert.Equal(123, 123, "they should be equal")

	// assert inequality
	// assert.NotEqual(123, 456, "they should not be equal")
}

func TestPhone(t *testing.T) {
	p1 := "081284026291"
	p2 := "+6281284026291"
	p3 := "6281284026191"
	p4 := "6681284026291"
	assert.Equal(t, p1, phonenumber.Parse(p1, "ID"))
	fmt.Println("p1: ", phonenumber.Parse(p1, "ID"))
	assert.Equal(t, p2, phonenumber.Parse(p2, "ID"))
	fmt.Println("p2: ", phonenumber.Parse(p2, "ID"))
	assert.Equal(t, p3, phonenumber.Parse(p3, "ID"))
	fmt.Println("p3: ", phonenumber.Parse(p3, "ID"))
	assert.Equal(t, p4, phonenumber.Parse(p4, "ID"))
	fmt.Println("p4: ", phonenumber.Parse(p4, "ID"))
}
