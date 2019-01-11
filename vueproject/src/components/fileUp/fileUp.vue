<template>
  <el-dialog
    title="文件上传"
    :visible.sync="show"
    width="30%"
    :before-close="handleClose">
    <el-upload
      class="upload-demo"
      :action="actionUrl"
      :on-change="handleChange"
      :headers="myHeaders"
      :before-upload="beforeAvatarUpload"
      :file-list="fileList">
      <el-button size="small" type="primary">点击上传</el-button>
      <div slot="tip" class="el-upload__tip">只能上传单个文件</div>
    </el-upload>
  </el-dialog>

</template>

<script>
  import { getToken } from '@/utils/auth'
  import {postFile} from  "@/api/up"
  export default {
    props:{
      path:{
        type:String,
        default:""
      }
    },
    computed:{
      url(){
        return this.$store.state.user.FileUrl
      },

    },
    data() {
      return {
        show:false,
        fileList: [],
        actionUrl:"",
        myHeaders: {
          token: getToken()
        }
      };
    },
    methods: {
      handleChange(file, fileList) {

       // console.log(fileList);
      },
      handleClose(done) {
         done();
        this.fileList=[]
      },
      beforeAvatarUpload (file){
        let path=this.url+this.path+"/"+file.name;
        console.log(path);
        let param = new FormData(); //创建form对象
        param.append('file',file,file.name);//通过append向form对象添加数据
        return new Promise((resolve, reject)=>{
          debugger
          postFile(path,param).then(res=>{
            resolve(file)
          }).catch(err=>{
            reject(err)
          })
        })


      }
    }
  }
</script>

<style scoped>

</style>
