<template>
  <div v-loading.fullscreen.lock="">
    <div class="gva-table-box">
      <warning-bar
        title="点击“文件名/备注”可以编辑文件名或者备注内容。"
      />
      <div class="gva-btn-list">
        <!-- <upload-common
          v-model:imageCommon="imageCommon"
          class="upload-btn"
          @on-success="getTableData"
        /> -->
        <!-- 去掉压缩上传 -->
        <!-- <upload-image
          v-model:imageUrl="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          class="upload-btn"
          @on-success="getTableData"
        /> -->

        <el-form ref="searchForm" :inline="true">  <!-- :model="search" -->
          <el-form-item label="图片分类">
            <el-select v-model="searchInfo.category_id" clearable placeholder="请选择">
              <el-option v-for="(item,key) in imgCategoryOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
          <el-form-item label="标签">
            <el-input v-model="searchInfo.tag" class="tag" placeholder="请输入标签" />
          </el-form-item>

          <el-form-item label="">
            <el-input v-model="searchInfo.keyword" class="keyword" placeholder="请输入文件名或备注" />
          </el-form-item>

          <el-form-item>
            <el-button size="small" type="primary" icon="search" @click="getTableData">查询</el-button>
            <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
        <!--普通上传-->
        <upload-common
          v-model:imageCommon="imageCommon"
          class="upload-btn"
          @on-success="getTableData"
        />

      </div>

      <el-table :data="tableData">
        <el-table-column label="预览" width="100">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.url" />
          </template>
        </el-table-column>

        <el-table-column label="类别" prop="categoryId" width="120">
            <template #default="scope">
              <el-select v-model="scope.row.categoryId" placeholder="未分类" size="small" @change="changeCategoryFunc(scope.row)">
                <el-option
                  v-for="item in imgCategoryOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                  :disabled="item.disabled"
                />
              </el-select>
            </template>
        </el-table-column>

        <!-- <el-table-column label="类别" prop="categoryId" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.categoryId,imgCategoryOptions) }}
            </template>
        </el-table-column> -->

        <el-table-column label="标签" prop="tag" width="200">
          <template #default="scope">
            <el-tag v-for="(value,index) in scope.row.tag.split(',')"
              :type="value === '足球' ? 'success' : ''"
              @click="editTagFunc(scope.row)"
              class="file-tag"
              disable-transitions
            >{{ value }}
            </el-tag>
          </template>
        </el-table-column>
        <!-- <el-table-column class="bot_list" width="100">
          <div class="share_icon"><a href="javascript:;" class="ico1" data-fx="wx1-btn" title="微信" data-spm-anchor-id="C73465.P82ADLVYUn7e.S57461.202"></a><a href="javascript:;" class="ico2" data-fx="wb1-btn" title="新浪微博" data-spm-anchor-id="C73465.P82ADLVYUn7e.S57461.203"></a><a href="javascript:;" class="ico3" data-fx="qz-btn" title="QQ空间" data-spm-anchor-id="C73465.P82ADLVYUn7e.S57461.204"></a><a href="javascript:;" class="ico4" data-fx="qq-btn" title="QQ" data-spm-anchor-id="C73465.P82ADLVYUn7e.S57461.205"></a></div>
        </el-table-column> -->
        <el-table-column label="文件名/备注" prop="name" width="180">
          <template #default="scope">
            <div class="name" @click="editFileNameFunc(scope.row)">{{ scope.row.name }}</div>
          </template>
        </el-table-column>

        <!-- <el-table-column label="标签" prop="tag" width="100">
          <template #default="scope">
            <el-tag
              :type="scope.row.tag === 'jpg' ? '' : 'success'"
              disable-transitions
            >{{ scope.row.tag }}
            </el-tag>
          </template>
        </el-table-column> -->
        <!-- <el-table-column label="分类" prop="categoryId" min-width="50" />-->

        <el-table-column label="链接" prop="url" min-width="450" /> 
        <el-table-column label="日期" prop="UpdatedAt" width="180">
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>        
        <el-table-column label="操作" width="160">
          <template #default="scope">
            <el-button size="small" icon="download" type="primary" link @click="downloadFile(scope.row)">下载</el-button>
            <el-button size="small" icon="delete" type="primary" link @click="deleteFileFunc(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :style="{ float: 'right', padding: '20px' }"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

import CustomPic from '@/components/customPic/index.vue'
//import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import warningBar from '@/components/warningBar/warningBar.vue'

