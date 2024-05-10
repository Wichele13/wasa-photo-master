<template>
  <div class="profile-container">
    <!-- Parte sinistra con le informazioni del profilo -->
    <div class="profile-info" :style="{ backgroundColor: leftColor }">
      <div class="profile-header">
        <h2>{{ username }}</h2>
      </div>
      <div class="profile-actions">
        <button @click="uploadPhoto">Carica Foto</button>
        <button @click="ChangeName">Rinomina Utente</button>
        <button @click="getMyStream">Mostra Stream</button>
      </div>
    </div>
    <!-- Parte destra con le foto del profilo -->
    <div class="profile-photos" :style="{ backgroundColor: rightColor }">
      <div class="photo-header">
        <h2>Le tue foto</h2>
      </div>
    </div>
    <!-- Modale -->
    <div class="modal-container" v-if="isModalOpen">
      <div class="modal-content">
        <span class="close" @click="closeModal">&times;</span>
        <h2>Cambia Username</h2>
        <form @submit.prevent="setMyUserName">
          <label for="newUsername">Nuovo Username:</label>
          <input type="text" id="newUsername" v-model="newUsername" required>
          <button type="submit">Salva</button>
        </form>
      </div>
    </div>
  </div>
</template>

  
    
<script>
export default {
    data: function () {
        return {
              
            leftColor: '#000000', // Colore bianco
            rightColor: '#006600', // Colore nero
            errormsg: null,
            isModalOpen: false,
            username: sessionStorage.getItem('username'),
            token: sessionStorage.getItem('token'),
            newUsername: "",
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

        async refresh() {
            await this.userProfile();
            await this.userPhotos();
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
                    this.errormsg = "Errore nei dati inseriti";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async uploadFile() {
          this.images = this.$refs.file.files[0]
        },
        async submitFile() {
          if (this.images === null) {
            this.errormsg = "Seleziona una foto da caricare"
          } else {
            try {
              let response = await this.$axios.put("/users/" + this.username + "/photo/" + Math.floor(Math.random() * 10000), this.images, {
                headers: {
                  Authorization: "Bearer " + sessionStorage.getItem("token")
                }
              })
              this.profile = response.data
              this.successmsg = "Foto caricata"
            } catch (e) {
              if (e.response && e.response.status === 400) {
                this.errormsg = "Errore nei dati inseriti";
                this.detailedmsg = null;
              } else if (e.response && e.response.status === 500) {
                this.errormsg = "Riprova più tardi";
                this.detailedmsg = e.toString();
              } else {
                this.errormsg = e.toString();
                this.detailedmsg = null;
              }
            }
          }
        },
            
        async setMyUserName() {

            if (this.newUsername.length <= 3){  
                this.errormsg = "Non lasciare il campo username vuoto!"
            } else {
                try {
                    let response = await this.$axios.put("/user/" + this.username + "/setusername", { username: this.newUsername }, {
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
                        this.errormsg = "Qualcosa è andato storto, per favore prova di nuovo controllando tutti i campi. Se il problema persiste, contattaci";
                        this.detailedmsg = null;
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "Qualcosa è andato storto :( ...riprova più tardi";
                        this.detailedmsg = e.toString();
                    } else {
                        this.errormsg = e.toString();
                        this.detailedmsg = null;
                    }
                }
            }
        },

        ChangeName(){
          this.isModalOpen = !this.isModalOpen 
          console.log(this.isModalOpen)

        },

        getMyStream() {
        // Logica per mostrare lo stream
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
                    this.errormsg = "Qualcosa è andato storto, per favore prova di nuovo controllando tutti i campi. Se il problema persiste, contattaci";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Qualcosa è andato storto :( ...riprova più tardi";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
    }
};
</script>

<style scoped>
.profile-container {
  display: flex;
  height: 100vh;


}

.profile-info {
  width: 33.33%; /* 1/3 della larghezza */
  background-color: #ffffff; /* Colore bianco */
  padding: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: white;
  border: 1px solid #fff; /* Aggiungi una bordatura bianca */
  border-radius: 3px; /* Angoli smussati */
}

.profile-photos {
  width: 66.67%; /* 2/3 della larghezza */
  background-color: #006600; /* Colore nero */
  padding: 20px;
  color: white;
  border: 1px solid #fff; /* Aggiungi una bordatura bianca */
  border-radius: 3px; /* Angoli smussati */
}

.profile-image img {
  width: 100px; /* Dimensioni dell'immagine del profilo */
  height: 100px;
  border-radius: 50%; /* Bordi arrotondati per l'immagine del profilo */
  margin-bottom: 10px;
}
.profile-actions {
  display: flex;
  flex-direction: column;
  align-items: center; /* Allineamento verticale dei bottoni */
}

.profile-actions button {
  margin-right: 10px;
  margin-bottom: 10px;
  align-items: center;
}

.photo-item {
  margin-bottom: 20px;
}

.photo-item img {
  width: 100%; /* La foto si adatta alla larghezza del suo contenitore */
}
button {
  background-color: leftColor; /* Green */
  border: none;
  color: black;
  padding: 10px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  transition-duration: 0.4s;
  cursor: pointer;
  border-radius: 10px;
}

/* Stili per il modale */
.modal-container {
  position: fixed; /* Posiziona il modale sopra il contenuto */
  z-index: 1; /* Imposta lo z-index in modo che il modale sia sopra gli altri elementi */
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto; /* Abilita lo scorrimento quando il contenuto è più grande del modale */
  background-color: rgba(0, 0, 0, 0.4); /* Sfondo scuro con trasparenza */
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto; /* Centra verticalmente e orizzontalmente il modale */
  padding: 20px;
  border: 1px solid #888;
  width: 80%; /* Larghezza del modale */
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

</style>