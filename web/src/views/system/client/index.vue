<template>
    <d2-container :class="{'page-compact':crud.pageOptions.compact}">
        <template slot="header">客户端</template>
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

        </d2-crud-x>
    </d2-container>
</template>

<script>
import { crudOptions } from './crud'
import { d2CrudPlus } from 'd2-crud-plus'
import api from '@/api'
export default {
  name: 'testPage',
  mixins: [d2CrudPlus.crud], // 最核心部分，继承d2CrudPlus.crud
  methods: {
    getCrudOptions () { return crudOptions(this) },
    pageRequest (query) {
      return api.SYS_CLIENT_LIST(query).then(function (ret) {
        ret.records = ret
        return ret
      })
    },
    addRequest (row) { return api.SYS_CLIENT_CREATE(row) },
    updateRequest (row) { return api.SYS_CLIENT_UPDATE(row, row.id) },
    delRequest (row) { return api.SYS_CLIENT_DELETE(row.id) }
  }
}
</script>
