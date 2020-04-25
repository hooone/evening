
import { EffectsCommandMap } from 'dva'
export interface <%- card.Name %> {
    <% card.Fields.forEach(
        field => {%>
    <%- field.Name +": "+((field.Type=="date"||field.Type=="datetime")?"Date":(field.Type=="int"||field.Type=="float")?"number":field.Type)+","%>
            <%}
    ) %>
}
export default {
    namespace: '<%- card.Name %>model',
    state: {
        "Data": <%- JSON.stringify(card.Data) %>,
        <% card.Actions.forEach(action=>{
            let needInput=false;
            action.Parameters.forEach(param=>{
                needInput=needInput||(param.IsEditable&&param.IsVisible)
            })
            if(needInput){%>
        <%- '"'+action.Name+'ModalVisible": false,'%>
        <%}})%>
    },
    reducers: {
        save(state: any, action: user[]) {
            return {
                "Data": action,
                <% card.Actions.forEach(action=>{
                    let needInput=false;
                    action.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })
                    if(needInput){%>
                <%- '"'+action.Name+'ModalVisible": false,'%>
                <%}})%>
            };
        },
        modalcancel(state: any) {
            return {
                "Data": state.Data,
                <% card.Actions.forEach(action=>{
                    let needInput=false;
                    action.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })
                    if(needInput){%>
                <%- '"'+action.Name+'ModalVisible": false,'%>
                <%}})%>
            };
        },
        <% card.Actions.forEach(action=>{
            let needInput=false;
            action.Parameters.forEach(param=>{
                needInput=needInput||(param.IsEditable&&param.IsVisible)
            })
            if(needInput){%>
        <%- action.Name+'ModalShow(state: any, action: '+card.Name+') {'%>
            <%- 'return {'%>
                <%- '"Data": state.Data,'%>
                <% card.Actions.forEach(act=>{
                    let needInput=false;
                    act.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })
                    if(needInput){%>
                <%- '"'+act.Name+'ModalVisible": '+(act.Name==action.Name?'true':'false')+','%>
                <%}})%>
            <%- '}'%>
        <%- '},'%>
        <%}})%>
    },
    effects: {
        <% card.Actions.forEach(
            action => {%>
        <%- "*"+action.Name+"(action: any, handler: EffectsCommandMap) {" %>

        <%- "},"%>
                <%}
        ) %>
    },
};

