<template> 
  <div class="input">
    <input v-if="isPrivate" type="password" :placeholder="title" />
    <input v-else type="text" :placeholder="title" />
  </div>
</template>

<style scoped>
input{
  border: none ; 
  background: none ; 
  outline: none ; 
  width: 100% ; 
  padding-bottom: 10px ;  
  font-size: 16px ; 
  font-weight: 500 ; 
  border-bottom: 2px solid black ; 
}
</style>

<script>
export default {
  props: ["title" , "isPrivate"]
}
</script>
