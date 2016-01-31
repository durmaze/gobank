package mountebank

import (
"strconv"
	. "github.com/durmaze/gobank/builders"
	"github.com/parnurzeal/gorequest"
)

func CreateImposter(imposter Imposter) {
	gorequest.New().Post("http://localhost:2525/imposters").Send(imposter).End()
}

func DeleteImposter(imposter Imposter) {
	imposterUri := "http://localhost:2525/imposters/" + strconv.Itoa(imposter.Port)

	gorequest.New().Delete(imposterUri).End()
}

func DeleteAllImposters() {
	imposterUri := "http://localhost:2525/imposters"

	gorequest.New().Delete(imposterUri).End()
}