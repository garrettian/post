package shared

type Proof struct {
	Identity     []byte
	Challenge    Challenge
	MerkleRoot   []byte
	ProofNodes   [][]byte
	ProvenLeaves [][]byte
}