package unit1a

type TXInput struct {
	Txid      []byte //引用之前的交易ID
	Vout      int    //引用的交易输出的索引值
	ScriptSig string //
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
