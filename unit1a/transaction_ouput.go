package unit1a

type TXOutput struct {
	value        int
	scriptPubkey string
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.scriptPubkey == unlockingData
}
