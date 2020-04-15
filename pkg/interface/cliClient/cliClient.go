package cliClient

import (
	"phone_activation/pkg/model"
	"phone_activation/pkg/usecase"

	"github.com/pkg/errors"
)

type Handler struct {
	listActivationDates usecase.IListActivationDatesUsecase
}

func NewHandler(listActivationDates usecase.IListActivationDatesUsecase) *Handler {
	return &Handler{listActivationDates}
}

func (h *Handler) GetFirstActivationDates() ([]model.Phone, error) {
	results, err := h.listActivationDates.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return results, nil
}
