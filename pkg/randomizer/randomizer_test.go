package randomizer

import (
	"testing"

	"github.com/dongri/phonenumber"
	"github.com/stretchr/testify/assert"
)

func TestPhone(t *testing.T) {
	p1 := "081284026291"
	p1want := "6281284026291"

	p2 := "+6281284026291"
	p2want := "6281284026291"

	p3 := "6681284026291"
	p3want := ""

	assert.Equal(t, p1want, phonenumber.Parse(p1, "ID"))
	assert.Equal(t, p2want, phonenumber.Parse(p2, "ID"))
	assert.Equal(t, p3want, phonenumber.Parse(p3, "ID"))

}
