package mountebank

import (
	. "github.com/durmaze/gobank/imposters"
	"github.com/parnurzeal/gorequest"
)

func CreateImposter(imposter Imposter) {
	gorequest.New().Post("http://localhost:2525/imposters").Send(imposter).End()
}

// func GetImposter(int port) {
// 	gorequest.New().Get("http://localhost:2525/imposters/4546").End()
// }