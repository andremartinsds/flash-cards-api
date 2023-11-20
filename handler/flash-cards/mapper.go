package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

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

func fromSaveFlashCardToResponse(f *schemas.FlashCards) FlashCardCreateResponseData {
	return FlashCardCreateResponseData{
		Id:    f.ID,
		Front: f.Front,
		Back:  f.Back,
		DecId: f.DecID,
	}
}

func fromListFlashCardToResponse(f []schemas.FlashCards) []FlashCardListResponseData {
	var flashCardListResponseCollection []FlashCardListResponseData
	for i := 0; i < len(f); i++ {
		flashCardListResponse := FlashCardListResponseData{
			Id:    f[i].ID,
			Front: f[i].Front,
			Back:  f[i].Back,
			DecId: f[i].DecID,
		}
		flashCardListResponseCollection = append(flashCardListResponseCollection, flashCardListResponse)
	}
	return flashCardListResponseCollection
}

func fromReadFlashCardToResponse(f *schemas.FlashCards) FlashCardReadResponseData {
	return FlashCardReadResponseData{
		Id:    f.ID,
		Front: f.Front,
		Back:  f.Back,
		DecId: f.DecID,
	}
}

func fromDeleteFlashCardToResponse(f *schemas.FlashCards) FlashCardDeleteResponseData {
	return FlashCardDeleteResponseData{
		Id:    f.ID,
		Front: f.Front,
		Back:  f.Back,
		DecId: f.DecID,
	}
}
