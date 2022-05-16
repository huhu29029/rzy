<template>
  <div>
    <h1>
      火花统一登录 - 跳转中...
    </h1>
  </div>
</template>

<script>
import { getUserInfo } from '@/api/user';
import { parseHash } from '@/utils/qs';

export default {
  data() {
    return {
      key: '',
    };
  },
  async created() {
    this.key = this.$route.query.key;
    // return
    this.$store.commit('user/setToken', this.key);
    const userRes = await getUserInfo();
    if (userRes.code === 0) {
      this.$store.commit('user/setUserInfo', userRes.data.user);
      await this.$store.dispatch('router/SetRouterList', userRes.data.user.authorityId);

      // 路由跳转
      const hash = window.sessionStorage.getItem('online-exam:redirect');
      const queryObj = parseHash(hash);
      const { redirect } = queryObj;
      if (redirect !== undefined) {
        this.$router.push(redirect);
      } else {
        this.$router.push({ name: this.$store.getters['user/userInfo'].authority.defaultRouter });
      }
    }
  },
};
</script>
