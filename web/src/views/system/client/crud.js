export const crudOptions = (vm) => { // vm即this
  return {
    format: {
      flatData: {
        disabled: false, // 启用数据扁平化
        symbol: '#', // 默认使用#号连接(不能用.号)， key需要配置成 user#gender
        deleteOnUnFlat: true // 是否在反扁平化后清理扁平化过的痕迹，默认true
      }
    },
    columns: [
      {
        title: '名称',
        key: 'name',
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          rules: [ // 表单校验规则
            { required: true, message: '客户端名称不能为空' }
          ],
          disabled: false// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: 'access_key',
        key: 'id',
        form: { // form表单的配置
          disabled: true // 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: 'access_secret',
        key: 'secret',
        form: { // form表单的配置
          disabled: true // 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: 'redirect',
        key: 'redirect',
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        disabled: true,
        form: { // form表单的配置
          value: 'http://localhost/auth/callback',
          disabled: false // 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '状态',
        key: 'revoked',
        type: 'radio',
        dict: {
          data: [
            {
              color: 'success',
              label: '正常',
              value: false
            },
            {
              color: 'danger',
              label: '禁用',
              value: true
            }
          ]
        },
        search: {
          disabled: true
        },
        form: {
          disabled: true
        }
      }
    ],
    viewOptions: {
      disabled: true, // 开启view
      componentType: 'row' // 查看时使用哪种组件展示【form=使用表单组件,row=使用行展示组件】
    },
    formOptions: { // 编辑对话框及el-form的配置
      labelWidth: '70px',
      labelPosition: 'left',
      saveLoading: false,
      gutter: 20,
      defaultSpan: 20, // 默认表单字段所占宽度
      draggable: false, // 是否支持表单对话框拖拽
      updateTableDataAfterEdit: false // 添加和删除提交后，是否直接更新本地table的数据
    },
    searchOptions: { // 查询配置参数
      form: { // 默认搜索参数
        category_name: '' // 请求列表默认会带上此处配置参数
      },
      show: true, // 是否显示搜索工具条
      disabled: false, // 是否禁用搜索工具条
      debounce: { // 自动查询防抖,debounce:false关闭自动查询
        wait: 500 // 延迟500毫秒
      }
    },
    options: { // d2-crud及el-table的配置参数
      stripe: false,
      border: true,
      highlightCurrentRow: false, // 是否高亮选中行
      rowKey: 'category_id'
    },
    pagination: { // 翻页配置,更多配置参考el-pagination
      currentPage: 1,
      pageSize: 10,
      total: 1,
      disabled: true,
      storage: true // 本地保存用户每页条数修改，刷新不会丢失该修改，false=关闭
    }
  }
}
