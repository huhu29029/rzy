<template>
  <div>
    <div>
      <div class="assessment">
        <p>试卷：{{ this.result.title }}</p>
        <p>得分：{{ this.result.score }} (总分：{{ this.result.totalScore }})</p>
      </div>
      <div class="assessment">
        <el-row>
          <el-col :span="18">
            <div>
              <el-button type="info">{{ this.questionIdx+1 }}/{{ this.questions.length }}</el-button>
              <span v-if="this.question.type === 1">[选择题]</span>
              <span v-else-if="this.question.type === 2">[判断题]</span>
              <span v-else>[多选题]</span>
            </div>
            <div class="title" v-html="question.title"></div>
            <div v-if="question.type !== MULTI">
              <el-radio-group v-model="answer">
                <el-radio
                  v-for="(option, idx) in question.options"
                  :label="idx"
                  :key="idx"
                  class="question-option"
                >
                  <div v-html="option.title"></div>
                </el-radio>
              </el-radio-group>
            </div>
            <div v-else>
              <el-checkbox-group v-model="answerList">
                <el-checkbox
                  v-for="(option, idx) in question.options"
                  :label="idx"
                  :key="idx"
                  class="question-option"
                >
                  <div v-html="option.title"></div>
                </el-checkbox>
              </el-checkbox-group>
            </div>
            <el-row>
              <el-button type="primary" @click="prevQuestion">上一题</el-button>
              <el-button type="primary" @click="nextQuestion">下一题</el-button>
            </el-row>
          </el-col>
          <el-col :span="6">
            <el-card class="card-box">
              <div class="card">
                <el-row>
                  <li
                    v-for="(ques, idx) in questions"
                    @click="toQuestion(idx)"
                    :key="idx"
                  >
                  {{ idx + 1 }}
                  </li>
                </el-row>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
      <div class="assessment">
        <el-button style="margin-top: 10px" type="primary" @click="() => this.$router.back(-1)">返回</el-button>
        <p>根据此次作答情况，推荐{{ recommendations.questionIds.length }}道题目</p>
        <el-button style="margin-top: 10px" type="primary" icon="el-icon-edit" @click="goRecommendations">去练习</el-button>
      </div>
      <div class="assessment">
        <p>或去涅槃之路平台学习基础知识</p>
        <el-button style="margin-top: 10px" type="primary" @click="toNiepan">去学习</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import { getResult } from '@/api/exam.js';
import { newXAPIData } from '@/api/xapi';

export default {
  data() {
    return {
      questions: [],
      question: {},
      questionIdx: 0,
      answerList: [],
      answer: '',
      MULTI: 3, // 题型：单选1 判断2 多选3
      result: {},
      recommendations: {},
      userOptions: new Map(),
    };
  },
  async created() {
    const { id } = this.$route.params;
    await this.getResult(id);
  },
  methods: {
    async getResult(id) {
      const res = await getResult(id);
      if (res.code === 0) {
        this.result = res.data.result;
        this.recommendations = res.data.recommendations;

        const questions = this.result.questions;
        for (let i = 0; i < questions.length; i++) {
          const question = questions[i];
          question.title = this.$md.render(question.title);
          question.options.map((option) => {
            option.title = this.$md.render(option.title);
            return option;
          });
          this.saveUserOption(question, i);
          questions[i] = question;
        }
        this.questions = questions;
        this.question = questions[this.questionIdx];
        this.getUserOption();
      } else {
        this.$message({
          type: 'error',
          message: '判卷失败',
        });
      }
    },
    prevQuestion() {
      if (this.questionIdx <= 0) {
        this.$message.info('已是第一题！');
        return;
      }
      this.questionIdx--;
      this.question = this.questions[this.questionIdx];

      this.getUserOption();
    },
    nextQuestion() {
      if (this.questionIdx >= this.questions.length - 1) {
        this.$message.info('已是最后一题！');
        return;
      }
      this.questionIdx++;
      this.question = this.questions[this.questionIdx];

      this.getUserOption();
    },
    toQuestion(idx) {
      this.questionIdx = idx;
      this.question = this.questions[this.questionIdx];

      this.getUserOption();
    },
    // 记录用户选项
    saveUserOption(question, questionIdx) {
      const { type, options } = question;
      if (type === this.MULTI) {
        const answerList = options.reduce((acc, cur, idx) => {
          if (cur.isChoosed) {
            acc.push(idx);
          }
          return acc;
        }, []);
        this.userOptions.set(questionIdx, answerList);
      } else {
        let answer = 0;
        for (let i = 0; i < options.length; i++) {
          if (options[i].isChoosed) {
            answer = i;
            break;
          }
        }
        this.userOptions.set(questionIdx, answer);
      }
    },
    // 恢复用户选项
    getUserOption() {
      const { type } = this.question;
      const idx = this.questionIdx;
      if (this.userOptions.get(idx) === undefined) {
        return;
      }
      if (type === this.MULTI) {
        this.answerList = this.userOptions.get(idx);
      } else {
        this.answer = this.userOptions.get(idx);
      }
    },
    goRecommendations() {
      this.$router.push({
        name: 'exam-exercises',
        params: {
          examTitle: this.examTitle,
          questions: this.recommendations.questionIds,
        },
      });
    },
    async toNiepan() {
      const url = 'https://www.ourspark.org/';
      // xapi: 数据埋点
      await newXAPIData({
        verb: 'study',
        object: {
          info: url,
        },
      });
      window.open(url);
    },
  },
};
</script>

<style scoped>
.question-option {
  display: flex;
  padding: 5px;
}
.assessment {
  margin-top: 30px;
  margin-left: 20px;
}
.assessment p {
  font-size: 1.3em;
  margin-top: 20px;
}
.title {
  font-size: 18px;
  padding: 15px;
  line-height: 25px;
}
.card li {
  margin-top: 5px;
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
.card .active{
  background-color: #ecf5ff;
}
</style>
