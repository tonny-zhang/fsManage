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
                    <el-upload
                        class="upload-demo"
                        action="https://jsonplaceholder.typicode.com/posts/"
                        multiple
                        :limit="3"
                        >
                        <el-button size="small" type="primary">点击上传</el-button>
                    </el-upload>
                    <el-button plain type="primary" size="mini" @click="handleUploadFile">上传文件</el-button>
                    <el-button plain type="primary" size="mini" @click="handleCreateFolder">新建目录</el-button>
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
                <a v-if=scope.row.isDir :href=scope.row.href><i class="tb-icon icon-folder"></i>{{scope.row.name}}</a>
                <span v-else @click="download(scope.row.path)"><i class="tb-icon icon-file"></i>{{scope.row.name}}</span>
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
    </div>
</template>
<script>
import {PARAMS, getList, deleteFile, downloadFile, createFolder} from "../api"

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
        handleUploadFile() {
            window.T = this
            this.$confirm('test')
        },
        handleCreateFolder() {
            const _this = this;
            _this.$prompt("请输入要创建的目录名（支持/多级目录）", {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
            }).then(({value}) => {
                if (value) {
                    value = value.trim();
                    if (value) {
                        if (value[0] != '/') {
                            value = dirCurrent + '/' + value;
                        }
                        createFolder(value).then(v => {
                            if (v.code == 200) {
                                _this.$message({
                                    type: 'success',
                                    message: v.msg
                                });
                                _this.refresh(dirCurrent)
                            } else {
                                _this.$message({
                                    type: 'error',
                                    message: v.msg,
                                });
                            }
                        })
                        return;
                    }
                }
                _this.dialog('请输入正确的目录');
            }).catch(() => {});
        },
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
                            _this.$message({
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
        download(e) {
            downloadFile(e)
        },
        refresh(dir) {
            const _this = this;
            getList(dir).then(v => {
                if (v.code == 200) {
                    let nav = v.data.nav;
                    v.data.list.forEach(vv => {
                        if (vv.isDir) {
                            vv.href = './?dir='+nav+'/'+vv.name
                        } else {
                            let path = nav+'/'+vv.name;
                            vv.href = './?file='+path;
                            vv.path = path;
                        }
                    })
                    console.log(v.data.list)
                    _this.filelist = v.data.list;

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
        if (PARAMS.file) {
            downloadFile(PARAMS.file)
        } else {
            this.refresh(PARAMS.dir)
        }
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