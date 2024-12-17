<template>
  <div class="release">
    <div class="container">
      <app-header/>  
      <div class="main">
        <img class="main_img" src="../assets/interstellar_2.jpg"/>
        <div class="content">
          <div class="year">
            2017
          </div>
          <div class="title">
           Interstellar reunion
          </div>
          <div class="genres">
            <div class="genre" v-for="item in list">
              {{item}}
            </div> 
          </div>
          <div class="text">
            Interstellar is a 2014 epic science fiction drama film directed by Christopher Nolan, who co-wrote the screenplay with his brother Jonathan Nolan. The screenplay had its origins in a script Jonathan developed in 2007 and was originally set to be directed by Steven Spielberg 
          </div>
          <div class="info">
            <div class="item" v-for="element in info">
              <img :src="element.src"/>
              <div class="label">
                {{element.label}}
              </div>
            </div>
            <div class="rating">
              <div class="dot">
              </div>
              <div class="label">
                Imdb: 
              </div> 
              <div style="font-size:18px ; font-weight:600 ;">
                9.1/10 
              </div>
            </div>
          </div> 
          <div class="button">
            <img src="../assets/play.svg"/>
            Watch later
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  .release{
    height: 100vh ;  
    background:linear-gradient(to bottom, rgba(0 , 0 , 0, 0.3), rgba(28, 27, 33, 0.7)), url('../assets/interstellar_1.jpg') ; 
    background-repeat: no-repeat ; 
    background-size: cover ; 
    color: white ; 
  }
  .container{ 
    width: 1650px ;  
    margin: 0 auto ; 
  }
  .main{
    display: flex ; 
    margin-top: 10vh ; 
    gap: 40px ;
  }
  .main_img{
    width: 600px ; 
    height: 765px ; 
    border-radius: 15px 15px 0 0  ;
  }
  .year{
    font-size: 20px ; 
    font-weight: 600 ; 
  }
  .title{
    margin: 40px 0 ;
    font-size: 60px ; 
    font-weight: 700 ; 
  }
  .genres{
    display: flex ; 
    gap: 24px ; 
    font-size: 18px ; 
  }
  .text{
    width:600px ; 
    margin: 40px 0 ; 
  }
  .info{
    display: flex ; 
    gap: 50px ; 
  }
  .item{
    display: flex ; 
    align-items: center ; 
    gap: 5px ; 
  }
  .rating{
    display: flex; 
    gap: 5px ; 
    align-items: center ; 
  }
  .dot{
    width:10px ; 
    height: 10px; 
    border-radius:50% ; 
    background: red ; 
  }
  .button{
    margin-top: 5vh ; 
    display: flex ; 
    gap: 10px; 
    color: rgba(255 , 255 , 255 , .8) ; 
    align-items: center ; 
    width: fit-content ; 
    padding: 10px 35px; 
    border-radius: 20px; 
    background: red ; 
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
import HeaderVue from '../components/Header.vue'
import src_1 from '../assets/clock.svg'
import src_2 from '../assets/subtitles.svg'
export default {
  data(){
    return{
      list: ['Science fiction' , 'Drama' , 'Action'] , 
      info: [
              {src: src_1 , label:'164 min' } , 
              {src: src_2 , label:'Subtitles'}
            ]
    }
  } , 
  components: {
    'app-header': HeaderVue  , 
  } 
}
</script>
