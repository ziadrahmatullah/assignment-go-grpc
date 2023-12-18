package apperror

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func (ce *CustomError) ToErrorRes() ErrorRes {
	return ErrorRes{
		Message: ce.Message,
	}
}

func (ce *CustomError) ToGrpcError() error {
	errorMapper := map[int]codes.Code{
		http.StatusInternalServerError: codes.Internal,
		http.StatusBadRequest:          codes.InvalidArgument,
		http.StatusUnauthorized:        codes.PermissionDenied,
	}
	return status.Error(errorMapper[ce.Code], ce.Message)
}

var (
	ErrFindUsersQuery       = NewCustomError(http.StatusInternalServerError, "find user query error")
	ErrFindUserDetailsQuery = NewCustomError(http.StatusInternalServerError, "find user details query error")
	ErrFindUserByIdQuery    = NewCustomError(http.StatusInternalServerError, "find user by id query error")
	ErrFindUserByEmail      = NewCustomError(http.StatusInternalServerError, "find user by email query error")
	ErrNewUserQuery         = NewCustomError(http.StatusInternalServerError, "new user query error")

	ErrUserNotFound           = NewCustomError(http.StatusBadRequest, "user not found")
	ErrEmailNotFound          = NewCustomError(http.StatusBadRequest, "email not found")
	ErrEmailALreadyUsed       = NewCustomError(http.StatusBadRequest, "email already used")
	ErrInvalidPasswordOrEmail = NewCustomError(http.StatusBadRequest, "invalid password or email")
	ErrInvalidEmail           = NewCustomError(http.StatusBadRequest, "invalid email")

	ErrFindWalletByIdQuery = NewCustomError(http.StatusInternalServerError, "find wallet by id query error")
	ErrNewWalletQuery      = NewCustomError(http.StatusInternalServerError, "new wallet query error")

	ErrInvalidPagination         = NewCustomError(http.StatusBadRequest, "invalid pagination")
	ErrFindListTransactionQuery  = NewCustomError(http.StatusInternalServerError, "find list transaction query error")
	ErrSortByTransactionQuery    = NewCustomError(http.StatusBadRequest, "wrong key for sorting")
	ErrSortTypeTrasacntionQueqry = NewCustomError(http.StatusBadRequest, "wrong sort type for sorting")
	ErrWrongStartDateFormat      = NewCustomError(http.StatusBadRequest, "wrong start date format")
	ErrWrongEndDateFormat        = NewCustomError(http.StatusBadRequest, "wrong end date format")

	ErrInvalidSourceOfFund = NewCustomError(http.StatusBadRequest, "invalid source of fund")
	ErrWalletNotFound      = NewCustomError(http.StatusBadRequest, "wallet not found")

	ErrFindBoxesQuery   = NewCustomError(http.StatusInternalServerError, "find boxes query error")
	ErrFindAttemptQuery = NewCustomError(http.StatusInternalServerError, "find attempt query error")
	ErrNewAttemptQuery  = NewCustomError(http.StatusInternalServerError, "new attempt query error")
	ErrFindBoxByIdQuery = NewCustomError(http.StatusInternalServerError, "find box by id query error")

	ErrNoAttemptLeft = NewCustomError(http.StatusBadRequest, "you have no attempt left")

	ErrBoxNotFound     = NewCustomError(http.StatusBadRequest, "box not found")
	ErrInvalidToken    = NewCustomError(http.StatusBadRequest, "invalid token")
	ErrApplyTokenQUery = NewCustomError(http.StatusInternalServerError, "")

	ErrCreateResetPassTokenQuery = NewCustomError(http.StatusInternalServerError, "create reset pass token query error")
	ErrTokenExpired              = NewCustomError(http.StatusBadRequest, "token expired")

	ErrGenerateHashPassword = NewCustomError(http.StatusInternalServerError, "couldn't generate hash password")
	ErrGenerateJWTToken     = NewCustomError(http.StatusInternalServerError, "can't generate jwt token")

	ErrPageNotFound             = NewCustomError(http.StatusBadRequest, "page not found")
	ErrInvalidSortFormat        = NewCustomError(http.StatusBadRequest, "invalid sorting format")
	ErrInvalidFilterFormat      = NewCustomError(http.StatusBadRequest, "invalid filter format")
	ErrCantTransferToYourWallet = NewCustomError(http.StatusBadRequest, "can't transfer to your wallet")
	ErrTxCommit                 = NewCustomError(http.StatusInternalServerError, "commit transaction error")
	ErrInsufficientBalance      = NewCustomError(http.StatusBadRequest, "insufficient balance")
	ErrInvalidWalletNumber      = NewCustomError(http.StatusBadRequest, "invalid wallet number")
	ErrInvalidAmount            = NewCustomError(http.StatusBadRequest, "invalid amount")
	ErrInvalidBody              = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrUnauthorize              = NewCustomError(http.StatusUnauthorized, "unauthorized")

	ErrInvalidAuthHeader = NewCustomError(http.StatusUnauthorized, "invalid auth header")
	ErrInvalidJWTToken   = NewCustomError(http.StatusUnauthorized, "invalid jwt token")

	ErrCantConnectDatabase       = NewCustomError(http.StatusInternalServerError, "can't connect database")
	ErrCantConnectToThirdParty   = NewCustomError(http.StatusInternalServerError, "can't connect to third party")
	ErrFailedToRequstThirdParty  = NewCustomError(http.StatusInternalServerError, "failed to request third party")
	ErrInvalidChildren           = NewCustomError(http.StatusBadRequest, "invalid children")
	ErrInvalidMonthlyIncome      = NewCustomError(http.StatusBadRequest, "invalid monthly income")
	ErrInvalidMonthlyExpense     = NewCustomError(http.StatusBadRequest, "invalid monthly expense")
	ErrInvalidFinancialAbilities = NewCustomError(http.StatusBadRequest, "invalid financial abilities")
	ErrInvalidMaritalStatus      = NewCustomError(http.StatusBadRequest, "invalid marital status")
)
