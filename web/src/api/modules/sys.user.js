
export default ({ service, request, serviceForMock, requestForMock, mock, faker, tools }) => ({
  /**
   * @description 登录
   * @param {Object} data 登录携带的信息
   */
  SYS_USER_LOGIN (data = {}) {
    // 接口请求
    return request({
      url: '/account/password',
      method: 'post',
      data
    })
  },

  SYS_USER_INFO (data = {}) {
    return request({
      url: '/account/admin',
      method: 'post',
      async: false,
      data
    })
  },
  // 管理员列表
  SYS_ADMIN_LIST (data = {}) {
    return request({
      url: '/admin/user',
      method: 'get',
      params: data
    })
  },
  // 管理员添加
  SYS_ADMIN_CREATE (data = {}) {
    return request({
      url: '/admin/user',
      method: 'post',
      data
    })
  },
  // 管理员更新
  SYS_ADMIN_UPDATE (data = {}, id) {
    return request({
      url: '/admin/user/' + id,
      method: 'put',
      data
    })
  },
  // 管理员删除
  SYS_ADMIN_DELETE (id) {
    return request({
      url: '/admin/user/' + id,
      method: 'delete'
    })
  },

  // 菜单列表
  SYS_MENU_LIST (data = {}) {
    return request({
      url: '/admin/menu',
      method: 'get',
      params: data
    })
  },
  // 获取所有菜单
  SYS_MENU_ALL () {
    return request({
      url: '/admin/menu/0',
      method: 'get',
      async: false
    })
  },
  SYS_MENU_ROLE_CHECK (data) {
    return request({
      url: '/admin/menu/check',
      method: 'post',
      async: false,
      data
    })
  },
  // 获取父及菜单
  SYS_MENU_PARENT_LIST () {
    return request({
      url: '/admin/menu/create',
      method: 'get'
    })
  },
  // 新增菜单
  SYS_MENU_CREATE (data = {}) {
    return request({
      url: '/admin/menu',
      method: 'post',
      data
    })
  },
  // 更新菜单
  SYS_MENU_UPDATE (data = {}, id) {
    return request({
      url: '/admin/menu/' + id,
      method: 'put',
      data
    })
  },
  // 删除菜单
  SYS_MENU_DELETE (id) {
    return request({
      url: '/admin/menu/' + id,
      method: 'delete'
    })
  },
  // 角色列表
  SYS_ROLE_LIST (data = {}) {
    return request({
      url: '/admin/role',
      method: 'get',
      params: data
    })
  },
  // 获取所有角色
  SYS_ROLE_ALL () {
    return request({
      url: '/admin/role/0',
      method: 'get'
    })
  },
  // 角色列表
  SYS_ROLE_CREATE (data = {}) {
    return request({
      url: '/admin/role',
      method: 'post',
      data
    })
  },
  // 更新角色
  SYS_ROLE_UPDATE (data = {}, id) {
    return request({
      url: '/admin/role/' + id,
      method: 'put',
      data
    })
  },
  // 删除角色
  SYS_ROLE_DELETE (id) {
    return request({
      url: '/admin/role/' + id,
      method: 'delete'
    })
  },
  // 客户端列表
  SYS_CLIENT_LIST (data = {}) {
    return request({
      url: '/oauth/clients',
      method: 'get'
    })
  },
  SYS_CLIENT_CREATE (data = {}) {
    return request({
      url: '/oauth/clients',
      method: 'post',
      data
    })
  },
  SYS_CLIENT_UPDATE (data = {}, id) {
    return request({
      url: '/oauth/clients/' + id,
      method: 'put',
      data
    })
  },
  SYS_CLIENT_DELETE (id) {
    return request({
      url: '/oauth/clients/' + id,
      method: 'delete'
    })
  }
})
