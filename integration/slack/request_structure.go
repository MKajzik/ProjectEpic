package slack

//RequestStruct export
type RequestStruct struct {
	Blocks []block `json:"blocks"`
}

//AddItem export
func (request *RequestStruct) AddItem(block block) []block {
	request.Blocks = append(request.Blocks, block)
	return request.Blocks
}

//NewRequest export
func NewRequest() RequestStruct {
	return RequestStruct{
		Blocks: []block{},
	}
}
