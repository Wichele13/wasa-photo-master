<script>
import ModaleCommenti from "../components/ModaleCommenti.vue";
import MessaggioErrore from '../components/MessaggioErrore.vue';

export default {
    components: { ModaleCommenti, MessaggioErrore},
    data: function () {
        return {
            errormsg: null,
            username: sessionStorage.getItem('username'),
            token: sessionStorage.getItem('token'),
            newUsername: "",
            images: null,
            profile: {
                requestId: 0,
                id: 0,
                username: "",
                followersCount: 0,
                followingCount: 0,
                photoCount: 0,
            },
            photoList: {
                requestUser: 0,
                identifier: 0,
                photos: [
                    {
                        id: 0,
                        userId: 0,
                        file: "",
                        date: "",
                        likesCount: 0,
                        commentsCount: 0,
                        likeStatus: null,
                        comment: "",
                    }
                ],
            },
            user: {
                id: 0,
                username: "",
            },
            comment: "",
            photoComments: {
                requestIdentifier: 0,
                photoIdentifier: 0,
                identifier: 0,
                comments: [
                    {
                        id: 0,
                        userId: 0,
                        photoId: 0,
                        photoOwner: 0,
                        ownerUsername: "",
                        username: "",
                        content: "",
                    }
                ],
            },
        }
    },
    methods: {

        async ViewSession() {
            this.$router.push({ path: '/session/'})
        },

        async refresh() {
            await this.userProfile();
            await this.userPhotos();
        },
        async uploadFile() {
			this.images = this.$refs.file.files[0]
		},

        async doLogout() {
			sessionStorage.removeItem("token")
			sessionStorage.removeItem("username")
			this.$router.push({ path: '/' })
		},
		async submitFile() {
			if (this.images === null) {
				this.errormsg = "Ricorda di selezionare prima una foto"
			} else {
				try {
					let response = await this.$axios.put("/users/" + this.username + "/photo/" + Math.floor(Math.random() * 10000), this.images, {
						headers: {
							Authorization: "Bearer " + sessionStorage.getItem("token")
						}
					})
					this.profile = response.data
					this.successmsg = "Foto caricata"
					this.refresh()

				} catch (e) {
					if (e.response && e.response.status === 400) {
						this.errormsg = "Errore nel caricamento della foto, riprova";
						this.detailedmsg = null;
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "Problemi con il server nel caricamento della foto, riprova più tardi";
						this.detailedmsg = e.toString();
					} else {
						this.errormsg = e.toString();
						this.detailedmsg = null;
					}
				}
			}
		},

        async userProfile() {
            try {
                let response = await this.$axios.get("/users/" + this.username + "/profile", {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.profile = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Errore nella richiesta, ricontrolla i dati inseriti e riprova";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Potrebbe esserci un problema con il server, riprova più tardi.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },

        async userPhotos() {
            try {
                let response = await this.$axios.get("/users/" + this.username + "/photo", {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.photoList = response.data
                for (let i = 0; i < this.photoList.photos.length; i++) {
                    this.photoList.photos[i].file = 'data:image/*;base64,' + this.photoList.photos[i].file
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Errore nella richiesta";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Errore interno al server. Riprova più tardi.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = "Nessuna foto trovata";
                    this.detailedmsg = null;
                }
            }
        },
        
        async deletePhoto(photoid) {
            try {
                let response = await this.$axios.delete("/users/" + this.username + "/photo/" + photoid, {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.refresh();
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Errore nella richiesta.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Errore interno al server. Riprova più tardi.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async changeName() {
            if (this.newUsername.length < 3){
                this.errormsg = "L'username deve essere lungo almeno 3 caratteri."
            } else if (this.newUsername === this.username) {
                this.errormsg = "Scegli uno username diverso da quello attuale."
            } else {
                try {
                    let response = await this.$axios.put("/user/" + this.username + "/setUsername", { username: this.newUsername }, {
                        headers: {
                            Authorization: "Bearer " + sessionStorage.getItem("token")
                        }
                    })
                    this.user = response.data
                    sessionStorage.setItem("username", this.user.username);
                    this.profile.username = this.user.username;
                    this.username = this.user.username;
                    this.$router.push({ path: '/users/' + this.user.username + '/profile' })
                    this.refresh()
                } catch (e) {
                    if (e.response && e.response.status === 400) {
                        this.errormsg = "Errore nella richiesta, ricontrolla i dati inseriti e riprova";
                        this.detailedmsg = null;
                    }else if (e.response && e.response.status === 500) {
                        this.errormsg = "Potrebbe esserci un problema con il server, o il nome utente non disponibile.";
                        this.detailedmsg = e.toString();
                    } else {
                        this.errormsg = e.toString();
                        this.detailedmsg = null;
                    }
                }
            }

        },
        async sendComment(username, photoid, comment) {
			if (comment.trim().length < 1) {
				this.errormsg = "Scrivi qualcosa prima di inviare il commento"
			} else {
				try {
					let response = await this.$axios.put("/users/" + username + "/photo/" + photoid + "/comment/" + Math.floor(Math.random() * 10000), { content: comment }, {
						headers: {
							Authorization: "Bearer " + sessionStorage.getItem("token")
						}
					})
					this.clear = response.data
					this.refresh()
                    window.location.reload();
				} catch (e) {
					if (e.response && e.response.status === 400) {
						this.errormsg = "Errore nell'invio del commento, riprova";
						this.detailedmsg = null;
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "Problemi con il server nell'invio del commento, riprova più tardi";
						this.detailedmsg = e.toString();
					} else {
						this.errormsg = e.toString();
						this.detailedmsg = null;
					}
				}
			}
		},
        async openLog(username, photoid) {
            try {
                let response = await this.$axios.get("/users/" + username + "/photo/" + photoid + "/comment", {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.photoComments = response.data;
                const modal = new bootstrap.Modal(document.getElementById('logviewer'));
                modal.show();
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Errore nella richiesta, ricontrolla i dati inseriti e riprova";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Potrebbe esserci un problema con il server, riprova più tardi.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async likePhoto(username, id) {
            try {
                let response = await this.$axios.put("/users/" + username + "/photo/" + id + "/like/" + Math.floor(Math.random() * 10000), {}, {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async deleteLike(username, id) {
            try {
                let response = await this.$axios.get("/users/" + username + "/photo/" + id + "/like", {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.like = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Errore nella richiesta";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Errore interno al server. Riprova più tardi.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

            try {
                let response = await this.$axios.delete("/users/" + username + "/photo/" + id + "/like/" + this.like.likeId, {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Errore nella richiesta";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Errore interno al server. Riprova più tardi.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
    },


    mounted() {
        this.userProfile()
        this.userPhotos()
    }
}
</script>

<template>
    <div class="container-fluid">
        <div class="row">
            <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar">
                <div class="position-sticky pt-3 sidebar-sticky">
                  <ul class="nav flex-column">
                    <li class="nav-item mb-2">
                      <button class="btn btn-primary w-100" @click="ViewSession">Home</button>
                    </li>
                    <li class="nav-item mb-2">
                      <input type="file" accept="image/*" class="btn btn-outline-primary w-100" @change="uploadFile" ref="file">
                    </li>
                    <li class="nav-item mb-2">
                      <button class="btn btn-success w-100" @click="submitFile">Carica</button>
                    </li>
                    <li class="nav-item">
                      <button class="btn btn-danger w-100" @click="doLogout">Logout</button>
                    </li>
                  </ul>
                </div>
              </nav>
            <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
                <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                    <h1 class="h2">Ecco il tuo profilo, {{ profile.username }}</h1>
                    <div class="p-4 text-black">
                        <div class="d-flex justify-content-end text-center py-1">
                            <div>
                                <p class="mb-1 h5">{{ profile.followersCount }}</p>
                                <p class="small text-muted mb-0">Seguaci</p>
                            </div>
                            <div class="px-3">
                                <p class="mb-1 h5">{{ profile.followingCount }}</p>
                                <p class="small text-muted mb-0">Seguiti</p>
                            </div>
                            <div>
                                <p class="mb-1 h5">{{ profile.photoCount }}</p>
                                <p class="small text-muted mb-0">Foto</p>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="input-group mb-3">
                    <input type="text" id="newUsername" v-model="newUsername" class="form-control"
                        placeholder="Cambia username..." aria-label="Recipient's username"
                        aria-describedby="basic-addon2">
                    <div class="input-group-append">
                        <button class="btn btn-success" type="button" @click="changeName">Conferma</button>
                </div>
                </div>
                <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
                <MessaggioErrore v-if="errormsg" :msg="errormsg"></MessaggioErrore>
                <ModaleCommenti id="logviewer" :log="photoComments" :token="token"></ModaleCommenti>
                <div class="row">
                    <div class="col-md-4" v-for="photo in photoList.photos" :key="photo.id">
                        <div class="card mb-4 shadow-sm">
                            <img class="card-img-top" :src="photo.file" alt="Card image cap">
                            <div class="card-body">
                                <RouterLink :to="'/users/' + profile.username + '/profile'" class="nav-link">
                                    <button type="button" class="btn btn-outline-primary">{{ profile.username }}</button>
                                </RouterLink>
                                <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
                                <div class="d-flex justify-content-between align-items-center">
                                    <p class="card-text">Likes : {{ photo.likesCount }}</p>
                                </div>
                                <div class="d-flex justify-content-between align-items-center">
                                    <p class="card-text">Commenti : {{ photo.commentsCount }}</p>
                                </div>
                                <p class="card-text">Caricata il : {{ photo.date }}</p>
                                <div class="input-group mb-3">
                                    <input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="Aggiungi un commento..."
                                        aria-describedby="basic-addon2">
                                    <div class="input-group-append">
                                        <button class="btn btn-primary" type="button" @click="sendComment(username, photo.id, photo.comment)">Invia</button>
                                    </div>
                                </div>
                                <div class="d-flex justify-content-between align-items-center">
                                    <div class="btn-group">
                                        <button type="button" class="btn btn-dark" @click="openLog(username, photo.id)">Commenti</button>
                                        <button type="button" v-if="photo.likeStatus == false" class="btn btn-primary" @click="likePhoto(username, photo.id)">Mi Piace</button>
                                        <button type="button" v-if="photo.likeStatus == true" class="btn btn-danger" @click="deleteLike(username, photo.id)">Non mi piace più</button>
                                        <button type="button" class="btn btn-sm btn btn-outline-danger" @click="deletePhoto(photo.id)">Elimina </button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </main>
        </div>
    </div>
</template>

<style>
</style>