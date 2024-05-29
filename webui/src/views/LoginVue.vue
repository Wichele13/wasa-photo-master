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
                    <button @click="doLogin" type="submit" :style="{ backgroundColor: buttonColor}" class="btn btn-primary">Login</button>
                </div>
                </form>
            </div>
    </div>
</template>




<script>
export default {
    data: function() {
    return {
      socialName: "WASAPhoto",
      description: "Rimani in contatto con i tuoi amici condividendo le foto dei momenti speciali, grazie a WASAPhoto! Potete caricare le vostre foto direttamente dal vostro PC, che saranno visibili a tutti coloro che vi seguono.",
      leftColor: '#000000',
      rightColor: '#13cf2f',
      buttonColor: '#1f1f1f',
      errormsg: null,
      username: "",
      profile: {
          id: 0,
          username: "",
      },
    };
  },
  methods: {
    async doLogin() {
            if (this.username.length <3){
                this.errormsg = "Lo username deve essere lungo almeno 3 caratteri";
            } else {
                try {
                    let response = await this.$axios.post("/session", { username: this.username })
                    this.profile = response.data
                    sessionStorage.setItem("token", this.profile.id);
                    sessionStorage.setItem("username", this.profile.username);
                    this.$router.push({ path: '/users/'+ sessionStorage.getItem('username')+'/profile' })
                } catch (e) {
                    if (e.response && e.response.status === 400) {
                        this.errormsg = "Qualcosa è andato storto";
                        this.detailedmsg = null;
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "Problema server, prova più tardi";
                        this.detailedmsg = e.toString();
                    } else {
                        this.errormsg = e.toString();
                        this.detailedmsg = null;
                    }
                }
            }

        }
    },
    mounted() {

    }


  }


</script>

<style>
/* Stili per il componente di login */
.login-container {
    display: flex;
    height: 100vh;
  }
  
  .social-description {
    width: 33.33%; /* 1/3 della larghezza */
    transition: background-color 0.5s ease; /* Aggiungi un effetto di transizione */
  }
  
  .login-form {
    width: 66.67%; /* 2/3 della larghezza */
    transition: background-color 0.5s ease; 
  }

  button {
    background-color: #4CAF50; /* Green */
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 4px 2px;
    transition-duration: 0.4s;
    cursor: pointer;
    border-radius: 10px;
  }
  

  
  .social-description, .login-form {
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    color: white;
    border: 1px solid #fff; /* Aggiungi una bordatura bianca */
    border-radius: 10px; /* Angoli smussati */
  }
  
  .social-description h2, .login-form h2 {
    margin-bottom: 20px;
  }
  
  .login-form input {
    width: 80%;
    margin-bottom: 8px;
  }

  @media (max-width: 768px) {
    .social-description, .login-form {
      width: 100%;
    }
  }
</style>