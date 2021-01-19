package proto

type PutState struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type GetState struct {
	Key string `json:"key,omitempty"`
}

type GetStateRequest struct {
	Key string `json:"key,omitempty"`
}

type GetHistoryRequest struct {
	Key string `json:"key,omitempty"`
}

type GetHistoryResponse struct {
	KeyModifications []*KeyModification `json:"key_modifications,omitempty"`
}

type KeyModification struct {
	TxId      string `json:"tx_id,omitempty"`
	Value     string `json:"value,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	IsDelete  bool   `json:"is_delete,omitempty"`
}