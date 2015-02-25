package badge

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHash(t *testing.T) {
	Convey("Badge's hash is correct multihash for IPFS object", t, func() {
		badge := NewBadge("foobar")
		So(badge.Hash, ShouldEqual, "QmTn6AULKTorC6unNZniyQUuWHnDmFugTyCBTFNCAsiVUz")
	})
}
