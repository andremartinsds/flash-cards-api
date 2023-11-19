package handler

import (
	"github.com/andremartinsds/flash-cards-api/schemas"
)

type DecCreateResponse struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}
type DecReadResponse struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}
type DecListResponse struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}

type DecDelete struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	UserId uint   `json:"userId"`
}

type ErrorDecResponse struct {
	Message string `json:"message"`
	Code    uint   `json:"code"`
}

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

func fromSaveDecToResponse(d *schemas.Dec) DecCreateResponse {
	return DecCreateResponse{
		Id:     d.ID,
		Name:   d.Name,
		UserId: d.UserID,
	}
}

func fromListDecToResponse(d []schemas.Dec) []DecListResponse {
	var decListResponseCollection []DecListResponse
	for i := 0; i < len(d); i++ {
		decListResponse := DecListResponse{
			Id:     d[i].ID,
			Name:   d[i].Name,
			UserId: d[i].ID,
		}
		decListResponseCollection = append(decListResponseCollection, decListResponse)
	}
	return decListResponseCollection
}

func fromReadDecToResponse(d *schemas.Dec) DecReadResponse {
	return DecReadResponse{
		Id:     d.ID,
		Name:   d.Name,
		UserId: d.ID,
	}
}

func fromDeleteDecToResponse(d *schemas.Dec) DecDelete {
	return DecDelete{
		Id:     d.ID,
		Name:   d.Name,
		UserId: d.ID,
	}
}
