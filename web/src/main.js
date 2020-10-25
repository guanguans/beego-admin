// Vue
import Vue from 'vue'
import i18n from './i18n'
import App from './App'
// 核心插件
import d2Admin from '@/plugin/d2admin'
// store
import store from '@/store/index'
import D2Crud from '@d2-projects/d2-crud'
import { d2CrudPlus } from 'd2-crud-plus'
import d2CrudX from 'd2-crud-x'
import { D2pFileUploader, D2pUploader, D2pTreeSelector } from 'd2p-extends'

// 菜单和路由设置
import router from './router'
import { frameInRoutes } from '@/router/routes'

// 核心插件
Vue.use(d2Admin)
Vue.use(D2Crud)

Vue.use(d2CrudX, { name: 'd2-crud-x' })

Vue.use(d2CrudPlus, {
  // 获取数据字典的请求方法
  commonOption () { // d2-crud option 全局设置，可以不用配置
    return {
      format: {
        response (ret) {
          // 这里默认配置是  return ret.data
          return ret
        },
        page: { // page接口返回的数据结构配置，
          request: { // 请求参数
            current: 'current', // 当前
            size: 'size'
          },
          response: { // 返回结果
            current: 'current', // 当前页码 ret.data.current
            size: 'size', // 每页条数，ret.data.size,
            // size: (data) => { return data.size }, //你也可以配置一个方法，自定义返回
            total: 'total', // 总记录数 ret.data.total
            records: 'records' // 列表数组 ret.data.records
          }
        }
      },
      formOptions: {
        defaultSpan: 12 // 默认的表单 span
      },
      options: {
        height: '100%' // 表格高度100%, 使用toolbar必须设置
      },
      pageOptions: {
        compact: true // 是否紧凑型页面
      },
      viewOptions: {
        disabled: false // 开启查看按钮
      }
    }
  }
})

Vue.use(D2pUploader)
Vue.use(D2pTreeSelector)

// 安装本扩展
Vue.use(D2pFileUploader, { d2CrudPlus })

new Vue({
  router,
  store,
  i18n,
  render: h => h(App),
  created () {
    // 处理路由 得到每一级的路由设置
    this.$store.commit('d2admin/page/init', frameInRoutes)
    // 设置顶栏菜单
    // this.$store.commit('d2admin/menu/headerSet', menuHeader)
    // // 初始化菜单搜索功能
    // this.$store.commit('d2admin/search/init', menuHeader)
  },
  mounted () {
    // 展示系统信息
    this.$store.commit('d2admin/releases/versionShow')
    // 用户登录后从数据库加载一系列的设置
    this.$store.dispatch('d2admin/account/load')
    // 获取并记录用户 UA
    this.$store.commit('d2admin/ua/get')
    // 初始化全屏监听
    this.$store.dispatch('d2admin/fullscreen/listen')
  },
  watch: {
    // 检测路由变化切换侧边栏内容
    '$route.matched': {
      handler (matched) {
        if (matched.length > 0) {
          const _side = this.$store.state.d2admin.menu.aside.filter(menu => menu.path === matched[0].path)
          this.$store.commit('d2admin/menu/asideSet', _side.length > 0 ? _side[0].children : [])
        }
      },
      immediate: true
    }
  }
}).$mount('#app')
