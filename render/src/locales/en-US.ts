const en_US = {
    <% locales.forEach(
        locale => {%>
    <%- locale.Name+': "'+(locale["en-US"]?locale["en-US"]:locale.Default)+'",' %>
            <%}
    ) %>
    
    actions:"actions",
    create:"create",
    read:"read",
    update:"update",
    delete:"delete",
    cancel:"cancel",
    confirm:"confirm",
}
export default en_US; 