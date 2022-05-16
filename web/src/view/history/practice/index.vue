<template>
  <div>
    <el-table :data="tableData" style="width: 100%" :cell-style="cellStyle">
      <el-table-column label="日期" width="180" align="center">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.date }}</span>
        </template>
      </el-table-column>
      <el-table-column label="做题情况" width="140" align="center">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.result }}</span>
        </template>
      </el-table-column>
      <el-table-column label="题干" prop="title" align="center">
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="140" align="center">
        <template slot-scope="scope">
          <el-button @click="resultDetail(scope.row)" type="text" size="small"
            >详情</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page.sync="pagination.currentPage"
      :page-sizes="[5, 10, 15, 20]"
      :page-size="pagination.pageSize"
      :style="{ float: 'right', padding: '20px' }"
      layout="total, sizes, prev, pager, next, jumper"
      :total="pagination.total">
    </el-pagination>
    <el-dialog
      :before-close="handleClose"
      :title="dialogViewTitle"
      :visible.sync="dialogViewVisible"
    >
      <div class="el-dialog__body">
        <div>
          <span v-if="this.question.type === 1">[选择题]</span>
          <span v-else-if="this.question.type === 2">[判断题]</span>
          <span v-else>[多选题]</span>
        </div>
        <div class="title" v-html="question.title"></div>
        <div class="resultShow">
          <span>正确答案：{{ questionAnswer }}</span>
          <span style="margin-left: 20px">你的答案：{{ questionChoosed }}</span>
          <span id="judge"> </span>
        </div>
        <div v-if="question.type !== MULTI">
          <el-radio-group v-model="answerResult">
            <el-radio
              v-for="option in question.options"
              :label="option.id"
              :key="option.id"
              class="question-option"
            >
              <div v-html="option.title"></div>
            </el-radio>
          </el-radio-group>
        </div>
        <div v-else>
          <el-checkbox-group v-model="answerResultMulti">
            <el-checkbox
              v-for="option in question.options"
              :label="option.id"
              :key="option.id"
              class="question-option"
            >
              <div v-html="option.title"></div>
            </el-checkbox>
          </el-checkbox-group>
        </div>
        <div class="analysis">解析：{{ question.analysis }}</div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getQuestionsHistory, getQuestionResult } from '@/api/questions.js';
import { formatTimeToStr } from '@/utils/date.js';
import { newXAPIData } from '@/api/xapi';

export default {
  data() {
    return {
      tableData: [],
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 1,
      },
      dialogViewVisible: false,
      dialogViewTitle: '做题详情',
      question: '',
      MULTI: 3, // 题型：单选1 判断2 多选3
      answerResult: '',
      answerResultMulti: [],
      questionChoosed: '',
      questionAnswer: '',
      result: 0,

      dialogPracticeVisible: false,
      dialogPracticeTitle: '练习',
    };
  },
  async created() {
    await this.getQuestionsHistory();
  },
  methods: {
    async resultDetail(row) {
      // xapi: 数据埋点
      await newXAPIData({
        verb: 'review',
        object: {
          info: `questionId: ${row.id}`,
        },
      });
      const res = await getQuestionResult(row.id);
      const { data } = res;
      const { result } = data;
      const question = result;
      question.title = this.$md.render(question.title);
      question.options.map((option) => {
        option.title = this.$md.render(option.title);
        return option;
      });
      const optionRes = [];
      const choose = [];
      const answer = [];
      let optionId = 1;
      if (question.isRight) {
        this.result = 1;
      } else {
        this.result = 0;
      }
      for (let j = 0; j < question.options.length; j++) {
        question.options[j].id = optionId;
        if (question.options[j].isChoosed) {
          optionRes.push(optionId);
          const num = 65 + j;
          const letter = String.fromCharCode(num);
          choose.push(letter);
        }
        if (question.options[j].isRight) {
          const num = 65 + j;
          const letter = String.fromCharCode(num);
          answer.push(letter);
        }
        optionId += 1;
      }
      this.question = question;
      this.$nextTick(() => {
        window.MathJax.typeset();
      });

      if (this.question.type === this.MULTI) {
        if (optionRes.length === 0) {
          this.questionChoosed = '空';
        } else {
          this.answerResultMulti = optionRes;
          let optionString = '';
          for (let i = 0; i < choose.length; i++) {
            optionString += choose[i];
          }
          this.questionChoosed = optionString;
        }
        let answerString = '';
        for (let i = 0; i < answer.length; i++) {
          answerString += answer[i];
        }
        this.questionAnswer = answerString;
      } else {
        if (optionRes.length === 0) {
          this.questionChoosed = '空';
        } else {
          this.answerResult = optionRes[0];
          this.questionChoosed = choose[0];
        }
        this.questionAnswer = answer[0];
      }
      this.dialogViewVisible = true;
      this.$nextTick(() => {
        const judge = document.getElementById('judge');
        if (this.result === 1) {
          // judge.innerHTML = 'test'
          judge.innerHTML = '(正确)';
          judge.style.color = '#71c647';
        } else {
          judge.innerHTML = '(错误)';
          judge.style.color = '#f56e6e';
        }
      });
    },
    cellStyle({ row, columnIndex }) {
      if (row.result === '对' && columnIndex === 1) {
        return 'color: #71c647';
      } if (row.result === '错' && columnIndex === 1) {
        return 'color: #f56e6e';
      }
      return 'color: #1a1a1b';
    },
    handleClose(done) {
      done();
    },
    async getQuestionsHistory(
      page = this.pagination.currentPage,
      size = this.pagination.pageSize,
    ) {
      const res = await getQuestionsHistory(page, size);
      const { data } = res;
      const { questionHistory, total } = data;
      this.pagination.total = total;
      const table = [];
      for (let i = 0; i < questionHistory.length; i++) {
        const question = questionHistory[i];
        const { questionResult } = question;
        const data = {};
        data.id = question.id;
        const time = formatTimeToStr(question.startTime);
        data.date = time;
        data.title = questionResult.title;
        if (questionResult.isRight) {
          data.result = '对';
        } else {
          data.result = '错';
        }
        table.push(data);
      }
      this.tableData = table;
    },
    async handleSizeChange(val) {
      this.pagination.pageSize = val;
      await this.getQuestionsHistory();
    },
    async handleCurrentChange(val) {
      this.pagination.currentPage = val;
      await this.getQuestionsHistory();
    },
  },
};
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
.analysis {
  padding: 10px;
}
.resultShow {
  height: 20px;
  border: 1px #409eff solid;
  padding: 5px;
  margin-bottom: 10px;
}
.title {
  font-size: 18px;
  padding: 15px;
  line-height: 25px;
}
.question-option {
  display: flex;
  padding: 5px;
}
.el-dialog__body{height: 60vh;overflow: auto;}
</style>
