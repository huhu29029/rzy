<template>
  <div>
      <el-form
        :model="examForm"
        ref="examForm"
      >
        <el-form-item label="试卷标题" required>
          <el-input
            autocomplete="off"
            placeholder="试卷标题"
            type="textarea"
            :autosize="{ minRows: 2, maxRows: 4 }"
            v-model="examForm.title"
          ></el-input>
        </el-form-item>
      </el-form>
      <div>
        <p>* 题目选择</p>
        <el-table :data="selectedQuestions" stripe style="width: 100%">
          <el-table-column label="题目ID" width="100" align="center">
            <template slot-scope="scope">
              <span style="margin-left: 10px">{{ scope.row }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <div>
                <el-button
                  size="small"
                  icon="el-icon-delete"
                  @click="delQuestion(scope.$index)"
                  >移除</el-button
                >
                <el-button
                  size="small"
                  icon="el-icon-view"
                  @click="viewQuestion(scope.row)"
                  >预览</el-button
                >
              </div>
            </template>
          </el-table-column>
        </el-table>
        <el-button
          size="mini"
          type="primary"
          icon="el-icon-edit"
          @click="addQuestion"
          style="margin-top: 5px"
          >添加题目</el-button
        >
      </div>
      <div style="margin-top: 30px">
        <el-button
          size="mini"
          type="success"
          icon="el-icon-view"
          @click="previewExam"
          style="margin-top: 5px"
          >预览试题</el-button
        >
        <el-button type="primary" @click="submit">确 定</el-button>
      </div>
      <el-dialog
        title="试题选择"
        fullscreen
        :visible.sync="selectQuestionDialogVisible"
        :before-close="closeSelectQuestionDialog"
        >
        <selectQuestion ref="selectQuestion" :defaultQuestions="selectedQuestions" />
      </el-dialog>
      <el-dialog
        title="试题预览"
        :visible.sync="previewQuestionVisible"
      >
        <question :id="questionId" />
      </el-dialog>
      <el-dialog
        title="试卷预览"
        :visible.sync="previewExamVisible"
      >
        <previewExam :defaultQuestionIds="selectedQuestions" />
      </el-dialog>
  </div>
</template>

<script>
import { addExam } from '@/api/exam.js';

import question from '@/components/question';
import selectQuestion from '../components/selectQuestion.vue';
import previewExam from '../components/previewExam.vue';

export default {
  data() {
    return {
      examForm: {
        title: '',
      },
      selectedQuestions: [], // 试题选择
      selectQuestionDialogVisible: false,
      previewQuestionVisible: false, // 单个试题预览
      questionId: 0,
      previewExamVisible: false, // 试卷预览
    };
  },
  components: {
    question,
    selectQuestion,
    previewExam,
  },
  methods: {
    addQuestion() {
      this.selectQuestionDialogVisible = true;
    },
    closeSelectQuestionDialog() {
      this.selectedQuestions = this.$refs.selectQuestion.getData();
      this.selectQuestionDialogVisible = false;
    },
    viewQuestion(id) {
      this.questionId = id;
      this.previewQuestionVisible = true;
    },
    delQuestion(idx) {
      this.selectedQuestions.splice(idx, 1);
    },
    previewExam() {
      this.previewExamVisible = true;
    },
    async submit() {
      if (this.examForm.title === '') {
        this.$message.error('标题不能为空！');
        return;
      }
      if (this.selectedQuestions.length === 0) {
        this.$message.error('试题不能为空！');
        return;
      }
      try {
        const res = await addExam({
          title: this.examForm.title,
          questionIds: this.selectedQuestions,
        });
        if (res.code !== 0) {
          this.$message.error('创建试卷失败！');
          return;
        }
        this.$message.success('创建试卷成功！,3s后跳转至试卷列表页...');
        setTimeout(() => {
          this.$router.push({
            name: 'm-exams',
          });
        }, 3000);
      } catch (err) {
        this.$message.error('创建试卷失败');
      }
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
</style>
