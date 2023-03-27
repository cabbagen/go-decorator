package provider

type MSCoreResponseType string;

type MSCoreResponseEnum struct {
	Code         int
	Message      string
}

type MSCoreResponse struct {
	Code         int             `json:"code"`
	Message      string          `json:"message"`
	Data         interface{}     `json:"data"`
}

var MSCoreResponseTypeMap = map[string]MSCoreResponseType{
	"FAILED": "FAILED",
	"SUCCESS": "SUCCESS",
	"FORBIDDEN": "FORBIDDEN",
	"UNAUTHORIZED": "UNAUTHORIZED",
	"VALIDATE_FAILED": "VALIDATE_FAILED",
}

var MSCoreResponseEnumMap = map[MSCoreResponseType]MSCoreResponseEnum {
	MSCoreResponseTypeMap["FORBIDDEN"]: MSCoreResponseEnum{100403, "没有相关权限"},
	MSCoreResponseTypeMap["UNAUTHORIZED"]: MSCoreResponseEnum {100401, "暂未登录或token已过期"},
	MSCoreResponseTypeMap["VALIDATE_FAILED"]: MSCoreResponseEnum {100400, "参数校验失败"},
	MSCoreResponseTypeMap["FAILED"]: MSCoreResponseEnum {100500, "接口操作失败"},
	MSCoreResponseTypeMap["SUCCESS"]: MSCoreResponseEnum {100200, "接口操作成功"},
}

func NewMSCoreResponse(responseType MSCoreResponseType, data interface{}, message string) MSCoreResponse {
	var response = MSCoreResponse{}

	response.Data = data
	response.Code = MSCoreResponseEnumMap[responseType].Code

	if len(message) > 0 {
		response.Message = message
	} else {
		response.Message = MSCoreResponseEnumMap[responseType].Message
	}
	return response
}
