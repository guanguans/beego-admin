<template>
  <div  v-loading="loading">
    <div v-for="(perm) in routes" >
      <el-checkbox  v-model="perm.permChecked" :indeterminate="perm.isIndeterminates"  @change="(value) => handleCheckAllChange(value, perm.title)">
        {{perm.title}}
      </el-checkbox>
      <div style="margin: 15px 0;"></div>
      <el-checkbox-group style="padding-left: 30px" v-model="perm.childrenPerms" @change="(value) => handleCheckedCitiesChange(value, perm)">
        <el-checkbox v-for="child in perm.children" :label="child.title" :key="child.path">{{child.title}}</el-checkbox>
      </el-checkbox-group>
    </div>
  </div>
</template>

<script>
import api from '@/api'
export default {
  name: 'Perms2',
  props: ['value'],
  data () {
    return {
      loading: false,
      routes: []
    }
  },
  mounted () {
    this.loading = true
    api.SYS_MENU_ALL().then(m => {
      this.loading = false
      this.routes = m
      if (this.value !== null && this.value !== '' && this.value.length > 0) {
        const json = this.value
        this.routes.forEach(function (item) {
          item.isIndeterminates = false
          item.permChecked = false
          item.childrenPerms = []
          var menuId = []
          json.forEach(function (check) {
            if (item.children !== undefined && item.children.length > 0) {
              // 字元素小于父级的字元素就是全选
              item.children.length > 0 ? item.children.forEach(function (ch) {
                if (check.menu_id === ch.id) {
                  menuId.push(ch.id)
                  item.childrenPerms.push(ch.title)
                }
              }) : []
              if (menuId.length < item.children.length) {
                item.isIndeterminates = true
                item.permChecked = false
                return item
              } else if (menuId.length === item.children.length) {
                item.isIndeterminates = false
                item.permChecked = true
                return item
              }
            } else {
              item.isIndeterminates = false
              item.permChecked = true
            }
          })
        })
      } else {
        this.routes.forEach(function (item) {
          item.isIndeterminates = false
          item.permChecked = false
          item.childrenPerms = []
        })
      }
    })
  },
  methods: {
    handleCheckAllChange (val, title) {
      this.routes = this.routes.map(function (item) {
        if (item.title === title && item.children !== undefined) {
          item.childrenPerms = val ? item.children.map(function (c) {
            return c.title
          }) : []
          item.isIndeterminates = false
          return item
        }
        item.isIndeterminates = false
        return item
      })
    },
    handleCheckedCitiesChange (value, parent) {
      this.routes = this.routes.map(function (item) {
        if (parent.title === item.title) {
          if (item.childrenPerms.length > 0 && item.childrenPerms.length < parent.children.length) {
            item.isIndeterminates = true
            item.permChecked = false
            return item
          } else if (item.childrenPerms.length > 0 && item.childrenPerms.length === parent.children.length) {
            item.isIndeterminates = false
            item.permChecked = true
            return item
          } else if (item.childrenPerms.length === 0) {
            item.isIndeterminates = false
            item.permChecked = false
            return item
          }
        }
        return item
      })
    }
  },
  watch: {
    routes: {
      handler (newName, oldName) {
        // 取出已选中的数据
        const parents = []
        var arr = newName.filter(function (item) {
          if (item.childrenPerms !== undefined && item.childrenPerms.length > 0) {
            parents.push(item.id)
            item.children.forEach(children => {
              if (item.childrenPerms.indexOf(children.title) !== -1) {
                parents.push(children.id)
              }
            })
          } else if (item.permChecked) {
            parents.push(item.id)
          }
        })
        this.$emit('input', parents)
      },
      deep: true
    }
  }
}
</script>

<style scoped>

</style>
