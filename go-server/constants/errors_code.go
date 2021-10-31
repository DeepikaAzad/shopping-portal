package constants

var ErrorCode = struct {
	INTERNAL_SERVER_ERROR string
	INVALID_REQUEST_ERROR string
	DUPLICATE_ERROR       string
	INVALID_TOKEN         string
	INVALID_USER          string
	NOT_FOUND             string
}{
	INTERNAL_SERVER_ERROR: "internal_server_error",
	INVALID_REQUEST_ERROR: "invalid_request",
	DUPLICATE_ERROR:       "duplicate_entity",
	INVALID_TOKEN:         "invalid_token",
	INVALID_USER:          "invalid_user",
	NOT_FOUND:             "not_found",
}
