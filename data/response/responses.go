package response

type Response struct {
	// StatusCode  int        `json:"statusCode"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
