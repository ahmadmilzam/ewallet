package v1

type SuccessResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}
