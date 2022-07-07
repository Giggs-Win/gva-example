<template>
  <div>
    <el-upload
        :action="`${path}/fileUploadAndDownload/upload`"
        :before-upload="checkFile"
        :headers="{ 'x-token': userStore.token }"
        :on-error="uploadError"
        :on-success="uploadSuccess"
        :show-file-list="false"
        class="upload-btn"
    >
      <el-button size="small" type="primary">普通上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isGif = file.type === 'image/gif'
  const isSvg = file.type === 'image/svg+xml'
  const isMp4 = file.type === 'video/mp4'
  //image/svg+xml

  //console.log(file.type)

  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isJPG && !isPng && !isGif && !isSvg && !isMp4) {
    ElMessage.error('上传支持 jpg, png, gif, svg, mp4')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('未压缩未见上传图片大小不能超过 2M, 请使用压缩上传')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG || isGif || isSvg || isMp4) && isLt2M
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}

</script>

<script>

export default {
  name: 'UploadCommon',
  methods: {

  }
}
</script>
