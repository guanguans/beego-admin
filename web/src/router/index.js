import Vue from 'vue'
import VueRouter from 'vue-router'

// 进度条
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

import store from '@/store/index'
import util from '@/libs/util.js'
import api from '@/api'
// 路由数据
import routes from './routes'

// fix vue-router NavigationDuplicated
const VueRouterPush = VueRouter.prototype.push
VueRouter.prototype.push = function push (location) {
  return VueRouterPush.call(this, location).catch(err => err)
}
const VueRouterReplace = VueRouter.prototype.replace
VueRouter.prototype.replace = function replace (location) {
  return VueRouterReplace.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

// 导出路由 在 main.js 里使用
const router = new VueRouter({
  routes
})

/**
 * 路由拦截
 * 权限验证
 */
router.beforeEach(async (to, from, next) => {
  // 确认已经加载多标签页数据 https://github.com/d2-projects/d2-admin/issues/201
  await store.dispatch('d2admin/page/isLoaded')
  // 确认已经加载组件尺寸设置 https://github.com/d2-projects/d2-admin/issues/198
  await store.dispatch('d2admin/size/isLoaded')
  // 进度条
  NProgress.start()
  // 关闭搜索面板
  store.commit('d2admin/search/set', false)
  // 验证当前路由所有的匹配中是否需要有登录验证的
  if (to.matched.some(r => r.meta.auth)) {
    // 这里暂时将cookie里是否存有token作为验证是否登录的条件
    // 请根据自身业务需要修改
    util.cookies.get('roles', '')
    util.cookies.get('routes', '')
    const token = util.cookies.get('token')
    if (token && token !== 'undefined') {
      // 登录之后获取菜单
      await api.SYS_USER_INFO().then(user => {
        util.cookies.set('roles', user.roles)
        router.app.$store.commit('d2admin/menu/headerSet', user.menus)
        router.app.$store.commit('d2admin/menu/asideSet', user.menus)
        router.app.$store.commit('d2admin/search/init', user.menus)
        util.cookies.set('routes', user.menus)
      })
      const routes = util.cookies.get('routes')
      const roles = util.cookies.get('roles')
      // 如果是超级管理员就不用验证任何路由
      if (roles.indexOf('super_admin') !== -1) {
        next()
      } else {
        if (routes) {
          var menus = JSON.parse(routes)
          var isTrue = menus.some(m => {
            if (m.path === to.path) {
              return true
            } else {
              if (m.children !== undefined && m.children.length > 0) {
                return m.children.some(c => {
                  if (c.path === to.path) {
                    return true
                  } else {
                    if (c.children !== undefined && c.children.length > 0) {
                      return c.children.some(c2 => {
                        if (c2.path === to.path) {
                          return true
                        }
                      })
                    }
                  }
                })
              }
            }
          })
          if (isTrue) {
            next()
          } else {
            // 没有权限跳转 404
            next({ name: '404' })
            NProgress.done()
          }
        } else {
          // 没有权限跳转 404
          next({ name: '404' })
          NProgress.done()
        }
      }
    } else {
      // 没有登录的时候跳转到登录界面
      // 携带上登陆成功之后需要跳转的页面完整路径
      next({
        name: 'login',
        query: {
          redirect: to.fullPath
        }
      })
      // https://github.com/d2-projects/d2-admin/issues/138
      NProgress.done()
    }
  } else {
    // 不需要身份校验 直接通过
    next()
  }
})

router.afterEach(to => {
  // 进度条
  NProgress.done()
  // 多页控制 打开新的页面
  store.dispatch('d2admin/page/open', to)
  // 更改标题
  util.title(to.meta.title)
})

// function menuIsPass(menus,to) {
//   return menus.some(m => {
//     if (m.path === to.path) {
//       return true
//     } else {
//       if (m.children !== undefined && m.children.length > 0) {
//         menuIsPass(m.children,to)
//       }
//     }
//   })
// }
export default router
