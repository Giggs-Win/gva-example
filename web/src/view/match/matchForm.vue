<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="分类:">
          <el-input v-model.number="formData.categoryId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="比赛名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="封面:">
          <el-input v-model="formData.cover" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="desc:">
          <el-input v-model="formData.desc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="作者id:">
          <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="page view:">
          <el-input v-model.number="formData.pv" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="点赞数:">
          <el-input v-model.number="formData.voteNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Match'
}
</script>

<script setup>
import {
  createMatch,
  updateMatch,
  findMatch
} from '@/api/match'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        categoryId: 0,
        name: '',
        cover: '',
        desc: '',
        userId: 0,
        sort: 0,
        pv: 0,
        voteNum: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMatch({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rematch
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createMatch(formData.value)
          break
        case 'update':
          res = await updateMatch(formData.value)
          break
        default:
          res = await createMatch(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
