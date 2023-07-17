package entities

type CreateProvider struct {
	DepositsId  int64
	ProviderRef string
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
