package handlers

import "errors"

var (
	ErrMissingParameters = errors.New("missing parameters")
	ErrLotNumberNotFound = errors.New("lot number not found")
)
