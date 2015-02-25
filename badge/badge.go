package badge

import (
	b58 "github.com/jbenet/go-base58"
	"github.com/jbenet/go-ipfs/merkledag"
)

type Badge struct {
	Hash  string
	Token string
}

func NewBadge(token string) Badge {
	node := merkledag.Node{}
	node.Data = []byte(token)
	multihash, _ := node.Multihash()
	encoded := b58.Encode(multihash)

	return Badge{encoded, token}
}
