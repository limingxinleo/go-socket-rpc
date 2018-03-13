package proto

type RequestACK struct {
	Service   string
	Method    string
	Arguments []interface{}
}

type ResponseACK struct {
	Success      bool
	Data         []interface{}
	ErrorCode    uint32
	ErrorMessage string
}

func (*RequestACK) JsonDecode() *RequestACK {

}
