package badge

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBadge(t *testing.T) {
	Convey("Hash and token are correct multihash for IPFS object", t, func() {
		badge := NewBadge("foobar")
		So(badge.Hash, ShouldEqual, "QmTn6AULKTorC6unNZniyQUuWHnDmFugTyCBTFNCAsiVUz")
		So(badge.Token, ShouldEqual, "foobar")
	})

	Convey("Hash and token are correct multihash for IPFS object", t, func() {
		badge := NewBadge("badge")
		So(badge.Hash, ShouldEqual, "QmNjhzd7LRJMgfdejdhyQNRSTxmkQLiTjPUjr8BqDY1MPa")
		So(badge.Token, ShouldEqual, "badge")
	})
}
