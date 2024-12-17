<template>
  <div class="header">
    <logotype/> 
    <div class="account">
      <router-link class="link" to="/authorization">{{user}}</router-link>
    </div>
  </div> 
</template>

<style scoped>
  .header{
    padding-top: 5vh ;  
    display: flex ; 
    justify-content: space-between ; 
  }
  .account{
    color: white ; 
    font-weight: 700 ; 
    font-size: 18px ; 
  }  
  .link{
    text-decoration: none ; 
    color: white ; 
  }
</style>

<script>
import LogotypeVue from './Logotype.vue'
export default {
  data(){
    return {
      user: "miko"
    }
  } , 
  components: {
    'logotype': LogotypeVue 
  }  ,
  mounted(){
    this.user = localStorage.getItem('username')
  }
}
</script>
