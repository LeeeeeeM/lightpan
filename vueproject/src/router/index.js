import Vue from 'vue'
import Router from 'vue-router'

import folder from '@/page/folder'
import login from '@/page/login'
import upload from '@/page/upload'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/login',
      name: 'login',
      component: login
    },{
      path: '/',
      name: 'folder',
      component: folder
    },{
      path: '/upload',
      name: 'upload',
      component: upload
    },
  ],
  mode:'history'
})
