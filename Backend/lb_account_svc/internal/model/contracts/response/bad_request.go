package response

type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
