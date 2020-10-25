import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

const meta = { auth: true }

export default {
  path: '/system',
  name: 'system',
  meta,
  redirect: { name: 'user' },
  component: layoutHeaderAside,
  children: (pre => [
    {
      path: '/system/role',
      name: 'role',
      component: _import('system/role'),
      meta: {
        title: '角色',
        ...meta
      }
    },
    {
      path: '/system/user',
      name: 'user',
      component: _import('system/user'),
      meta: {
        title: '管理员',
        ...meta
      }
    },
    {
      path: '/system/menu',
      name: 'menu',
      component: _import('system/menu'),
      meta: {
        title: '菜单权限',
        ...meta
      }
    },
    {
      path: '/system/client',
      name: 'client',
      component: _import('system/client'),
      meta: {
        title: '客户端',
        ...meta
      }
    }
  ])('system')
}
