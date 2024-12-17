<template> 
  <div class="input">
    <input v-if="isPrivate" type="password" :value="modelValue"  @input="$emit('update:modelValue', $event.target.value)" :placeholder="title" />
    <input v-else type="text"   :value="modelValue"  @input="$emit('update:modelValue', $event.target.value)"  :placeholder="title" />
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
  props: ["title" , "isPrivate" , "modelValue"]
}
</script>
