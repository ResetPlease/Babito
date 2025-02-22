package models

import "github.com/gin-gonic/gin"

const AuthHeaderKey = "Authorization"
const ParamItemName = "item"

var ErrorBadRequest = ErrorResponse{
	Errors: "Bad request",
}

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

var ErrorUserNotExist = ErrorResponse{
	Errors: "User not exist",
}

var ErrorNotEnoughtFunds = ErrorResponse{
	Errors: "Not enought funds",
}

var ErrorProductNotFound = ErrorResponse{
	Errors: "Product not found",
}

var MessageOK = gin.H{
	"Message": "OK",
}
