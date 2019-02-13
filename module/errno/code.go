package errno

var (
	OK                  = &Errno{Code: 0, Message: "OK."}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrDataExist  = &Errno{Code: 20003, Message: "Data existed."}

	ErrRecordNotFound  = &Errno{Code: 20101, Message: "The record was not found."}
	ErrUserAuthFailed  = &Errno{Code: 20102, Message: "Username does not exist or password was wrong."}
	ErrTokenAuthFailed = &Errno{Code: 20103, Message: "Fail to authorize your token."}
	ErrLogoutFailed    = &Errno{Code: 200104, Message: "Fail to logout."}
)
