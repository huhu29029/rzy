<template>
  <div>
    <el-empty v-if="this.question === undefined" description="暂无题目"></el-empty>
    <el-row v-else>
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
            <el-button type="primary" plain @click="prevQuestion"
              >上一题</el-button
            >
            <el-button @click="submit" type="primary" style="margin-left: 50px"
              >提交</el-button
            >
            <el-divider></el-divider>
            <div>
              <h6 style="margin: 3px">所属知识点</h6>
              <el-tag
                v-for="tag in question.tags"
                :key="tag.id"
                effect="plain"
                >
                {{ `${tag.layer1}/${tag.layer2}/${tag.layer3}/${tag.layer4}` }}
              </el-tag>
            </div>
            <div style="margin-top: 10px">
              <h6 style="margin: 3px">知识点资源推荐</h6>
              <el-tooltip
                class="item"
                effect="dark"
                content="去往涅槃之路平台学习相关知识点!"
                placement="bottom"
              >
                <el-button
                  type="info"
                  size="small"
                  @click="toNiepan"
                  icon="el-icon-position"
                  >涅槃之路平台</el-button
                >
              </el-tooltip>
            </div>
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
          <div class="resultShow">
            <span>正确答案：{{ selectedAnswer.rightAnswerStr }}</span>
            <span style="margin-left: 20px">你的答案：{{ selectedAnswer.answerStr }}</span>
            <span class="judgeCorrect" v-if="selectedAnswer.isRight">（正确）</span>
            <span class="judgeWrong" v-else>（错误）</span>
          </div>
          <div class="analysis">解析：{{ question.analysis }}</div>
          <div>
            <el-button
              @click="continuePractice"
              type="primary"
              >继续练习</el-button
            >
            <el-divider></el-divider>
            <div>
              <h6 style="margin: 3px">所属知识点</h6>
              <el-tag
                v-for="tag in question.tags"
                :key="tag.id"
                effect="plain"
                >
                {{ `${tag.layer1}/${tag.layer2}/${tag.layer3}/${tag.layer4}` }}
              </el-tag>
            </div>
            <div style="margin-top: 10px">
              <h6 style="margin: 3px">资源推荐</h6>
              <el-tooltip
                class="item"
                effect="dark"
                content="去往涅槃之路平台学习相关知识点!"
                placement="bottom"
              >
                <el-button
                  type="info"
                  size="small"
                  @click="toNiepan"
                  icon="el-icon-position"
                  >涅槃之路平台</el-button
                >
              </el-tooltip>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <el-card>
          <div slot="header">
            <span>本次练习</span>
          </div>
          <div>题目：{{ statistic.questionsCount}}个</div>
          <div style="margin-top: 5px">时间：{{ statistic.totalTime }}分钟</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {
  getRandomQuestion,
  getQuestionWithKT,
  getQuestionWithWrong,
  getChapterRandomQuestion,
  submitAnswer,
} from '@/api/questions';
import { newXAPIData } from '@/api/xapi';

const dayjs = require('dayjs');
const relativeTime = require('dayjs/plugin/relativeTime');

dayjs.extend(relativeTime);

const MODE_OBJ = {
  random: 'random', // 随机模式
  chapters: 'chapters', // 章节模式
  kt: 'kt', // 算法模式
  wrong: 'wrong', // 错题重练模式
};

