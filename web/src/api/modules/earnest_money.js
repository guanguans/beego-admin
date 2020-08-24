
export default ({ service, request, serviceForMock, requestForMock, mock, faker, tools }) => ({

  // 在途保证金列表
  SYS_EARNEST_MONEY_ONLINE_LIST (data = {}) {
    return request({
      url: '/admin/earnest/money/online',
      method: 'get',
      params: data
    })
  },
  // 担保账户保证金记录
  SYS_EARNEST_MONEY_GUARANTEE_LIST (data = {}) {
    return request({
      url: '/admin/earnest/money/guarantee',
      method: 'get',
      params: data
    })
  }
})
