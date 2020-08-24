<template>
  <d2-container :class="{'page-compact':crud.pageOptions.compact}">
    <template slot="header">用户担保保证金列表</template>
    <d2-crud-x
      ref="d2Crud"
      v-bind="_crudProps"
      v-on="_crudListeners">

      <div slot="header">
        <crud-search ref="search" :options="crud.searchOptions" @submit="handleSearch"  />
        <el-button size="small" type="primary"  @click="addRow"><i class="el-icon-plus"/> 新增</el-button>
        <crud-toolbar :search.sync="crud.searchOptions.show"
                      :compact.sync="crud.pageOptions.compact"
                      :columns="crud.columns"
                      @refresh="doRefresh()"
                      @columns-filter-changed="handleColumnsFilterChanged"/>
      </div>
      <template slot="expandSlot" slot-scope="scope">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="资金：">
            <span>{{scope.row.earnest_money}}</span>
          </el-form-item>
          <el-form-item label="资金余额：">
            <span>{{scope.row.earnest_log_balance}}</span>
          </el-form-item>
          <el-form-item label="资金说明：">
            <span>{{scope.row.money_desc}}</span>
          </el-form-item>
          <el-form-item label="业务类型：">
            <span>{{scope.row.business_desc}}</span>
          </el-form-item>
          <el-form-item label="备注：">
            <span>{{scope.row.desc}}</span>
          </el-form-item>
          <el-form-item label="操作备注：">
            <span>{{scope.row.action_desc}}</span>
          </el-form-item>
        </el-form>
      </template>
    </d2-crud-x>
  </d2-container>
</template>

<script>
import { crudOptions } from './crud'
import { d2CrudPlus } from 'd2-crud-plus'
import api from '@/api'
export default {
  name: 'online',
  mixins: [d2CrudPlus.crud],
  methods: {
    getCrudOptions () { return crudOptions(this) },
    pageRequest (query) { return api.SYS_EARNEST_MONEY_GUARANTEE_LIST(query) }
  }
}
</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
