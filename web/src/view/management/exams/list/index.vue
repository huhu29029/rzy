<template>
  <div>
    <el-button type="primary" @click="addExam">添加试卷</el-button>
    <el-table :data="exams" style="width: 100%">
      <el-table-column label="试卷ID" width="200" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="试卷标题" align="left">
        <template slot-scope="scope">
          <i class="el-icon-document"></i>
          <span>{{ scope.row.title }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="350" align="center">
        <template slot-scope="scope">
          <el-button
            @click="editExam(scope.row.id)"
            size="small"
            type="primary"
            icon="el-icon-edit"
            >编辑</el-button
          >
          <el-button
            size="small"
            type="danger"
            icon="el-icon-delete"
            @click="delExam(scope.row.id)"
            >删除</el-button
          >
          <el-button
            size="small"
            type="success"
            icon="el-icon-view"
            @click="previewExam(scope.row.id)"
            >预览</el-button
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
      :before-close="closeQuestionDialog"
      title="试卷详情"
      :visible.sync="dialogQuestionVisible"
    >
    <div>
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
          <el-button
            type="primary"
            style="margin-left: 100px"
            id="return"
            @click="closeQuestionDialog"
            >返回</el-button
          >
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
    </el-dialog>
  </div>
</template>

<script>
import {
  getExamsListByAdmin,
  getExam,
  delExam,
} from '@/api/exam.js';

import question from '@/components/question';

export default {
  data() {
    return {
      exams: [],
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 1,
      },
      dialogQuestionVisible: false,
      question: {},
      questionIdx: 0,
      questionId: 0,
      questionIds: [],
      MULTI: 3, // 题型：单选1 判断2 多选3
    };
  },
  components: {
    question,
  },
  async created() {
    await this.getExamsList();
  },
  methods: {
    // 编辑区
    async getExamsList(
      page = this.pagination.currentPage,
      size = this.pagination.pageSize,
    ) {
      const res = await getExamsListByAdmin(page, size);
      const { data } = res;
      const { exams, total } = data;
      this.pagination.total = total;
      this.exams = exams;
    },
    async delExam(id) {
      try {
        await this.$confirm('确定删除?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        });
        const res = await delExam(id);
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '操作成功',
          });
          await this.getExamsList();
        } else {
          this.$message({
            type: 'error',
            message: '操作失败',
          });
        }
      } catch (err) {
        this.$message({
          type: 'info',
          message: '已取消',
        });
      }
    },
    addExam() {
      this.$router.push({
        name: 'exam-new',
      });
    },
    editExam(id) {
      this.$router.push({
        name: 'exam-edit',
        query: {
          id,
        },
      });
    },
    // 预览区
    async previewExam(id) {
      const res = await getExam(id);
      const { data } = res;
      const { exam } = data;
      const { questions } = exam;

      this.questionIds = questions.reduce((acc, cur) => {
        acc.push(cur.id);
        return acc;
      }, []);
      this.questionIdx = 0;
      this.questionId = this.questionIds[0];

      this.dialogQuestionVisible = true;
    },
    async closeQuestionDialog() {
      this.dialogQuestionVisible = false;
      this.questionIds = [];
    },
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
    // 试卷列表分页
    async handleSizeChange(val) {
      this.pagination.pageSize = val;
      await this.getExamsList();
    },
    async handleCurrentChange(val) {
      this.pagination.currentPage = val;
      await this.getExamsList();
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
