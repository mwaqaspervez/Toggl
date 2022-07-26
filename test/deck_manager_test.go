package test

import (
	"Toggl/manager"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestValidateAndGetCards_WithEmptyInput(t *testing.T) {

	values, err := manager.ValidateAndGetCards(make([]string, 0))

	assert.NotNil(t, values)
	assert.Nil(t, err)
	assert.Equal(t, 52, len(values))
}

func TestValidateAndGetCards_WithCards(t *testing.T) {

	cards := "2D,3C,2H,TD"
	values, err := manager.ValidateAndGetCards(strings.Split(cards, ","))

	assert.NotNil(t, values)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(values))
}

func TestValidateAndGetCards_WithInvalidCards(t *testing.T) {

	cards := "2D,3C,1A,TD"
	values, err := manager.ValidateAndGetCards(strings.Split(cards, ","))

	assert.NotNil(t, err)
	assert.Nil(t, values)
	assert.Equal(t, 401, err.ResponseCode)
}

func TestCreateAndShuffle_WithoutShuffle(t *testing.T) {

	values, _ := manager.ValidateAndGetCards(make([]string, 0))
	deck := manager.CreateAndShuffle(false, values)

	assert.NotNil(t, deck)
	assert.Equal(t, 52, len(deck))
}

func TestCreateAndShuffle_WithShuffle(t *testing.T) {

	values, _ := manager.ValidateAndGetCards(make([]string, 0))
	deck := manager.CreateAndShuffle(true, values)

	assert.NotNil(t, deck)
	assert.Equal(t, 52, len(deck))
}

func TestCreateAndShuffle_WithCards(t *testing.T) {

	cards := "2D,3C,TD"
	values, _ := manager.ValidateAndGetCards(strings.Split(cards, ","))
	deck := manager.CreateAndShuffle(true, values)

	assert.NotNil(t, deck)
	assert.Equal(t, 3, len(deck))
}

func TestDrawRandomCard_withCount(t *testing.T) {

	values, _ := manager.ValidateAndGetCards(make([]string, 0))
	createDeck := manager.CreateAndShuffle(true, values)
	deck, cards := manager.DrawRandomCards(createDeck, 5)

	assert.NotNil(t, deck)
	assert.NotNil(t, cards)
	assert.Equal(t, 5, len(deck))
}
