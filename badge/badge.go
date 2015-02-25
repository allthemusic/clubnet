package badge

import (
	b58 "github.com/jbenet/go-base58"
	"github.com/jbenet/go-ipfs/merkledag"
)

type Badge struct {
	Token string
}

func NewBadge(token string) Badge {
	return Badge{Token: token}
}

func (b *Badge) Hash() (h string, err error) {
	node := merkledag.Node{}
	node.Data = []byte(b.Token)
	multihash, err := node.Multihash()
	if err != nil {
		return
	}
	return b58.Encode(multihash), nil
}
