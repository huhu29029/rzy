<template>
  <div>
    <el-button @click="addQuestion" type="primary">添加题目</el-button>
    <el-table :data="tableData" style="width: 100%">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form
            v-for="option in props.row.options"
            :key="option.id"
            label-position="left"
            inline
            class="table-expand"
          >
            <el-form-item label="选项">
              <span>{{ option.title }}</span>
            </el-form-item>
            <el-form-item label="正误">
              <span>{{ option.isRight }}</span>
            </el-form-item>
          </el-form>
          <el-form
            v-for="tag in props.row.tags"
            :key="tag.id"
            label-position="left"
            inline
            class="table-expand"
          >
            <el-form-item label="标签">
              <span>{{ [tag.layer1, tag.layout, tag.layer3, tag.layer4].join('/') }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label="题目 ID" prop="id"> </el-table-column>
      <el-table-column label="题干" prop="title"> </el-table-column>
      <el-table-column label="解析" prop="analysis"> </el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button
            @click="editQuestion(scope.row.id)"
            size="small"
            type="primary"
            icon="el-icon-edit"
            >编辑</el-button
          >
          <el-button
            size="small"
            type="danger"
            icon="el-icon-delete"
            @click="delQuestion(scope.row.id)"
            >删除</el-button
          >
          <el-button
            size="small"
            type="info"
            icon="el-icon-view"
            @click="viewQuestion(scope.row.id)"
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
    <!-- <el-dialog
      title="确定删除？"
      :visible.sync="delDialogVisible"
      width="20%">
      <span slot="footer" class="dialog-footer">
        <el-button @click="delDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="delDialogVisible = false">确 定</el-button>
      </span>
    </el-dialog> -->

    <el-dialog
      :before-close="handleClose"
      :title="dialogTitle"
      :visible.sync="dialogFormVisible"
    >
      <el-form
        :model="form"
        label-position="top"
        label-width="85px"
        ref="questionForm"
      >
        <el-form-item label="章节" prop="analysis">
          <el-select
            v-model="form.chapterId"
            filterable
            placeholder="请选择"
            @change="getChapterValue"
            clearable
            size="medium"
          >
            <el-option
              v-for="item in optionsChapterId"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="知识点类型" prop="analysis">
          <el-select
            v-model="form.tags[0]"
            filterable
            placeholder="请选择"
            @change="getKnowledgeValue"
            clearable
            size="medium"
            style="width: 590px"
          >
            <el-option
              v-for="item in optionsKnowledge"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="题干(markdown格式)" prop="title">
          <el-input
            autocomplete="off"
            placeholder="题干"
            type="textarea"
            :autosize="{ minRows: 2 }"
            v-model="form.title"
          ></el-input>
        </el-form-item>
        <el-form-item label="解析" prop="analysis">
          <el-input
            autocomplete="off"
            placeholder="解析"
            type="textarea"
            :autosize="{ minRows: 2, maxRows: 4 }"
            v-model="form.analysis"
          ></el-input>
        </el-form-item>
        <el-form-item label="题目类型" prop="analysis">
          <el-select v-model="form.type" placeholder="请选择">
            <el-option :value="1" label="单选"></el-option>
            <el-option :value="2" label="判断"></el-option>
            <el-option :value="3" label="多选"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div>
        <el-button
          size="small"
          type="primary"
          icon="el-icon-edit"
          @click="addOption(form)"
          >添加选项</el-button
        >
        <el-table :data="form.options" stripe style="width: 100%">
          <el-table-column prop="isRight" label="参数类型" width="180">
            <template slot-scope="scope">
              <el-select v-model="scope.row.isRight" placeholder="请选择">
                <el-option :value="true" label="正确"></el-option>
                <el-option :value="false" label="错误"></el-option>
              </el-select>
            </template>
          </el-table-column>
          <el-table-column prop="title" label="选项" width="400">
            <template slot-scope="scope">
              <div>
                <el-input
                  type="textarea"
                  :autosize="{ minRows: 2, maxRows: 4 }"
                  v-model="scope.row.title"
                ></el-input>
              </div>
            </template>
          </el-table-column>
          <el-table-column>
            <template slot-scope="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="el-icon-delete"
                  @click="delOption(form.options, scope.$index)"
                  >删除</el-button
                >
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
    <el-dialog
      :before-close="handleClose"
      :title="dialogViewTitle"
      :visible.sync="dialogViewVisible"
    >
      <question :id="questionId" />
    </el-dialog>
  </div>
</template>

<script>
import {
  getQuestions,
  getQuestion,
  addQuestion,
  editQuestion,
  delQuestion,
  getTags,
  getChapterId,
} from '@/api/questions.js';

import question from '@/components/question';

export default {
  data() {
    return {
      question: '',
      questionId: 0,
      MULTI: 3, // 题型：单选1 判断2 多选3
      answerList: [],
      answer: '',
      tableData: [], // 展示
      dialogTitle: '操作', // 增删改
      dialogViewTitle: '题目预览', // 增删改
      dialogFormVisible: false,
      delDialogVisible: false,
      dialogViewVisible: false,
      form: {
        title: '',
        analysis: '',
        type: 1,
        tags: [],
        options: [],
        chapterId: null,
      },
      isEdit: false, // 新建 or 编辑
      pagination: {
        currentPage: 1,
        pageSize: 5,
        total: 1,
      },
      optionsKnowledge: [],
      optionsChapterId: [],
      optionValue: '',
      tags: 0,
    };
  },
  components: {
    question,
  },
  async created() {
    await this.getQuestions();
    await this.optionsKnowledgeDeal();
    await this.optionsChapterIdDeal();
  },
  methods: {
    async optionsKnowledgeDeal() {
      const res = await getTags();
      const { data } = res;
      const { tags } = data;
      for (let i = 0; i < tags.length; i++) {
        const data = {};
        const label = `${tags[i].layer1
        } / ${
          tags[i].layer2
        } / ${
          tags[i].layer3
        } / ${
          tags[i].layer4}`;
        data.value = tags[i].id;
        data.label = label;
        this.optionsKnowledge.push(data);
      }
    },
    async optionsChapterIdDeal() {
      const res = await getChapterId();
      const { data } = res;
      const { chapters } = data;
      for (let i = 0; i < chapters.length; i++) {
        const data = {};
        data.value = chapters[i].id;
        data.label = chapters[i].title;
        this.optionsChapterId.push(data);
      }
    },
    getKnowledgeValue(val) {
      let obj = {};
      obj = this.optionsKnowledge.find((item) => item.value === val);
      this.form.tags[0] = obj.value;
    },
    getChapterValue(val) {
      let obj = {};
      obj = this.optionsChapterId.find((item) => item.value === val);
      this.form.chapterId = obj.value;
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
    async addQuestion() {
      this.form = {
        title: '',
        analysis: '',
        type: 1,
        tags: [],
        options: [],
        chapterId: null,
      };
      this.dialogFormVisible = true;
      this.isEdit = false;
    },
    async editQuestion(id) {
      const res = await getQuestion(id);
      const { data } = res;
      const { question } = data;
      this.form = question;
      if (this.form.tags.length !== 0) {
        const tagsId = this.form.tags[0].id;
        const tags = [tagsId];
        this.form.tags = tags;
      } else {
        this.form.tags = [];
      }
      if (this.form.chapterId === 0) {
        this.form.chapterId = null;
      }
      this.dialogFormVisible = true;
      this.isEdit = true;
    },
    async delQuestion(id) {
      this.delDialogVisible = true;
      try {
        await this.$confirm('确定删除?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        });
        const res = await delQuestion(id);
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '操作成功',
          });
          await this.getQuestions();
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
    async viewQuestion(id) {
      this.questionId = id;
      this.dialogViewVisible = true;
    },
    handleClose(done) {
      done();
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.tags = 0;
    },
    async enterDialog() {
      let res;
      if (this.isEdit) {
        res = await editQuestion(this.form.id, this.form);
      } else {
        res = await addQuestion(this.form);
      }

      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '操作成功',
        });
        await this.getQuestions();
      } else {
        this.$message({
          type: 'error',
          message: '操作失败',
        });
      }

      this.dialogFormVisible = false;
      this.tags = 0;
      this.optionValue = '';
    },
    addOption(form) {
      form.options.push({
        title: '',
        isRight: true,
      });
    },
    delOption(options, index) {
      options.splice(index, 1);
    },
    async changePage(page) {
      this.pagination.page = page;
      await this.getQuestions();
    },
    async handleSizeChange(val) {
      this.pagination.pageSize = val;
      await this.getQuestions();
    },
    async handleCurrentChange(val) {
      this.pagination.currentPage = val;
      await this.getQuestions();
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
.table-expand {
  font-size: 0;
}
.table-expand label {
  width: 90px;
  color: #99a9bf;
}
.table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>
