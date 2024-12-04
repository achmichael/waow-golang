package common

type ResponseSuccess struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
}

type Message struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
