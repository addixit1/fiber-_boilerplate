package config

import (
	"github.com/addixit1/fiber-boilerplate/internal/lib/locale"
)

// APIResponse represents a standard API response
type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Type       string      `json:"type"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// ListResponse represents a list response with pagination
type ListResponse struct {
	StatusCode int         `json:"statusCode"`
	Type       string      `json:"type"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Total      int64       `json:"total,omitempty"`
	Page       int         `json:"page,omitempty"`
	Limit      int         `json:"limit,omitempty"`
}

// buildResponse creates a response with localized message
func buildResponse(statusCode int, responseType string, data interface{}, lang string) APIResponse {
	return APIResponse{
		StatusCode: statusCode,
		Type:       responseType,
		Message:    locale.Get(lang, responseType),
		Data:       data,
	}
}

// ========================
// SUCCESS RESPONSES
// ========================

// Default success response
func Default(lang string) APIResponse {
	return buildResponse(OK, TYPE_DEFAULT, nil, lang)
}

// Details returns success with data
func Details(data interface{}, lang string) APIResponse {
	return buildResponse(OK, TYPE_DEFAULT, data, lang)
}

// List returns list response with data
func List(data interface{}, lang string) APIResponse {
	return buildResponse(OK, TYPE_DEFAULT, data, lang)
}

// ListWithPagination returns list with pagination info
func ListWithPagination(data interface{}, total int64, page, limit int, lang string) ListResponse {
	return ListResponse{
		StatusCode: OK,
		Type:       TYPE_DEFAULT,
		Message:    locale.Get(lang, TYPE_DEFAULT),
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
	}
}

// Login success response
func Login(data interface{}, lang string) APIResponse {
	return buildResponse(OK, TYPE_LOGIN, data, lang)
}

// Signup success response
func Signup(data interface{}, lang string) APIResponse {
	return buildResponse(CREATED, TYPE_SIGNUP, data, lang)
}

// Profile returns profile data
func Profile(data interface{}, lang string) APIResponse {
	return buildResponse(OK, TYPE_PROFILE, data, lang)
}

// UpdateProfile success
func UpdateProfile(lang string) APIResponse {
	return buildResponse(UPDATED, TYPE_UPDATE_PROFILE, nil, lang)
}

// Logout success
func Logout(lang string) APIResponse {
	return buildResponse(OK, TYPE_LOGOUT, nil, lang)
}

// ChangePassword success
func ChangePassword(lang string) APIResponse {
	return buildResponse(OK, TYPE_CHANGE_PASSWORD, nil, lang)
}

// ResetPassword success
func ResetPassword(lang string) APIResponse {
	return buildResponse(OK, TYPE_RESET_PASSWORD, nil, lang)
}

// SendOTP success
func SendOTP(lang string) APIResponse {
	return buildResponse(OK, TYPE_SEND_OTP, nil, lang)
}

// VerifyOTP success with data
func VerifyOTP(data interface{}, lang string) APIResponse {
	return buildResponse(OK, TYPE_VERIFY_OTP, data, lang)
}

// ========================
// ERROR RESPONSES
// ========================

// Error returns custom error response
func Error(message string, lang string, statusCode ...int) APIResponse {
	code := BAD_REQUEST
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	return APIResponse{
		StatusCode: code,
		Type:       TYPE_ERROR,
		Message:    message,
	}
}

// UnauthorizedAccess error
func UnauthorizedAccess(lang string) APIResponse {
	return buildResponse(UNAUTHORIZED, TYPE_UNAUTHORIZED_ACCESS, nil, lang)
}

// InternalServerError error
func InternalServerError(lang string) APIResponse {
	return buildResponse(INTERNAL_SERVER_ERROR, TYPE_INTERNAL_SERVER_ERROR_TYPE, nil, lang)
}

// BadToken error
func BadToken(lang string) APIResponse {
	return buildResponse(UNAUTHORIZED, TYPE_BAD_TOKEN, nil, lang)
}

// TokenExpired error
func TokenExpired(lang string) APIResponse {
	return buildResponse(UNAUTHORIZED, TYPE_TOKEN_EXPIRED, nil, lang)
}

// SessionExpired error
func SessionExpired(lang string) APIResponse {
	return buildResponse(UNAUTHORIZED, TYPE_SESSION_EXPIRED, nil, lang)
}

// IncorrectPassword error
func IncorrectPassword(lang string) APIResponse {
	return buildResponse(BAD_REQUEST, TYPE_INCORRECT_PASSWORD, nil, lang)
}

// EmailNotRegistered error
func EmailNotRegistered(lang string) APIResponse {
	return buildResponse(BAD_REQUEST, TYPE_EMAIL_NOT_REGISTERED, nil, lang)
}

// EmailAlreadyExists error
func EmailAlreadyExists(lang string) APIResponse {
	return buildResponse(BAD_REQUEST, TYPE_EMAIL_ALREADY_EXIST, nil, lang)
}

// EmailNotVerified error
func EmailNotVerified(lang string, statusCode ...int) APIResponse {
	code := BAD_REQUEST
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	return buildResponse(code, TYPE_EMAIL_NOT_VERIFIED, nil, lang)
}

// UserNotFound error
func UserNotFound(lang string) APIResponse {
	return buildResponse(BAD_REQUEST, TYPE_USER_NOT_FOUND, nil, lang)
}

// InvalidOTP error
func InvalidOTP(lang string) APIResponse {
	return buildResponse(BAD_REQUEST, TYPE_INVALID_OTP, nil, lang)
}

// OTPExpired error
func OTPExpired(lang string) APIResponse {
	return buildResponse(BAD_REQUEST, TYPE_OTP_EXPIRED, nil, lang)
}

// BlockedUser error
func BlockedUser(lang string) APIResponse {
	return buildResponse(ACCESS_FORBIDDEN, TYPE_BLOCKED, nil, lang)
}

// DeactivatedUser error
func DeactivatedUser(lang string) APIResponse {
	return buildResponse(ACCESS_FORBIDDEN, TYPE_DEACTIVATED, nil, lang)
}
