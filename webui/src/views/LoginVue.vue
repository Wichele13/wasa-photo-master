<template>
    <div class="login-container">
      <!-- Parte sinistra con la descrizione del social -->
        <div class="social-description" :style="{ backgroundColor: leftColor }">
            <h2>{{ socialName }}</h2>
            <p>{{ description }}</p>
        </div>
        <!-- Parte destra con il modulo di accesso -->
            <div class="login-form" :style="{ backgroundColor: rightColor }">
                <h2>Login</h2>
                <form @submit.prevent="login">
                <div class="form-group">                   
                    <input type="text" placeholder="Inserisci username..." class="form-control" id="username" v-model="username" required>
                    <button type="submit" :style="{ backgroundColor: buttonColor}" class="btn btn-primary">Login</button>
                </div>
                </form>
            </div>
    </div>
  </template>





<script>
import 'webui/src/assets/loginVue.css'
export default {
    data() {
    return {
      socialName: "WASAPhoto",
      description: "Rimani in contatto con i tuoi amici condividendo le foto dei momenti speciali, grazie a WASAPhoto! Potete caricare le vostre foto direttamente dal vostro PC, che saranno visibili a tutti coloro che vi seguono.",
      username: "",
      leftColor: '#000000',
      rightColor: '#13cf2f',
      buttonColor: '#1f1f1f'
    };
  },
  methods: {


  }
};

</script>