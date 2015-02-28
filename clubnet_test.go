package clubnet

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMain(t *testing.T) {
	Convey("it's alive!!", t, func() {
		So(1, ShouldEqual, 1)
	})
}
