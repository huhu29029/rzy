<template>
  <div>
    <el-table :data="tableData" style="width: 100%">
      <el-table-column label="日期" width="250" align="center">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.date }}</span>
        </template>
      </el-table-column>
      <el-table-column label="试卷标题" align="left">
        <template slot-scope="scope">
          <i class="el-icon-document"></i>
          <span style="margin-left: 10px">{{ scope.row.title }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="240" align="center">
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
          <el-button type="info">{{ this.order }}/{{ this.numbers }}</el-button>
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
        <!-- button区 -->
        <el-row>
          <el-button
            v-if="order === 1"
            type="primary"
            disabled
            @click="resultLast"
            >上一题</el-button
          >
          <el-button v-else type="primary" @click="resultLast"
            >上一题</el-button
          >
          <el-button
            v-if="order === numbers"
            type="primary"
            disabled
            @click="resultNext"
            >下一题</el-button
          >
          <el-button v-else type="primary" @click="resultNext"
            >下一题</el-button
          >
          <el-button
            type="primary"
            style="margin-left: 100px"
            id="return"
            @click="close"
            >返回</el-button
          >
        </el-row>
        <div class="card">
          <el-row>
            <li
              v-for="(id, i) in idList"
              :id="id"
              :key="id"
              @click="liOnclickRes(i + 1)"
            >
              {{ i + 1 }}
            </li>
          </el-row>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getUserInfoExamHistory } from '@/api/questions.js';
import { getResult } from '@/api/exam.js';
import { formatTimeToStr } from '@/utils/date.js';

export default {
  data() {
    return {
      id: 0,
      tableData: [],
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 1,
      },
      dialogViewVisible: false,
      dialogViewTitle: '做题详情',
      idList: [],
      order: 1,
      numbers: 0,
      question: '',
      questions: [],
      MULTI: 3, // 题型：单选1 判断2 多选3
      answerResult: '',
      answerResultMulti: [],
      questionsChoosed: [],
      questionChoosed: '',
      questionsAnswer: [],
      questionAnswer: '',
      result: [], // 每题做对做错，做对为1，做错为0
      optionsChoosed: [],
    };
  },
  async created() {
    this.id = sessionStorage.getItem('userInfoId');
    await this.getExamHistory();
  },
  methods: {
    async resultDetail(row) {
      // 关闭查看考试信息功能
      // this.$message({
      //   type: 'info',
      //   message: '考试后不允许查看考试信息'
      // })
      const res = await getResult(row.id);
      const { data } = res;
      const { result } = data;
      const { questions } = result;
      let optionId = 100;
      this.numbers = questions.length;
      this.order = 1;
      let questionId = 1;
      for (let i = 0; i < questions.length; i++) {
        const question = questions[i];
        if (question.isRight) {
          this.result.push(1);
        } else {
          this.result.push(0);
        }
        const optionRes = [];
        const choose = [];
        const answer = [];
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
        this.optionsChoosed.push(optionRes);
        this.questionsChoosed.push(choose);
        this.questionsAnswer.push(answer);
        question.title = this.$md.render(question.title);
        question.options.map((option) => {
          option.title = this.$md.render(option.title);
          return option;
        });
        questions[i] = question;
        this.idList.push(questionId);
        questionId += 1;
      }
      this.questions = questions;
      this.question = questions[0];
      this.dialogViewVisible = true;
      this.$nextTick(() => {
        this.showResult();
      });
    },
    async resultLast() {
      this.question = this.questions[this.order - 2];
      this.order -= 1;
      await this.renderMathJax();
      await this.showResult();
    },
    async resultNext() {
      this.question = this.questions[this.order];
      this.order += 1;
      await this.renderMathJax();
      await this.showResult();
    },
    async liOnclickRes(order) {
      this.order = order;
      this.question = this.questions[order - 1];
      await this.renderMathJax();
      await this.showResult();
    },
    async showResult() {
      const { order } = this;
      this.answerResultMulti = [];
      for (let i = 0; i < this.numbers; i++) {
        if (this.result[i] === 1) {
          document.getElementById(this.idList[i]).className = 'correct';
        } else {
          document.getElementById(this.idList[i]).className = 'incorrect';
        }
      }
      if (this.result[order - 1] === 1) {
        document.getElementById(this.idList[order - 1]).className = 'currentCorrect';
      } else {
        document.getElementById(this.idList[order - 1]).className = 'currentIncorrect';
      }
      const option = this.optionsChoosed[order - 1];
      const optionChose = this.questionsChoosed[order - 1];
      const answer = this.questionsAnswer[order - 1];
      if (this.question.type === this.MULTI) {
        if (option.length === 0) {
          this.questionChoosed = '空';
        } else {
          this.answerResultMulti = option;
          let optionString = '';
          for (let i = 0; i < optionChose.length; i++) {
            optionString += optionChose[i];
          }
          this.questionChoosed = optionString;
        }
        let answerString = '';
        for (let i = 0; i < answer.length; i++) {
          answerString += answer[i];
        }
        this.questionAnswer = answerString;
      } else {
        if (option.length === 0) {
          this.questionChoosed = '空';
        } else {
          this.answerResult = option[0];
          this.questionChoosed = optionChose[0];
        }
        this.questionAnswer = answer[0];
      }
      const res = this.result[order - 1];
      const judge = document.getElementById('judge');
      if (res === 1) {
        judge.innerHTML = '(正确)';
        judge.style.color = '#71c647';
      } else {
        judge.innerHTML = '(错误)';
        judge.style.color = '#f56e6e';
      }
    },
    async getExamHistory(
      page = this.pagination.currentPage,
      size = this.pagination.pageSize,
    ) {
      const res = await getUserInfoExamHistory(this.id, page, size);
      const { data } = res;
      const { examHistory, total } = data;
      this.pagination.total = total;
      const table = [];
      for (let i = 0; i < examHistory.length; i++) {
        const exam = examHistory[i];

        table.push({
          id: exam.examId,
          date: formatTimeToStr(exam.startTime),
          title: exam.examTitle,
        });
      }
      this.tableData = table;
    },
    async renderMathJax() {
      this.$nextTick(() => {
        window.MathJax.typeset();
      });
    },
    handleClose(done) {
      this.idList = [];
      done();
    },
    async close() {
      this.dialogViewVisible = false;
      this.idList = [];
    },
    async handleSizeChange(val) {
      this.pagination.pageSize = val;
      await this.getExamHistory();
    },
    async handleCurrentChange(val) {
      this.pagination.currentPage = val;
      await this.getExamHistory();
    },

  },
};
</script>
<style scoped>
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
.card li {
  margin-top: 1px;
  margin-left: 10px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  width: 30px;
  height: 30px;
  color: #606266;
  text-align: center;
  line-height: 30px;
  float: left;
  cursor: pointer;
  background-color: white;
}
.card li:hover {
  background-color: #ecf5ff;
}
.card .active {
  background-color: #ecf5ff;
}
.card .passive {
  background-color: white;
}
.card .correct {
  background-color: #f0f9eb;
}
.card .currentCorrect {
  background-color: #71c647;
}
.card .incorrect {
  background-color: #fef0f0;
}
.card .currentIncorrect {
  background-color: #f56e6e;
}
.el-dialog__body{height: 60vh;overflow: auto;}
</style>
