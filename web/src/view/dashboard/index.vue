<template>
  <div class="big">
    <el-row>
      <div class="card">
        <el-col :xs="24" :lg="16" :md="16">
          <div class="car-left">
            <el-row>
              <div>
                <el-col :xs="4" :md="3" :lg="3">
                  <span class="card-img">
                    <img src="~@/assets/online_exam.png" alt="" />
                  </span>
                </el-col>
                <el-col :xs="20" :lg="12" :md="12">
                  <div class="text">
                    <h4>练练看（beta版）</h4>
                    <p style="margin-top: 10px">JavaScript/微信小程序</p>
                  </div>
                </el-col>
              </div>
            </el-row>
          </div>
        </el-col>
      </div>
    </el-row>
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
    <el-row>
      <div class="card" style="width: 400px" title="涅槃之路平台">
        <el-button
          type="text"
          size="small"
          style="float: right"
          @click="goToUrl"
          >访问</el-button>
        <h6>学习平台推荐：</h6>
        <el-col :xs="24" :lg="16" :md="16">
          <div class="car-left">
            <el-row>
              <div>
                <el-col :xs="4" :md="3" :lg="3">
                  <span class="card-img">
                    <img src="~@/assets/spark.png" alt="" />
                  </span>
                </el-col>
                <el-col :xs="20" :lg="12" :md="12" style="margin-left: 60px">
                  <div class="text" style="width: 300px">
                    <h4>涅槃之路平台</h4>
                    <p style="margin-top: 10px">
                      提供JavaScript、微信小程序基础知识学习资料
                    </p>
                  </div>
                </el-col>
              </div>
            </el-row>
          </div>
        </el-col>
      </div>
    </el-row>
    <el-row>
        <p class="tips">
          遇到问题/提出建议可联系：
        <br>
        微信：Z2937443168
        </p>
    </el-row>
  </div>
</template>

<script>
import { newXAPIData } from '@/api/xapi';

export default {
  name: 'Dashboard',
  data() {
    return {
      toolCards: [
        {
          label: '练习',
          icon: 'el-icon el-icon-edit',
          name: 'questions',
          color: '#ff9c6e',
        },
        {
          label: '考试',
          icon: 'el-icon el-icon-document',
          name: 'exam',
          color: '#69c0ff',
        },
        // {
        //   label: '学习计划',
        //   icon: 'el-icon el-icon-sunrise',
        //   name: 'studyplan',
        //   color: '#EE6666',
        // },
        // {
        //   label: '错题记录',
        //   icon: 'el-icon el-icon-paperclip',
        //   name: 'myerror',
        //   color: '#66CD00',
        // },
        // {
        //   label: '资料共享',
        //   icon: 'el-icon el-icon-message',
        //   name: 'sharing',
        //   color: '#A3D1D1',
        // },
        {
          label: '个人信息',
          icon: 'el-icon el-icon-user',
          name: 'p-information',
          color: '#b37feb',
        },
        {
          label: '做题历史',
          icon: 'el-icon el-icon-time',
          name: 'p-history',
          color: '#ffd668',
        },
      ],
    };
  },
  methods: {
    async goToUrl() {
      const url = 'https://www.ourspark.org/';
      // xapi: 数据埋点
      await newXAPIData({
        verb: 'study',
        object: {
          info: url,
        },
      });
      window.open(url);
    },
    toTarget(name) {
      this.$router.push({ name });
    },
  },
};
</script>

<style lang="scss" scoped>
.big {
  margin: 70px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;
  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;
    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }
  .mid {
    width: 100%;
    height: 380px;
    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }
  .tips {
    text-align: center;
    margin-top: 80px;
  }
}
</style>
