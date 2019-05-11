package response

type AccessDenied struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
