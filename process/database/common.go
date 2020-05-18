package database

import (
	"encoding/json"
	"fmt"

	"github.com/ElrondNetwork/elrond-go/core/indexer"
)

func convertObjectToBlock(obj object) (*indexer.Block, string, error) {
	h1 := obj["hits"].(object)["hits"].([]interface{})
	if len(h1) == 0 {
		return nil, "", errCannotFindBlockInDb
	}
	h2 := h1[0].(object)["_source"]

	h3 := h1[0].(object)["_id"]
	blockHash := fmt.Sprint(h3)

	marshalizedBlock, _ := json.Marshal(h2)
	var block indexer.Block
	err := json.Unmarshal(marshalizedBlock, &block)
	if err != nil {
		return nil, "", errCannotUnmarshalBlock
	}

	return &block, blockHash, nil
}

func convertObjectToTransactions(obj object) ([]indexer.Transaction, error) {
	hits, ok := obj["hits"].(object)
	if !ok {
		return nil, errCannotGetTxsFromBody
	}

	txs := make([]indexer.Transaction, 0)
	for _, h1 := range hits["hits"].([]interface{}) {
		h2 := h1.(object)["_source"]

		var tx indexer.Transaction
		marshalizedTx, _ := json.Marshal(h2)
		err := json.Unmarshal(marshalizedTx, &tx)
		if err != nil {
			continue
		}

		h3 := h1.(object)["_id"]
		txHash := fmt.Sprint(h3)
		tx.Hash = txHash

		txs = append(txs, tx)
	}
	return txs, nil
}
