package main

type HttpResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cookie  string `json:"cookie"`
	Host    string `json:"host"`
}

func (h *HttpResult) Fail(status int, Message string) {
	h.Status = status
	h.Message = Message
}

func (h *HttpResult) Success(ck, host string) {
	h.Status = 200
	h.Message = "success"
	h.Cookie = ck
	h.Host = host
}
