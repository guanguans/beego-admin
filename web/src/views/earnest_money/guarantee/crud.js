export const crudOptions = (vm) => { // vm即this
  return {
    expandRow: { // 或者直接传true,不显示title，不居中
      title: '详情',
      align: 'center'
    },
    columns: [
      {
        title: 'ID',
        key: 'id',
        width: 80,
        form: { // form表单的配置
          disabled: true // 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '流水号',
        key: 'earnest_money_sn',
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: false // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '用户ID',
        key: 'user_id',
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: false // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '手机号',
        key: 'phone',
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: false // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '账户类型',
        key: 'account_type',
        type: 'radio',
        dict: {
          data: [
            {
              color: 'primary',
              label: '普通商户',
              value: 1
            },
            {
              color: 'success',
              label: '高级商户',
              value: 2
            }
          ]
        },
        search: {
          disabled: false
        },
        form: {
          disabled: true
        }
      },
      {
        title: '资金',
        key: 'earnest_money',
        type: 'input',
        width: 100,
        disabled: true,
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '资金余额',
        key: 'earnest_log_balance',
        type: 'input',
        width: 100,
        disabled: true,
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '资金说明',
        key: 'money_desc',
        type: 'input',
        disabled: true,
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '支付类型',
        key: 'pay_type',
        type: 'radio',
        dict: {
          data: [
            {
              color: 'primary',
              label: '支付宝',
              value: 1
            },
            {
              color: 'success',
              label: '微信',
              value: 2
            }
          ]
        },
        search: {
          disabled: false
        },
        form: {
          disabled: true
        }
      },
      {
        title: '业务类型',
        key: 'business_desc',
        disabled: true,
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '备注',
        key: 'desc',
        type: 'input',
        disabled: true,
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '操作备注',
        key: 'action_desc',
        disabled: true,
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '支付时间',
        key: 'pay_time',
        type: 'datetime',
        component: {
          props: {
            // 行展示组件使用的dayjs，格式化标准与el-datepicker不一样
            format: 'YYYY-MM-DD HH:mm:ss'
          }
        },
        search: { // 查询配置，默认启用查询
          component: {
            name: 'daterange'
          },
          disabled: false // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
        }
      },
      {
        title: '入账时间',
        key: 'created_at',
        type: 'input',
        search: { // 查询配置，默认启用查询
          disabled: true // 【可选】true禁止查询,默认为false
        },
        form: { // form表单的配置
          disabled: true// 禁止添加输入与修改输入【可选】默认false
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
      rowKey: 'id'
    },
    rowHandle: {
      remove: { disabled: true, show: false }
    },
    pagination: { // 翻页配置,更多配置参考el-pagination
      currentPage: 1,
      pageSize: 10,
      total: 1,
      storage: true // 本地保存用户每页条数修改，刷新不会丢失该修改，false=关闭
    }
  }
}
