<template>
    <el-row>
      <el-col :span="18">
        <div v-if="!showAnalysis">
          <div>
            <span v-if="this.question.type === 1">[选择题]</span>
            <span v-else-if="this.question.type === 2">[判断题]</span>
            <span v-else>[多选题]</span>
          </div>
          <div class="title" v-html="question.title"></div>
          <div v-if="question.type !== MULTI">
            <el-radio-group v-model="answer">
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
            <el-checkbox-group v-model="answerList">
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
          <div class="submit">
            <el-button v-if="firstRemark === 1" type="primary" plain disabled
              >上一题</el-button
            >
            <el-button v-else type="primary" plain @click="lastQuestion"
              >上一题</el-button
            >
            <el-button @click="submit" type="primary" style="margin-left: 50px"
              >提交</el-button
            >
            <el-divider></el-divider>
            <el-tooltip
              class="item"
              effect="dark"
              content="点击去往涅槃之路平台学习!"
              placement="bottom"
            >
              <el-button
                type="info"
                size="small"
                style="margin-left: 10px"
                @click="jumpPlatform"
                icon="el-icon-thumb"
                >涅槃之路平台</el-button
              >
            </el-tooltip>
          </div>
        </div>
        <div v-else class="analysis-wrapper">
          <div class="title" v-html="question.title"></div>
          <div>
            <div
              v-for="option in question.options"
              :key="option.id"
              class="question-option"
            >
              <div
                v-html="option.title"
                :class="option.isRight ? 'right' : ''"
              ></div>
            </div>
          </div>
          <div class="resultShow" v-if="showLastQuestion">
            <span>正确答案：{{ questionAnswer }}</span>
            <span style="margin-left: 20px">你的答案：{{ questionChoosed }}</span>
            <!-- <span id="judge"></span> -->
            <span class="judgeCorrect" v-if="this.LastJudge">（正确）</span>
            <span class="judgeWrong" v-else>（错误）</span>
          </div>
          <div class="analysis">解析：{{ question.analysis }}</div>
          <div>
            <el-button
              v-if="!showLastQuestion"
              @click="continuePractice"
              type="primary"
              >继续练习</el-button
            >
            <el-button
              v-else
              type="primary"
              plain
              @click="nextQuestion"
              style="margin-top: 10px"
              >下一题</el-button
            >
            <el-divider></el-divider>
            <el-button
              type="info"
              size="small"
              style="margin-left: 10px"
              @click="jumpPlatform"
              >知识点学习资源</el-button
            >
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <el-card>
          <div slot="header">
            <span>错题再练</span>
          </div>
          <div class="card_tips">错题共：{{ questionsTotal }}个</div>
          <div class="card_tips">当前第：{{ questionNum+1 }}题</div>
        </el-card>
      </el-col>
    </el-row>

</template>

<script>
import {
  getWrongQuestion,
  getQuestion,
  // delWrongQuestion
} from '@/api/questions';
import { newXAPIData } from '@/api/xapi';

