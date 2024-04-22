package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func (rt *_router) Close() error {
	return nil
}


func (rt *_router) liveness(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if err := rt.db.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
