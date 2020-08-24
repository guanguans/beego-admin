import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

const meta = { auth: true }

export default {
  path: '/earnest_money',
  name: 'earnest_money',
  meta,
  redirect: { name: 'online' },
  component: layoutHeaderAside,
  children: (pre => [
    {
      path: '/earnest_money/guarantee',
      name: 'online',
      component: _import('earnest_money/guarantee'),
      meta: {
        title: '担保账户记录',
        ...meta
      }
    }
  ])('earnest_money')
}
