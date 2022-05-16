<template>
  <div>
    <el-card class="box-card">
    <div slot="header" class="clearfix">
        <span>资料下载</span>
    </div>
    <div class="text item">
    <div>
         <p>
            此处为同学们提供了网盘链接，网盘内有大量课本、ppt、题解等资料供同学们查看下载。
            文件按照专业-学科-资料分类呈现，同学们可按需选择对应文件夹查找资料并下载，下载无需注册登录。
            网盘内容将根据同学们上传的资料时常更新。
            点击“查看”并输入提取码进入网盘。
        </p>
        <el-button
          type="text"
          size="small"
          style="float: right"
          @click="goToUrl"
          >查看（提取码：kea6tw）</el-button
        >
    </div>
    </div>
    </el-card>
     <el-card class="box-card">
    <div slot="header" class="clearfix">
        <span>资料上传</span>
    </div>
    <div class="text item">
    <div>
         <p>
            如果您有好的学习资料愿意分享，欢迎将资料发送至邮箱：</p><p>buptsharing@163.com</p><p>
            我们将在审核后上传在网盘中，供所有平台用户浏览下载。
            感谢您对本平台的支持！
        </p>
    </div>
    </div>
    </el-card>
    <el-card class="box-card">
    <div slot="header" class="clearfix">
        <span>意见反馈</span>
    </div>
    <div class="text item">
    <div>
         <p>
            如果您没能在网盘中找到您需要的资料，请将您的需求告诉我们。
            如果您发现网盘资料中的错误和纰漏，也欢迎您告诉我们。
            联系方式：
        </p>
        <p>
            微信：Z2937443168
        </p>
    </div>
    </div>
    </el-card>
  </div>
</template>

<script>
export default {
  methods: {
    goToUrl() {
      window.open('https://cowtransfer.com/s/8e21bc1c31f64a');
    },
  },
};
</script>

<style scoped>
  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }

  .box-card {
    width: 100%;
    box-shadow: none;
    border-color: white;
  }
  span {
      font-weight: bold;
  }
  p {
    text-indent: 2em;
    font-weight: normal;
    font-size: 16px;
    line-height: 22px;
    padding-bottom: 4px;
  }
</style>
