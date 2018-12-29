import Vue from 'vue'
import Router from 'vue-router'

import up from '@/page/up'
import login from '@/page/login'
import list from '@/page/list'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/login',
      name: 'login',
      component: login
    },{
      path: '/up',
      name: 'up',
      component: up
    },{
      path: '/',
      name: 'list',
      component: list
    },
  ],
  mode:'history'
})
