import { defineConfig } from 'umi';

export default defineConfig({
  routes: [
    { exact: true, path: '/', redirect: '/home' },
    { exact: false, path: '/login', component: '@/pages/login' },
    { exact: false, path: '/', component: '@/pages/index' },
  ],
  dva: {
    skipModelValidate: true,
    immer: true,
    hmr: true,
  },
  locale: {
    default: 'zh-CN',
    antd: true,
    title: false,
    baseNavigator: true,
    baseSeparator: '-',
  },
  publicPath:"/static/"
});
