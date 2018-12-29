import axios from 'axios'
 
export const mhost = 'http://127.0.0.1:9003/';
export const fhost = 'http://127.0.0.1:9002/';
 
function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
}

function $fetch(method,url,data){
    return new Promise((reslove,reject)=>{
        axios({
            method,
            url,
            data:data,
            headers:{
                token: getCookie('token')
            },
            credentials: 'includes' 
        }).then(res=>{
            let body = res.data
            reslove(body)
        }).catch(err=>{
            reject(err)
        })

    })
}
 
// 注册登录
export const signin = data => $fetch('post',mhost+'login',data)

export const postfile = (who,folder,name,data )=> $fetch('post',fhost+who+"/"+folder+"/"+name,data)
 