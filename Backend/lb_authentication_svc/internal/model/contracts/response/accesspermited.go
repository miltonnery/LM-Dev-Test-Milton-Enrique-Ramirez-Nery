package response

type AccessPermitted struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Tkn     string `json:"tkn"`
}
