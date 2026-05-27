package entity

//go:generate go tool oapi-codegen -config ../../api/generate-model.config.yaml -o model.gen.go ../../api/api.yaml
func NewSuccessResponse(data interface{}) BaseSuccess {
	return BaseSuccess{
		Code:    "00",
		Data:    data,
		Message: "success",
	}
}
