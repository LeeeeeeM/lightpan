<template>
  <div class="conmentContainer">
       <div class="ListHead">
          <div>
            <el-button type="primary" icon="el-icon-upload" @click="openFileUpAlert" size="small">上传</el-button>
            <el-button type="primary" icon="el-icon-circle-plus" size="small">新建文件夹</el-button>
            <el-button-group v-show="multipleSelection.length>0">
              <el-button type="primary" icon="el-icon-share" size="small">分享</el-button>
              <el-button type="primary" icon="el-icon-download" size="small">下载</el-button>
              <el-button type="primary" icon="el-icon-delete" size="small">删除</el-button>
              <el-button type="primary" size="small">重命名</el-button>
              <el-button type="primary" size="small">复制到</el-button>
              <el-button type="primary" size="small">移动到</el-button>
            </el-button-group>
          </div>
         <div style="position: relative">
           <input type="text" class="listiput" placeholder="请搜索您的文件">
            <div class="inputsearch"></div>
         </div>
       </div>
       <div class="ListNav">
         <div class="navBox">

           <a @click="back()" class="navback" v-show="navList.length>1">返回上一级</a>
           <el-breadcrumb separator="/">
             <el-breadcrumb-item v-for="(item,i) in navList" :key="i">
               <a v-if="i!=(navList.length-1)" >{{item.Name}}</a>
               <span v-else>{{item.Name}}</span>
             </el-breadcrumb-item>
           </el-breadcrumb>
         </div>
         <div style="font-size: 14px">
           已全部加载，共12个
         </div>
       </div>
       <div style="flex:1">
         <el-table
           ref="multipleTable"
           :data="tableData"
           height="100%"
           width="100%"
           tooltip-effect="dark"
           @selection-change="handleSelectionChange"
           style="width: 100%"
           @row-click="RowClick"
         >
           <el-table-column
             type="selection"
             width="55">
           </el-table-column>
           <el-table-column
             label="文件名称"
             >
             <template slot-scope="scope">
               <div class="rowimg">
                 <img v-if="!scope.row.isFile" src="@/css/img/dir.png" alt="">
                 <img v-if="scope.row.isFile" src="@/css/img/file.png" alt="">
                 {{scope.row.Name}}
               </div>
             </template>
           </el-table-column>
           <el-table-column
             prop="size"
             label="大小"
             width="80">
           </el-table-column>
           <el-table-column
             prop="Createtime"
             label="日期"
             width="180"
             show-overflow-tooltip>
           </el-table-column>
           <el-table-column
             fixed="right"
             label="操作"
             width="120">
             <template slot-scope="scope">
               <i class="icon iconfont icon-ai-share aaa" title="分享"></i>
               <i class="icon iconfont icon-xiazai aaa" title="下载"></i>
               <el-dropdown>
                  <span class="el-dropdown-link">
                    <i class="icon iconfont icon-gengduo aaa" title="更多"></i>
                  </span>
                 <el-dropdown-menu slot="dropdown">
                   <el-dropdown-item>移动到</el-dropdown-item>
                   <el-dropdown-item>复制到</el-dropdown-item>
                   <el-dropdown-item>重命名</el-dropdown-item>
                   <el-dropdown-item>删除</el-dropdown-item>
                 </el-dropdown-menu>
               </el-dropdown>

             </template>
           </el-table-column>
         </el-table>
       </div>
    <fileUp :path="path" ref="fileUpAlert"></fileUp>
  </div>
</template>

<script>
  import fileUp from "@/components/fileUp/fileUp"
  import {getFile} from "../api/up";

  export default {
    components:{
      fileUp
    },
    computed:{
      // path(){
      //   return
      // }
    },
    mounted(){
     this.loadFile('/')
    },
    data() {
      return {
        path:"/wujun/test",
        tableData: [

       ],
        multipleSelection: [],
        navList:[
          {
            Name:"全部文件",
            Path:"/"
          }
        ]
      }
    },

    methods: {
      handleSelectionChange(val) {
        this.multipleSelection = val;
      },
      RowClick(row){
        if(!row.isFile){
          var Path=row.Path
          this.loadFile(Path);
          this.navList.push({
            name:row.Name,
            Path:Path
          })
        }

      },
      back(){
        if( this.navList.length>1){
          this.$message('返回成功，刷新表格');
          this.navList.pop()
        }else{
          this.$message({
            message: '已经是根目录了！',
            type: 'warning'
          });
        }
      },
      openFileUpAlert(){
        this.$refs.fileUpAlert.show=true;
      },
      loadFile(path){
        var _this=this;
        getFile(path).then(res=>{
          console.log(res);
          _this.tableData=[];
         var files= res.Result.Files;
          var Childfolder= res.Result.Childfolder;
            for(var key in Childfolder){
              _this.tableData.push({
                Name:key,
                Path:"/"+key+"/",
                isFile:false
              })
            }
            for(var i in files){
              var  json=files[i];
              json.isFile=true
              _this.tableData.push(json)
            }
        })
      }
    }
  }
</script>
<style scoped>
    .ListHead{
      height: 50px;
      padding: 0 10px;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
    }
  .listiput{
    padding-left: 10px;
    padding-right: 40px;
    height: 32px;
    width: 180px;
     background-color: #f1f2f4;
    border-radius: 16px;
     border: none;
    outline: none;
    position: relative;
  }
    .inputsearch{
    position: absolute;
      right: 6px;
      top: 2px;
    width: 32px;
    height: 32px;
    background-image: url("../css/img/search.png");
  }


    .ListNav{
      height: 32px;
      padding: 0 10px;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
    }
   .navBox{
     display: flex;
     flex-direction: row;
     align-items: center;
   }
  .navback{
    font-weight:700;
    font-size: 14px;
    color:#303133;
     margin: 0 5px;
    padding: 0 5px;
    height: 16px;
    line-height: 16px;
    border-right: 1px solid #303133;
    cursor: pointer;
  }
  .navback:hover{
     text-decoration: underline;
  }
  .aaa{
    font-size: 18px;
    cursor: pointer;
  }
  .aaa:hover{
    color: #409eff;
  }
  .rowimg{
    display: flex;
    align-items: center;
  }
  .rowimg>img{
    width: 20px;
    height: 20px;
    margin-right:  10px;
  }
</style>
