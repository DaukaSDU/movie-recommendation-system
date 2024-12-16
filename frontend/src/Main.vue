
<template>
  <div class="main"> 
    <release/>
    <list/>
    <app-footer/>
  </div>
</template>

<style >
@import url('https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');

*{
  margin: 0 ; 
  padding: 0 ; 
  border-size: border-box ; 
  font-family: 'Poppins' , sans-serif ; 
}
</style>

<script>
import ReleaseVue from './sections/Release.vue'
import ListVue from './sections/List.vue'
import FooterVue from './sections/Footer.vue'
export default {
  components: {
    'release' : ReleaseVue , 
    'list': ListVue , 
    'app-footer': FooterVue 
  }
}
</script>
