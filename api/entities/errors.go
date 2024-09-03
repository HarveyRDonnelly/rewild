package entities

import (
	"fmt"
)

type Error struct {
	StatusCode  int    `json:"status_code"`
	ErrorCode   int    `json:"error_code"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func EntityNotExistsException(entityName string) Error {
	return Error{
		StatusCode: 400,
		ErrorCode: 2000,
		Title: fmt.Sprintf("Entity '%s' does not exist", entityName),
		Description: "An entity of type '%s' in the request does not exist.",
	}
}

func EntityNotFoundException(entityName string) Error {
	return Error{
		StatusCode: 404,
		ErrorCode: 2001,
		Title: fmt.Sprintf("Entity '%s' cannot be found", entityName),
		Description: fmt.Sprintf("An entity of type '%s' in the request could not be found.", entityName),
	}
}

func RequestBodyCastFailureException() Error {
	return Error{
		StatusCode: 400,
		ErrorCode: 2002,
		Title: fmt.Sprintf("Request body cast failed", entityName),
		Description: fmt.Sprintf("The request body could not be cast to the request variable schema.", entityName),
	}
}