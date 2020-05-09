export default {
<% card.Actions.forEach(act=>{ %>
    <%- '"POST /'+card.Name+'/'+act.Name+'": {'%>
        <%- 'Success: true,'%>
        <% if(act.Type==="READ"){ %>
        <%- 'Data: '+JSON.stringify(card.Data)%>
        <%}else{%>
        <%- 'Data: []'%>
        <%}%>
    <%- '},'%>
<%})%>
}