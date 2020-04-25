const zh_CN = {
    <% locales.forEach(
        locale => {%>
    <%- locale.Name+': "'+(locale["zh-CN"]?locale["zh-CN"]:locale.Default)+'",' %>
            <%}
    ) %>
    
    actions:"操作",
    create:"添加",
    read:"查询",
    update:"修改",
    delete:"删除",
    cancel:"取消",
    confirm:"确认",
}
export default zh_CN; 