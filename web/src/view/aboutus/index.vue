<template>
<!-- 只user，preacher和admin都没写上 -->
  <div>
    <el-card class="box-card">
    <div slot="header" class="clearfix">
        <span>项目介绍与项目特色</span>
    </div>
    <div class="text item">
    <div>
         <p>
            题库是巩固学习、自我检验的好方法。线上题库更有着题量大、更新快、操作便捷、高效的特点。
            但目前市面上存在更多的是在线课程教学平台，缺少题库，已有的题库系统也仅仅实现了题目的展示，无法针对不同学生自身情况给予个性化的指引。
        </p>
        <p>
            大学不同于高中，更加重视自学能力的培养。
            但是自学能力的建立不在一朝一夕，在这个过渡阶段，导引和工具的作用是重要的，甚至决定了一个学生相当长一段时间内的学习状态。
        </p>
        <p>
            本项目则旨在解决以上问题，开发出一个通用于各学科的题库系统，为学生提供一个练习和自我检验的系统，
            并通过集成知识追踪/推荐算法为学生提供个性化学习方案，从而提升学习效率。
            同时也为老师提供组卷、试卷分析功能，使得老师精细化分析学生学习情况，从而动态调整教学计划，实现教-学-教的闭环。
        </p>
    </div>
    </div>
    </el-card>
     <el-card class="box-card">
    <div slot="header" class="clearfix">
        <span>宣传口号</span>
    </div>
    <div class="text item">
    <div>
         <p>
            为大学新生提供学习导引是我们的初衷，为各高校建立资源共享是我们的原则，所谓布道者：给予帮助，创造可能！
            我们愿意用自己的努力，创建一个导论学习的新平台。
        </p>
    </div>
    </div>
    </el-card>
    <el-card class="box-card">
    <div slot="header" class="clearfix">
        <span>团队介绍与企业相关</span>
    </div>
    <div class="text item">
    <div>
         <p>
            北京邮电大学“面向新生自学编程的在线题库设计与实现”大创项目组
        </p>
        <p>
            北京邮电大学（Beijing University of Posts and Telecommunications），简称北邮，有“信息黄埔”的美誉
        </p>
    </div>
    </div>
    </el-card>
  </div>
</template>

<script>
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
    font-family: Helvetica, Tahoma, Arial, STXihei, "华文细黑", "Microsoft YaHei", "微软雅黑", SimSun, "宋体", Heiti, "黑体", sans-serif;
  }
</style>
