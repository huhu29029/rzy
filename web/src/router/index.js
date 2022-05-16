import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

// 获取原型对象上的push函数
const originalPush = Router.prototype.push;
// 修改原型对象中的push方法
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch((err) => err);
};

const baseRouters = [{
  path: '/',
  redirect: '/login',
}, {
  path: '/login',
  name: 'login',
  component: () => import('@/view/login/login.vue'),
}, {
  path: '/oauth',
  name: 'oauth',
  component: () => import('@/view/oauth/index.vue'),
  children: [{
    path: 'ourspark',
    name: 'oauth-ourspark',
    component: () => import('@/view/oauth/ourspark/index.vue'),
  }],
}, {
  path: '/layout',
  name: 'layout',
  meta: {
    title: '底层layout',
  },
  component: () => import('@/view/layout/index.vue'),
  children: [
    {
      path: 'dashboard',
      name: 'dashboard',
      meta: {
        title: '主页',
      },
      component: () => import('@/view/dashboard/index.vue'),
    },
    {
      path: 'information',
      name: 'p-information',
      meta: {
        title: '个人信息',
      },
      component: () => import('@/view/person/information/person.vue'),
    },
    {
      path: 'aboutus',
      name: 'aboutus',
      meta: {
        title: '关于我们',
      },
      component: () => import('@/view/aboutus/index.vue'),
    },
    // {
    //   path: 'sharing',
    //   name: 'sharing',
    //   meta: {
    //     title: '资料共享',
    //   },
    //   component: () => import('@/view/sharing/index.vue'),
    // },
    {
      path: 'history',
      name: 'history',
      meta: {
        title: '历史',
      },
      redirect: {
        name: 'p-history',
      },
      component: () => import('@/view/history/index.vue'),
      children: [
        {
          path: 'guide',
          name: 'p-history',
          meta: {
            title: '做题历史',
          },
          component: () => import('@/view/history/guide/index.vue'),
        },
        {
          path: 'practice',
          name: 'h-practice',
          meta: {
            title: '练习历史',
          },
          component: () => import('@/view/history/practice/index.vue'),
        },
        {
          path: 'exam',
          component: () => import('@/view/history/exam/index.vue'),
          name: 'h-exam',
          meta: {
            title: '考试',
          },
          redirect: {
            name: 'h-exam-list',
          },
          children: [
            {
              path: 'list',
              name: 'h-exam-list',
              meta: {
                title: '考试历史',
              },
              component: () => import('@/view/history/exam/list/index.vue'),
            },
            {
              path: 'detail',
              name: 'h-exam-detail',
              meta: {
                title: '考试详情',
              },
              component: () => import('@/view/history/exam/detail/index.vue'),
            },
          ],
        },
      ],
    }, {
      path: 'about',
      name: 'about',
      meta: {
        title: '关于我们',
      },
      component: () => import('@/view/about/index.vue'),
    }, {
      path: '404',
      name: '404',
      meta: {
        title: '页面未找到',
      },
      component: () => import('@/view/error/index.vue'),
    }, {
      path: 'questions',
      name: 'questions',
      meta: {
        title: '练习',
      },
      redirect: {
        name: 'questions-guide',
      },
      component: () => import('@/view/questions/index.vue'),
      children: [{
        path: 'guide',
        name: 'questions-guide',
        meta: {
          title: '答题指引',
        },
        component: () => import('@/view/questions/guide/index.vue'),
      }, {
        // path: 'practice/:questionId',
        path: 'practice',
        name: 'questions-practice',
        meta: {
          title: '练习模式',
        },
        component: () => import('@/view/questions/practice/index.vue'),
      }],
    }, {
      path: 'exam',
      name: 'exam',
      meta: {
        title: '考试',
      },
      redirect: {
        name: 'exam-guide',
      },
      component: () => import('@/view/exam/index.vue'),
      children: [{
        path: 'guide',
        name: 'exam-guide',
        meta: {
          title: '考试模式',
        },
        component: () => import('@/view/exam/guide/index.vue'),
      }, {
        path: 'paper',
        name: 'exam-paper',
        meta: {
          title: '试卷',
        },
        component: () => import('@/view/exam/paper/index.vue'),
      }, {
        path: 'exercises',
        name: 'exam-exercises',
        meta: {
          title: '考后小练',
        },
        component: () => import('@/view/exam/exercises/index.vue'),
      },
      ],
    }, {
    //   path: 'studyplan',
    //   name: 'studyplan',
    //   meta: {
    //     title: '学习计划',
    //   },
    //   redirect: {
    //     name: 'studyplan-guide',
    //   },
    //   component: () => import('@/view/studyplan/index.vue'),
    //   children: [{
    //     path: 'guide',
    //     name: 'studyplan-guide',
    //     meta: {
    //       title: '学习计划',
    //     },
    //     component: () => import('@/view/studyplan/guide/index.vue'),
    //   }, {
    //     path: 'dailyplan',
    //     name: 'studyplan-dailyplan',
    //     meta: {
    //       title: '每日计划',
    //     },
    //     component: () => import('@/view/studyplan/dailyplan/index.vue'),
    //   }],
    // }, {
    //   path: 'myerror',
    //   name: 'myerror',
    //   meta: {
    //     title: '错题记录',
    //   },
    //   redirect: {
    //     name: 'myerror-guide',
    //   },
    //   component: () => import('@/view/myerror/index.vue'),
    //   children: [{
    //     path: 'guide',
    //     name: 'myerror-guide',
    //     meta: {
    //       title: '错题模式',
    //     },
    //     component: () => import('@/view/myerror/guide/index.vue'),
    //   }, {
    //     path: 'list',
    //     name: 'myerror-list',
    //     meta: {
    //       title: '错题列表',
    //     },
    //     component: () => import('@/view/myerror/list/index.vue'),
    //   }, {
    //     path: 'mistakes',
    //     name: 'myerror-mistakes',
    //     meta: {
    //       title: '错题集',
    //     },
    //     component: () => import('@/view/myerror/mistakes/index.vue'),
    //   }],
    // }, {
      path: 'management',
      name: 'management',
      meta: {
        title: '管理',
      },
      redirect: {
        name: 'm-questions',
      },
      component: () => import('@/view/management/index.vue'),
      children: [
        {
          path: 'questions',
          name: 'm-questions',
          meta: {
            title: '试题管理',
          },
          component: () => import('@/view/management/questions/index.vue'),
        },
        {
          path: 'users',
          name: 'm-users',
          meta: {
            title: '用户管理',
          },
          redirect: {
            name: 'm-guide',
          },
          component: () => import('@/view/management/users/index.vue'),
          children: [
            {
              path: 'guide',
              name: 'm-guide',
              meta: {
                title: '用户',
              },
              component: () => import('@/view/management/users/guide/index.vue'),
            },
            {
              path: 'userInfo',
              name: 'm-userInfo',
              meta: {
                title: '用户详情',
              },
              component: () => import('@/view/management/users/userInfo/index.vue'),
            },
            {
              path: 'userExam',
              name: 'm-userExam',
              meta: {
                title: '用户试卷历史',
              },
              component: () => import('@/view/management/users/userInfo/exam/index.vue'),
            },
            {
              path: 'userPractice',
              name: 'm-userPractice',
              meta: {
                title: '用户练习历史',
              },
              component: () => import('@/view/management/users/userInfo/practice/index.vue'),
            },
          ],
        },
        {
          path: 'exams',
          name: 'm-exams',
          meta: {
            title: '试卷管理',
          },
          redirect: {
            name: 'exam-list',
          },
          component: () => import('@/view/management/exams/index.vue'),
          children: [
            {
              path: '',
              name: 'exam-list',
              meta: {
                title: '试卷',
              },
              component: () => import('@/view/management/exams/list/index.vue'),
            },
            {
              path: 'new',
              name: 'exam-new',
              meta: {
                title: '新增试卷',
              },
              component: () => import('@/view/management/exams/new/index.vue'),
            },
            {
              path: 'edit',
              name: 'exam-edit',
              meta: {
                title: '试卷编辑',
              },
              component: () => import('@/view/management/exams/edit/index.vue'),
            },
          ],
        },
      ],
    },
  ],
}, {
  path: '*',
  redirect: '/layout/404',
}];

// 需要通过后台数据来生成的组件
const createRouter = () => new Router({
  routes: baseRouters,
});

const router = createRouter();

export default router;
