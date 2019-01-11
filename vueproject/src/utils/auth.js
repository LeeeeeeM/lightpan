import Cookies from 'js-cookie'

const TokenKey = 'token'

const NameKey = 'Name'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}


export function getName() {
  return Cookies.get(NameKey)
}

export function setName(Name) {
  return Cookies.set(NameKey, Name)
}

export function removeName() {
  return Cookies.remove(NameKey)
}
