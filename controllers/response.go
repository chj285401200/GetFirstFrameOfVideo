package controllers

type Response struct {
	Code        int        `json:"code" description:"return code"`
	Msg         string     `json:"msg" description:"return message"`
	Data        ResultType `json:"data" description:"return data"`
	Data2       ResultType `json:"data2" description:"return data2"`
}

type ResultType interface{}

func NewResponse() *Response {
	return &Response{200, "",nil, nil}
}
