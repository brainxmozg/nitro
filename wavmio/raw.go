package wavmio

func getLastBlockHash(output []byte)
func readInboxMessage(offset uint32, output []byte) uint32
func advanceInboxMessage()
func resolvePreImage(hash []byte, offset uint32, output []byte) uint32
func setLastBlockHash([]byte)