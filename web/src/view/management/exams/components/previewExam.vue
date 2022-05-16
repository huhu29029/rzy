<template>
  <div>
    <el-empty v-if="questionIds.length === 0" description="无试题"></el-empty>
    <div v-else>
     <div>
        <el-button type="info">{{questionIdx+1}}/{{questionIds.length}}</el-button>
        <span v-if="question.type === 1">[选择题]</span>
        <span v-else-if="question.type === 2">[判断题]</span>
        <span v-else>[多选题]</span>
      </div>
      <question :id="questionId" />
      <el-row>
        <el-button type="primary" @click="prevQuestion">上一题</el-button>
        <el-button type="primary" @click="nextQuestion">下一题</el-button>
      </el-row>
     <div class="card">
        <el-row>
          <li
            v-for="(_, idx) in questionIds"
            :id="idx"
            :key="idx"
            @click="toQuestion(idx)"
          >
            {{ idx + 1 }}
          </li>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import question from '@/components/question';

export default {
  name: 'previewExam',
  data() {
    return {
      question: {},
      questionId: 0,
      questionIdx: 0,
      questionIds: [],
    };
  },
  components: {
    question,
  },
  props: ['defaultQuestionIds'],
  watch: {
    defaultQuestionIds(newV) {
      this.questionIds = newV;
      this.questionIdx = 0;
      this.questionId = this.questionIds[0];
    },
  },
  created() {
    this.questionIds = this.defaultQuestionIds;
    this.questionIdx = 0;
    this.questionId = this.questionIds[0];
  },
  methods: {
    prevQuestion() {
      if (this.questionIdx === 0) {
        this.$message('已经是第一题！');
        return;
      }
      this.questionIdx--;
      this.questionId = this.questionIds[this.questionIdx];
    },
    nextQuestion() {
      if (this.questionIdx === this.questionIds.length - 1) {
        this.$message('已经是最后一题！');
        return;
      }
      this.questionIdx++;
      this.questionId = this.questionIds[this.questionIdx];
    },
    toQuestion(idx) {
      this.questionIdx = idx;
      this.questionId = this.questionIds[idx];
    },
  },
};
</script>

<style>
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
</style>
