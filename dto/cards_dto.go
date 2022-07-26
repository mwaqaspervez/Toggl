package dto

type CreateDeckDTO struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type OpenDeckDTO struct {
	ID        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

type DrawCardDTO struct {
	Cards []Card
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}
type Deck struct {
	ID       string `json:"id"`
	Cards    []Card `json:"cards"`
	Shuffled bool   `json:"shuffled"`
}

type Error struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
