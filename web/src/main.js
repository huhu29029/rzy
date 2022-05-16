import Vue from 'vue';
import * as Sentry from '@sentry/vue';
import { BrowserTracing } from '@sentry/tracing';

import {
  Button,
  Link,
  Select,
  Dialog,
  Form,
  Input,
  FormItem,
  Option,
  Loading,
  Message,
  Container,
  Card,
  Dropdown,
  DropdownMenu,
  DropdownItem,
  Row,
  Col,
  Menu,
  Submenu,
  MenuItem,
  Aside,
  Main,
  Badge,
  Header,
  Tabs,
  Breadcrumb,
  BreadcrumbItem,
  Scrollbar,
  Avatar,
  TabPane,
  Divider,
  Table,
  TableColumn,
  Cascader,
  Checkbox,
  CheckboxGroup,
  Pagination,
  Tag,
  Drawer,
  Tree,
  Popover,
  Switch,
  Collapse,
  CollapseItem,
  Tooltip,
  DatePicker,
  InputNumber,
  Steps,
  Upload,
  Progress,
  Radio,
  RadioGroup,
  MessageBox,
  Empty,
} from 'element-ui';
import VueClipboard from 'vue-clipboard2';
import uploader from 'vue-simple-uploader';
import App from './App.vue';
// let md = require('markdown-it')()

//  按需引入element

// 引入封装的router
import router from '@/router/index';

// time line css
import 'timeline-vuejs/dist/timeline-vuejs.css';

import '@/permission';
import { store } from '@/store/index';

// 路由守卫
import Bus from '@/utils/bus.js';

import { auth } from '@/directive/auth';

const md = require('markdown-it')().use(require('markdown-it-mathjax')());

Vue.use(VueClipboard);

Vue.use(Button);
Vue.use(Link);
Vue.use(Select);
Vue.use(Dialog);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Input);
Vue.use(Option);
Vue.use(Container);
Vue.use(Card);
Vue.use(Dropdown);
Vue.use(DropdownMenu);
Vue.use(DropdownItem);
Vue.use(Row);
Vue.use(Col);
Vue.use(Menu);
Vue.use(Submenu);
Vue.use(MenuItem);
Vue.use(Aside);
Vue.use(Main);
Vue.use(Badge);
Vue.use(Header);
Vue.use(Tabs);
Vue.use(Breadcrumb);
Vue.use(BreadcrumbItem);
Vue.use(Avatar);
Vue.use(TabPane);
Vue.use(Divider);
Vue.use(Table);
Vue.use(TableColumn);
Vue.use(Checkbox);
Vue.use(Cascader);
Vue.use(Tag);
Vue.use(Pagination);
Vue.use(Drawer);
Vue.use(Tree);
Vue.use(CheckboxGroup);
Vue.use(Popover);
Vue.use(InputNumber);
Vue.use(Switch);
Vue.use(Collapse);
Vue.use(CollapseItem);
Vue.use(Tooltip);
Vue.use(DatePicker);
Vue.use(Steps);
Vue.use(Upload);
Vue.use(Progress);
Vue.use(Scrollbar);
Vue.use(Loading.directive);
Vue.use(Radio);
Vue.use(RadioGroup);
Vue.use(Empty);

Vue.prototype.$loading = Loading.service;
Vue.prototype.$message = Message;
Vue.prototype.$confirm = MessageBox.confirm;
Dialog.props.closeOnClickModal.default = false;
Vue.config.productionTip = false;
Vue.use(Bus);
// 按钮权限指令
auth(Vue);
Vue.use(uploader);

// render markdown to html
Vue.prototype.$md = md;

// Sentry
Sentry.init({
  Vue,
  dsn: 'https://ec9ad48975134a18944f1a04a18a9628@o338003.ingest.sentry.io/6313617',
  integrations: [
    new BrowserTracing({
      routingInstrumentation: Sentry.vueRouterInstrumentation(router),
      tracingOrigins: ['exam.feel.ac.cn'],
    }),
  ],
  tracesSampleRate: 1.0,
});

export default new Vue({
  render: (h) => h(App),
  router,
  store,
}).$mount('#app');
