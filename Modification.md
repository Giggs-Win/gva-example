
## 自动生成代码

### 2. 安装 Gin-Vue-Admin
#### 2.1 简易MD
* 标题 # DCS WORLD
* 无序列表 *
* 有序列表 1.|2.|3.
* 超链接 [百度](http://www.baidu.com)
* 图片 ![logo](https://www.gin-vue-admin.com/img/coding__isometric.svg)
* 引用 >
* 粗体 **我是粗休**
* 斜体 *我是斜休*
* 粗斜体 ***我是粗斜休***
* 下划线 <u>下划线</u>
* 换行 两空格+回车
* 分割线 ---
* 转义 \
* 代码块  
```
    package main
    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
```

* 表格  
    |左对齐|中对齐|右对齐| 默认 |
    | :---- | :---: | ---: |  --- |
    | 数据   | 表名 | 说明 | DESC |
---
* git new project
```
    echo "# gva-example" >> README.md
    git init
    git add README.md
    git commit -m "first commit"
    git branch -M main
    git remote add origin https://github.com/Giggs-Win/gva-example.git
    git push -u origin main
```
* git push an existing project 
```
    git remote add origin https://github.com/Giggs-Win/gva-example.git
    git branch -M main
    git push -u origin main
```
---

#### 2.2 准备工作
* 安装Node.js
* 安装Go
```
#设置代理
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```
* 安装MySQL 这里安装的是XAMPP
* SQL访问工具 HeidiSQL
* 安装 Redis
* 安装 VSCode

#### 2.3 安装(VSCode)
1. [GVA 官网 https://www.gin-vue-admin.com/](https://www.gin-vue-admin.com/)
2. [下载代码]Git clone https://github.com/flipped-aurora/gin-vue-admin.git
* 以下为Server
---
3. 使用 go mod 并安装go依赖包 go generate
4. go build -o server main.go (windows编译命令为go build -o server.exe main.go )
5. 运行./server (windows运行命令为 server.exe)
* 以下为Web
---
6. npm install
7. npm run dev(npm run serve)


### 3. 系统数据库
#### 3.1 系统原始表及数据
| 数据   | 表名 | 说明 | DESC |
| :----: | --- | --- |  --- |
| --| authority_menu           | 视图 | 新版去掉了 |
|342| casbin_rule              |          | 基于角色的访问控制 |
| 2 | sys_users                | 系统用户 |  | 
| 4 | sys_user_authority       | 用户<->角色 | 用户权限关系表|
| 0 | jwt_blacklists           |  |  |
| 92| sys_apis                 | API列表 |  |
| 3 | sys_authorities          | 系统角色表 | |
| 5 | sys_data_authority_id    |  | 
| 0 | sys_authority_btns       |  |  |
| 44| sys_authority_menus      | 菜单<->角色 |  |
| 0 | sys_auto_codes           | 生成模块表 | 自动生成代码模块 |
| 0 | sys_auto_code_histories  | 模块生成记录 |  |
| 27| sys_base_menus           | 菜单表 | |
| 0 | sys_base_menu_btns       |  | 
| 0 | sys_base_menu_parameters |  | 
| 6 | sys_dictionaries         | 字典表 |   |
| 25| sys_dictionary_details   | 字典名细表 |  | 
| 0 | sys_operation_records    | 操作记录表 |  |
| 2 | exa_file_upload_and_downloads | 上传文件表 |  |
| 0 | exa_customers            | 例：用户表 |  |
| 0 | exa_files                |  | 
| 0 | exa_file_chunks          |  | 

#### 3.2 改为适全自已习惯
```sql
    #SQL
    // sys_users
    delete from sys_users where id=2;

    // sys_user_authority
    delete from sys_user_authority where sys_user_id=2;
    delete from sys_user_authority where sys_authority_authority_id='8881';
    delete from sys_user_authority where sys_authority_authority_id='9528';
    update sys_user_authority set sys_authority_authority_id=1 where sys_authority_authority_id='888';

    // sys_authority_menus
    delete from sys_authority_menus where sys_authority_authority_id <>'888'
    update sys_authority_menus set sys_authority_authority_id=1 where sys_authority_authority_id='888';

    //casbin_rule
    delete from casbin_rule where V0<>'888';
    update casbin_rule set v0=1 where v0='888';

    //sys_authorities
    delete from sys_authorities where authority_id<>'888';
    update from sys_authorities set authority_id=1 where sys_authorities='888';
 
    //sys_data_authority_id
    delete from sys_data_authority_id where data_authority_id_authority_id<>'888';
    update sys_data_authority_id set sys_authority_authority_id=1,data_authority_id_authority_id=1 where data_authority_id_authority_id='888';

    //exa_file_upload_and_downloads
    delete from exa_file_upload_and_downloads where id=2;
```

#### 3.3 几个小改动
    1. 菜单排序
    2. 



### 4. 目录结构
#### 4.1 server
```shell
├── api                         --api组
│   └── v1                      --v1版本接: v1版本接口
├── config                      --配置包: config.yaml对应的配置结构体
├── core                        --核心文件: 核心组件(zap, viper, server)的初始化
├── docs                        --swagger文档目录: swagger文档目录
├── global                      --全局对象:	全局对象
├── initialize                  --初始化:router,redis,gorm,validator, timer的初始化
│   └── internal                --初始化内部函数: gorm 的 longger 自定义,在此文件夹的函数只能由 initialize 层进行调用
├── middleware                  --中间件层: 用于存放 gin 中间件代码
├── model                       --模型层: 模型对应数据表
│   ├── request                 --入参结构体: 接收前端发送到后端的数据
│   └── response                --出参结构体: 返回给前端的数据结构体
├── packfile                    --静态文件打包: 静态文件打包
├── resource                    --静态资源文件夹: 负责存放静态文件
│   ├─ autocode_template        --自动生成代码模板
│   │        ├── server 
│   │        |      ├── api.go
│   │        |      ├── model.go
│   │        |      ├── request.go
│   │        |      ├── router.go
│   │        |      └── service.go
│   │        └──── web
│   │               ├── api.js
│   │               ├── form.vue
│   │               └── table.vue
│   ├── excel                   --excel导入导出默认路径: excel导入导出默认路径
│   ├── page                    --表单生成器: 表单生成器 打包后的dist
│   └── template                --模板:	模板文件夹,存放的是代码生成器的模板
├── router                      --路由层: 路由层
├── service                     --service层: 存放业务逻辑问题
├── source                      --source层: 存放初始化数据的函数
└── utils                       --工具包: 工具函数封装
    ├── timer                   --timer: 定时器接口封装
    └── upload                  --upload: oss接口封装
```

### 4.2 web目录结构
```
web
 ├── babel.config.js
 ├── Dockerfile
 ├── favicon.ico
 ├── index.html                 -- 主页面
 ├── limit.js                   -- 助手代码
 ├── package.json               -- 包管理器代码
 ├── src                        -- 源代码
 │   ├── api                    -- api 组
 │   ├── App.vue                -- 主页面
 │   ├── assets                 -- 静态资源
 │   ├── components             -- 全局组件
 │   ├── core                   -- gva 组件包
 │   │   ├── config.js          -- gva网站配置文件
 │   │   ├── gin-vue-admin.js   -- 注册欢迎文件
 │   │   └── global.js          -- 统一导入文件
 │   ├── directive              -- v-auth 注册文件
 │   ├── main.js                -- 主文件
 │   ├── permission.js          -- 路由中间件
 │   ├── pinia                  -- pinia 状态管理器，取代vuex
 │   │   ├── index.js           -- 入口文件
 │   │   └── modules            -- modules
 │   │       ├── dictionary.js
 │   │       ├── router.js
 │   │       └── user.js
 │   ├── router                 -- 路由声明文件
 │   │   └── index.js
 │   ├── style                      -- 全局样式
 │   │   ├── base.scss              --
 │   │   ├── basics.scss            --
 │   │   ├── element_visiable.scss  -- 此处可以全局覆盖 element-plus 样式
 │   │   ├── iconfont.css           -- 顶部几个icon的样式文件
 │   │   ├── main.scss              -- import: basics.scss,iconfont.css
 │   │   ├── mobile.scss            -- import: basics.scss
 │   │   └── newLogin.scss
 │   ├── utils                  -- 方法包库
 │   │   ├── asyncRouter.js     -- 动态路由相关
 │   │   ├── btnAuth.js         -- 动态权限按钮相关
 │   │   ├── bus.js             -- 全局mitt声明文件
 │   │   ├── date.js            -- 日期相关
 │   │   ├── dictionary.js      -- 获取字典方法 
 │   │   ├── downloadImg.js     -- 下载图片方法
 │   │   ├── format.js          -- 格式整理相关
 │   │   ├── image.js           -- 图片相关方法
 │   │   ├── page.js            -- 设置页面标题
 │   │   ├── request.js         -- 统一请求文件
 │   │   └── stringFun.js       -- 字符串文件
 |   ├── view                   -- 主要view代码
 |   |   ├── about              -- 关于我们
 |   |   ├── dashboard          -- 面板
 |   |   ├── error              -- 错误
 |   |   ├── example            -- 上传案例
 |   |   ├── iconList           -- icon列表
 |   |   ├── init               -- 初始化数据  
 |   |   ├── layout             -- layout约束页面 
 |   |   |   ├── aside          -- 侧边栏
 |   |   |   ├── bottomInfo     -- bottomInfo
 |   |   |   ├── screenfull     -- 全屏设置
 |   |   |   ├── setting        -- 系统设置
 |   |   |   └── index.vue      -- base 约束
 |   |   ├── login              -- 登录 
 |   |   ├── person             -- 个人中心 
 |   |   ├── superAdmin         -- 超级管理员操作
 |   |   ├── system             -- 系统检测页面
 |   |   ├── systemTools        -- 系统配置相关页面
 |   |   └── routerHolder.vue   -- page 入口页面 
 ├── vite.config.js             -- vite 配置文件
 └── yarn.lock
```

### *. 其它知识点
#### 5.1 VUE生命周期的变化
| Vue2.x        | Vue3            | Note  |
|  ----         |  -----          | ----- | 
| beforeCreate  | setup()         |
| created       | setup()         |
| beforeMount   | onBeforeMount   |
| mounted       | onMounted       |
| beforeUpdate  | onBeforeUpdate  |
| updated       | onUpdated       |
| beforeDestroy | onBeforeUnmount |
| destroyed     | onUnmounted     |
| errorCaptured | onErrorCaptured |


### 5.2 从npm环境下新建工程
- npm init -y
- npm i vite -D
- npm i @vitejs/plugin-vue -D
- npm i unplugin-vue-components -D          //按需引用插件
- npm i unplugin-auto-import -D             //按需引用import
- npm i vite-auto-import-resolvers -D       //按需引入项目API
- npm i vite-plugin-pages -D                //自动文件路由


#### 5.3 富文本
- npm install @wangeditor/editor --save
- npm install @wangeditor/editor-for-vue@next --save
- 在components 里增加一个组件在 editor 目录 index.vue
    ```
    <template>
    <div style="border: 1px solid #ccc">
        <Toolbar
        style="border-bottom: 1px solid #ccc"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        :mode="mode"
        />
        <Editor
        style="height: 500px; overflow-y: hidden;"
        v-model="valueHtml"
        :defaultConfig="editorConfig"
        :mode="mode"
        @onCreated="handleCreated"
        @onChange="handleChange"
        />
    </div>
    </template>

    <script>
    import '@wangeditor/editor/dist/css/style.css' // 引入 css

    import { onBeforeUnmount, ref, toRefs, shallowRef, onMounted } from 'vue'
    import { Editor, Toolbar } from '@wangeditor/editor-for-vue'

    export default {
    name:'WEditor',
    components: { Editor, Toolbar },
    props:{
        desc:String
    },
    setup(props, {emit}) {
        // 编辑器实例，必须用 shallowRef
        const editorRef = shallowRef()

        // 内容 HTML
        //const valueHtml = ref('<h1>hello</h1>')
        //var { desc } = toRefs(props)
        //const valueHtml = desc

        const valueHtml = ref('')
        //console.log(props)
        // 模拟 ajax 异步获取内容
        onMounted(() => {
        setTimeout(() => {
            valueHtml.value = props.desc
            //valueHtml.value = desc.value
        }, 100)
        })

        const toolbarConfig = {}
        const editorConfig = { placeholder: '请输入内容...' }

        // 组件销毁时，也及时销毁编辑器
        onBeforeUnmount(() => {
        const editor = editorRef.value
        if (editor == null) return
        editor.destroy()
        })

        const handleCreated = (editor) => {
        editorRef.value = editor // 记录 editor 实例，重要！
        }

        const handleChange = (editor) => {
        //console.log('S-change1:', editor.getHtml())
        //console.log('change2:', valueHtml)
        // emit('desc-change',editor.getHtml())
        emit('update:desc', editor.getHtml())
        };

        return {
        editorRef,
        valueHtml,
        mode: 'simple', // 或 'simple'default
        toolbarConfig,
        editorConfig,
        handleCreated,
        handleChange
        };
    }
    }
    </script>   
    ```
- 使用
    ```
    <el-form-item label="比赛简介:">
        <WEditor v-model:desc="formData.desc"></WEditor>
    </el-form-item>
    ```



#### 5.4 ES6 =>箭头函数


