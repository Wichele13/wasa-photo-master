package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

// FUNZIONE CHE AGGIUNGE IL LIKE
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like
	var user User
	username := ps.ByName("username")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	photoid, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	likeid, err := strconv.ParseUint(ps.ByName("likeId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	like.LikeId = likeid
	token := getToken(r.Header.Get("Authorization"))
	like.UserIdentifier = token
	like.PhotoIdentifier = photoid
	like.PhotoOwner = user.Id
	dblike, err := rt.db.SetLike(like.LikeToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	like.LikeFromDatabase(dblike)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(like)
}

// FUNZIONE CHE RIMUOVE IL LIKE
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var like Like
	username := ps.ByName("username")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	photoid, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	likeid, err := strconv.ParseUint(ps.ByName("likeId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	like.LikeId = likeid
	token := getToken(r.Header.Get("Authorization"))
	like.UserIdentifier = token
	like.PhotoIdentifier = photoid
	like.PhotoOwner = user.Id
	err = rt.db.RemoveLike(like.LikeToDatabase())
	if errors.Is(err, database.ErrLikeDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", likeid).Error("Impossibile rimuovere il Mi Piace")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// FUNZIONE CHE RESTITUISCE TUTTI I LIKE DI UNA FOTO
func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photo Photo
	var requestUser User
	token := getToken(r.Header.Get("Authorization"))
	requestUser.Id = token
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requestUser.FromDatabase(dbrequestuser)
	username := ps.ByName("username")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	photoid, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.Id = photoid
	dbphoto, err := rt.db.CheckPhoto(photo.PhotoToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.PhotoFromDatabase(dbphoto)
	like, err := rt.db.GetLike(photo.Id, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(like)
}
