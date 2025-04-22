package domain

type HttpResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Title    string      `json:"title"`
	Status   int         `json:"status"`
	Instance interface{} `json:"Instance,omitempty"`
}
