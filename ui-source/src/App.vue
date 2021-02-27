<template>
  <div>
    <FileList :fileList="fileList" :nav="nav"></FileList>
  </div>
</template>

<script>
import FileList from './components/FileList'
export default {
  name: 'App',
  data() {
    return {
        fileList: [{
          name: "a.js",
          size: "15M"
        }],
        nav: [{
          name: "a",
          href: "/?a",
        }, {
          name: "b",
          href: "/?a/b",
        }]
    }
  },
  components: {
    FileList
  }
}
</script>

<style>
*{
  margin: 0;
  padding: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 20px;
}
ul, li{
  list-style: none;
  text-align: left;
}
</style>
