import { login, register, getUserInfo } from '@/api/user';
import { jsonInBlacklist } from '@/api/jwt';
import router from '@/router/index';
import { parseHash } from '@/utils/qs';

const user = {
  namespaced: true,
  state: {
    userInfo: {
      uuid: '',
      nickName: '',
      headerImg: '',
      authority: '',
    },
    token: '',
  },
  mutations: {
    setUserInfo(state, userInfo) {
      // 这里的 `state` 对象是模块的局部状态
      state.userInfo = userInfo;
    },
    setToken(state, token) {
      // 这里的 `state` 对象是模块的局部状态
      state.token = token;
    },
    NeedInit(state) {
      state.userInfo = {};
      state.token = '';
      sessionStorage.clear();
      router.push({ name: 'init', replace: true });
    },
    LoginOut(state) {
      state.userInfo = {};
      state.token = '';
      sessionStorage.clear();
      router.push({ name: 'login', replace: true });
      window.location.reload();
    },
    ResetUserInfo(state, userInfo = {}) {
      state.userInfo = {
        ...state.userInfo,
        ...userInfo,
      };
    },
  },
  actions: {
    async LoginIn({ commit, dispatch, getters }, loginInfo) {
      const res = await login(loginInfo);
      if (res.code === 0) {
        commit('setToken', res.data.token);
        const userRes = await getUserInfo();
        if (userRes.code === 0) {
          commit('setUserInfo', userRes.data.user);
          await dispatch('router/SetRouterList', userRes.data.user.authorityId, { root: true });

          // 路由跳转
          const { hash } = document.location;
          const queryObj = parseHash(hash);
          const { redirect } = queryObj;
          if (redirect !== undefined) {
            router.push(redirect);
          } else {
            router.push({ name: getters.userInfo.authority.defaultRouter });
          }
        }
        return true;
      }
      return false;
    },
    async Register({ commit, dispatch, getters }, registerInfo) {
      const res = await register(registerInfo);
      if (res.code === 0) {
        commit('setToken', res.data.token);
        const userRes = await getUserInfo();
        if (userRes.code === 0) {
          commit('setUserInfo', userRes.data.user);
          await dispatch('router/SetRouterList', userRes.data.user.authorityId, { root: true });

          // 路由跳转
          const { hash } = document.location;
          const queryObj = parseHash(hash);
          const { redirect } = queryObj;
          if (redirect !== undefined) {
            router.push(redirect);
          } else {
            router.push({ name: getters.userInfo.authority.defaultRouter });
          }
        }
        return true;
      }
      return false;
    },
    async LoginOut({ commit }) {
      const res = await jsonInBlacklist();
      if (res.code === 0) {
        commit('LoginOut');
      }
    },
  },
  getters: {
    userInfo(state) {
      return state.userInfo;
    },
    token(state) {
      return state.token;
    },

  },
};

export {
  user,
};
