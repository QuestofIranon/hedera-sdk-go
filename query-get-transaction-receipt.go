package hedera

// #include "hedera-query-get-transaction-receipt.h"
import "C"

type QueryGetTransactionReceipt struct {
	Query
}

type TransactionReceipt struct {
	// TODO: Make enum for [Status]
	Status    uint8
	AccountID *AccountID
	// unsupported: ContractID *C.HederaContractId
	// unsupported: FileID *C.HederaFileId
}

func newQueryGetTransactionReceipt(client *Client, transactionID TransactionID) QueryGetTransactionReceipt {
	return QueryGetTransactionReceipt{
		Query{C.hedera_query__get_transaction_receipt__new(client.inner, transactionID.c())}}
}

func (query QueryGetTransactionReceipt) Answer() (TransactionReceipt, error) {
	var answer C.HederaQueryGetTransactionReceiptAnswer
	err := C.hedera_query__get_transaction_receipt__answer(query.inner, &answer)
	if err != 0 {
		return TransactionReceipt{}, hederaError(err)
	}

	receipt := TransactionReceipt{Status: uint8(answer.status)}

	if answer.account_id != nil {
		accountID := accountIDFromC(*answer.account_id)
		receipt.AccountID = &accountID
	}

	return receipt, nil
}
