<template>
  <div>
    <el-table :data="tableData" style="width: 100%">
      <el-table-column label="id" width="150" align="center">
        <template slot-scope="scope">
          <!-- <i class="el-icon-time"></i> -->
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户名" width="250" align="center">
        <template slot-scope="scope">
          <!-- <i class="el-icon-document"></i> -->
          <span style="margin-left: 10px">{{ scope.row.userName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="昵称" width="250" align="center">
        <template slot-scope="scope">
          <!-- <i class="el-icon-document"></i> -->
          <span style="margin-left: 10px">{{ scope.row.nickName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="角色" width="250" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.authorityId === 1" style="margin-left: 10px">{{
            role[0]
          }}</span>
          <span
            v-else-if="scope.row.authorityId === 2"
            style="margin-left: 10px"
            >{{ role[1] }}</span
          >
          <span v-else style="margin-left: 10px">{{ role[2] }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="userInfo(scope.row)">详情</el-button>
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
import { getUserList } from '@/api/user.js';

export default {
  data() {
    return {
      tableData: [],
      pagination: {
        currentPage: 1,
        pageSize: 10,
        total: 1,
      },
      role: ['管理员', '用户', '布道师'],
    };
  },
  async created() {
    await this.getUserList();
  },
  methods: {
    async getUserList(
      page = this.pagination.currentPage,
      size = this.pagination.pageSize,
    ) {
      const userRes = await getUserList(page, size);
      const { data } = userRes;
      const { users, total } = data;
      this.pagination.total = total;
      this.tableData = users;
    },
    userInfo(row) {
      const { id } = row;
      sessionStorage.setItem('userInfoId', id);
      sessionStorage.setItem('userInfoPage', this.pagination.page);
      sessionStorage.setItem('userInfoSize', this.pagination.size);
      this.$router.push({
        name: 'm-userInfo',
      });
    },
    async handleSizeChange(val) {
      this.pagination.pageSize = val;
      await this.getUserList();
    },
    async handleCurrentChange(val) {
      this.pagination.currentPage = val;
      await this.getUserList();
    },
  },
};
</script>
