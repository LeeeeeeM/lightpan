// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
//element ui
import ElementUI from 'element-ui';
import './css/theme/index.css';
import '@/icons' // icon

import './css/index.css'

import '@/permission' // permission control
import uploader from 'vue-simple-uploader'
Vue.use(uploader)
Vue.use(ElementUI);

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
