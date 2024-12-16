
import { createRouter, createWebHistory } from 'vue-router'
import Home from './Main.vue'  
import Authorization from './Authorization.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/authorization' , 
    name: 'Authorization' ,
    component: Authorization
  }
]

const router = createRouter({
  history: createWebHistory(),  
  routes
})

export default router
