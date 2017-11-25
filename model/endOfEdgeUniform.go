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
func (uniform EndOfEdgeUniform)LbpUniform() []byte{
	byteArr := []byte{
		3,6,12,24,48,96,192,65,
	}
	return byteArr
}

func(uniform EndOfEdgeUniform)GetName() string {
	return "End of Edge Uniform"
}