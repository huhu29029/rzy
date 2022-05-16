<template>
  <div>
    <div class="title">练习模式</div>
    <div class="desc">本模式下可查看提示、详解</div>
    <el-divider></el-divider>
    <div class="mode">
      <el-card>
        <div slot="header">
          <span style="font-weight: bold">算法推荐模式</span>
          <el-button type="text" @click="practiceWithKT"> 开始练习 </el-button>
        </div>
        <ul style="padding-left:14px">
          <li>根据练习历史分析知识点掌握情况</li>
          <li>针对薄弱点推荐题目</li>
          <li>提高学习效率</li>
        </ul>
      </el-card>
    </div>
    <div class="mode">
      <el-card>
        <div slot="header">
          <span style="font-weight: bold">随机练习模式</span>
          <el-button type="text" @click="practice"> 开始练习 </el-button>
        </div>
        <ul style="padding-left:14px">
          <li>随机选题</li>
          <li>覆盖全部题目</li>
          <li>无限练习</li>
        </ul>
      </el-card>
    </div>
    <div class="mode">
      <el-card>
        <div slot="header">
          <span style="font-weight: bold">章节练习模式</span>
          <el-button type="text" @click="chooseChapter"> 开始练习 </el-button>
        </div>
        <ul style="padding-left:14px">
          <li>章节选择</li>
          <li>随机选题</li>
          <li>无限练习</li>
        </ul>
      </el-card>
    </div>
    <div class="mode">
      <el-card>
        <div slot="header">
          <span style="font-weight: bold">错题再练模式</span>
          <el-button type="text" @click="practiceWrongQuestions"> 进入练习 </el-button>
        </div>
        <ul style="padding-left:14px">
          <li>强化记忆知识点</li>
          <li>针对性训练</li>
          <li>快速练习</li>
        </ul>
      </el-card>
    </div>
    <el-dialog
      :before-close="handleClose"
      title="请选择对应章节"
      :visible.sync="dialogFormVisible"
    >
      <el-form
        label-position="top"
        label-width="85px"
      >
        <p style="margin-bottom: 15px; font-weight: bold; font-size: 12px">
          *提示：不选择章节，直接`确定`则默认所有章节！
        </p>
        <el-checkbox-group v-model="selectedChapters">
          <el-checkbox
            v-for="chapter in chapters"
            :label="chapter.id"
            :key="chapter.id"
            class="question-option"
          >
            <p>{{ chapter.title }}</p></el-checkbox
          >
        </el-checkbox-group>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button @click="chapterPractice" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getChapterId,
} from '@/api/questions.js';

export default {
  data() {
    return {
      chapters: [],
      selectedChapters: [],
      dialogFormVisible: false,
    };
  },
  methods: {
    async chooseChapter() {
      const res = await getChapterId();
      this.chapters = res.data.chapters;
      this.dialogFormVisible = true;
    },
    handleClose() {
      this.dialogFormVisible = false;
      this.selectedChapters = [];
    },
    async chapterPractice() {
      let { selectedChapters } = this;
      if (selectedChapters.length === 0) {
        selectedChapters = this.chapters.reduce((acc, cur) => {
          acc.push(cur.id);
          return acc;
        }, []);
      }
      this.$router.push({
        name: 'questions-practice',
        query: {
          chapters: JSON.stringify(selectedChapters),
        },
      });
    },
    async practice() {
      this.$router.push({
        name: 'questions-practice',
      });
    },
    async practiceWithKT() {
      this.$router.push({
        name: 'questions-practice',
        query: {
          kt: true,
        },
      });
    },
    async practiceWrongQuestions() {
      this.$router.push({
        name: 'questions-practice',
        query: {
          isWrong: true,
        },
      });
    },
  },
};
</script>

<style scoped>
ul, ol, li {
  list-style-type: disc;
}
li {
  margin-bottom: 4px;
}
.title {
  font-size: 20px;
  font-weight: bold;
  padding: 10px;
}
.desc {
  padding: 8px;
}
.mode {
  margin-top: 20px;
  margin-left: 20px;
  float: left;
}
.question-option {
  display: flex;
  padding: 5px;
}
</style>
