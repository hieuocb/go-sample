package models

type ConsumeBalanceRequest struct {
	PushNotifyId      string `json:"pushNotifyId"`
	CoreBankRefNum    string `json:"coreBankRefNum"`
	BankAccount       string `json:"bankAccount"`
	TransactionAmount int64  `json:"transactionAmount"`
	Narrative         string `json:"narrative"`
	TransactionType   string `json:"transactionType"`
	TransactionTime   int64  `json:"transactionTime"`
	AdditionalField1  string `json:"additionalField1"`
	AdditionalField2  string `json:"additionalField2"`
	AdditionalField3  string `json:"additionalField3"`
	AdditionalField4  string `json:"additionalField4"`
}

type ConsumeBalanceResponse struct {
	Error            string `json:"error"`
	ErrorMsg         string `json:"errorMsg"`
	Details          string `json:"details"`
	RefTransactionId string `json:"refTransactionId"`
}
