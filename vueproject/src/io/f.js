// 注册登录
export const postfile = (who,folder,name,data )=> $fetch('post',who+"/"+folder+"/"+name,data)