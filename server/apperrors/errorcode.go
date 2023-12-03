package apperrors

type ErrCode string

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}

const (
	Unkwon           ErrCode = "UOOO"
	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataField  ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
)
