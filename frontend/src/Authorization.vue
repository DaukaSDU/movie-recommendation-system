<template>
  <div class="authorization">
    <div class="container">
      <logotype/> 
      <div class="content">
        <div class="login" v-if="isLogin">
          <login @changeToRegister="changeComponent"/> 
        </div>
        <div class="register" v-else>
          <register @changeToLogin="changeComponent"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.authorization{
  background: url('./assets/authorization.jpg') ; 
  background-repeat: no-repeat ; 
  background-size: cover ; 
  height: 100vh ; 
}
.container{
  width: 1650px ; 
  margin: 0 auto ; 
  padding-top: 5vh ; 
}
.content{
  margin-top: 10vh ;  
}

@media (max-width:1750px){
  .container{
    width: 1500px ; 
  } 
}
@media (max-width: 1650px){
  .container{
    width: 1350px ; 
  }
}
</style>

<script>
import LogotypeVue from './components/Logotype.vue'
import LoginVue from './components/Login.vue'
import RegisterVue from './components/Register.vue'
export default {
  components: {
    'logotype': LogotypeVue ,
    'login': LoginVue , 
    'register': RegisterVue 
  }, 
  data(){
    return{
      isLogin: true 
    }
  } , 
  methods: {
    changeComponent(){
     this.isLogin = !this.isLogin ; 
    }
  }
}
</script>
