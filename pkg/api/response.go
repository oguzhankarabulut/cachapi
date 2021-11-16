package api

type singleResponse struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

func NewSingleResponse(k string, v interface{}) *singleResponse {
	return &singleResponse{
		Key: k,
		Value: v,
	}
}


