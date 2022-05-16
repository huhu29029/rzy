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
  </div>
</template>

<script>
import { getExamHistory } from '@/api/questions.js';
import { formatTimeToStr } from '@/utils/date.js';
import { newXAPIData } from '@/api/xapi';

export default {
  data() {
    return {
      tableData: [],
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 1,
      },
    };
  },
  async created() {
    await this.getExamHistory();
  },
  methods: {
    async resultDetail(row) {
      // xapi: 数据埋点
      await newXAPIData({
        verb: 'review',
        object: {
          info: `examId: ${row.id}`,
        },
      });
      this.$router.push({
        name: 'h-exam-detail',
        params: {
          id: row.id,
        },
      });
    },
    async getExamHistory(
      page = this.pagination.currentPage,
      size = this.pagination.pageSize,
    ) {
      const res = await getExamHistory(page, size);
      const { data } = res;
      const { examHistory, total } = data;
      this.pagination.total = total;
      const table = [];
      for (let i = 0; i < examHistory.length; i++) {
        const exam = examHistory[i];
        const data = {};
        data.id = exam.examId;
        const time = formatTimeToStr(exam.startTime);
        data.date = time;
        data.title = exam.examTitle;
        table.push(data);
      }
      this.tableData = table;
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
