<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/online_exam.png" alt="" />
        </div>
        <div class="header">
          <a href="/">
            <span class="title">练练看（beta版）</span>
          </a>
        </div>
      </div>
      <div class="main" v-if="isLogin">
        <el-form
          :model="loginForm"
          :rules="rules"
          ref="loginForm"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="loginForm.password"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%"
              >登 录</el-button
            >
          </el-form-item>
        </el-form>
        <!-- <el-button type="text" style="color: white" @click="isLogin = !isLogin">注册</el-button> -->
        <el-button
          type="text"
          style="margin-left: 10px; color: white"
          @click="oauthLogin('ourspark')">
          火花统一登录
        </el-button>
      </div>
      <div class="main" v-else>
        <el-form
          :model="registerForm"
          :rules="registerRules"
          ref="registerForm"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="registerForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="registerForm.password"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
          </el-form-item>
          <!-- TODO: Vue 或 element UI 更新节点导致注册和登录表单混淆，无法校验再次输入密码的内容-->
          <!-- <el-form-item prop="confirm">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请再次输入密码"
              v-model="registerForm.confirm"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
          </el-form-item> -->
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%">注 册</el-button>
          </el-form-item>
        </el-form>
        <el-button type="text" style="color: white" @click="isLogin = !isLogin">登录</el-button>
      </div>

      <div class="footer">
        <div class="copyright">
           Copyright &copy; 2021 - {{ new Date().getFullYear() }}
           <a href="https://beian.miit.gov.cn" target="_blank" style="margin-left: 5px">京ICP备18044922号-3</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'Login',
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error('请输入正确的用户名（长度大于5且小于12）'));
      }
      return callback();
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error('请输入正确的密码'));
      }
      return callback();
    };
    const checkConfirm = (rule, value, callback) => {
      if (value !== 'this.registerForm.password') {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      isLogin: true,
      isMobile: false,
      lock: 'lock',
      loginForm: {
        username: '',
        password: '',
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }],
      },
      registerForm: {
        username: '',
        password: '',
        confirm: '',
      },
      registerRules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }],
        confirm: [{ validator: checkConfirm, trigger: 'blur' }],
      },
      logVerify: '',
      picPath: '',
    };
  },
  mounted() {
    const screenWidth = document.body.clientWidth;
    if (screenWidth < 1000) {
      this.isMobile = true;
    }
  },
  created() {
    this.alert();
  },
  methods: {
    ...mapActions('user', ['LoginIn', 'Register']),
    alert() {
      if (this.isMobile) {
        this.$message.info('提示：请使用PC端打开网页以获得更好的浏览体验！');
      }
    },
    async login() {
      await this.LoginIn(this.loginForm);
    },
    async register() {
      const { username, password } = this.registerForm;
      const data = {
        username,
        password,
      };
      await this.Register(data);
    },
    async submitForm() {
      if (this.isLogin) {
        this.$refs.loginForm.validate(async (v) => {
          if (v) {
            await this.login();
            return true;
          }
          this.$message({
            type: 'error',
            message: '请正确填写登录信息',
            showClose: true,
          });
          return false;
        });
      } else {
        this.$refs.registerForm.validate(async (v) => {
          if (v) {
            await this.register();
            return true;
          }
          this.$message({
            type: 'error',
            message: '请正确填写注册信息',
            showClose: true,
          });
          return false;
        });
      }
    },
    changeLock() {
      if (this.lock === 'lock') {
        this.lock = 'unlock';
      } else {
        this.lock = 'lock';
      }
    },
    oauthLogin(thridPlatform) {
      window.sessionStorage.setItem('online-exam:redirect', document.location.hash);

      const oauthOrigin = process.env.VUE_APP_OAUTH_OURSPARK_ORIGIN || 'https://oauth.staging.feel.ac.cn:8443';
      const oauthAppid = process.env.VUE_APP_OAUTH_OURSPARK_APPID || 'exam';

      const OURSPARK_URL = `${oauthOrigin}/oauth?redirect_uri=${document.location.origin}/api/oauth/ourspark&appid=${oauthAppid}`;
      switch (thridPlatform) {
        case 'ourspark': {
          document.location.href = OURSPARK_URL;
          break;
        }
        default: {
          document.location.href = OURSPARK_URL;
        }
      }
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/style/login.scss";
</style>
