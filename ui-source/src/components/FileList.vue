<template>
    <div>
        <div>
            <label>当前位置：</label>
            <ul class="nav">
                <li v-for="(n, i) in nav" :key="i">
                    <a :href=n.href>{{n.name}}</a>
                    <span v-if="i < nav.length - 1">/</span>
                </li>
            </ul>
        </div>
        
        <table class="filelist">
            <tr>
                <th width="40%">文件名</th>
                <th width="40%">文件大小</th>
                <th></th>
            </tr>
            <tr v-for="file in fileList" :key="file">
                <td>{{file.name}}</td>
                <td>{{file.size}}</td>
                <td><button>下载</button></td>
            </tr>
        </table>
    </div>
</template>
<script>
export default {
    name: "FileList",
    props: {
        fileList: Array,
        nav: Array
    }
}
</script>
<style scoped>
.nav{
    display: inline-block;
    text-align: left;
}
.nav li{
    display: inline-block;
}
.nav a{
    display: inline-block;
    padding: 0 5px;
    text-decoration: none;
}
.filelist{
    width: 100%;
    text-align: left;
}
</style>