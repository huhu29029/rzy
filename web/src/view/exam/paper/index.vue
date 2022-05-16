<template>
    <div v-if="!showExam">
      <p>{{ this.message }}</p>
      <el-button type="primary" @click="() => this.$router.back(-1)">返回</el-button>
    </div>
    <el-row v-else>
      <el-col :span="18">
        <div class="exam-title" id="title"></div>
        <div>
          <el-button type="info">{{ this.questionIdx+1 }}/{{ this.questions.length }}</el-button>
          <span v-if="this.question.type === 1">[选择题]</span>
          <span v-else-if="this.question.type === 2">[判断题]</span>
          <span v-else>[多选题]</span>
        </div>
        <div class="title" v-html="question.title"></div>
        <div v-if="question.type !== MULTI">
          <el-radio-group v-model="answer" @change="saveUserOption">
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
          <el-checkbox-group v-model="answerList" @change="saveUserOption">
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
                :class="(userOptions.has(ques.id) || question.id === ques.id) ? 'active' : ''"
              >
              {{ idx + 1 }}
              </li>
            </el-row>
          </div>
          <div align="center">
            <el-button
              type="primary"
              @click="submit">
              交卷
            </el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
</template>

<script>
import { submitAnswer, getExam } from '@/api/exam';
import { newXAPIData } from '@/api/xapi';
import { mapToObj, objToMap } from '@/utils/map-obj';

export default {
  data() {
    return {
      examTitle: '',
      question: {},
      questionIdx: 0,
      questions: [],

      answer: '', // 当前非多选题目的用户选项
      answerList: [], // 当前多选题目的用户选项
      userOptions: new Map(), // 用户答卷

      startTime: new Date(),
      MULTI: 3, // 题型：单选1 判断2 多选3

      showExam: true, // 是否可考试
      message: '',
    };
  },
  async created() {
    const userOptions = sessionStorage.getItem(`exam_${this.$route.query.examId}`);
    if (userOptions !== null && userOptions !== '{}') {
      try {
        await this.$confirm('检测到存在答题记录，是否恢复？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        });
        this.userOptions = objToMap(JSON.parse(userOptions));
      } catch (err) {
        this.$message({
          type: 'info',
          message: '已清除上次答题记录',
        });
        sessionStorage.clear();
      }
    }

    await this.getExam(this.$route.query.examId);

    this.getUserOption();
  },
  mounted() {
    window.addEventListener('beforeunload', this.persistUserOptions);
  },
  destroyed() {
    window.removeEventListener('beforeunload', this.persistUserOptions);
  },
  methods: {
    persistUserOptions() {
      sessionStorage.setItem(`exam_${this.$route.query.examId}`, JSON.stringify(mapToObj(this.userOptions)));
    },
    async getExam(id) {
      const res = await getExam(id);
      if (res.code !== 0) {
        this.showExam = false;
        this.message = res.msg;
        return;
      }
      const exam = res.data.exam;
      const { questions } = exam;

      for (let i = 0; i < questions.length; i++) {
        const question = questions[i];
        question.title = this.$md.render(question.title);
        question.options.map((option) => {
          option.title = this.$md.render(option.title);
          return option;
        });
        questions[i] = question;
      }

      this.examTitle = exam.title;
      this.questions = questions;
      this.questionIdx = 0;
      this.question = questions[this.questionIdx];
      await this.renderMathJax();
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
    saveUserOption() {
      const { type, id } = this.question;
      if (type === this.MULTI) {
        this.userOptions.set(id, this.answerList);
      } else {
        this.userOptions.set(id, this.answer);
      }
    },
    // 恢复用户选项
    getUserOption() {
      const { type, id } = this.question;
      if (this.userOptions.get(id) === undefined) {
        return;
      }
      if (type === this.MULTI) {
        this.answerList = this.userOptions.get(id);
      } else {
        this.answer = this.userOptions.get(id);
      }
    },
    // 数学公式
    async renderMathJax() {
      this.$nextTick(() => {
        window.MathJax.typeset();
      });
    },
    // 提交button的函数
    async submit() {
      if (this.userOptions.size < this.questions.length) {
        try {
          await this.$confirm('还有题目未做完，是否继续提交？', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          });
        } catch (err) {
          this.$message({
            type: 'info',
            message: '已取消',
          });
          return;
        }
      }

      const questions = [];
      for (let i = 0; i < this.questions.length; i++) {
        const curQues = this.questions[i];
        const { id } = curQues;
        let options = [];
        const userOption = this.userOptions.get(id);
        if (userOption !== undefined) {
          options = typeof userOption === 'number' ? [userOption] : userOption;
        }
        questions.push({
          id,
          options,
        });
      }
      const postData = {
        examId: Number(this.$route.query.examId),
        questions,
        startTime: this.startTime,
        endTime: new Date(),
      };
      const res = await submitAnswer(postData);
      if (res.code === 0) {
        // xapi: 数据埋点
        await newXAPIData({
          verb: 'examine',
          object: {
            info: `exam_id: ${this.examId}`,
          },
        });
        const id = res.data.id;
        await this.getResult(id);
      } else {
        this.$message({
          type: 'error',
          message: '提交失败',
        });
      }
    },
    async getResult(id) {
      this.$router.push({
        name: 'h-exam-detail',
        params: {
          id,
        },
      });
    },
  },
};
</script>

<style scoped>
.exam-title {
  font-size: 20px;
  font-weight: bold;
  padding: 17px;
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
.colon {
  font-size: 30px;
}
.el-divider--horizontal {
  margin: 0px;
}
</style>
