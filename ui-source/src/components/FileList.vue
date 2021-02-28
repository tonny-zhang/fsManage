<template>
    <div>
        <div>
            <label>当前位置：</label>
            
            <el-tag v-if="nav.length == 0" size="mini" type="danger">目录不存在</el-tag>
            <template v-else style="display:block">
                <ul class="nav">
                    <li v-for="(n, i) in nav" :key="i">
                        <a :href=n.href>{{n.name}}</a>
                        <span v-if="i < nav.length - 1">/</span>
                    </li>
                </ul>
                <div style="margin-top: 5px">
                    <el-button plain type="primary" size="mini" @click="uploadFile">上传文件</el-button>
                    <el-button plain type="primary" size="mini" @click="createFolder">新建目录</el-button>
                </div>
            </template>
        </div>
        <el-table
            :data="filelist"
            style="width: 100%"
            >
            <el-table-column
            prop="name"
            label="文件名"
            sortable>
            <template #default="scope">
                <i :class="['tb-icon', scope.row.isDir?'icon-folder':'icon-file']"></i>
                <a v-if=scope.row.isDir :href=scope.row.href>{{scope.row.name}}</a>
                <span v-else>{{scope.row.name}}</span>
            </template>
            </el-table-column>
            <el-table-column
            prop="size"
            label="文件大小"
            sortable
            width="100">
            </el-table-column>
            <el-table-column
            width="50"
            label="操作">
                <template #default="scope">
                    <el-button type="text" size="small" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <table v-show="false" class="filelist">
            <tr>
                <th width="60%">文件名</th>
                <th width="20%">文件大小</th>
                <th>操作</th>
            </tr>
            <tr v-for="file in filelist" :key="file">
                <td :class="'type-'+(file.isDir?'folder':'file')">
                    <a v-if=file.isDir :href=file.href>{{file.name}}</a>
                    <span v-else>{{file.name}}</span>
                </td>
                <td>{{file.size}}</td>
                <td>
                    <button v-if=file.isDir>打包下载</button>
                    <button v-else>下载</button>
                </td>
            </tr>
        </table>
    </div>
</template>
<script>
import {PARAMS, getList, deleteFile} from "../api"

let dirCurrent;

export default {
    name: "FileList",
    data() {
        return {
            filelist: [],
            nav: [],
        }
    },
    methods: {
        uploadFile() {
            window.T = this
            console.log(this)
            this.$confirm('test')
        },
        createFolder() {},
        sortFile(a, b, col) {
            if (col == "name") {
                return a[col].localeCompare(b[col])
            } else {
                return parseFloat(a[col]) - parseFloat(b[col])
            }
        },
        dialog(msg) {
            this.$alert(msg, "提示", {
                confirmButtonText: '确定',
                type: 'warning'
            });
        },
        handleDelete(index, row) {
            const _this = this;
            if (!row.isDir) {
                _this.$confirm("确定要删除此文件吗？", "提示", {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    deleteFile(dirCurrent+'/'+row.name).then((result) => {
                        if (result.code == 200) {
                            this.$message({
                                type: 'success',
                                message: '成功删除'
                            });
                            _this.refresh(dirCurrent)
                        } else {
                            _this.dialog(result.msg)
                        }
                    })
                }).catch(() => {});
            } else {
                _this.dialog("暂时不支持删除目录")
            }
        },
        refresh(dir) {
            const _this = this;
            getList(dir).then(v => {console.log(v)
                if (v.code == 200) {
                    let nav = v.data.nav;
                    let list = [];
                    v.data.list.forEach(vv => {
                        if (vv.isDir) {
                            vv.href = './?dir='+nav+'/'+vv.name
                        }
                        list.push(vv);
                    })
                    _this.filelist = list

                    let param = '';
                    let navList = [];
                    v.data.nav.split('/').forEach(v => {
                        navList.push({
                            name: v,
                            href: './?dir='+ param + v
                        });
                        param += v + '/'
                    })
                    _this.nav = navList

                    dirCurrent = v.data.nav;
                    console.log(_this)
                } else {
                    _this.filelist = []
                }
            }).catch(() => {
                _this.fileList = []
            })
        }
    },
    mounted: function() {
        this.refresh(PARAMS.dir)
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
a{
    text-decoration: none;
}
a, a:visited{
    color: #000;
}
.tb-icon{
    margin-right: 4px;
    background-position: 0;
    background-repeat: no-repeat;
    padding-right: 6px;
    width: 20px;
    height: 20px;
    display: inline-block;
    vertical-align: middle;
}
.tb-icon.icon-file{
    background-image: url(../assets/file.svg);
}
.tb-icon.icon-folder{
    background-image: url(../assets/folder.svg);
}


td{
    padding: 3px;
}
tr:nth-child(2n){
    background-color: rgba(230, 230, 230, 0.4);
    border-radius: 5px;
}
.type-file::before,
.type-folder::before{
    content: ' ';
    background-position: 0;
    background-repeat: no-repeat;
    padding-right: 6px;
    width: 20px;
    height: 20px;
    display: inline-block;
        vertical-align: middle;
    
}
.type-file::before{
background-image: url(../assets/file.svg);
}
.type-folder::before{
    background-image: url(../assets/folder.svg);
}
</style>