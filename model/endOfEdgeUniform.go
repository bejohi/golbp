package model

// The lbp uniform for EndOfEdges.
type EndOfEdgeUniform struct {
}

// LbpUniform returns an array of 8 bytes which represents the lbp uniform "End of Edge",
// where 2 1's are direct neighbours.
// 0b00000011 == 3
// 0b00000110 == 6
// 0b00001100 == 12
// 0b00011000 == 24
// 0b00110000 == 48
// 0b01100000 == 96
// 0b11000000 == 192
// 0b10000001 == 65
func (uniform EndOfEdgeUniform)GetUniformList() *[]byte{
	byteArr := []byte{
		3,6,12,24,48,96,192,65,
	}
	return &byteArr
}

// GetName returns a spellable name of the EndOfEdgeUniform
func (uniform EndOfEdgeUniform)GetName() string {
	return "End of Edge Uniform"
}

// IsByteInUniform checks whether the given byte is in the uniform list, or not.
func (uniform EndOfEdgeUniform)IsByteInUniform(toProof byte) bool{
	for _, uniform := range *(uniform.GetUniformList()) {
		if uniform == toProof{
			return true
		}
	}

	return false
}