package entity

func NewSuccessResponse(data interface{}) BaseSuccess {
	return BaseSuccess{
		Code:    "00",
		Data:    data,
		Message: "success",
	}
}
