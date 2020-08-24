import util from '@/libs/util.js'

export default {
  namespaced: true,
  mutations: {
    /**
     * @description 显示版本信息
     * @param {Object} state state
     */
    versionShow () {
      const processTitle = process.env.VUE_APP_TITLE || 'beego-admin'
      util.log.capsule(processTitle, 'v1.0.0', 'danger')
    }
  }
}
