<template>
  <div>
    <div class="title">学习计划</div>
    <div class="desc">动态规划刷题指南</div>
    <el-divider></el-divider>
    <div class="modes">
      <el-card>
        <div slot="header">
          <span style="font-weight: bold">导论验收</span>
          <el-button type="text" @click="chooseDate"> 开始练习 </el-button>
        </div>
        <ul style="padding-left:14px">
          <li>每日习题</li>
          <li>针对导论</li>
        </ul>
      </el-card>
    </div>
    <el-dialog
      :before-close="handleClose"
      :title="dialogTitle"
      :visible.sync="dialogFormVisible"
    >
    <el-radio-group v-model="dateNum">
        <el-radio
            v-for="option in options"
            :label="option.id"
            :key="option.id"
            class="question-option"
            >
            <div v-html="option.label"></div>
        </el-radio>
    </el-radio-group>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="dailyplan" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialogFormVisible: false,
      dialogTitle: '请选择练习阶段',
      dateNum: 1,
      options: [],
      postData: [],
      remark: 0,
    };
  },
  methods: {
    async chooseDate() {
      for (let i = 1; i < 5; i++) {
        const data = {};
        data.id = i;
        data.label = `day${i.toString()}`;
        this.options.push(data);
      }
      this.dialogFormVisible = true;
    },
    closeDialog() {
      this.options = [];
      this.dialogFormVisible = false;
    },
    handleClose(done) {
      this.options = [];
      done();
    },
    async dailyplan() {
      if (this.options.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择一项进行练习！',
        });
      }
      this.$router.push({
        name: 'studyplan-dailyplan',
        query: {
          dateNum: this.dateNum,
        },
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
