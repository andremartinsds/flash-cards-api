package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

func fromRequestDecToDec(d CreateDecRequest) *schemas.Dec {
	return &schemas.Dec{
		Name:   d.Name,
		UserID: d.UserID,
	}
}
func fromUpdateRequestDecToDec(d UpdateDecRequest, dec *schemas.Dec) *schemas.Dec {

	if d.Name != "" {
		dec.Name = d.Name
	}
	return dec
}

func fromSaveDecToResponse(d *schemas.Dec) CreateDecResponseData {
	return CreateDecResponseData{
		Id:     d.ID,
		Name:   d.Name,
		UserId: d.UserID,
	}
}

func fromListDecToResponse(d []schemas.Dec) []ListDecResponseData {
	var decListResponseCollection []ListDecResponseData
	for i := 0; i < len(d); i++ {
		decListResponse := ListDecResponseData{
			Id:     d[i].ID,
			Name:   d[i].Name,
			UserId: d[i].ID,
		}
		decListResponseCollection = append(decListResponseCollection, decListResponse)
	}
	return decListResponseCollection
}

func fromReadDecToResponse(d *schemas.Dec) DecReadResponseData {
	return DecReadResponseData{
		Id:     d.ID,
		Name:   d.Name,
		UserId: d.ID,
	}
}

func fromUpdateDecToResponse(d *schemas.Dec) DecUpdateResponseData {
	return DecUpdateResponseData{
		Id:     d.ID,
		Name:   d.Name,
		UserId: d.ID,
	}
}

func fromDeleteDecToResponse(d *schemas.Dec) DeleteDecResponseData {
	return DeleteDecResponseData{
		Id:     d.ID,
		Name:   d.Name,
		UserId: d.ID,
	}
}
