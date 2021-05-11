package errorcode

import "fmt"

type ErrorCodeStruct struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}
type GeneralErrStruct struct {
	ERR_400 ErrorCodeStruct
	ERR_401 ErrorCodeStruct
	ERR_403 ErrorCodeStruct
	ERR_404 ErrorCodeStruct
	ERR_500 ErrorCodeStruct
}

var (
	GeneralErr = GeneralErrStruct{
		ERR_400: ErrorCodeStruct{
			ErrorCode: "ERR_400",
			Message:   "Dữ liệu gửi không hợp lệ",
		},
		ERR_401: ErrorCodeStruct{
			ErrorCode: "ERR_401",
			Message:   "Hết hạn đăng nhập",
		},
		ERR_403: ErrorCodeStruct{
			ErrorCode: "ERR_403",
			Message:   "Bạn không có quyền truy cập việc này",
		},
		ERR_404: ErrorCodeStruct{
			ErrorCode: "ERR_404",
			Message:   "Yêu cầu dữ liệu không thấy",
		},
		ERR_500: ErrorCodeStruct{
			ErrorCode: "ERR_500",
			Message:   "Lỗi kết nối cơ sở dữ liệu",
		},
	}
)

func CustomErr(errorCode string, err error) ErrorCodeStruct {
	return ErrorCodeStruct{
		ErrorCode: errorCode,
		Message:   fmt.Sprintf("%s", err),
	}
}
