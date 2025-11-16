package RandomAPI

import (
	"math/rand"
	"net/http"
	"strconv"
)

type RandomAPI struct{}

func NewRandomAPI(router *http.ServeMux) {
	h := &RandomAPI{}
	router.HandleFunc("/random", h.Random())
}

func (h *RandomAPI) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		randomNumber := rand.Intn(6) + 1
		w.Write([]byte(strconv.Itoa(randomNumber)))
	}
}
