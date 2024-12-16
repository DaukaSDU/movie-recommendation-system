<template>
  <div class="footer">
    <div class="container">
      <div class="cheaters">
        <div class="cheater" v-for="item in list">
          {{item}}
        </div>
      </div>
      <logotype class="logo"/> 
    </div>
  </div>
</template>

<style scoped>
.footer{
  padding: 10vh 0 ; 
  background: #1D1E23 ; 
}
.container{
  width: fit-content;  
  margin: 0 auto ; 
}
.cheaters{
  display: flex ; 
  gap: 50px ; 
  color: white ;
  font-size: 16px ;
  font-weight: 500 ; 
  text-transform: uppercase ; 
}
.logo{
  width: 100% ; 
}
</style>

<script>
import LogotypeVue from '../components/Logotype.vue'
export default {
  components: {
    'logotype': LogotypeVue 
  } , 
  data(){
    return{
      list: ["Daulet" , "Asset" , "Miras"] 
    }
  }
}
</script>
