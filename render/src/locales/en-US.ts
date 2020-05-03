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
    min:" min",
    max:" max",
}
export default en_US; 