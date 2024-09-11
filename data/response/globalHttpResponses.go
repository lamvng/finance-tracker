package response

type Response struct {
	Code   uint        `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
