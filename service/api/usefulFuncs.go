package api

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func getToken(message string) uint64 {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	stringToken := re.FindAllString(message, -1)
	token, _ := strconv.Atoi(stringToken[0])
	return uint64(token)
}

func (rt *_router) Close() error {
	return nil
}

func (rt *_router) liveness(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if err := rt.db.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
