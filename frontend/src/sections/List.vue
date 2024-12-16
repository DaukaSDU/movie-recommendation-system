<template>
  <div class="list">
    <div class="container">
      <div class="main">
        <div class="item" v-for="(item , index) in items" :key="item.id" @mouseleave="setVisibility(index , false)" @mouseenter="setVisibility(index , true)">
          <img :src="item.src"/>
          <div class="label"> 
            {{item.name}} 
          </div>
          <transition name="fade">
            <div class="front" v-show="isVisible[index]">
               <img src="../assets/plus.svg"/>
            </div>
          </transition>
        </div>
      </div>
      <div class="footer">
        <div class="line">
        </div>
        <div class="button">
          Show more
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.list{
  background:  #18191E ;  
  padding-top: 10vh ; 
  color: white ; 
}
.container{
  width: 1650px ; 
  margin: 0 auto ; 
}
.header{
  display: flex ;
  justify-content: space-between ; 
}
.main{
  display: grid ; 
  grid-template-columns: repeat(4 , 1fr) ; 
  gap: 30px ; 
}
.main img{
  width: 100% ;
  height: 600px ; 
}
.item{
  position: relative ;
}
.label{
  font-size: 20px ; 
  font-weight: 600 ; 
  text-transform: uppercase ; 
}
.front{
  cursor: pointer ; 
  width:100% ; 
  height: calc(100% - 35px ); 
  position: absolute ; 
  top: 0 ; 
  left: 0 ; 
  display: flex ; 
  justify-content: center ; 
  align-items: center ; 
  background: rgba(0,0,0,.4) ; 
}
.front img{
  width: 35px ; 
  height: 35px ; 
}
.fade-enter-active , 
.fade-leave-active{
  transition: .4s all ease ; 
}
.fade-enter-from , 
.fade-leave-to {
  opacity: 0 ; 
}
.footer{
  display: flex ; 
  padding: 5vh 0 ;  
  gap: 30px ;
  align-items: center ; 
}
.line{
  width: 100% ;  
  height: 2px; 
  background:rgba(255 , 255 , 255 , .8) ;  
}
.button{
  display: flex ; 
  width: 170px ; 
  height: 40px ; 
  justify-content: center ; 
  align-items: center ; 
  color: rgba(255 , 255 , 255 , .8) ; 
  border-radius: 20px; 
  background: red ; 
}
</style>

<script>
import src_1 from '../assets/list_1.jpg' 
import src_2 from '../assets/list_2.jpg' 
import src_3 from '../assets/list_3.jpg' 
import src_4 from '../assets/list_4.jpg' 
import src_5 from '../assets/list_5.jpg' 
import src_6 from '../assets/list_6.jpg' 
import src_7 from '../assets/list_7.jpg' 
import src_8 from '../assets/list_8.jpg' 

export default {
  data(){
    return {
      items: [
        {id: 1 , name: "Harry Potter" , src: src_1} , 
        {id: 2 , name: "Corps Bride" , src: src_2} , 
        {id: 3 , name: "Moana" , src: src_3} , 
        {id: 4 , name: "Cars" , src: src_4} , 
        {id: 5 , name: "Star Wars" , src: src_6} , 
        {id: 6 , name: "Iron Man" , src: src_5 } ,
        {id: 7 , name: "Avatar" , src: src_7} , 
        {id: 8 , name: "Up" , src: src_8}
      ] , 
      isVisible: [false , false , false , false , false , false , false , false]
    }
  } , 
  methods:{
    setVisibility(index , k){
      this.isVisible[index] = k 
    }
  }
}
</script>
