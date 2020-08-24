<template>
  <d2-container>
    <template slot="header">角色列表</template>
    <d2-crud
      ref="admin"
      :columns="columns"
      :options="options"
      :loading="loading"
      :pagination="pagination"
      :data="data"
      add-title="新增角色"
      :add-template="addTemplate"
      :form-options="formOptions"
      :rowHandle="rowHandle"
      :add-rules="addRules"
      :edit-rules="editRules"
      edit-title="修改角色"
      :edit-template="editTemplate"
      @row-add="handleAdd"
      @row-edit="handleRowEdit"
      @row-remove="handleRowRemove"
      @dialog-cancel="handleDialogCancel"
      @dialog-open="handleDialogOpen"
      @pagination-current-change="paginationCurrentChange">
      <el-input style="width: 25%" slot="header" v-model="searchQuery.keyword" placeholder="角色名称"></el-input>
      <el-button slot="header" style="margin-bottom: 5px;margin-left: 10px;" @click="search">搜索</el-button>
      <el-button slot="header" type="primary" style="margin-bottom: 5px" @click="Create">新增</el-button>
    </d2-crud>
  </d2-container>
</template>

<script>
import Perms from './components/perms'
import api from '@/api'
export default {
  name: 'role',
  data () {
    var checkRoleName = (rule, value, callback) => {
      const reg = /[A-Za-z]/
      if (!reg.test(value)) {
        callback(new Error('角色名称必须是英文字母'))
      } else {
        callback()
      }
    }
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
          key: 'id',
          width: '50'
        },
        {
          title: '角色',
          key: 'name',
          width: '180'
        },
        {
          title: '角色名称',
          key: 'desc',
          width: '180'
        },
        {
          title: '创建时间',
          key: 'created_at'
        }
      ],
      // 添加表单
      addTemplate: {
        name: {
          title: '角色',
          value: ''
        },
        desc: {
          title: '角色名称',
          value: ''
        },
        perms: {
          title: '权限选择',
          value: [],
          component: {
            name: Perms
          }
        }
      },
      // 更新表单
      editTemplate: {
        name: {
          title: '角色',
          value: ''
        },
        desc: {
          title: '角色名称',
          value: ''
        },
        perms: {
          title: '权限选择',
          value: [],
          component: {
            name: Perms
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
        name: [{ required: true, message: '请输入角色', trigger: 'blur' }, { validator: checkRoleName, trigger: 'blur' }],
        desc: [{ required: true, message: '请输入角色名称', trigger: 'blur' }]
      },
      editRules: {
        name: [{ required: true, message: '请输入角色', trigger: 'blur' }, { validator: checkRoleName, trigger: 'blur' }],
        desc: [{ required: true, message: '请输入角色名称', trigger: 'blur' }]
      }
    }
  },
  mounted () {
    // 获取列表数据
    this.GetList()
  },
  methods: {
    // 获取数据列表
    GetList () {
      this.loading = true
      this.searchQuery.page = this.pagination.currentPage
      this.searchQuery.limit = this.pagination.pageSize
      api.SYS_ROLE_LIST(this.searchQuery).then(res => {
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
      api.SYS_ROLE_CREATE(row).then(res => {
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
      api.SYS_ROLE_UPDATE(row, row.id).then(res => {
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
      api.SYS_ROLE_DELETE(row.id).then(res => {
        this.$message({
          message: '删除成功',
          type: 'success'
        })
        done()
        this.GetUsersList()
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
        // row.password = ''
        // if (row.roles.length > 0) {
        //   row.roles = Array.from(row.roles.map((res,key) => {
        //     return res.id || row.roles[key]
        //   }))
        // }
      }
    }
  }
}
</script>

<style scoped>

</style>