export default {
  data() {
    return {
      question: undefined,
      selectedAnswer: {
        value: '',
        isRight: false,
      }, // 已做题目的答案
      answer: '', // 用于单选
      answerList: [], // 用于多选
      startTime: new Date(),
      MULTI: 3, // 题型：单选1 判断2 多选3
      showAnalysis: false,
      niepan: {
        jsUrl: '602df4b15165a14f347bb489',
        mpUrl: '602df4b15165a14f347bb48b',
      },
      prevQuestions: [], // 做题记录
      nextQuestions: [], // 未做题记录
      chapters: undefined, // 章节练习模式
      statistic: {
        questionsCount: 0,
        startTime: (new Date()).getTime(),
        totalTime: 0,
      },
      mode: MODE_OBJ.random, // 练习模式
    };
  },
  async created() {
    document.addEventListener('visibilitychange', this.visibilitychange);
    const { chapters, kt, isWrong } = this.$route.query;
    if (chapters !== undefined) {
      this.chapters = JSON.parse(chapters);
      this.mode = MODE_OBJ.chapters;
    } else if (kt !== undefined) {
      this.mode = MODE_OBJ.kt;
    } else if (isWrong !== undefined) {
      this.mode = MODE_OBJ.wrong;
    }

    await this.getQuestion();
    await this.testTimer();
  },
  destroyed() {
    document.removeEventListener('visibilitychange', this.visibilitychange);
  },
  methods: {
    async toNiepan() {
      const { question } = this;
      const { tags } = question;

      let url = 'https://www.ourspark.org/point/course_content/';
      if (tags.length === 0) {
        this.$message.info('该题目暂无学习资源！');
        return;
      }
      const knowlegde = tags[0];
      const { layer1, layer2, layer3 } = knowlegde;
      if (layer1 === 'JavaScript') {
        url = `${url}${this.niepan.jsUrl}`;
      } else if (layer1 === '微信小程序') {
        url = `${url}${this.niepan.mpUrl}`;
      }
      url = `${url},${layer2},${layer3}`;

      // xapi: 数据埋点
      await newXAPIData({
        verb: 'study',
        object: {
          info: url,
        },
      });

      window.open(url);
    },
    async getQuestion() {
      let res;
      switch (this.mode) {
        case MODE_OBJ.chapters:
          res = await getChapterRandomQuestion({
            chapterIds: this.chapters,
          });
          break;
        case MODE_OBJ.kt:
          res = await getQuestionWithKT();
          break;
        case MODE_OBJ.wrong:
          res = await getQuestionWithWrong();
          break;
        default:
          res = await getRandomQuestion();
          break;
      }
      if (res.code === 1000) {
        this.$message.info('已无题目！, 3s后返回上页。');
        setTimeout(() => {
          this.$router.back(-1);
        }, 2000);
        return;
      } if (res.code !== 0) {
        this.$message.error('获取题目出错，请返回重试！');
      }
      const { data } = res;
      const { question } = data;

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
      const { id, type, options } = this.question;
      // xapi: 数据埋点
      await newXAPIData({
        verb: `practice_${this.mode}`,
        object: {
          info: `question_id: ${id}`,
        },
      });
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
      const postData = {
        questionId: id,
        options: answer,
        startTime: this.startTime,
        endTime: new Date(),
      };
      this.options = answer;
      await submitAnswer(postData);

      const isRight = JSON.stringify(rightAnswer) === JSON.stringify(answer);
      const rightAnswerStr = options.reduce((acc, cur, idx) => {
        if (cur.isRight) {
          acc += String.fromCharCode(65 + idx);
        }
        return acc;
      }, '');
      const answerStr = options.reduce((acc, cur, idx) => {
        if (answer.indexOf(cur.id) !== -1) {
          acc += String.fromCharCode(65 + idx);
        }
        return acc;
      }, '');
      this.prevQuestions.push({
        question: this.question,
        answer: {
          isRight,
          answerStr,
          rightAnswerStr,
        },
      });
      this.statistic.questionsCount += 1;

      this.answer = '';
      this.answerList = [];
      if (isRight) {
        await this.getQuestion();
      } else {
        this.showAnalysis = true;
        this.selectedAnswer = {
          isRight,
          answerStr,
          rightAnswerStr,
        };
        this.$nextTick(() => {
          window.MathJax.typeset();
        });
      }
    },
    async continuePractice() {
      this.showAnalysis = false;
      if (this.nextQuestions.length !== 0) {
        this.prevQuestions.push({
          question: this.question,
          answer: this.selectedAnswer,
        });
        this.question = this.nextQuestions.shift();
      } else {
        await this.getQuestion();
      }
    },
    prevQuestion() {
      if (this.prevQuestions.length === 0) {
        this.$message.info('当前为第一题！');
        return;
      }
      this.nextQuestions.unshift(this.question);
      const { question, answer } = this.prevQuestions.pop();
      this.question = question;
      this.selectedAnswer = answer;
      this.showAnalysis = true;
    },
    nextQuestion() {
      this.prevQuestions.push(this.question);
      this.question = this.nextQuestions.pop();
    },
    async testTimer() {
      this.statistic.totalTime = dayjs((new Date()).getTime() - this.statistic.startTime).minute();
      setInterval(() => {
        this.statistic.totalTime = dayjs((new Date()).getTime() - this.statistic.startTime).minute();
      }, 1000);
    },
    async visibilitychange() {
      const { id } = this.question;
      if (document.visibilityState === 'visible') {
        // xapi: 数据埋点
        await newXAPIData({
          verb: 'practice_visible',
          object: {
            info: `question_id: ${id}`,
          },
        });
      } else {
        // xapi: 数据埋点
        await newXAPIData({
          verb: 'practice_invisible',
          object: {
            info: `question_id: ${id}`,
          },
        });
      }
    },
  },
};
</script>

<style scoped>
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
