package api

type Response struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

func (o *Response) OK(data interface{}) {
	o.Result = "OK"
	o.Data = data
}

func (o *Response) Error(result string, data interface{}) {
	o.Result = result
	o.Data = data
}
