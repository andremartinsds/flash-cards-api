package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

type FlashCardCreateResponse struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type FlashCardReadResponse struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type FlashCardListResponse struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}

type FlashCardDelete struct {
	Id    uint   `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}

func fromRequestFlashCardToFlashCardModel(f CreateFlashCardRequest) *schemas.FlashCards {
	return &schemas.FlashCards{
		Front: f.Front,
		Back:  f.Back,
		DecID: f.DecId,
	}
}

func fromUpdateRequestFlashCardToFlashCardModel(f UpdateFlashCardRequest, flashCard *schemas.FlashCards) *schemas.FlashCards {

	if f.Front != "" {
		flashCard.Front = f.Front
	}
	if f.Back != "" {
		flashCard.Back = f.Back
	}

	return flashCard
}

func fromSaveFlashCardToResponse(f *schemas.FlashCards) FlashCardCreateResponse {
	return FlashCardCreateResponse{
		Id:    f.ID,
		Front: f.Front,
		Back:  f.Back,
		DecId: f.DecID,
	}
}

func fromListFlashCardToResponse(f []schemas.FlashCards) []FlashCardListResponse {
	var flashCardListResponseCollection []FlashCardListResponse
	for i := 0; i < len(f); i++ {
		flashCardListResponse := FlashCardListResponse{
			Id:    f[i].ID,
			Front: f[i].Front,
			Back:  f[i].Back,
			DecId: f[i].DecID,
		}
		flashCardListResponseCollection = append(flashCardListResponseCollection, flashCardListResponse)
	}
	return flashCardListResponseCollection
}

func fromReadFlashCardToResponse(f *schemas.FlashCards) FlashCardReadResponse {
	return FlashCardReadResponse{
		Id:    f.ID,
		Front: f.Front,
		Back:  f.Back,
		DecId: f.DecID,
	}
}

func fromDeleteFlashCardToResponse(f *schemas.FlashCards) FlashCardDelete {
	return FlashCardDelete{
		Id:    f.ID,
		Front: f.Front,
		Back:  f.Back,
		DecId: f.DecID,
	}
}
