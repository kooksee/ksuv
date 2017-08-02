package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"net/http"
)

func TestAll(t *testing.T) {
	http.NewRequest()

	Convey("test /api/programs", t, func() {
		Convey("post /api/programs", func() {
			fmt.Println("ok")
		})
	})
}
