<template>
  <div class="authorization">
    <div class="container">
      <logotype/> 
      <div class="content">
        <div class="login" v-if="isLogin">
          <login @change="changeComponent"/> 
        </div>
        <div class="register" v-else>
          <register @change=" changeComponent"/>
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