export default {
  data() {
    return {
      question: '',
      questionId: 0,
      answer: '',
      answerList: [],
      MULTI: 3,
      questionNum: 0,
      questionsTotal: 0,
      MAX_question: 100, // 请求题目数量
      showLastQuestion: false,
      questionAnswer: '',
      questionChoosed: '',
      LastJudge: false,
      firstRemark: 0,
      questionRemark: '',
      showAnalysis: false,
      questionList: [],
      lastQuestionId: '',
      nextQuestionId: '',
    };
  },
  async created() {
    this.firstRemark = 1;
    await this.getWrongQuestion();
  },

  methods: {
    async getWrongQuestion() {
      const res = await getWrongQuestion(1, this.MAX_question);
      const { data } = res;
      const { questionHistory, total } = data;
      if (total < 0) {
        this.$message({
          type: 'warning',
          message: '还没有错题哦，先去练习吧！',
        });
        this.$router.push({
          name: 'questions-guide',
        });
      } else {
        for (let i = 0; i < questionHistory.length; i++) {
          const question = questionHistory[i];
          this.questionList[i] = question.questionId;
        }
        this.questionId = this.questionList[this.questionNum];
        this.questionsTotal = questionHistory.length;
        await this.getQuestion();
      }
    },
    async getQuestion() {
      this.questionId = this.questionList[this.questionNum];
      let res = {};
      res = await getQuestion(Number(this.questionId));
      const { data } = res;
      const { question } = data;
      if (this.showLastQuestion) {
        if (this.questionRemark === 'last') {
          this.showAnalysis = true;
          const { options } = question;
          let answer = '';
          let choose = '';
          for (let i = 0; i < options.length; i++) {
            if (options[i].isRight) {
              const num = 65 + i;
              const letter = String.fromCharCode(num);
              answer += letter;
            }
            for (let j = 0; j < this.options.length; j++) {
              if (this.options[j] === options[i].id) {
                const num = 65 + i;
                const letter = String.fromCharCode(num);
                choose += letter;
              }
            }
          }
          if (this.lastRes) {
            this.LastJudge = true;
          } else {
            this.LastJudge = false;
          }
          this.questionAnswer = answer;
          this.questionChoosed = choose;
        } else {
          this.showLastQuestion = false;
          this.showAnalysis = false;
        }
      }
      if (this.questionRemark === 'continue') {
        this.showAnalysis = false;
      }
      question.title = this.$md.render(question.title);
      question.options.map((option) => {
        option.title = this.$md.render(option.title);
        return option;
      });
      this.question = question;
      this.$nextTick(() => {
        window.MathJax.typeset();
      });
    },
    async submit() {
      if (this.questionNum + 1 >= this.questionsTotal) {
        this.$message({
          type: 'success',
          message: '已完成所有错题练习！',
        });
        setTimeout(() => {
          this.$router.push({
            name: 'myerror-guide',
          });
        }, 5000);
      }
      // xapi: 数据埋点
      await newXAPIData({
        verb: 'repractice',
        object: {
          info: `questionId: ${this.question.id}`,
        },
      });
      const { type, options } = this.question;
      const rightAnswer = options.reduce((acc, option) => {
        if (option.isRight) {
          acc.push(option.id);
        }
        return acc;
      }, []);

      let answer = [];
      if (type === this.MULTI) {
        answer = this.answerList;
      } else if (this.answer !== '') {
        answer = [this.answer];
      }
      this.options = answer;
      this.answer = '';
      this.answerList = [];
      if (JSON.stringify(rightAnswer) === JSON.stringify(answer)) {
        this.lastQuestionId = this.question.id;
        this.lastRes = true;
        this.firstRemark = 0;
        await this.nextQuestion();
      } else {
        this.showAnalysis = true;
        this.$nextTick(() => {
          window.MathJax.typeset();
        });
      }
    },
    async lastQuestion() {
      this.nextQuestionId = this.question.id;
      this.questionNum -= 1;
      this.questionRemark = 'last';
      this.showLastQuestion = true;
      await this.getQuestion();
    },
    async nextQuestion() {
      this.lastQuestionId = this.question.id;
      this.questionNum += 1;
      this.questionRemark = 'next';
      this.firstRemark = 0;
      await this.getQuestion();
    },
    async continuePractice() {
      this.lastRes = false;
      this.lastQuestionId = this.question.id;
      this.questionNum += 1;
      this.questionRemark = 'continue';
      this.firstRemark = 0;
      await this.getQuestion();
    },
    async jumpPlatform() {
      const { question } = this;
      const { tags } = question;
      let url = 'https://www.ourspark.org/point/course_content/';
      if (tags.length !== 0) {
        const knowlegde = tags[0];
        const { layer1 } = knowlegde;
        if (layer1 === 'JavaScript') {
          url = `${url
            + this.urlJavaScript
          },${
            knowlegde.layer2
          },${
            knowlegde.layer3}`;
        } else if (layer1 === '微信小程序') {
          url = `${url
            + this.urlWeChat
          },${
            knowlegde.layer2
          },${
            knowlegde.layer3}`;
        }
        // xapi: 数据埋点
        await newXAPIData({
          verb: 'study',
          object: {
            info: url,
          },
        });
        window.open(url);
      } else {
        this.$message.info('该题目暂无学习资源！');
      }
    },
  },
};
</script>

<style scoped>
.card_tips {
  font-size: 14px;
  padding: 2px;
}
.title {
  font-size: 18px;
  padding: 20px;
}
.question-option {
  display: flex;
  padding: 5px;
}
.submit {
  margin-top: 10px;
}
.judgeCorrect {
  color: #71c647;
}
.judgeWrong {
  color: #f56e6e;
}

.analysis-wrapper .question-option {
  display: flex;
  padding: 5px;
  margin-left: 15px;
}
.analysis-wrapper .right {
  color: green;
}
.analysis-wrapper .right:before {
  content: "\2714";
  color: #008100;
  position: absolute;
  left: 0px;
}
.analysis-wrapper .analysis {
  padding: 10px;
}
.resultShow {
  height: 20px;
  border: 1px #409eff solid;
  padding: 5px;
  margin-bottom: 10px;
}
</style>

<style>
img {
  max-width: 100%;
}
</style>
