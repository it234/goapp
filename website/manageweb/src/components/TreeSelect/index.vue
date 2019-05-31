<template>
  <el-select
    :placeholder="props.placeholder"
    :value="valueTitle"
    :clearable="clearable"
    @clear="clearHandle"
  >
    <el-option :value="valueTitle" :label="valueTitle" class="options">
      <el-tree
        id="tree-option"
        ref="selectTree"
        :accordion="accordion"
        :data="options"
        :props="props"
        :node-key="props.value"
        :default-expanded-keys="defaultExpandedKey"
        @node-click="handleNodeClick"
      />
    </el-option>
  </el-select>
</template>

<script>
export default {
  name: 'SelectTree',
  props: {
    props: {
      type: Object,
      default: () => {
        return {
          value: 'id',
          label: 'title',
          children: 'children'
        }
      }
    },
    options: {
      type: Array,
      default: () => { return [] }
    },
    value: {
      type: Number,
      default: () => { return null }
    },
    clearable: {
      type: Boolean,
      default: () => { return true }
    },
    accordion: {
      type: Boolean,
      default: () => { return false }
    }
  },
  data() {
    return {
      valueId: this.value,
      valueTitle: '',
      defaultExpandedKey: []
    }
  },
  watch: {
    value() {
      this.valueId = this.value
      this.initHandle()
    }
  },
  mounted() {
    this.initHandle()
  },
  methods: {
    initHandle() {
      if (this.valueId) {
        this.valueTitle = this.$refs.selectTree.getNode(this.valueId).data[this.props.label]
        this.$refs.selectTree.setCurrentKey(this.valueId)
        this.defaultExpandedKey = [this.valueId]
      }
      this.initScroll()
    },
    initScroll() {
      this.$nextTick(() => {
        const scrollWrap = document.querySelectorAll('.el-scrollbar .el-select-dropdown__wrap')[0]
        const scrollBar = document.querySelectorAll('.el-scrollbar .el-scrollbar__bar')
        scrollWrap.style.cssText = 'margin: 0px; max-height: none; overflow: hidden;'
        scrollBar.forEach(ele => {
          ele.style.width = 0
        })
      })
    },
    handleNodeClick(node) {
      this.valueTitle = node[this.props.label]
      this.valueId = node[this.props.value]
      this.$emit('getValue', this.valueId)
      this.defaultExpandedKey = []
    },
    clearHandle() {
      this.valueTitle = ''
      this.valueId = null
      this.defaultExpandedKey = []
      this.clearSelected()
      this.$emit('getValue', null)
    },
    clearSelected() {
      const allNode = document.querySelectorAll('#tree-option .el-tree-node')
      allNode.forEach((element) => element.classList.remove('is-current'))
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-scrollbar .el-scrollbar__view .el-select-dropdown__item {
  height: auto;
  max-height: 274px;
  padding: 0;
  overflow: hidden;
  overflow-y: auto;
}
.el-select-dropdown__item.selected {
  font-weight: normal;
}
ul li >>> .el-tree .el-tree-node__content {
  height: auto;
  padding: 0 20px;
}
.el-tree-node__label {
  font-weight: normal;
}
.el-tree >>> .is-current .el-tree-node__label {
  color: #409eff;
  font-weight: 700;
}
.el-tree >>> .is-current .el-tree-node__children .el-tree-node__label {
  color: #606266;
  font-weight: normal;
}
/* 开发禁用 */
/* .el-tree-node:focus>.el-tree-node__content{
    background-color:transparent;
    background-color: #f5f7fa;
    color: #c0c4cc;
    cursor: not-allowed;
  }
  .el-tree-node__content:hover{
    background-color: #f5f7fa;
  } */
</style>
