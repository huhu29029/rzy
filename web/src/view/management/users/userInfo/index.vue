<template>
  <div>
    <el-row>
      <el-col :span="6">
        <div class="fl-left avatar-box">
          <div class="user-card">
            <div
              class="user-headpic-update"
              :style="{
                'background-image': `url(${
                  userInfo.headerImg &&
                  userInfo.headerImg.slice(0, 4) !== 'http'
                    ? path + userInfo.headerImg
                    : userInfo.headerImg
                })`,
                'background-repeat': 'no-repeat',
                'background-size': 'cover',
              }"
            ></div>
            <div class="user-personality">
              <p class="nickname">{{ userInfo.nickName }}</p>
            </div>
            <div class="user-information">
              <ul>
                <li>
                  <i class="el-icon-user"></i>{{ userInfo.nickName }}
                  <p
                    style="
                      margin-top: 5px;
                      margin-bottom: 5px;
                      font-weight: bold;
                    "
                  >
                    uuid:
                  </p>
                  <div>
                    <p>{{ userInfo.uuid }}</p>
                    <el-link
                      style="margin-top: 5px; float: right"
                      type="primary"
                      :underline="false"
                      class="copy-btn"
                      v-clipboard:copy="userInfo.uuid"
                    >
                      复制uuid
                    </el-link>
                  </div>
                  <div
                    style="
                      margin-top: 20px;
                      margin-bottom: 5px;
                      font-weight: bold;
                    "
                  >
                    <span> 角色： </span>
                    <span v-if="userInfo.authorityId === 1">{{ role[0] }}</span>
                    <span v-else-if="userInfo.authorityId === 2">{{
                      role[1]
                    }}</span>
                    <span v-else>{{ role[2] }}</span>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="18">
        <div class="user-addcount">
          <el-tabs v-model="activeName">
            <el-tab-pane label="做题历史" name="second">
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
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { getUserList } from '@/api/user';

const path = process.env.VUE_APP_BASE_API;
export default {
  name: 'Person',
  data() {
    return {
      path,
      toolCards: [
        {
          label: '练习历史',
          icon: 'el-icon el-icon-time',
          name: 'm-userPractice',
          color: '#ff9c6e',
        },
        {
          label: '考试历史',
          icon: 'el-icon el-icon-time',
          name: 'm-userExam',
          color: '#69c0ff',
        },
      ],
      userInfo: {
        uuid: '',
        userName: '',
        nickName: '',
        headerImg: '',
        authority: '',
        authorityId: '',
        id: '',
      },
      id: 0,
      page: 0,
      size: 0,
      activeName: 'second',
      showPassword: false,
      pwdModify: {},
      uuid: 0,
      authorityId: 0,
      role: ['管理员', '用户', '布道师'],
    };
  },
  async created() {
    await this.getParams();
    const usersRes = await getUserList(this.page, this.size);
    const { data } = usersRes;
    const { users } = data;
    for (let i = 0; i < users.length; i++) {
      if (users[i].id === this.id) {
        this.userInfo = users[i];
      }
    }
  },

  methods: {
    async getParams() {
      const paramsId = sessionStorage.getItem('userInfoId');
      const page = sessionStorage.getItem('userInfoPage');
      const size = sessionStorage.getItem('userInfoSize');
      this.id = Number(paramsId);
      this.page = Number(page);
      this.size = Number(size);
    },
    toTarget(name) {
      this.$router.push({ name });
    },
  },
};
</script>
<style lang="scss">
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.avatar-box {
  box-shadow: -2px 0 20px -16px;
  width: 80%;
  height: 100%;
  .user-card {
    min-height: calc(90vh - 200px);
    padding: 30px 20px;
    text-align: center;
    .el-avatar {
      border-radius: 50%;
    }
    .user-personality {
      padding: 24px 0;
      text-align: center;
      p {
        font-size: 16px;
      }
      .nickname {
        font-size: 26px;
      }
      .person-info {
        margin-top: 6px;
        font-size: 14px;
        color: #999;
      }
    }
    .user-information {
      width: 100%;
      height: 100%;
      text-align: left;
      ul {
        display: inline-block;
        height: 100%;
        li {
          i {
            margin-right: 8px;
          }
          padding: 20px 0;
          font-size: 16px;
          font-weight: 400;
          color: #606266;
        }
      }
    }
  }
}
.user-addcount {
  ul {
    li {
      .title {
        padding: 10px;
        font-size: 18px;
        color: #696969;
      }
      .desc {
        font-size: 16px;
        padding: 0 10px 20px 10px;
        color: #a9a9a9;
        a {
          color: rgb(64, 158, 255);
          float: right;
        }
      }
      border-bottom: 2px solid #f0f2f5;
    }
  }
}
.user-headpic-update {
  width: 120px;
  height: 120px;
  line-height: 120px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  border-radius: 20px;
  &:hover {
    color: #fff;
    background: linear-gradient(
        to bottom,
        rgba(255, 255, 255, 0.15) 0%,
        rgba(0, 0, 0, 0.15) 100%
      ),
      radial-gradient(
          at top center,
          rgba(255, 255, 255, 0.4) 0%,
          rgba(0, 0, 0, 0.4) 120%
        )
        #989898;
    background-blend-mode: multiply, multiply;
    .update {
      color: #fff;
    }
  }
  .update {
    height: 120px;
    width: 120px;
    text-align: center;
    color: transparent;
  }
}
</style>
