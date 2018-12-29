<template>
  <div class="hello">
    <div>
      <h2>post text</h2>
      <textarea v-model="text" type='textarea'></textarea>
      <button @click="postText">提交</button>

    </div>
    <div>
      <h2>post file</h2>
     <input type="file" ref="file"/>
     <button @click="getFile ">提交</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'up',
  data () {
    return {
      msg: 'Welcome to folder',
      text:"",
    }
  },
  methods: {
    postText () {
    　　console.log(this.text)
    },
    getFile () {
       let file=this.$refs.file
    　　console.log(this.$refs.file.files)
        if (file.length === 0) {
          return
        }
        let reader = new FileReader()
        reader.readAsArrayBuffer(file[0])
        reader.onload = function (e) {
          console.log('文件内容')
          console.log(e.target.result)
        }
 
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
