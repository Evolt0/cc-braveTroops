package proto

type PutState struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type GetState struct {
	Key string `json:"key,omitempty"`
}