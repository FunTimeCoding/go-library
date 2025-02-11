package gorilla

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"log"
	"net/http"
)

func echo(
	w http.ResponseWriter,
	r *http.Request,
) {
	c, upgradeFail := upgrader.Upgrade(w, r, nil)

	if upgradeFail != nil {
		log.Print("upgrade:", upgradeFail)

		return
	}

	defer errors.LogClose(c)

	for {
		messageType, message, readFail := c.ReadMessage()

		if readFail != nil {
			log.Println("read:", readFail)

			break
		}

		log.Printf("recv: %s", message)

		if e := c.WriteMessage(messageType, message); e != nil {
			log.Println("write:", e)

			break
		}
	}
}
