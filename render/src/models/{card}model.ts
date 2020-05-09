import { EffectsCommandMap, SubscriptionAPI } from 'dva'
import moment from 'moment';
import { AJAX, CommonResult } from '@/util';
export interface <%- card.Name %> {
    <% card.Fields.forEach(
        field => {%>
    <%- field.Name +": "+((field.Type=="date"||field.Type=="datetime")?"Date":(field.Type=="int"||field.Type=="float")?"number":field.Type)+","%>
            <%}
    ) %>
}
const dateFormat = 'YYYY-MM-DD';
const datetimeFormat = 'YYYY-MM-DD HH:mm:ss';
export default {
    namespace: '<%- card.Name %>model',
    state: {
        "Data": [],
        <% card.Actions.forEach(action=>{
            let needInput=false;
            action.Parameters.forEach(param=>{
                needInput=needInput||(param.IsEditable&&param.IsVisible)
            })
            if(needInput){%>
        <%- '"'+action.Name+'ModalVisible": false,'%>
        <%- '"'+action.Name+'FormData": ['%>
        <% action.Parameters.forEach(param=>{%>
            <%- '{'%>
                <%- '"name": "'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '",'%>
                <%if (param.Field.Type === 'int') {%>
                <%- '"value": '+(param.Default==""?"0":param.Default) %>
                <%} else if (param.Field.Type === 'datetime') {%>
                <%- '"value": moment(new Date(new Date().getTime() - ('+(param.Default==""?"0":param.Default)+')), datetimeFormat)' %>
                <%} else if (param.Field.Type === 'date') {%>
                <%- '"value": moment(new Date(new Date().getTime() - ('+(param.Default==""?"0":param.Default)+')), dateFormat)' %>
                <%}else if (param.Field.Type === 'bool') {%> 
                <%- '"value": "'+(param.Default=="true"?"true":"false")+'"' %>
                <%}else {%>
                <%- '"value": "'+param.Default+'"' %>
                <%}%>
               
            <%- '},'%>
        <% })%>
        <%- '],'%>
        <%}})%>
    },
    reducers: {
        save(state: any, action: any) {
            return {
                "Data": action.Data,
                <% card.Actions.forEach(action=>{
                    let needInput=false;
                    action.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })
                    if(needInput){%>
                <%- '"'+action.Name+'ModalVisible": false,'%>
                <%- '"'+action.Name+'FormData": state.'+(action.Name)+'FormData,'%>
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
                <%- '"'+action.Name+'FormData": state.'+(action.Name)+'FormData,'%>
                <%}})%>
            };
        },
        <% card.Actions.forEach(action=>{
            let needInput=false;
            action.Parameters.forEach(param=>{
                needInput=needInput||(param.IsEditable&&param.IsVisible)
            })
            if(needInput){%>
        <%- action.Name+'ModalShow(state: any, action: any) {'%>
            <%- 'return {'%>
                <%- '"Data": state.Data,'%>
                <% card.Actions.forEach(act=>{
                    let needInput=false;
                    act.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })
                    if(needInput){%>
                <%- '"'+act.Name+'ModalVisible": '+(act.Name==action.Name?'true':'false')+','%>
                <% if(act.Name==action.Name&&(action.Type==="UPDATE"||action.Type==="DELETE")) {%>
                <%- '"'+act.Name+'FormData": ['%>
                <% action.Parameters.forEach(param=>{%>
                    <%- '{'%>
                        <%- '"name": "'+param.Field.Name+'",'%>
                        <%if (param.Field.Type === 'int') {%>
                        <%- '"value": action.record.'+(param.Field.Name) %>
                        <%} else if (param.Field.Type === 'datetime') {%>
                        <%- '"value": moment(new Date(action.record.'+(param.Field.Name)+'.replace("T", " ").replace("-", "/").replace("-", "/").replace("-", "/").replace("-", "/")), datetimeFormat)' %>
                        <%} else if (param.Field.Type === 'date') {%>
                        <%- '"value": moment(new Date(action.record.'+(param.Field.Name)+'.replace("T", " ").replace("-", "/").replace("-", "/").replace("-", "/").replace("-", "/")), dateFormat)' %>
                        <%}else if (param.Field.Type === 'bool') {%> 
                        <%- '"value": action.record.'+(param.Field.Name)+'? "true": "false"' %>
                        <%}else {%>
                        <%- '"value": action.record.'+(param.Field.Name) %>
                        <%}%>
                        
                    <%- '},'%>
                <% })%>
                <%- '],'%>
                <% }else { %>
                <%- '"'+act.Name+'FormData": state.'+(act.Name)+'FormData,'%>
                <%}}})%>
            <%- '}'%>
        <%- '},'%>
        <%- action.Name+'ModalChange(state: any, action: any) {'%>
            <%- 'return {'%>
                <%- '"Data": state.Data,'%>
                <% card.Actions.forEach(act=>{
                    let needInput=false;
                    act.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })
                    if(needInput){%>
                <%- '"'+act.Name+'ModalVisible": '+(act.Name==action.Name?'true':'false')+','%>
                <% if(act.Name==action.Name) {%>
                <%- '"'+act.Name+'FormData": action.Fields,'%>
                <% }else { %>
                <%- '"'+act.Name+'FormData": state.'+(act.Name)+'FormData,'%>
                <%}}})%>
            <%- '}'%>
        <%- '},'%>
        <%}})%>
    },
    effects: {
        <% card.Actions.forEach(
            action => {
            let needInput=false;
            action.Parameters.forEach(param=>{
                needInput=needInput||(param.IsEditable&&param.IsVisible)
            })%>
        <% if(needInput){ %>
        <%- "*"+action.Name+"(action: any, handler: EffectsCommandMap) {" %>
            <%- 'let fields = yield handler.select((state: any) => state.'+card.Name+'model.'+action.Name+'FormData);' %>
            <%- 'let formdata = {' %>
            <% action.Parameters.forEach((param,idx)=>{%>
                <%if (param.Field.Type === 'int') {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": fields['+idx+'].value,' %>
                <%} else if (param.Field.Type === 'datetime') {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": fields['+idx+'].value.format(datetimeFormat),' %>
                <%} else if (param.Field.Type === 'date') {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": fields['+idx+'].value.format(dateFormat),' %>
                <%}else if (param.Field.Type === 'bool') {%> 
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": fields['+idx+'].value,' %>
                <%}else {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": fields['+idx+'].value,' %>
                <%}%>
            <% })%>
            <%- '}' %>
            <%- 'console.log(formdata)' %>
            <%- 'let result: CommonResult = yield AJAX("/'+card.Name+'/'+action.Name+'", formdata);' %>
            <%- 'if (!result.Success) {' %>
                <%- 'return' %>
            <%- '}' %>
            <% if(action.Type==="READ"){ %>
            <%- 'yield handler.put({ type: "save", Data: result.Data });' %>
            <% }else {%>
            <%- 'yield handler.put({ type: "Read" });' %>
            <% } %>
        <%- "},"%>
        <% }else{ %>
        <%- "*"+action.Name+"(action: any, handler: EffectsCommandMap) {" %>
            <%- 'let formdata = {' %>
            <% action.Parameters.forEach((param,idx)=>{%>
                <%if (param.Field.Type === 'int') {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": action.record.'+param.Field.Name+',' %>
                <%} else if (param.Field.Type === 'datetime') {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": action.record.'+param.Field.Name+',' %>
                <%} else if (param.Field.Type === 'date') {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": action.record.'+param.Field.Name+',' %>
                <%}else if (param.Field.Type === 'bool') {%> 
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": action.record.'+param.Field.Name+',' %>
                <%}else {%>
                <%- '"'+param.Field.Name+((param.Compare === 'gt' || param.Compare === 'ge')?'_min':((param.Compare === 'lt' || param.Compare === 'le')?'_max':''))+ '": action.record.'+param.Field.Name+',' %>
                <%}%>
            <% })%>
            <%- '}' %>
            <%- 'console.log(formdata)' %>
            <%- 'let result: CommonResult = yield AJAX("/'+card.Name+'/'+action.Name+'", formdata);' %>
            <%- 'if (!result.Success) {' %>
                <%- 'return' %>
            <%- '}' %>
            <% if(action.Type==="READ"){ %>
            <%- 'yield handler.put({ type: "save", Data: result.Data });' %>
            <% }else {%>
            <%- 'yield handler.put({ type: "Read" });' %>
            <% } %>
        <%- "},"%>
        <% }%>
                <%}
        ) %>
    },
    subscriptions: {
        setup(handler: SubscriptionAPI) {
            handler.dispatch({ type: "Read" });
        },
    }
};

