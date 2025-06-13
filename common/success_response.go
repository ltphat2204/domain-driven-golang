package common

type successResponse struct {
	Successs   bool        `json:"success"`
	Data       interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *successResponse {
	return &successResponse{
		Successs:   true,
		Data:       data,
	}
}

func NewSimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{
		Successs: true,
		Data:     data,
	}
}
