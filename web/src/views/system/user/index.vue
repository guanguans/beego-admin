<template>
  <d2-container>
    <template slot="header">管理员列表</template>
    <d2-crud
      ref="admin"
      :columns="columns"
      :options="options"
      :loading="loading"
      :pagination="pagination"
      :data="data"
      add-title="新增管理员"
      :add-template="addTemplate"
      :form-options="formOptions"
      :rowHandle="rowHandle"
      :add-rules="addRules"
      :edit-rules="editRules"
      edit-title="修改管理员"
      :edit-template="editTemplate"
      @row-add="handleAdd"
      @row-edit="handleRowEdit"
      @row-remove="handleRowRemove"
      @dialog-cancel="handleDialogCancel"
      @dialog-open="handleDialogOpen"
      @pagination-current-change="paginationCurrentChange">
      <el-input style="width: 25%" slot="header" v-model="searchQuery.keyword" placeholder="用户名/手机号"></el-input>
      <el-button slot="header" style="margin-bottom: 5px;margin-left: 10px;" @click="search">搜索</el-button>
      <el-button slot="header" type="primary" style="margin-bottom: 5px" @click="Create">新增</el-button>
    </d2-crud>
  </d2-container>
</template>

<script>
import api from '@/api'
export default {
  name: 'user',
  data () {
    return {
      searchQuery: {
        keyword: ''
      },
      loading: true,
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 0
      },
      // 表格数据
      columns: [
        {
          title: 'ID',
          key: 'uid',
          width: '50'
        },
        {
          title: '用户名',
          key: 'username',
          width: '180'
        },
        {
          title: '真实姓名',
          key: 'real_name',
          width: '180'
        },
        {
          title: '手机号',
          key: 'phone'
        },
        {
          title: '邮箱',
          key: 'email'
        },
        {
          title: '上次登录时间',
          key: 'last_login_at'
        },
        {
          title: '创建时间',
          key: 'created_at'
        }
      ],
      // 添加表单
      addTemplate: {
        username: {
          title: '用户名',
          value: ''
        },
        real_name: {
          title: '真实姓名',
          value: ''
        },
        password: {
          title: '密码',
          value: '',
          component: {
            type: 'password',
            showPassword: true
          }
        },
        phone: {
          title: '手机号',
          value: ''
        },
        email: {
          title: '邮箱',
          value: ''
        },
        roles: {
          title: '角色',
          value: '',
          component: {
            name: 'el-select',
            options: [],
            multiple: true
          }
        }
      },
      // 更新表单
      editTemplate: {
        username: {
          title: '用户名',
          value: ''
        },
        real_name: {
          title: '真实姓名',
          value: ''
        },
        password: {
          title: '密码',
          value: '',
          component: {
            type: 'password',
            showPassword: true,
            placeholder: '为空代表不修改密码'
          }
        },
        phone: {
          title: '手机号',
          value: ''
        },
        email: {
          title: '邮箱',
          value: ''
        },
        roles: {
          title: '角色',
          value: '',
          component: {
            name: 'el-select',
            options: [],
            multiple: true
          }
        }
      },
      // 数据操作按钮
      rowHandle: {
        columnHeader: '操作按钮',
        edit: {
          icon: 'el-icon-edit',
          text: '编辑',
          size: 'small'
        },
        remove: {
          icon: 'el-icon-delete',
          size: 'small',
          fixed: 'right',
          confirm: true
        }
      },
      //  操作的表单设置
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false,
        gutter: 20
      },
      // 管理员数据
      data: [],
      // 表格个性化设置
      options: {
        border: true
      },
      // 验证规则
      addRules: {
        username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
        phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
        roles: [{ required: true, message: '请选择角色', trigger: 'change' }]
      },
      editRules: {
        username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
        roles: [{ required: true, message: '请选择角色', trigger: 'change' }]
      }
    }
  },
  mounted () {
    // 获取列表数据
    this.GetList()
    this.GetRolesList()
  },
  methods: {
    // 获取角色
    async GetRolesList () {
      api.SYS_ROLE_ALL().then(res => {
        var optionArr = res.map(function (item) {
          return {
            label: item.desc,
            value: item.id
          }
        })
        this.addTemplate.roles.component.options = optionArr
        this.editTemplate.roles.component.options = optionArr
      })
    },
    // 获取数据列表
    GetList () {
      this.loading = true
      this.searchQuery.page = this.pagination.currentPage
      this.searchQuery.limit = this.pagination.pageSize
      api.SYS_ADMIN_LIST(this.searchQuery).then(res => {
        this.loading = false
        this.data = res.list
        this.pagination.total = res.count
      })
    },
    // 分页切换
    paginationCurrentChange (currentPage) {
      this.pagination.currentPage = currentPage
      this.GetList()
    },
    // 新增数据
    Create () {
      this.$refs.admin.showDialog({
        mode: 'add'
      })
    },
    // 保存新增数据
    handleAdd (row, done) {
      this.formOptions.saveLoading = true
      api.SYS_ADMIN_CREATE(row).then(res => {
        this.$message({
          message: '新增成功',
          type: 'success'
        })
        done()
        this.GetList()
        this.formOptions.saveLoading = false
      }).catch(() => {
        this.formOptions.saveLoading = false
      })
    },
    // 修改
    handleRowEdit ({ index, row }, done) {
      this.formOptions.saveLoading = true
      api.SYS_ADMIN_UPDATE(row, row.uid).then(res => {
        this.$message({
          message: '更新成功',
          type: 'success'
        })
        done()
        this.GetList()
        this.formOptions.saveLoading = false
      }).catch(() => {
        this.formOptions.saveLoading = false
      })
    },
    // 删除数据
    handleRowRemove ({ index, row }, done) {
      api.SYS_ADMIN_DELETE(row.uid).then(res => {
        this.$message({
          message: '删除成功',
          type: 'success'
        })
        done()
        this.GetList()
      })
    },
    // 搜索
    search () {
      this.GetList()
    },
    // 取消操作弹窗
    handleDialogCancel (done) {
      done()
    },
    // 打开弹窗前置操作
    handleDialogOpen ({ mode, row }) {
      if (mode === 'edit') {
        row.password = ''
        if (row.roles.length > 0) {
          row.roles = Array.from(row.roles.map((res, key) => {
            return res.id || row.roles[key]
          }))
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
