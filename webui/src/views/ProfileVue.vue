<template>
    <div class="profile-container">
      <!-- Parte sinistra con le informazioni del profilo -->
      <div class="profile-info" :style="{ backgroundColor: leftColor }">
        <div class="profile-header">
          
          <h2><a :href="userProfileLink">{{ username }}</a></h2>
        </div>
        <div class="profile-actions">
          <button @click="uploadPhoto">Carica Foto</button>
          <button @click="renameUser">Rinomina Utente</button>
          <button @click="showStream">Mostra Stream</button>
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
        { url: "C:\\Users\\fontm\\Desktop\\images.jpg" },
        { url: "photo2.jpg" },
        { url: "photo3.jpg" }
    ],
    leftColor: '#000000', // Colore bianco
    rightColor: '#006600' // Colore nero
    };
    },
    methods: {
        uploadPhoto() {
        // Logica per caricare una nuova foto
        },
        renameUser() {
        // Logica per rinominare l'utente
        },
        showStream() {
        // Logica per mostrare lo stream
        },
        deletePhoto() {
        // Logica per eliminare una foto
        
        }
    }
};
</script>
