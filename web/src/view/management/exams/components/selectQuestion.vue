<template>
  <div>
    {{ Array.from(questionSet.values())  }}
    <el-table
      ref="questions"
      :data="tableData"
      style="width: 100%">
      <el-table-column label="题目 ID" prop="id"> </el-table-column>
      <el-table-column label="题干" prop="title"> </el-table-column>
      <el-table-column label="查看">
        <template slot-scope="scope">
          <el-button
            size="small"
            icon="el-icon-view"
            @click="viewQuestion(scope.row.id)"
            >查看</el-button
          >
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="选中">
        <template slot-scope="scope">
          <questionSwitch
            @change="selectQuestion"
            :isChoosed="questionSet.has(scope.row.id)"
            :id="scope.row.id"
          />
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
      append-to-body
    >
      <question :id="questionId" />
    </el-dialog>
  </div>
</template>

<script>
import {
  getQuestions,
} from '@/api/questions.js';

import question from '@/components/question';
import questionSwitch from './questionSwitch.vue';

export default {
  data() {
    return {
      question: '',
      questionId: 0,
      questionSet: new Set(), // 存储id
      tableData: [], // 展示
      dialogViewTitle: '题目预览', // 增删改
      dialogViewVisible: false,
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 1,
      },
    };
  },
  props: {
    defaultQuestions: {},
  },
  watch: {
    defaultQuestions() {
      this.questionSet.clear();
      for (let i = 0; i < this.defaultQuestions.length; i++) {
        this.questionSet.add(this.defaultQuestions[i]);
      }
    },
  },
  components: {
    question,
    questionSwitch,
  },
  async created() {
    this.questionSet.clear();
    for (let i = 0; i < this.defaultQuestions.length; i++) {
      this.questionSet.add(this.defaultQuestions[i]);
    }
    await this.getQuestions();
  },
  methods: {
    getData() {
      return Array.from(this.questionSet);
    },
    async getQuestions(
      page = this.pagination.currentPage,
      size = this.pagination.pageSize,
    ) {
      const res = await getQuestions(page, size);
      const { data } = res;
      const { questions, total } = data;
      this.tableData = questions;
      this.pagination.total = total;
    },
    handleClose(done) {
      done();
    },
    async handleSizeChange(val) {
      this.pagination.pageSize = val;
      await this.getQuestions();
      this.reToggleSelection();
    },
    async handleCurrentChange(val) {
      this.pagination.currentPage = val;
      await this.getQuestions();
      this.reToggleSelection();
    },
    selectQuestion(data) {
      const { id, isChoosed } = data;
      if (isChoosed) {
        this.questionSet.add(id);
      } else {
        this.questionSet.delete(id);
      }
    },
    viewQuestion(id) {
      this.dialogViewVisible = true;
      this.questionId = id;
    },
  },
};
</script>

<style>
.title {
  font-size: 18px;
  padding: 20px;
}
.question-option {
  display: flex;
  padding: 5px;
}
.analysis {
  padding: 10px;
  line-height: 20px;
}
</style>
