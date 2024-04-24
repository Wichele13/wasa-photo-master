package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// FUNZIONE CHE ESEGUE IL LOGIN
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dbuser, err := rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// FUNZIONE CHE SETTA IL NOME DELL'UTENTE
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	username := ps.ByName("username")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := getToken(r.Header.Get("Authorization"))
	user.Id = token
	dbuser, err := rt.db.SetUsername(user.ToDatabase(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// FUNZIONE CHE RESTITUISCE IL PROFILO DELL'UTENTE CHE RICHIEDE
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User
	var profile Profile
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
	profile.RequestId = token
	profile.Id = user.Id
	profile.Username = user.Username
	followersCount, err := rt.db.GetFollowersCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowersCount = followersCount
	followingCount, err := rt.db.GetFollowingsCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowingCount = followingCount
	photoCount, err := rt.db.GetPhotosCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.PhotoCount = photoCount
	profile.BanStatus, err = rt.db.GetBanStatus(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowStatus, err = rt.db.GetFollowStatus(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.CheckIfBanned, err = rt.db.CheckIfBanned(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(profile)
}

// FUNZIONE CHE RESTITUISCE LO STREAM DELL'UTENTE CHE RICHIED
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photoList database.Stream
	token := getToken(r.Header.Get("Authorization"))
	username := ps.ByName("username")
	user.Id = token
	user.Username = username
	// questo è l'utente che richiede lo stream
	dbuser, err := rt.db.CheckUser(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	photos, err := rt.db.GetMyStream(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photoList.Identifier = token
	photoList.Photos = photos
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photoList)
}
