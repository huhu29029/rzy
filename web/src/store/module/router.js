const adminRouterList = [{
  children: [
    {
      path: 'dashboard',
      name: 'dashboard',
      meta: {
        title: '主页',
        icon: 'house',
      },
    }, {
      path: 'questions',
      name: 'questions',
      meta: {
        title: '练习',
        icon: 'edit',
      },
    }, {
      path: 'exam',
      name: 'exam',
      meta: {
        title: '考试',
        icon: 'document',
      },
    }, {
    //   path: 'sharing',
    //   name: 'sharing',
    //   meta: {
    //     title: '资料共享',
    //     icon: 'message',
    //   },
    // }, {
    //   path: 'studyplan',
    //   name: 'studyplan',
    //   meta: {
    //     title: '学习计划',
    //     icon: 'sunrise',
    //   },
    // }, {
    //   path: 'myerror',
    //   name: 'myerror',
    //   meta: {
    //     title: '错题记录',
    //     icon: 'paperclip',
    //   },
    // }, {
      path: 'history',
      name: 'p-history',
      meta: {
        title: '做题历史',
        icon: 'time',
      },
    },
    {
      path: 'information',
      name: 'p-information',
      meta: {
        title: '个人信息',
        icon: 'user-solid',
      },
    },
    {
      path: 'management',
      name: 'management',
      meta: {
        title: '管理',
        icon: 's-management',
      },
      children: [
        {
          path: 'questions',
          name: 'm-questions',
          meta: {
            title: '试题管理',
            icon: 'notebook-1',
          },
        },
        {
          path: 'exams',
          name: 'm-exams',
          meta: {
            title: '试卷管理',
            icon: 'document-copy',
          },
        },
        {
          path: 'users',
          name: 'm-users',
          meta: {
            title: '用户管理',
            icon: 'user',
          },
        },
      ],
    },
    {
      path: 'aboutus',
      name: 'aboutus',
      meta: {
        title: '关于我们',
        icon: 'info',
      },
    },
  ],
},
];

const userRouterList = [{
  children: [
    {
      path: 'dashboard',
      name: 'dashboard',
      meta: {
        title: '主页',
        icon: 'house',
      },
    }, {
      path: 'questions',
      name: 'questions',
      meta: {
        title: '练习',
        icon: 'edit',
      },
    }, {
      path: 'exam',
      name: 'exam',
      meta: {
        title: '考试',
        icon: 'document',
      },
    }, {
    //   path: 'sharing',
    //   name: 'sharing',
    //   meta: {
    //     title: '资料共享',
    //     icon: 'message',
    //   },
    // }, {
    //   path: 'studyplan',
    //   name: 'studyplan',
    //   meta: {
    //     title: '学习计划',
    //     icon: 'sunrise',
    //   },
    // }, {
    //   path: 'myerror',
    //   name: 'myerror',
    //   meta: {
    //     title: '错题记录',
    //     icon: 'paperclip',
    //   },
    // }, {
      path: 'history',
      name: 'p-history',
      meta: {
        title: '做题历史',
        icon: 'time',
      },
    },
    {
      path: 'information',
      name: 'p-information',
      meta: {
        title: '个人信息',
        icon: 'user-solid',
      },
    },
    {
      path: 'aboutus',
      name: 'aboutus',
      meta: {
        title: '关于我们',
        icon: 'info',
      },
    },
  ],
}];

const preacherRouterList = [{
  children: [
    {
      path: 'dashboard',
      name: 'dashboard',
      meta: {
        title: '主页',
        icon: 'house',
      },
    }, {
      path: 'questions',
      name: 'questions',
      meta: {
        title: '练习',
        icon: 'edit',
      },
    }, {
      path: 'exam',
      name: 'exam',
      meta: {
        title: '考试',
        icon: 'document',
      },
    }, {
    //   path: 'sharing',
    //   name: 'sharing',
    //   meta: {
    //     title: '资料共享',
    //     icon: 'message',
    //   },
    // }, {
    //   path: 'studyplan',
    //   name: 'studyplan',
    //   meta: {
    //     title: '学习计划',
    //     icon: 'sunrise',
    //   },
    // }, {
    //   path: 'myerror',
    //   name: 'myerror',
    //   meta: {
    //     title: '错题记录',
    //     icon: 'paperclip',
    //   },
    // }, {
      path: 'history',
      name: 'p-history',
      meta: {
        title: '做题历史',
        icon: 'time',
      },
    },
    {
      path: 'information',
      name: 'p-information',
      meta: {
        title: '个人信息',
        icon: 'user-solid',
      },
    },
    {
      path: 'management',
      name: 'management',
      meta: {
        title: '管理',
        icon: 's-management',
      },
      children: [{
        path: 'questions',
        name: 'm-questions',
        meta: {
          title: '试题管理',
          icon: 'notebook-1',
        },
      },
      {
        path: 'exams',
        name: 'm-exams',
        meta: {
          title: '试卷管理',
          icon: 'document-copy',
        },
      },
      {
        path: 'aboutus',
        name: 'aboutus',
        meta: {
          title: '关于我们',
          icon: 'info',
        },
      },
      ],
    },
  ],
}];

const typeMap = {
  ADMIN: 1,
  USER: 2,
  PREACHER: 3,
};

const router = {
  namespaced: true,
  state: {
    routerList: [],
  },
  mutations: {
    setRouterList(state, type) {
      if (type === typeMap.ADMIN) {
        state.routerList = adminRouterList;
      } else if (type === typeMap.USER) {
        state.routerList = userRouterList;
      } else {
        state.routerList = preacherRouterList;
      }
    },
  },
  actions: {
    SetRouterList({ commit }, authority) {
      commit('setRouterList', authority);
      return true;
    },
  },
  getters: {
    routerList(state) {
      return state.routerList;
    },
  },
};

export {
  router,
};
