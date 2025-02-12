package models

import "github.com/gin-gonic/gin"

var ErrorEmptyRequestBody = ErrorResponse{
	Errors: "Empty request body",
}

var ErrorWrongDataFormat = ErrorResponse{
	Errors: "Wrong data format",
}

var ErrorMissingRequiredField = ErrorResponse{
	Errors: "Missing required field",
}

var ErrorInternalServerError = ErrorResponse{
	Errors: "Internal server error",
}

var ErrorUnauthorized = ErrorResponse{
	Errors: "Unauthorized",
}

var MessageOK = gin.H{
	"Message": "OK",
}
