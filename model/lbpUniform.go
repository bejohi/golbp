package model

// LbpUniform provides methods to work with lbp uniform patterns.
type LbpUniform interface {
	GetUniformList() *[]byte
	IsByteInUniform(toProof byte) bool
	GetName() string
}