//import { formatDate } from '@/utils/format'
import { downloadImage } from '@/utils/downloadImg'
import { useUserStore } from '@/pinia/modules/user'
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'

import { getFileList, deleteFile, editFileName, editTag, changeCategory } from '@/api/fileUploadAndDownload'


const path = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

const imageUrl = ref('')
const imageCommon = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(30)
const searchInfo = ref({})
const tableData = ref([])

// const accountRef = Ref<InstanceTyp<typeof LoginAccount>>()

// const InputRef = ref<InstanceType<typeof ElInput>>()

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// var arrTag="png,svg,jpg,jpeg"
// console.log(arrTag.split(","))


// 查询
const getTableData = async() => {
  console.log(searchInfo.value)
  const table = await getFileList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  
  if (table.code === 0) {
    tableData.value = table.data.list
    console.log(table.data.list)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const imgCategoryOptions = ref([])

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    imgCategoryOptions.value = await getDictFunc('imgCategory')
    // console.log('------------')
    console.log(imgCategoryOptions.value)
}

// 获取需要的字典 可能为空 按需保留
setOptions()


const deleteFileFunc = async(row) => {
  ElMessageBox.confirm('此操作将永久文件, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(async() => {
      const res = await deleteFile(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!',
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除',
      })
    })
}

const downloadFile = (row) => {
  if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    downloadImage(row.url, row.name)
  } else {
    debugger
    downloadImage(path.value + '/' + row.url, row.name)
  }
}

const changeCategoryFunc = async (row) => {
  const res = await changeCategory(row)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '分类更新成功' })
    getTableData()
  }
}

/**
 * 编辑Tag
 * @param row
 * @returns {Promise<void>}
 */
const editTagFunc = async(row) => {
  ElMessageBox.prompt('请输入标鉴,标鉴之间,分开(eg:足球,英超)', '编辑', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '不能为空',
    inputValue: row.tag
  }).then(async({ value }) => {
    row.tag = value
    const res = await editTag(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '编辑成功!',
      })
      getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '取消修改'
    })
  })
}

/**
 * 编辑文件名或者备注
 * @param row
 * @returns {Promise<void>}
 */
const editFileNameFunc = async(row) => {
  ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '不能为空',
    inputValue: row.name
  }).then(async({ value }) => {
    row.name = value
    //console.log("===editFileNameFunc====",row)
    const res = await editFileName(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '编辑成功!',
      })
      getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '取消修改'
    })
  })
}
</script>

<script>

export default {
  name: 'Upload',
}
</script>
<style scoped>
/* .bot_list .share_icon {
    height: 26px;
    position: absolute;
    bottom: 0;
    right: 1px;
}
.bot_list .share_icon a:first-child {
    margin-left: 0;
}
.bot_list .share_icon .ico1 {
    background: url(https://p1.img.cctvpic.com/photoAlbum/templet/common/DEPA1563326600105320/share1.png) 0 center no-repeat;
}
.bot_list .share_icon a {
    width: 26px;
    height: 26px;
    margin-left: 10px;
    display: inline-block;
    float: left;
    padding-left: 0;
}
.bot_list .share_icon .ico2 {
    background: url(https://p1.img.cctvpic.com/photoAlbum/templet/common/DEPA1563326600105320/share1.png) -33px center no-repeat;
}
.bot_list .share_icon a {
    width: 26px;
    height: 26px;
    margin-left: 10px;
    display: inline-block;
    float: left;
    padding-left: 0;
}
.bot_list .share_icon .ico3 {
    background: url(https://p1.img.cctvpic.com/photoAlbum/templet/common/DEPA1563326600105320/share1.png) -68px center no-repeat;
}
.bot_list .share_icon a {
    width: 26px;
    height: 26px;
    margin-left: 10px;
    display: inline-block;
    float: left;
    padding-left: 0;
}
.bot_list .share_icon .ico4 {
    background: url(https://p1.img.cctvpic.com/photoAlbum/templet/common/DEPA1563326600105320/share1.png) -103px center no-repeat;
}
.bot_list .share_icon a {
    width: 26px;
    height: 26px;
    margin-left: 10px;
    display: inline-block;
    float: left;
    padding-left: 0;
} */


.name {
  cursor: pointer;
}
.file-tag{
   margin-right: 5px;
}

.upload-btn + .upload-btn {
  margin-left: 12px;
}
</style>
