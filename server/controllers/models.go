package controllers

//go:generate ffjson $GOFILE
type ResultError struct {
	Error  struct {
		Description string `json:"description" example:"the request cannot be fulfilled because something bad happen"`
	} `json:"error"`
}