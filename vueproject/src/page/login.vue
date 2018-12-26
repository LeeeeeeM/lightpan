<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <input
    v-model="user.id"
    placeholder="用户id"
    >
    <input
    v-model="user.password"
    placeholder="密码"
    type='password'
    >
    <button @click="loginform">登录</button>
  </div>
</template>

<script>
import {
  login }from '@/io/serverapi'
export default {
  name: 'login',
  data () {
    return {
        user: {
        id:'',
        password:''
      },
      msg: 'Welcome to login'
    }
  },
  methods: {
     async loginform() {
      try {
        const res = await login()
        if (res.code === 0) {
          console.log(res)
        } else {
          throw new Error('获取数据失败')
        }
      } catch (err) {
        console.log('获取数据失败', err)
      }
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
