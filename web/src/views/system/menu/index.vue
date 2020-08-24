<template>
  <div>
    <d2-container>
    <template slot="header">菜单权限列表</template>
    <el-button @click="handAdd" type="primary" size="small" style="margin-bottom: 5px">新增</el-button>
    <el-table
      :data="data"
      v-loading="loading"
      style="width: 100%;margin-bottom: 20px;"
      row-key="id"
      border
      default-expand-all
      :tree-props="{children: 'children'}">

      <el-table-column
        prop="id"
        label="ID"
        align="center"
        width="100">
      </el-table-column>
      <el-table-column
        prop="title"
        label="菜单名称"
        align="center"
        width="250"
      >
      </el-table-column>
      <el-table-column
        prop="icon"
        align="center"
        label="icon图标"
        width="100"
      >

        <template slot-scope="scope">
          <d2-icon :name="scope.row.icon"/>
        </template>
      </el-table-column>
      <el-table-column
        prop="path"
        align="center"
        label="路由"
        width="500"
      >
      </el-table-column>
      <el-table-column
        align="left"
        label="操作"
      >
        <template slot-scope="scope">
          <el-button @click="handEdit(scope.row)" plain size="small" icon="el-icon-edit">操作</el-button>
          <el-button @click="handleRowRemove(scope.row)" type="danger" size="small" icon="el-icon-delete">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </d2-container>
    <el-dialog :title="edit === false ? '添加菜单' : '编辑菜单'" :visible.sync="dialogFormVisible">

    <el-form :model="form">
      <el-form-item label="父级菜单" :label-width="formOptions.labelWidth">
        <el-cascader
          v-model="form.parent_id"
          :options="parentMenu"
          filterable
          placeholder="默认顶级菜单"
          :props="{ checkStrictly: true, value:'id', label:'title'}"
          clearable>
        </el-cascader>
      </el-form-item>
      <el-form-item label="菜单名称" :label-width="formOptions.labelWidth">
        <el-input v-model="form.menu_name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="菜单ICON" :label-width="formOptions.labelWidth">
        <d2-icon-select v-model="form.icon"/>
      </el-form-item>
      <el-form-item label="菜单路由" :label-width="formOptions.labelWidth">
        <el-input v-model="form.menu_path" placeholder="菜单路由为项目中的path路径" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="是否展示" :label-width="formOptions.labelWidth">
        <el-radio v-model="form.is_show"   label="yes">是</el-radio>
        <el-radio v-model="form.is_show" label="no">否</el-radio>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button v-if="edit === false" :loading="formOptions.saveLoading" type="primary" @click="handleAdd">确 定</el-button>
      <el-button v-else :loading="formOptions.saveLoading" type="primary" @click="handleRowEdit">更新</el-button>
    </div>
  </el-dialog>
  </div>
</template>

<script>
import api from '@/api'
export default {
  name: 'menus',
  data () {
    return {
      searchQuery: {
        keyword: ''
      },
      dialogFormVisible: false,
      loading: true,
      edit: false,
      edit_id: 0,
      parentMenu: [],
      // 表格数据
      columns: [
        {
          title: 'ID',
          key: 'id',
          width: '50'
        },
        {
          title: '菜单名称',
          key: 'title',
          width: '180'
        },
        {
          title: 'icon图标',
          key: 'title',
          width: '180'
        },
        {
          title: '路由',
          key: 'path'
        }
      ],
      //  操作的表单设置
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false,
        gutter: 20
      },
      form: {
        menu_name: '',
        menu_path: '',
        icon: '',
        is_show: 'yes',
        parent_id: 0
      },
      // 管理员数据
      data: []
    }
  },
  mounted () {
    // 获取列表数据
    this.GetList()
  },
  methods: {
    // 获取父级菜单
    GetParentMenu () {
      api.SYS_MENU_PARENT_LIST().then(res => {
        this.parentMenu = res
      })
    },
    // 获取数据列表
    GetList () {
      this.loading = true
      api.SYS_MENU_LIST().then(res => {
        this.loading = false
        this.data = res.list
      })
    },
    // 保存新增数据
    handleAdd (row, done) {
      this.formOptions.saveLoading = true
      if (this.form.parent_id.length > 1) {
        this.form.parent_id = this.form.parent_id.pop()
      } else if (this.form.parent_id.length === 1) {
        this.form.parent_id = this.form.parent_id.shift()
      }
      api.SYS_MENU_CREATE(this.form).then(res => {
        this.$message({
          message: '新增成功',
          type: 'success'
        })
        this.GetList()
        this.dialogFormVisible = false
        this.formOptions.saveLoading = false
      }).catch(() => {
        this.dialogFormVisible = false
        this.formOptions.saveLoading = false
      })
    },
    handAdd () {
      this.GetParentMenu()
      this.dialogFormVisible = true
      this.form.parent_id = 0
      this.form.menu_path = ''
      this.form.menu_name = ''
      this.form.is_show = 'yes'
      this.form.icon = ''
      this.edit = false
      this.edit_id = 0
    },
    handEdit (row) {
      this.GetParentMenu()
      this.edit = true
      this.edit_id = row.id
      this.dialogFormVisible = true
      this.form.parent_id = row.parent_id
      this.form.menu_path = row.path
      this.form.menu_name = row.title
      this.form.is_show = row.is_show
      this.form.icon = row.icon
    },
    // 修改
    handleRowEdit () {
      this.formOptions.saveLoading = true
      if (this.form.parent_id.length > 1) {
        this.form.parent_id = this.form.parent_id.pop()
      } else if (this.form.parent_id.length === 1) {
        this.form.parent_id = this.form.parent_id.shift()
      }
      api.SYS_MENU_UPDATE(this.form, this.edit_id).then(res => {
        this.$message({
          message: '更新成功',
          type: 'success'
        })
        this.edit = false
        this.edit_id = 0
        this.GetList()
        this.dialogFormVisible = false
        this.formOptions.saveLoading = false
      }).catch(() => {
        this.edit = false
        this.edit_id = 0
        this.dialogFormVisible = false
        this.formOptions.saveLoading = false
      })
    },
    // 删除数据
    handleRowRemove (row) {
      this.$confirm('确认删除么？', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }).then(() => {
        api.SYS_MENU_DELETE(row.id).then(res => {
          this.$message({
            message: '删除成功',
            type: 'success'
          })
          this.GetList()
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
  .title-group {
    margin-top: 20px;
    margin-bottom: 10px;
    &:first-child {
      margin-top: 0px;
    }
    .title {
      margin: 0px;
    }
    .sub-title {
      margin: 0px;
      color: $color-text-sub;
      font-size: 10px;
    }
  }
</style>
