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