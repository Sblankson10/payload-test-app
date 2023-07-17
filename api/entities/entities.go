package entities

type CreateProvider struct {
	DepositsId  int64
	ProviderRef string
}

type Provider struct {
	DepositsId  int64  `json:"deposits_id"`
	ProviderRef string `json:"provider_ref"`
	Created     string `json:"created_at"`
	Modified    string `json:"modified_at"`
}

type IncomingPayload struct {
	ProfileID         int64
	Msisdn            int64
	Amount            float32
	ReferenceID       int64
	Reference         string
	Name              string
	ProviderRef       string
	TransactionTypeId int64
	DepositsId        int64
}
