<template>
    <div class="profile-container">
      <!-- Parte sinistra con le informazioni del profilo -->
      <div class="profile-info" :style="{ backgroundColor: leftColor }">
        <div class="profile-header">
          
          <h2><a :href="userProfile">{{ username }}</a></h2>
        </div>
        <div class="profile-actions">
          <button @click="uploadPhoto">Carica Foto</button>
          <button @click="setMyUserName">Rinomina Utente</button>
          <button @click="getMyStream">Mostra Stream</button>
        </div>
      </div>
      <!-- Parte destra con le foto del profilo -->
      <div class="profile-photos" :style="{ backgroundColor: rightColor }">
        <div class="photo-header">
          <h2>Le tue foto</h2>
        </div>
        <div v-for="(photo, index) in photos" :key="index" class="photo-item">
          <img :src="photo.url" :alt="'Photo ' + (index + 1)">
          <button @click="deletePhoto()">Elimina</button>
        </div>
      </div>
    </div>
</template>
  
    
<script>

import 'webui/src/assets/loginVue.css'
export default {
    data() {
    return {
    username: "Nome Utente",
    userProfileLink: "https://esempio.com/profilo", // Link del profilo dell'utente
    photos: [

    ],
    leftColor: '#000000', // Colore bianco
    rightColor: '#006600' // Colore nero
    };
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
                        Authorization: "Bearer " + localStorage.getItem("token")
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
                  Authorization: "Bearer " + localStorage.getItem("token")
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
            if (length(this.newUsername) <= 3){
                this.errormsg = "Non lasciare il campo username vuoto!"
            } else {
                try {
                    let response = await this.$axios.put("/user/" + this.username + "/setusername", { username: this.newUsername }, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("token")
                        }
                    })
                    this.user = response.data
                    localStorage.setItem("username", this.user.username);
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
        getMyStream() {
        // Logica per mostrare lo stream
        },
        async deletePhoto(photoid) {
            try {
                let response = await this.$axios.delete("/users/" + this.username + "/photo/" + photoid, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
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
