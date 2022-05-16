<template>
  <div>
    <div class="title">错题模式</div>
    <div class="desc">本模式下可查看提示、详解</div>
    <el-divider></el-divider>
    <!-- <div class="modes"> -->
    <!--   <el-card> -->
    <!--     <div slot="header"> -->
    <!--       <span style="font-weight: bold">历史错题合集</span> -->
    <!--       <el-button type="text" @click="errorList"> 进入列表 </el-button> -->
    <!--     </div> -->
    <!--     <ul style="padding-left:14px"> -->
    <!--       <li>展示所有错题</li> -->
    <!--       <li>时间顺序排序</li> -->
    <!--       <li>可删除错题</li> -->
    <!--     </ul> -->
    <!--   </el-card> -->
    <!-- </div> -->
    <div class="modes">
      <el-card>
        <div slot="header">
          <span style="font-weight: bold">错题再练模式</span>
          <el-button type="text" @click="errorBook"> 进入练习 </el-button>
        </div>
        <ul style="padding-left:14px">
          <li>强化记忆知识点</li>
          <li>针对性训练</li>
          <li>快速练习</li>
        </ul>
      </el-card>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    async errorBook() {
      this.$router.push({
        name: 'myerror-mistakes',
      });
    },
    async errorList() {
      this.$router.push({
        name: 'myerror-list',
      });
    },
  },
};
</script>

<style scoped>
ul, ol, li {
  list-style-type: disc;
}
li {
  margin-bottom: 4px;
}
.title {
  font-size: 20px;
  font-weight: bold;
  padding: 10px;
}
.desc {
  padding: 8px;
}
.modes {
  margin-top: 20px;
  margin-left: 20px;
  float: left;
}
.question-option {
  display: flex;
  padding: 5px;
}
</style>
