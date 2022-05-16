<template>
  <div>
    <div class="title">考试模式</div>
    <div class="desc">本模式下无解析，用于用户自测</div>
    <el-divider></el-divider>
    <div class="modes" v-for="exam in exams" :key="exam.id">
      <el-card>
        <div slot="header">
          <span style="font-weight: bold">{{ exam.title }}</span>
          <el-button type="text" @click="practice(exam.id)" id="exam.id">
            开始练习
          </el-button>
        </div>
        <ul style="padding-left:14px">
          <li>基础知识</li>
          <li>程序语法</li>
        </ul>
      </el-card>
    </div>
  </div>
</template>

<script>
import { getExamsList } from '@/api/exam';

export default {
  data() {
    return {
      exams: [],
    };
  },
  async created() {
    const res = await getExamsList();
    this.exams = res.data.exams;
  },
  methods: {
    practice(id) {
      this.$router.push({
        name: 'exam-paper',
        query: {
          examId: id,
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
.modes {
  margin-top: 20px;
  margin-left: 20px;
  display: inline-block;
}
</style>
