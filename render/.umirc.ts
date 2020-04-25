import { defineConfig } from 'umi';

export default defineConfig({
  routes: [
    { exact: true, path: '/', redirect: '/home' },
    <% folders.forEach(folder => {%>
      <%  if (!folder.IsFolder) {
        if (folder.Pages !== null && folder.Pages.length === 1) {%>
    <%- "{ exact: true, path: '/" + folder.Pages[0].Name + "', component: '@/pages/" + folder.Pages[0].Name + "' }," %>
          <%  }
      }else {
          if(folder.Pages !== null) {
  folder.Pages.forEach(
    node => {%>
    <%- "{ exact: true, path: '/" + folder.Name + "/" + node.Name + "', component: '@/pages/" + node.Name + "' }," %>
      <% }
  )
}
           } %>
  <% }); %>
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
});