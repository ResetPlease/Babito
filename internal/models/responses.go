package models

import "github.com/gin-gonic/gin"

var ErrorEmptyRequestBody = gin.H{
	"Error": "Empty request body",
}

var ErrorWrongDataFormat = gin.H{
	"Error": "Wrong data format",
}

var ErrorMissingRequiredField = gin.H{
	"Error": "Missing required field",
}

var ErrorInternalServerError = gin.H{
	"Error": "Internal server error",
}
