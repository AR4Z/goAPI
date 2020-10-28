package models

type Transaction struct {
	Uid      string 			 `json:"uid,omitempty"`
	Id       string 			 `json:"id,omitempty"`
	BuyerId  string              `json:"buyer_id,omitempty"`
	Ip       string              `json:"ip,omitempty"`
	Device   string              `json:"device,omitempty"`
	Products []map[string]string `json:"products,omitempty"`
	DType    []string            `json:"dgraph.type,omitempty"`
}
