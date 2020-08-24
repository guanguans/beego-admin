<template>
  <div  v-loading="loading">
    <el-tree
      style="padding-top: 10px"
      class="filter-tree"
      :data="data"
      :props="defaultProps"
      node-key="id"
      show-checkbox
      :default-checked-keys="checked"
      @check="getCheckedKeys"
      default-expand-all
      :filter-node-method="filterNode"
      ref="tree">
    </el-tree>
  </div>
</template>

<script>
import api from '@/api'
export default {
  props: ['value'],
  data () {
    return {
      checked: [],
      loading: false,
      data: [],
      defaultProps: {
        children: 'children',
        label: 'title'
      }
    }
  },
  mounted () {
    this.loading = true
    api.SYS_MENU_ALL().then(m => {
      this.loading = false
      this.data = m
      if (this.value.length > 0) {
        var checked = this.value.map(function (item) {
          return item.menu_id
        })
        api.SYS_MENU_ROLE_CHECK({ check: checked }).then(res => {
          this.checked = res
        })
      }
    })
  },
  methods: {
    filterNode (value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    getCheckedKeys () {
      var checked = this.$refs.tree.getCheckedNodes(false, true)
      var keys = checked.map(function (item) {
        return item.id
      })
      this.$emit('input', keys)
    }
  }
}
</script>
