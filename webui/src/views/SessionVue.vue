<script>
import ModaleCommenti from "../components/ModaleCommenti.vue";
import MessaggioSuccesso from "../components/MessaggioSuccesso.vue";
import MessaggioErrore from "../components/MessaggioErrore.vue";

export default {
	components: { ModaleCommenti, MessaggioSuccesso, MessaggioErrore },
	data: function () {
		return {
			errormsg: null,
			successmsg: null,
			detailedmsg: null,
			username: sessionStorage.getItem('username'),
			token: sessionStorage.getItem('token'),
			loading: false,
			some_data: null,
			images: null,
			image: null,
			clear: null,
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
			comment: "",
			stream: {
				identifier: 0,
				photoStream: [
					{
						id: 0,
						userId: 0,
						username: "",
						file: "",
						date: "",
						likeCount: 0,
						commentCount: 0,
						comment: "",
						likeStatus: null,
					}
				],
			},
			searchUserUsername: "",
			like: {
				likeId: 0,
				identifier: 0,
				photoIdentifier: 0,
				photoOwner: 0,
			},
			profile: {
				requestId: 0,
				id: 0,
				username: "",
				followersCount: 0,
				followingCount: 0,
				photoCount: 0,
				followStatus: null,
				banStatus: null,
			},
		}
	},
	methods: {
		async refresh() {
			this.getStream()
		},
		async uploadFile() {
			this.images = this.$refs.file.files[0]
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
		
		async getStream() {
			try {
				let response = await this.$axios.get("/user/" + this.username + "/stream", {
					headers: {
						Authorization: "Bearer " + sessionStorage.getItem("token")
					}
				})
				this.stream = response.data
				for (let i = 0; i < this.stream.photoStream.length; i++) {
					this.stream.photoStream[i].file = 'data:image/*;base64,' + this.stream.photoStream[i].file
				}
			} catch (e) {
				if (e.response && e.response.status === 400) {
					this.errormsg = "Problemi con l'ottenimento dello stream, riprova più tardi";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Problemi con il server nell'ottenimento dello stream, riprova più tardi";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = "Segui qualcuno per avere uno stream di foto!";
					this.detailedmsg = null;
				}
			}
		},

		async SearchUser() {
			if (this.searchUserUsername === this.username) {
				this.errormsg = "Non puoi cercare te stesso"
			} else if (this.searchUserUsername.length < 3) {
				this.errormsg = "Inserisci almeno 3 caratteri per cercare un utente"
			} else {
				try {
					let response = await this.$axios.get("users/" + this.searchUserUsername + "/profile", {
						headers: {
							Authorization: "Bearer " + sessionStorage.getItem("token")
						}
					})
					this.profile = response.data
					this.$router.push({ path: '/users/' + this.searchUserUsername + '/view' })
				} catch (e) {
					if (e.response && e.response.status === 400) {
						this.errormsg = "Problemi con la ricerca dell'utente";
						this.detailedmsg = null;
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "Questo utente non esiste in WASAPhoto";
						this.detailedmsg = e.toString();
					} else {
						this.errormsg = e.toString();
						this.detailedmsg = null;
					}
				}
			}
		},

		async sendComment(username, photoid, comment) {
			if (comment === "") {
				this.errormsg = "Scrivi qualcosa prima di inviare il commento!"
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
					this.errormsg = "Errore nell'ottenimento dei commenti";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Problemi con il server, riprova più tardi";
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
					this.errormsg = "Errore nel like della foto";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Problemi con il server, riprova più tardi";
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
					this.errormsg = "Errore nella cancellazione del like";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Problemi con il server, riprova più tardi";
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
					this.errormsg = "Errore nella rimozione del like";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Problemi con il server, riprova più tardi";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
		},
		async doLogout() {
			sessionStorage.removeItem("token")
			sessionStorage.removeItem("username")
			this.$router.push({ path: '/' })
		},
		async ViewProfile() {
			this.$router.push({ path: '/users/' + this.username + '/profile' })
		},
	},
	mounted() {
		this.getStream()
	}
}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar">
				<div class="position-sticky pt-3 sidebar-sticky">
					<ul class="nav flex-column">

						<div class="btn-group me-2" style="flex-direction: column;">

							<button class="btn btn-primary" type="button" @click="ViewProfile">Profilo</button>
							<input type="file" accept="image/*" class="btn btn-outline-primary"  @change="uploadFile" ref="file">
							<button class="btn btn-success" type="button" @click="submitFile">Carica</button>
							<button class="btn btn-danger" type="button" @click="doLogout">Logout</button>
						</div>
					</ul>
				</div>
			</nav>
			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
					<h1 class="h2">Bentornato, {{ this.username }}</h1>

				</div>
				<div class="input-group mb-3">
					<input type="text" id="searchUserUsername" v-model="searchUserUsername" class="form-control"
						placeholder="Cerca un altro utente..." aria-label="Recipient's username"
						aria-describedby="basic-addon2">
					<div class="input-group-append">
						<button class="btn btn-primary" type="button" @click="SearchUser">Cerca</button>
					</div>
				</div>
				<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
				<MessaggioErrore v-if="errormsg" :msg="errormsg"></MessaggioErrore>
				<MessaggioSuccesso v-if="successmsg" :msg="successmsg"></MessaggioSuccesso>
				<ModaleCommenti id="logviewer" :log="photoComments" :token="token"></ModaleCommenti>
				<div class="row">
					<div class="col-md-4" v-for="photo in stream.photoStream" :key="photo.id">
						<div class="card mb-4 shadow-sm">	
							<img class="card-img-top" :src="photo.file" alt="Card image cap">
							<div class="card-body">
								<RouterLink :to="'/users/' + photo.username + '/view'" class="nav-link">
									<button type="button" class="btn btn-outline-primary">{{ photo.username }}</button>
								</RouterLink>
								<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
								<div class="d-flex justify-content-between align-items-center">
									<p class="card-text">Mi Piace : {{ photo.likeCount }}</p>
								</div>
								<div class="d-flex justify-content-between align-items-center">
									<p class="card-text">Commenti : {{ photo.commentCount }}</p>
								</div>
								<p class="card-text">Caricata il : {{ photo.date }}</p>
								<div class="input-group mb-3">
									<input type="text" id="comment" v-model="photo.comment" class="form-control"
										placeholder="Aggiungi un commento..." aria-label="Recipient's username"
										aria-describedby="basic-addon2">
									<div class="input-group-append">
										<button class="btn btn-primary" type="button"
											@click="sendComment(photo.username, photo.id, photo.comment)">Invia</button>
									</div>
								</div>
								<div class="d-flex justify-content-between align-items-center">
									<div class="btn-group">
										<button type="button" class="btn btn-dark" @click="openLog(photo.username, photo.id)">Commenti</button>
										<button type="button" v-if="photo.likeStatus == false" class="btn btn-primary"
											@click="likePhoto(photo.username, photo.id)">Mi Piace</button>
										<button type="button" v-if="photo.likeStatus == true" class="btn btn-danger"
											@click="deleteLike(photo.username, photo.id)">Non mi piace più</button>
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
