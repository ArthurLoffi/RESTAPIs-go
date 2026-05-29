package error

type ErrorSyntax struct {
	ErrorCode int `json:"Status"`
	ErrorDescription string `json:"Description"`
}

func FormatedError(httpStatus int, err string) ErrorSyntax {
	return ErrorSyntax{
        ErrorCode: httpStatus,
        ErrorDescription: err,
    }
}