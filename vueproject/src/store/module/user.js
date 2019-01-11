import { login, logout, getInfo } from '@/api/login'
import { getToken, setToken, removeToken ,setName,removeName} from '@/utils/auth'

const user = {
  state: {
    token:getToken(),
    name: '',
    avatar: '',
    roles: [],
    FileUrl:"http://47.107.102.188:9002",
    CommontUrl:"http://47.107.102.188:9003"
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    }
  },

  actions: {
    // 登录
    Login({ commit }, userInfo) {
      const username = userInfo.id.trim()
      return new Promise((resolve, reject) => {
        login(username, userInfo.password).then(response => {
           const tokenStr = response.Result.Token
          setToken(tokenStr);
          setName(username)
          commit('SET_TOKEN', tokenStr)
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
      //  return new Promise((resolve, reject) => {
      //    setToken("123");
      //    commit('SET_TOKEN', "123")
      //    resolve()
      //  })
    },

    // 获取用户信息
    GetInfo({ commit, state }) {
      // return new Promise((resolve, reject) => {
      //   getInfo().then(response => {
      //     const data = response.data
      //     if (data.roles && data.roles.length > 0) { // 验证返回的roles是否是一个非空数组
      //       commit('SET_ROLES', data.roles)
      //     } else {
      //       reject('getInfo: roles must be a non-null array !')
      //     }
      //     commit('SET_NAME', data.username)
      //     commit('SET_AVATAR', data.icon)
      //     resolve(response)
      //   }).catch(error => {
      //     reject(error)
      //   })
      // })

      return new Promise((resolve, reject) => {
        commit('SET_ROLES', [1,2,3])
        resolve()
      })
    },

    // 登出
    LogOut({ commit, state }) {
      // return new Promise((resolve, reject) => {
      //   logout(state.token).then(() => {
      //     commit('SET_TOKEN', '')
      //     commit('SET_ROLES', [])
      //     removeToken()
      //     resolve()
      //   }).catch(error => {
      //     reject(error)
      //   })
      // })

      return new Promise((resolve, reject) => {
            commit('SET_TOKEN', '')
            //commit('SET_ROLES', [])
            removeToken()
            removeName()
            resolve()
      })
    },

    // 前端 登出
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        removeToken()
        removeName()
        resolve()
      })
    }
  }
}

export default user
