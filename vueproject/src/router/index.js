import Vue from 'vue'
import Router from 'vue-router'

import up from '@/page/up'
import login from '@/page/login'
import list from '@/page/list'
import home from '@/page/home'
import upload from '@/components/upload/upload'


Vue.use(Router)

export default new Router({
  routes: [
    {path:'/',component:login},
    {path:'/login',component:login},
    {path: '/404', component: () => import('@/page/404')},
    {path:'/upload',component:upload},
    {path:"/home",component:home,meta:{auth:true},children:[
        {path:'',meta:{auth:true},component:list},
        {path:'/list',mate:{auth:true},component:list},
        {path:'/up',mate:{auth:true},component:up},
      ]},
    {path: '*', redirect: '/404'}
  ],
  mode:'history'
})
