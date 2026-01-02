package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"log"
)

func PostFile(
	name string,
	url string,
	user string,
	password string,
) {
	f := system.Open(name)
	defer func() {
		fmt.Println("Close file")
		errors.PanicClose(f)
	}()

	r := NewPostBytes(url, f)
	r.SetBasicAuth(user, password)
	r.Header.Set(constant.ContentType, constant.Form)
	r.ContentLength = system.FileStat(f).Size()
	s := Send(Client(), r)
	defer func() {
		fmt.Println("Close response body")
		errors.PanicClose(s.Body)
	}()

	if !ResponseOkay(s) {
		log.Panicf("upload %d: %s", s.StatusCode, ReadString(s))
	}
}
