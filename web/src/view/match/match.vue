<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="分类">
          <!-- <el-input v-model="searchInfo.category_id" placeholder="搜索条件" /> -->
          <el-select v-model="searchInfo.category_id" clearable placeholder="请选择">
            <el-option v-for="(item,key) in cateItems" :key="key" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="比赛名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="分类" prop="categoryId" width="120" />
        <el-table-column align="left" label="比赛名称" prop="name" width="120" />
        <el-table-column align="left" label="封面" prop="cover" width="120" />
        <el-table-column align="left" label="desc" prop="desc" width="120" />
        <el-table-column align="left" label="作者id" prop="userId" width="120" />
        <el-table-column align="left" label="排序" prop="sort" width="120" />
        <el-table-column align="left" label="page view" prop="pv" width="120" />
        <el-table-column align="left" label="点赞数" prop="voteNum" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateMatchFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="分类:">
          <el-input v-model.number="formData.categoryId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="比赛名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="封面:">
          <div style="display:inline-block" @click="openHeaderChange">
            <img v-if="formData.cover" class="header-img-box" :src="(formData.cover && formData.cover.slice(0, 4) !== 'http')?path+formData.cover:formData.cover">
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
          <!-- <el-input v-model="formData.cover" clearable placeholder="请输入" /> -->
        </el-form-item>
        <el-form-item label="比赛简介:">
          <!-- <WEditor :desc="formData.desc" @desc-change="updateDesc"></WEditor> -->
          <WEditor v-model:desc="formData.desc"></WEditor>
          <!-- <span v-html="formData.desc"></span> -->
          
          <!-- <el-input v-model="formData.desc" clearable placeholder="请输入" /> -->
        </el-form-item>
        <el-form-item label="排序:">
          <el-input-number v-model.number="formData.sort" :min="0" :max="20" placeholder="1" />
          <!-- <el-input v-model.number="formData.sort" clearable placeholder="请输入" /> -->
        </el-form-item>
        <!-- <el-form-item label="作者:">
          <el-input v-model.number="formData.user_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="浏览数:">
          <el-input v-model.number="formData.pv" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="点赞数:">
          <el-input v-model.number="formData.votenum" clearable placeholder="请输入" />
        </el-form-item> -->
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="formData" :target-key="`cover`" />
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
  deleteMatch,
  deleteMatchByIds,
  updateMatch,
  findMatch,
  getMatchList
} from '@/api/match'

// 全量引入格式化工具 请按需保留
import { useUserStore } from '@/pinia/modules/user'
import { getMatchCategoryList } from '@/api/matchCategory'
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'

import WEditor from '@/components/editor/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'

import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
// import Editor from '../../components/editor/index.vue'

const path = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()
const router = useRouter()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  category_id: 0,
  category_name: '',
  name: '',
  cover: '',
  desc: '',
  user_id: 0,
  user_name:'',
  sort: 0,
  pv: 0,
  votenum: 0,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const cateItems = ref([])

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMatchList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============
const getCategoryItems = async() => {
  //cateItems.value = await getMatchCategoryList({ page: 1, pageSize: 999 })
  const res = await getMatchCategoryList({ page: 1, pageSize: 50 })
  // console.log('============')
  // console.log(res)
  cateItems.value=res.data.list
}

getCategoryItems()

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMatchFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteMatchByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMatchFunc = async(row) => {
    const res = await findMatch({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rematch
        dialogFormVisible.value = true
    }
}


const viewMatchDetail = (row) => {
  router.push({
    name: 'viewMatchDetail',
    params: {
      id: row.ID,
    },
  })
}

// const previewMatchFunc = async(row) => {
//     const res = await findMatch({ ID: row.ID })
//     console.log(res)
//     type.value = 'update'
//     if (res.code === 0) {
//         formData.value = res.data.rematch
//         dialogFormVisible.value = true
//     }
// }


// 删除行
const deleteMatchFunc = async (row) => {
    const res = await deleteMatch({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        category_id: 0,
        category_name: '',
        name: '',
        cover: '',
        desc: '',
        userId: 0,
        sort: 0,
        pv: 0,
        voteNum: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      formData.value.user_id=userStore.userInfo.ID
      res = await createMatch(formData.value)
      break
    case 'update':
      formData.value.pv+=1
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
    closeDialog()
    getTableData()
  }
}
</script>


<style lang="scss">
.match-dialog {
  .header-img-box {
  width: 400px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 5px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}

</style>
