package manager

import (
	"Toggl/dto"
	"math/rand"
)

var values = []string{"2", "3", "4", "5", "6", "7",
	"8", "9", "TEN", "JACK", "QUEEN", "KING", "ACE"}

var valuesMap = map[string]string{
	"2": values[0],
	"3": values[1],
	"4": values[2],
	"5": values[3],
	"6": values[4],
	"7": values[5],
	"8": values[6],
	"9": values[7],
	"T": values[8],
	"J": values[9],
	"Q": values[10],
	"K": values[11],
	"A": values[12],
}

var suits = []string{"HEARTS", "DIAMONDS", "CLUBS", "SPADES"}
var suitsKMap = map[string]string{
	"H": suits[0],
	"D": suits[1],
	"C": suits[2],
	"S": suits[3],
}

// ValidateAndGetCards validates and returns the values for the cards
func ValidateAndGetCards(cards []string) ([]string, *dto.Error) {
	if cards == nil || len(cards) == 0 {
		return defaultCards()
	} else {
		values := make([]string, 0)
		for i := 0; i < len(cards); i++ {
			suit, ok := suitsKMap[string(cards[i][1])]
			if !ok {
				return nil, &dto.Error{ResponseCode: 401, Message: "Invalid suit provided " + suit}
			}
			values = append(values, cards[i])
		}
		return values, nil
	}
}

func defaultCards() ([]string, *dto.Error) {
	defaultValues := make([]string, 0)
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			value := string(values[j][0]) + string(suits[i][0])
			defaultValues = append(defaultValues, value)
		}
	}
	return defaultValues, nil
}

// CreateAndShuffle creates a deck of card from the values provided,
// shuffles the cards if the input is set to true
func CreateAndShuffle(shuffled bool, values []string) []dto.Card {
	var cards []dto.Card
	for i := 0; i < len(values); i++ {
		code := string(values[i][1])
		card := dto.Card{
			Value: valuesMap[string(values[i][0])],
			Suit:  suitsKMap[code],
			Code:  string(values[i][0]) + string(suitsKMap[code][0]),
		}
		cards = append(cards, card)

	}
	if shuffled {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}
	return cards
}

// DrawRandomCards Draws random cards from the deck and returns the updates cards slice
func DrawRandomCards(cards []dto.Card, count int) ([]dto.Card, []dto.Card) {
	drawCards := make([]dto.Card, 0)
	for i := 0; i < count; i++ {
		index := rand.Intn(len(cards))
		drawCards = append(drawCards, cards[index])
		cards = removeIndex(cards, index)
	}
	return drawCards, cards
}

func removeIndex(s []dto.Card, index int) []dto.Card {
	ret := make([]dto.Card, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
