<template>
  <div>
    <div class="title" v-html="question.title"></div>
    <div v-if="question.type !== MULTI">
      <el-radio-group v-model="question.answer">
        <el-radio
          v-for="option in question.options"
          :label="option.id"
          :key="option.id"
          class="question-option"
          :value="option.isRight ? option.id : '' "
        >
          <div v-html="option.title"></div>
        </el-radio>
      </el-radio-group>
    </div>
    <div v-else>
      <el-checkbox-group v-model="question.answer">
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
</template>

<script>
import { getQuestion } from '@/api/questions.js';

export default {
  name: 'question',
  data() {
    return {
      question: {},
      MULTI: 3, // 题型：单选1 判断2 多选3
    };
  },
  props: {
    id: {
      required: true,
    },
  },
  async created() {
    this.question = await this.getQuestion(this.id);
  },
  watch: {
    async id(newId) {
      this.question = await this.getQuestion(newId);
    },
  },
  methods: {
    async getQuestion(id) {
      const res = await getQuestion(id);
      const { data } = res;
      const { question } = data;
      question.title = this.$md.render(question.title);
      question.options.map((option) => {
        option.title = this.$md.render(option.title);
        return option;
      });

      if (question.type === this.MULTI) {
        question.answer = question.options.reduce((acc, cur) => {
          if (cur.isRight) {
            acc.push(cur.id);
          }
          return acc;
        }, []);
      } else {
        question.answer = question.options.reduce((acc, cur) => {
          if (cur.isRight) {
            acc = cur.id;
          }
          return acc;
        }, 0);
      }
      return question;
    },
  },
};
</script>

<style scoped>
.title {
  font-size: 18px;
  padding: 10px;
  line-height: 25px;
}
.question-option {
  padding: 5px;
  display: flex;
}
</style>
