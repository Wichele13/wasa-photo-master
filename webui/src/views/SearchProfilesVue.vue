<script>
import ModaleCommenti from "../components/ModaleCommenti.vue";
import MessaggioSuccesso from "../components/MessaggioSuccesso.vue";
import MessaggioErrore from "../components/MessaggioErrore.vue";

export default {
    components: { ModaleCommenti, MessaggioErrore, MessaggioSuccesso},
    data: function () {
        return {
            errormsg: null,
            successmsg: null,
            username: sessionStorage.getItem('username'),
            token: sessionStorage.getItem('token'),
            profile: {
                requestId: 0,
                id: 0,
                username: "",
                followersCount: 0,
                followingCount: 0,
                photoCount: 0,
                followStatus: null,
                banStatus: null,
                checkIfBanned: null,
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
                    }
                ],
            },
            comment: "",
            follow: {
                followId: 0,
                followedId: 0,
                userId: 0,
                banStatus: 0,
            },
            ban: {
                banId: 0,
                bannedId: 0,
                userId: 0,
            },
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
                        comment: "",
                        username: "",
                        content: "",
                    }
                ],
            },
        }
    },
    methods: {
        async refresh() {
            await this.userProfile()
            await this.userPhotos()
        },
        async ViewProfile() {
			this.$router.push({ path: '/users/' + this.username + '/profile' })
		},
        async ViewSession() {
            this.$router.push({ path: '/session/'})
        },
		async doLogout() {
			sessionStorage.removeItem("token")
			sessionStorage.removeItem("username")
			this.$router.push({ path: '/' })
		},

        async userProfile() {
            try {
                let response = await this.$axios.get("/users/" + this.$route.params.username + "/profile", {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.profile = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Profilo non trovato, controlla di aver inserito correttamente il nome utente";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async userPhotos() {
            try {
                let response = await this.$axios.get("/users/" + this.$route.params.username + "/photo", {
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
                    this.errormsg = "Problema nel caricamento delle foto, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = "Questo utente non ha ancora postato nulla"
                    this.detailedmsg = null;
                }
            }
        },
        async banUser(username) {
            try {
                let response = await this.$axios.put("/users/" + username + "/ban/" + Math.floor(Math.random() * 10000), {}, {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Problema nel bloccare l'utente, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async unbanUser(username) {
            try {
                let response = await this.$axios.get("/users/" + username + "/ban", {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.ban = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Problema nello sbloccare l'utente, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

            try {
                let response = await this.$axios.delete("/users/" + username + "/ban/" + this.ban.banId, {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Problema nello sbloccare l'utente, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async followUser(username) {


            try {
                
                let response = await this.$axios.put("/users/" + username + "/follow/" + Math.floor(Math.random() * 10000), {}, {
                    
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                        
                    }
                    
                })
                
                this.clear = response.data
                this.refresh()
                this.successmsg = "Utente" + username + "seguito"
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Problema nel seguire l'utente, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async unfollowUser(username) {
            try {
                let response = await this.$axios.get("/users/" + username + "/follow", {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.follow = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Problema nel non seguire più l'utente, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

            try {
                let response = await this.$axios.delete("/users/" + username + "/follow/" + this.follow.followId, {
                    headers: {
                        Authorization: "Bearer " + sessionStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Problema nel non seguire più l'utente, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

        },
        async sendComment(username, photoid, comment) {
            if (comment === "") {
                this.errormsg = "Aggiungi almeno un carattere al commento!"
            } else {
                try {
                    let response = await this.$axios.put("/users/" + username + "/photo/" + photoid + "/comment/" + Math.floor(Math.random() * 10000), { content: comment }, {
                        headers: {
                            Authorization: "Bearer " + sessionStorage.getItem("token")
                        }
                    })
                    this.clear = response.data
                    this.refresh()
                } catch (e) {
                    if (e.response && e.response.status === 400) {
                        this.errormsg = "";
                        this.detailedmsg = null;
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "Problema con il server, riprova più tardi";
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
                    this.errormsg = "Problema nel caricamento dei commenti, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
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
                    this.errormsg = "Problema nel mettere mi piace alla foto, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
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
                    this.errormsg = "Problema nel togliere il mi piace alla foto, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
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
                    this.errormsg = "Problema nel togliere il mi piace alla foto, riprova più tardi";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Problema con il server, riprova più tardi";
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
                        <button class="btn btn-primary" type="button" @click="ViewSession">Home</button>
                        <button class="btn btn-primary" type="button" @click="ViewProfile">Profilo</button>
                        <button class="btn btn-danger" type="button" @click="doLogout">Logout</button>


                    </ul>
                </div>
            </nav>
            <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
                <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom" v-if="profile.checkIfBanned">
                    <div class="alert alert-danger" role="alert">
                        <p>L'utente @{{ profile.username }} ti ha bloccato.</p>
                        <hr>
                        <p class="mb-0"></p>
                    </div>
                </div>

                <div>
                    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                        <h1 class="h2">Profilo di {{ profile.username }}</h1>
                        <div v-if="!profile.checkIfBanned" class="p-4 text-black">
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
                        <div class="form-group row">
                            <div class="col-md-6">
                                <button type="button" v-if="!profile.followStatus" class="btn btn-outline-primary" @click="followUser(profile.username)">Segui</button>
                                <button type="button" v-if="profile.followStatus" class="btn btn-primary" @click="unfollowUser(profile.username)">Non seguire più</button>
                            </div>
                            <div class="col-md-6">
                                <button type="button" v-if="!profile.banStatus" class="btn btn-outline-danger" @click="banUser(profile.username)">Blocca</button>
                                <button type="button" v-if="profile.banStatus" class="btn btn-outline-danger" @click="unbanUser(profile.username)">Sblocca</button>
                            </div>
                        </div>
                    </div>

                    <MessaggioSuccesso v-if="successmsg" :msg="successmsg"></MessaggioSuccesso>
                    <MessaggioErrore v-if="errormsg" :msg="errormsg"></MessaggioErrore>

                    <ModaleCommenti id="logviewer" :log="photoComments" :token="token"></ModaleCommenti>
                    <div v-if="!profile.checkIfBanned" class="row">
                        <div class="col-md-4" v-for="photo in photoList.photos" :key="photo.id">
                            <div class="card mb-4 shadow-sm">
                                <img class="card-img-top" :src="photo.file" alt="Card image cap">
                                <div class="card-body">
                                    <RouterLink :to="'/users/' + profile.username + '/view'" class="nav-link">
                                        <button type="button" class="btn btn-outline-primary">{{ profile.username }}</button>
                                    </RouterLink>
                                    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
                                    <div class="d-flex justify-content-between align-items-center">
                                        <p class="card-text">Mi Piace : {{ photo.likesCount }}</p>
                                    </div>
                                    <div class="d-flex justify-content-between align-items-center">
                                        <p class="card-text">Commenti : {{ photo.commentsCount }}</p>
                                    </div>
                                    <p class="card-text">Caricata il : {{ photo.date }}</p>
                                    <div class="input-group mb-3">
                                        <input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="Aggiungi un commento..." aria-label="Recipient's username" aria-describedby="basic-addon2">
                                        <div class="input-group-append">
                                            <button class="btn btn-primary" type="button" @click="sendComment(profile.username, photo.id, photo.comment)">Invia</button>
                                        </div>
                                    </div>
                                    <div class="d-flex justify-content-between align-items-center">
                                        <div class="btn-group">
                                            <button type="button" class="btn btn-dark" @click="openLog(profile.username, photo.id)">Commenti</button>
                                            <button type="button" v-if="!photo.likeStatus" class="btn btn-primary" @click="likePhoto(profile.username, photo.id)">Mi Piace</button>
                                            <button type="button" v-if="photo.likeStatus" class="btn btn-danger" @click="deleteLike(profile.username, photo.id)">Non mi piace più</button>
                                        </div>
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