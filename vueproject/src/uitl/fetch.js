
// options中包含着数据
export function axiosfetch(options) {
   
    return new Promise((resolve, reject) => {
      var token = window.localStorage.getItem('token') ? window.localStorage.getItem('token') : "";
      var cid = window.localStorage.getItem('X-CID') ? window.localStorage.getItem('X-CID') : "";
      // var language = window.localStorage.getItem('language') ? window.localStorage.getItem('language') : "";
       
      var language = tools.getCookie('language')?tools.getCookie('language'): navigator.language;
      language = language == "en-US" ? "en" : language ;
      debug.log(language)
      var params = tools.deepClone(options.params);//深拷贝
      var sign_str = tools.sign(params); //签名
      const instance = axios.create({
        baseURL: ajaxUrl,
        timeout: 30000,
        //instance创建一个axios实例，可以自定义配置，可在 axios文档中查看详情
        //所有的请求都会带上这些配置，比如全局都要用的身份信息等。
        headers: { //所需的信息放到header头中
          // 'Content-Type': 'application/json',
          "Authorization": token,
          "X-CID":cid,
          "X-LOCALE":language,
          "X-SIGN":sign_str
        },
        // timeout: 30 * 1000 //30秒超时
      });
       
      let httpDefaultOpts = { //http默认配置
        method:options.method,
        url: options.url,
        timeout: 600000,
        params:Object.assign(params),
        data:qs.stringify(Object.assign(params)),
        // headers: options.method=='get'?{
        //   'X-Requested-With': 'XMLHttpRequest',
        //   "Accept": "application/json",
        //   "Content-Type": "application/json; charset=UTF-8",
        //   "Authorization": token
        // }:{
        //   'X-Requested-With': 'XMLHttpRequest',
        //   'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
        //   "Authorization": token
        // }
      }
   
      if(options.method=='get'){ //判断是get请求还是post请求
        delete httpDefaultOpts.data
      }else{
        delete httpDefaultOpts.params
      }
      instance(httpDefaultOpts)
        .then(response => {//then 请求成功之后进行什么操作
          debug.log('参数:')
          debug.log(options.params)
          debug.log('响应:')
          debug.log(response)
          debug.log(response.data.errno)
          if(response.data.errno == 401){
            // alert(response.data.errmsg)
            // window.location.href = window.location.protocol + "//" +window.location.host + '/#/login'
            return
          }
          if(response.data.errno == 400){
            reject(response)
            return
          }
          resolve(response)//把请求到的数据发到引用请求的地方
        })
        .catch(error => {
          // console.log('请求异常信息=>：' + options.params + '\n' + '响应' + error)
          debug.log('请求异常信息=>：' + options.params + '\n' + '响应' + error)
          reject(error)
        })
    })
  }