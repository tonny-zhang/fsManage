<template>
    <div>
        <el-dialog 
            v-model="show"
            @closed="this.$emit('close')" fullscreen>
            <el-upload
                class="upload-demo"
                drag
                :show-file-list=false
                :action=uploadURL
                :data={dir:dir}
                :on-success="handleSuccess"
                :on-progress="handleProgress"
                multiple>
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
            </el-upload>
        </el-dialog>
        <slot></slot>
    </div>
</template>
<script>
import {UPLOAD_URL} from '../api'
export default {
    name: "upload",
    props: {
        isShow: Boolean,
        dir: String
    },
    watch: {
        isShow(val) {
            this.show = val;
        },
    },
    data() {
        return {
            uploadURL: UPLOAD_URL,
            show: false
        }
    },
    methods: {
        handleSuccess(res) {
            const _this = this;
            if (res.code == 200) {
                _this.$message({
                    type: 'success',
                    message: res.msg
                });
                _this.$emit('uploaded');
            } else {
                _this.$message({
                    type: 'error',
                    message: res.msg,
                });
            }
        },
        handleProgress() {
            
        },
    },
    emits: ['close', 'uploaded']
}
</script>
<style>
.el-upload-dragger{
    width: 100% !important;
}
.el-upload{
    width: 100%;
}
</style>