<template>
  <div class="login">
    <div class="container">
      <div class="title">
        Login
      </div>
      <app-input v-model="username" title="Username" :isPrivate="false" />
      <div class="br_a"></div>
      <app-input v-model="password" title="Password" :isPrivate="true" />
      <div class="br_a"></div>
      <div @click="login" class="button">
        Login 
      </div>
      <div class="footer">
        Don't have an account? 
        <div @click="this.$emit('changeToRegister')" class="toRegister">
          Register 
        </div>
      </div>
    </div> 
  </div>
</template>

<style scoped>
.login{
  width: 450px ; 
  margin: 0 auto ; 
  background: rgba(255, 255, 255, 0.1); 
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2); 
  border-radius: 15px; 
  border: 1px solid white ; 
  padding: 40px;
  text-align: center ; 
}
.container{
  width: calc(100% - 50px) ; 
  margin: 0 auto ; 
}
.title{
  font-size: 34px ; 
  font-weight: 700 ; 
  margin: 2vh 0 5vh ;  
}
.br_a{
  height: 5vh ; 
}
.button{
  display: flex ; 
  height: 55px ; 
  font-size: 18px ; 
  font-weight: 600 ; 
  border-radius: 8px ; 
  align-items: center ; 
  justify-content: center ; 
  color: white ; 
  background: #19283B; 
}
.footer{
  display: flex ; 
  gap: 5px; 
  width: fit-content ; 
  margin: 3vh auto ;
}
.toRegister{
  font-weight: 500 ; 
  cursor: pointer ; 
}
</style>

<script>
import InputVue from './Input.vue'
export default {
  data (){
    return {
      username: "" , 
      password: ""
    }
  } , 
  components: {
    'app-input': InputVue 
  },
  methods: {
    async login() {
      const response = await fetch("http://localhost:8000/auth/sign-in", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: this.username,
          password: this.password,
        }),
      });


      if (response.ok) {
        const result = await response.json();
        console.log("User's token:", result);
        this.goMain()
        localStorage.setItem('username', this.username);
      } else {
        const error = await response.json();
        console.log("Error: ", error);
      }
    } , 
    goMain(){
      this.$router.push('/');
    }
  }
}
</script>
