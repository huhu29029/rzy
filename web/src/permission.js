import router from './router/index';
import { store } from '@/store/index';
import getPageTitle from '@/utils/page';
import { newXAPIData } from '@/api/xapi';

let asyncRouterFlag = 0;

const whiteList = ['login', 'init', 'oauth-ourspark'];
router.beforeEach(async (to, from, next) => {
  const token = store.getters['user/token'];
  // 在白名单中的判断情况
  // 修改网页标签名称
  document.title = getPageTitle(to.meta.title);
  if (whiteList.indexOf(to.name) > -1) {
    if (token) {
      next({ name: store.getters['user/userInfo'].authority.defaultRouter });
    } else {
      next();
    }
  } else {
    // 不在白名单中并且已经登陆的时候
    if (token) {
      // xapi: 数据埋点
      await newXAPIData({
        verb: 'navigate',
        object: {
          info: to.fullPath,
        },
      });
      next();
      // 添加flag防止多次获取动态路由和栈溢出
      if (!asyncRouterFlag && store.getters['router/routerList'].length === 0) {
        asyncRouterFlag++;
        const user = store.getters['user/userInfo'];
        const { authorityId } = user;
        await store.dispatch('router/SetRouterList', authorityId);

        next({ ...to, replace: true });
      } else {
        next();
      }
    }
    // 不在白名单中并且未登陆的时候
    if (!token) {
      next({
        name: 'login',
        query: {
          redirect: document.location.hash.split('#')[1],
        },
      });
    }
  }
});
