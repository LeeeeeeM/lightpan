<template>
  <el-dialog
    title="文件上传"
    :visible.sync="show"
    width="30%"
    :before-close="handleClose">

    <el-form ref="form"  label-width="80px">
      <el-form-item label="是否公开:"  style="text-align: left">
        <el-switch v-model="flag"></el-switch>
      </el-form-item>
      <el-form-item label="文件上传:" style="text-align: left">
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
      </el-form-item>
    </el-form>


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
        },
        flag:false
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
        var _this=this;
        let path=this.url+this.path+"/"+file.name;
        console.log(path);
        let param = new FormData(); //创建form对象
        param.append('file',file,file.name);//通过append向form对象添加数据
        param.append('pub',_this.flag);//通过append向form对象添加数据
        // return new Promise((resolve, reject)=>{
        //   debugger
        //   postFile(path,param).then(res=>{
        //     resolve(file)
        //   }).catch(err=>{
        //     reject(err)
        //   })
        // })
          postFile(path,param).then(res=>{});
          return true;

      }
    }
  }
</script>

<style scoped>

</style>
