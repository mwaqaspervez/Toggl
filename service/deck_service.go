package service

import (
	"Toggl/dto"
	"Toggl/manager"
	"github.com/google/uuid"
)

type DeckService interface {
	CreateDeck(shuffled bool, cards []string) error
	OpenDeck(id string) error
	DrawCard(id string, count int) error
}

type Decks map[string]*dto.Deck

var deckMap Decks = make(map[string]*dto.Deck)

func CreateDeck(shuffled bool, inputCards []string) (*dto.CreateDeckDTO, *dto.Error) {
	values, err := manager.ValidateAndGetCards(inputCards)
	if err != nil {
		return nil, err
	}

	cards := manager.CreateAndShuffle(shuffled, values)
	id := uuid.New().String()
	deckMap[id] = &dto.Deck{ID: id, Shuffled: shuffled, Cards: cards}
	return &dto.CreateDeckDTO{ID: id, Shuffled: shuffled, Remaining: len(cards)}, nil
}

func OpenDeck(id string) (*dto.OpenDeckDTO, *dto.Error) {
	deck, ok := deckMap[id]
	if !ok {
		return nil, &dto.Error{ResponseCode: 401, Message: "could not find deck with the given id"}
	}

	return &dto.OpenDeckDTO{ID: id, Shuffled: deck.Shuffled, Remaining: len(deck.Cards), Cards: deck.Cards}, nil
}

func DrawCard(id string, count int) (*dto.DrawCardDTO, *dto.Error) {
	deck, ok := deckMap[id]
	if !ok {
		return nil, &dto.Error{ResponseCode: 401, Message: "could not find deck with the given id"}
	}

	if len(deck.Cards) < count {
		return nil, &dto.Error{ResponseCode: 401, Message: "cannot draw cards as the draw value is greater than cards available"}
	}

	cards, deckCards := manager.DrawRandomCards(deck.Cards, count)
	deck.Cards = deckCards
	return &dto.DrawCardDTO{Cards: cards}, nil
}
