package badge

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewBadge(t *testing.T) {
	Convey("NewBadge returns a new Badge with the Token field set", t, func() {
		badge := NewBadge("token")
		So(badge.Token, ShouldEqual, "token")
	})
}

func TestHash(t *testing.T) {
	var (
		badge Badge
		hash  string
		err   error
	)

	Convey("Hash is correct Base58-encoded SHA2-256 hash for IPFS object with no Links and where Data is the token", t, func() {
		testCases := []struct {
			token string
			hash  string
		}{
			{"foobar", "QmTn6AULKTorC6unNZniyQUuWHnDmFugTyCBTFNCAsiVUz"},
			{"badge", "QmNjhzd7LRJMgfdejdhyQNRSTxmkQLiTjPUjr8BqDY1MPa"},
			{"foo1", "QmUqAtDZnfyXDhnn8LyDYkfYm5hwiQz9ogqB5uEPzCqcYL"},
			{"foo2", "QmPT4eiQUvgHp9YTNWuzba623qjV82kPGMRqpWGc3UQ2Ws"},
			{"\n", "QmYLkUvueFBGBxttJJ3dLvKMFdgJh2ePei11Yq3w3EMfwj"},
			{"…_…Ú", "QmPjp94ggmrgSTaScPC5aNnT9SeurJ4D8iXUsCeTE8Ku9y"},
		}
		for _, testCase := range testCases {
			badge = NewBadge(testCase.token)
			hash, err = badge.Hash()
			So(err, ShouldEqual, nil)
			So(hash, ShouldEqual, testCase.hash)
		}
	})
}
