<template>
    <div class="main">
         <div class="head">
                <!--<div class="headleft">-->
                  <!--<div class="titleBox">-->
                    <!--<div class="title">-->
                      <!--<i class="icon iconfont icon-wangpanjihuo" style="font-size: 40px"></i>-->
                      <!--<span style="font-size: 24px">LightPan</span>-->
                    <!--</div>-->
                    <!--&lt;!&ndash;<div style="font-size: 18px;cursor: pointer" @click="openflag=!openflag" :title="openflag?'展开':'关闭'" :class="{'icon':true,'iconfont':true,'icon-zhankai':openflag,'icon-shouqi':!openflag}">&ndash;&gt;-->

                    <!--&lt;!&ndash;</div>&ndash;&gt;-->
                  <!--</div>-->
                  <!--<div class="selctItemBox">-->
                       <!--<div class="selctItem selctItemActive">网盘</div>-->
                    <!--<div class="selctItem">分享</div>-->
                    <!--<el-badge :value="1"  type="primary">-->
                      <!--<div class="selctItem">找资源</div>-->
                    <!--</el-badge>-->
                    <!--<div class="selctItem">更多</div>-->
                  <!--</div>-->
                <!--</div>-->
               <div class="user" style="padding-left: 80px">
                 <img class="userimg" src="@/css/img/user.png" alt="">
                 <el-dropdown class="username" style="color:#409eff">
                       <span class="el-dropdown-link">
                         admin<i class="el-icon-arrow-down el-icon--right"></i>
                         </span>
                   <el-dropdown-menu slot="dropdown">
                     <el-dropdown-item>
                       <i class="icon iconfont icon-wode"></i>
                       消息
                     </el-dropdown-item>
                     <el-dropdown-item>
                       <i class="icon iconfont icon-zhuce"></i>
                       注册
                     </el-dropdown-item>
                     <el-dropdown-item disabled>
                       <i class="icon iconfont icon-xinxi"></i>
                       信息
                     </el-dropdown-item>
                     <el-dropdown-item divided>
                       <i class="icon iconfont icon-h"></i>
                       退出
                     </el-dropdown-item>
                   </el-dropdown-menu>
                 </el-dropdown>
               </div>
         </div>
         <div class="container" style="flex:1">
             <div class="left" :style="{'width':openflag?'auto':'250px'}">

               <el-menu
                 default-active="2"
                 class="el-menu-vertical-demo"
                 :collapse="openflag"
                 background-color="#141a48"
                 text-color="#fff"
                 active-text-color="#ffd04b">

                 <el-menu-item index="0">
                   <i class="icon iconfont icon-CombinedShape lefticon"></i>
                   <span slot="title">首页</span>
                 </el-menu-item>
                 <el-submenu index="1">
                   <template slot="title">
                     <i class="icon iconfont icon-wenjian lefticon"></i>
                     <span>全部文件</span>
                   </template>
                     <el-menu-item index="1-1">
                       <i class="icon iconfont icon-tupian"></i>
                       图片
                     </el-menu-item>
                     <el-menu-item index="1-2">
                       <i class="icon iconfont icon-wenjian1"></i>
                       文件
                     </el-menu-item>
                   <el-menu-item index="1-3">
                     <i class="icon iconfont icon-shipin-tianchong"></i>
                     视频
                   </el-menu-item>
                 </el-submenu>
                 <el-menu-item index="2">
                   <i class="icon iconfont icon-ai-share lefticon"></i>
                   <span slot="title">分享</span>
                 </el-menu-item>
                 <el-menu-item index="3">
                   <i class="icon iconfont icon-huishouzhan lefticon"></i>
                   <span slot="title">回收站</span>
                 </el-menu-item>

               </el-menu>
             </div>
             <div class="viewCenter">
               <transition name="slide-fade">
                   <router-view/>
               </transition>
             </div>
         </div>
    </div>
</template>

<script>
    export default {
        data(){
          return{
            isCollapse:true,
            openflag:false
          }
        }
    }
</script>

<style scoped>

    .main{
      width: 100%;
      height: 100%;
      display: flex;
      flex-direction: column;
    }
  .container{
    display: flex;
    flex-direction: row;
  }
  .viewCenter{
  flex:1
  }
  .head{
    height: 80px;
    background-color: #141a48;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    color:#409eff;
  }
  .titleBox{
    height: 80px;
    width: 250px;
    color:#409eff;
    display: flex;
    flex-direction: row;

    align-items: center;
  }
  .title{
    height: 80px;
    color:#409eff;
    display: flex;
    padding: 0 30px;
    flex-direction: column;
   justify-content: center;
    align-items: left;
  }
  .user{
    margin-right: 40px;
    display: flex;
    flex-direction: row;
    height: 80px;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    position: relative;
  }
  .userimg{
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }
  .username{
    padding: 0 10px;
  }
  .userbox{
    position: absolute;
    bottom: -80px;
    right: 0;
    left: 0;
    height: 80px;
    border: 1px solid #aaa;
    border-bottom: none;
    background-color: #fff;

    text-align: center;
  }
  .headleft{
    height:80px;
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .selctItemBox{

    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .selctItem{
    font-size: 18px;
    font-weight: bold;
    padding: 5px 0;
    margin: 0 20px;
    color: #fff;
    border-bottom: 5px solid rgba(0,0,0,0);
    cursor: pointer;
  }
  .selctItemActive{
    color: #409eff!important;
    border-bottom: 5px solid #409eff!important;
  }
  .left{
    max-width: 250px;
    text-align: left;
    height: 100%;
    background-color: #141a48;
  }
  .lefticon{
    vertical-align: middle;
    margin-right: 5px;
    width: 24px;
    text-align: center;
    font-size: 18px;
  }
    .slide-fade{
      position: absolute;left:0;right: 0;
    }
    .slide-fade-enter-active {
      transition: all 1.2s ease;
    }
    .slide-fade-leave-active {

      transition: all .1s cubic-bezier(2.0, 0.5, 0.8, 1.0);
    }
    .slide-fade-enter, .slide-fade-leave-to
    {
      left:0;right: 0;
      transform: translateX(50px);
      opacity: 0;
    }
</style>
