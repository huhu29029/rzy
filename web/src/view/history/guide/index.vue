<template>
  <div>
    <div class="shadow">
      <el-row :gutter="20">
        <el-col
          :span="4"
          v-for="(card, key) in toolCards"
          :key="key"
          @click.native="toTarget(card.name)"
          :xs="8"
        >
          <el-card shadow="hover" class="grid-content">
            <i :class="card.icon" :style="{ color: card.color }"></i>
            <p>{{ card.label }}</p>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      toolCards: [
        {
          label: '练习历史',
          icon: 'el-icon el-icon-time',
          name: 'h-practice',
          color: '#ff9c6e',
        },
        {
          label: '考试历史',
          icon: 'el-icon el-icon-time',
          name: 'h-exam',
          color: '#69c0ff',
        },
      ],
    };
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name });
    },
  },
};
</script>
