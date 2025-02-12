package models

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
