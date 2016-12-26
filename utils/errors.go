package utils

const (
	// TypeValidationError is used when data cannot be validated
	TypeValidationError = "VALIDATION_ERROR"

	// TypeInternalServerError is used for unexpected internal error, such as no database access
	TypeInternalServerError = "INTERNAL_SERVER_ERROR"

	// TypeServiceUnavailableError is used a service is unavailable
	TypeServiceUnavailableError = "SERVICE_UNAVAILABLE_ERROR"

	// TypeAuthenticationError is used if authentication is not successful
	TypeAuthenticationError = "AUTHENTICATION_ERROR"

	// TypeAuthorizationError is used if authorization is not granted for that data
	TypeAuthorizationError = "AUTHORIZATION_ERROR"

	// TypeParameterError is used for wrong parameter value
	TypeParameterError = "PARAMETER_ERROR"

	// TypeNotFoundError is used for not found errors
	TypeNotFoundError = "NOT_FOUND"
)

// Errors struct for API error
type Errors struct {
	Message string
	Type    string
	Tip     string
}

var (
	//  --------------------- 400 ---------------------------------------

	// MsgBadParameter is the error definition for an invalid paramater
	MsgBadParameter = Errors{
		Type:    TypeValidationError,
		Message: "A least one paramater is invalid",
		Tip:     "Please check your parameters or payload",
	}

	//  --------------------- 403 ---------------------------------------

	// MsgTokenMalformed is the error definition for malformed token
	MsgTokenMalformed = Errors{
		Type:    TypeAuthorizationError,
		Message: "Authorization header is not valid",
		Tip:     "Make sure that the header has a valid format",
	}

	// MsgTokenIsUnknown token can't be found in the database
	MsgTokenIsUnknown = Errors{
		Type:    TypeAuthenticationError,
		Message: "Token is unknown",
		Tip:     "Make sure that the token is correct",
	}

	// MsgMissingAuth is the error definition for the missing token
	MsgMissingAuth = Errors{
		Type:    TypeAuthenticationError,
		Message: "Missing Authorization",
		Tip:     "Add Authorization Header with the token to your request",
	}

	// MsgTokenIsRevoked is the error definition for revoked token
	MsgTokenIsRevoked = Errors{
		Type:    TypeAuthorizationError,
		Message: "Token has been revoked",
		Tip:     "Create a new access token and then retry",
	}

	// MsgTokenHasExpired is the error definition for expired token
	MsgTokenHasExpired = Errors{
		Type:    TypeAuthorizationError,
		Message: "Token has expired",
		Tip:     "Refresh the access token and then retry",
	}

	//  --------------------- 404 ---------------------------------------

	// MsgEndpointDoesNotExist is the error definition for user not found
	MsgEndpointDoesNotExist = Errors{
		Type:    TypeNotFoundError,
		Message: "Endpoint doesn't exist",
		Tip:     "Please check your url",
	}

	// MsgEntityDoesNotExist is the error definition for Entity not found
	MsgEntityDoesNotExist = Errors{
		Type:    TypeNotFoundError,
		Message: "Entity doesn't exist",
		Tip:     "Please check the id",
	}

	//  --------------------- 500 ---------------------------------------

	// MsgInternalServerError is the error definition for internal server error
	MsgInternalServerError = Errors{
		Type:    TypeInternalServerError,
		Message: "Error occured on server",
		Tip:     "Please light a candle and pray for our devops",
	}

	// MsgServiceUnavailable is the error definition for service unavailable
	MsgServiceUnavailable = Errors{
		Type:    TypeServiceUnavailableError,
		Message: "Service is currently unavailable",
		Tip:     "Please light a candle and pray for our devops",
	}
)
