package test

import (
	"Toggl/service"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreateDeck_WithDefaultValues(t *testing.T) {

	deck, err := service.CreateDeck(false, make([]string, 0))

	assert.Nil(t, err)
	assert.NotNil(t, deck)
	assert.False(t, deck.Shuffled)
	assert.NotNil(t, deck.ID)
	assert.Equal(t, 52, deck.Remaining)
}

func TestCreateDeck_WithShuffledTrue(t *testing.T) {

	deck, err := service.CreateDeck(true, make([]string, 0))

	assert.Nil(t, err)
	assert.NotNil(t, deck)
	assert.True(t, deck.Shuffled)
	assert.NotNil(t, deck.ID)
	assert.Equal(t, 52, deck.Remaining)
}

func TestCreateDeck_WithCards(t *testing.T) {
	cards := "4C,2D,5H,JC"
	deck, err := service.CreateDeck(false, strings.Split(cards, ","))

	assert.Nil(t, err)
	assert.NotNil(t, deck)
	assert.False(t, deck.Shuffled)
	assert.NotNil(t, deck.ID)
	assert.Equal(t, 4, deck.Remaining)
}

func TestCreateDeck_WithInvalidCards(t *testing.T) {
	cards := "4C,2D,5H,JA"
	deck, err := service.CreateDeck(false, strings.Split(cards, ","))

	assert.Nil(t, deck)
	assert.NotNil(t, err)
	assert.Equal(t, 401, err.ResponseCode)
	assert.NotNil(t, err.Message)
}

func TestOpenDeck_WithValidId(t *testing.T) {

	newDeck, _ := service.CreateDeck(false, make([]string, 0))

	deck, err := service.OpenDeck(newDeck.ID)

	assert.Nil(t, err)
	assert.NotNil(t, deck)
	assert.Equal(t, newDeck.ID, deck.ID)
	assert.False(t, deck.Shuffled)
	assert.Equal(t, 52, deck.Remaining)
	assert.NotNil(t, deck.Cards)
	assert.Equal(t, 52, len(deck.Cards))
}

func TestOpenDeck_WithInvalidId(t *testing.T) {

	service.CreateDeck(false, make([]string, 0))

	deck, err := service.OpenDeck("asdv-123123azx-x")

	assert.NotNil(t, err)
	assert.Nil(t, deck)
	assert.Equal(t, 401, err.ResponseCode)
	assert.NotNil(t, err.Message)
}

func TestDrawCard_WithValidId(t *testing.T) {

	newDeck, _ := service.CreateDeck(false, make([]string, 0))

	deck, err := service.DrawCard(newDeck.ID, 4)

	assert.Nil(t, err)
	assert.NotNil(t, deck)
	assert.Equal(t, 4, len(deck.Cards))

	openDeck, err := service.OpenDeck(newDeck.ID)
	assert.Nil(t, err)
	assert.NotNil(t, openDeck)
	assert.Equal(t, newDeck.ID, openDeck.ID)
	assert.False(t, openDeck.Shuffled)
	assert.Equal(t, 48, openDeck.Remaining)
	assert.NotNil(t, openDeck.Cards)
	assert.Equal(t, 48, len(openDeck.Cards))
}

func TestDrawCard_WithInvalidId(t *testing.T) {

	service.CreateDeck(false, make([]string, 0))

	deck, err := service.DrawCard("12312-123123", 4)

	assert.NotNil(t, err)
	assert.Nil(t, deck)
	assert.Equal(t, 401, err.ResponseCode)
	assert.NotNil(t, err.Message)
}

func TestDrawCard_WithInvalidCount(t *testing.T) {

	newDeck, _ := service.CreateDeck(false, make([]string, 0))

	deck, err := service.DrawCard(newDeck.ID, 60)

	assert.NotNil(t, err)
	assert.Nil(t, deck)
	assert.Equal(t, 401, err.ResponseCode)
	assert.NotNil(t, err.Message)
}
