import axios from 'axios'

export const mhost = 'http://127.0.0.1:9003/';
export const fhost = 'http://127.0.0.1:9002/';
import { getToken } from '@/utils/auth'


function $fetch(method,url,data){
    let token=window.localStorage.getItem('token');
    return new Promise((reslove,reject)=>{
        axios({
            method,
            url,
            headers:{
                'token':getToken(),
            },
            mode:'cors',
            data:data,
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

export const postfile = (who,folder,name,data )=> $fetch('post',fhost+who+"/"+folder+"/"+name,data);
