import axios from 'axios'
 
export const baseurl = 'http://127.0.0.1:9003/';
let $axios = axios.create({
  baseURL: baseurl,
});
function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
}

function $fetch(method,url,data){
    return new Promise((reslove,reject)=>{
        $axios({
            method,
            url,
            data:data,
            headers:{
                token: getCookie('token')
            }
        }).then(res=>{
            let body = res.data
            reslove(body)
        }).catch(err=>{
            reject(err)
        })

    })
}
 
// 注册登录
export const signin = data => $fetch('post','login',data)
 