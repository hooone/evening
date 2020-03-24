import { List, Divider, Button, Modal } from 'antd';
import { connect, useIntl } from 'umi';
import { EyeOutlined, EyeInvisibleOutlined, TagOutlined, CodeOutlined, ExclamationCircleOutlined } from '@ant-design/icons';

const { confirm } = Modal;

const ActionsList = ({ actionList, dispatch }) => {
    const intl = useIntl();
    function createAction() {
        dispatch({
            type: 'actionInfoConfig/show',
            action: {
                Id: 0,
                CardId: actionList.cardId,
                Name: "",
                Type: "UPDATE",
                Seq: 0,
                DoubleCheck: false,
            },
        });
        dispatch({
            type: 'draw/subOpen',
            title: intl.formatMessage(
                {
                    id: 'create',
                }
            ),
        });
    }
    function updateAction(action) {
        console.log(action)
        dispatch({
            type: 'actionInfoConfig/show',
            action: action,
        });
        dispatch({
            type: 'draw/subOpen',
            title: action.Text,
        });
    }
    function updateParameter(action) {
        dispatch({
            type: 'parameterInfoConfig/show',
            parameters: action.Parameters,
            cardId: actionList.cardId,
        });
        dispatch({
            type: 'draw/subOpen',
            title: action.Text,
        });
    }
    function deleteAction(action) {
        confirm({
            title:
                (<span><span>
                    {intl.formatMessage(
                        {
                            id: 'confirm',
                        }
                    )}
                </span>
                    <span>
                        {intl.formatMessage(
                            {
                                id: 'delete',
                            }
                        )}
                    ?
                    </span>
                </span>)
            ,
            icon: <ExclamationCircleOutlined />,
            onOk() {
                dispatch({
                    type: 'actionList/deleteAction',
                    action: action,
                });
            },
            onCancel() {

            },
        });

    }


    function findAttribute(ele, attr) {
        if (!ele)
            return '';
        while (!ele.getAttribute(attr)) {
            if (ele.id === 'root')
                return '';
            else {
                ele = ele.parentElement;
            }
            if (!ele)
                return '';
        }
        return ele.getAttribute(attr);
    }
    function onDragStart(e) {
        let real = e.target;
        if (real.tagName === "LI" && real.classList.contains("ant-list-item")) {
            e.dataTransfer.setData("id", findAttribute(real, "data-id"));
            e.dataTransfer.setData("type", findAttribute(real, "data-type"));
        }
    }
    function getRealDOM(s) {
        let real = s;
        if (real.tagName === "path") {
            real = real.parentElement
        }
        if (real.tagName === "svg") {
            real = real.parentElement
        }
        if (real.tagName === "EM") {
            real = real.parentElement
        }
        if (real.tagName === "A") {
            real = real.parentElement
        }
        if (real.tagName === "I") {
            real = real.parentElement
        }
        if (real.tagName === "LI" && !real.classList.contains("ant-list-item")) {
            real = real.parentElement
        }
        if (real.tagName === "UL") {
            real = real.parentElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement
        }
        if (real.tagName === "H4") {
            real = real.parentElement
        }
        if (real.tagName === "DIV") {
            real = real.parentElement
        }
        if (real.tagName === "DIV") {
            real = real.parentElement
        }
        if (real.tagName === "DIV") {
            real = real.parentElement
        }
        return real
    }
    function onDrop(e) {
        e.preventDefault()
        let formdata = {
            Source: 0,
            Target: 0,
            Position: 0,
        }
        let real = getRealDOM(e.target)
        real.classList.add("dragEnterList")
        if (real.tagName === "LI" && real.classList.contains("ant-list-item")) {
            real.classList.remove("dragEnterList")
            formdata.Source = e.dataTransfer.getData("id");
            formdata.Target = findAttribute(real, "data-id")
            let sourceType = e.dataTransfer.getData("type");
            if (e.clientY - 71 - real.offsetTop < real.clientHeight / 2) {
                formdata.Position = 1;
            }
            else {
                formdata.Position = 2;
            }
            dispatch({
                type: 'actionList/updateSeq',
                move: formdata,
                cardId: actionList.cardId
            });
        }

    }
    function onDragOver(e) {
        let real = getRealDOM(e.target);
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            if (e.clientY - 71 - real.offsetTop < real.clientHeight / 2) {
                real.classList.add("dragTop")
                real.classList.remove("dragInner")
                real.classList.remove("dragBottom")
            }
            else {
                real.classList.remove("dragTop")
                real.classList.remove("dragInner")
                real.classList.add("dragBottom")
            }
        }
        e.preventDefault()
    }
    function onDragEnter(e) {
        let real = getRealDOM(e.target);
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            real.classList.add("dragEnterList")
        }
        e.preventDefault()
    }
    function onDragLeave(e) {
        let real = e.target
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            real.classList.remove("dragEnterList")
        }
        e.preventDefault()
    }

    return <List
        dataSource={actionList.actions}
        footer={<Button type="dashed" block onClick={() => { createAction() }}>新增</Button>}
        renderItem={item => (
            <List.Item
                className="drawer-list"
                draggable="true"
                onDragStart={(event) => { onDragStart(event) }}
                onDrop={(event) => { onDrop(event) }}
                onDragEnter={(event) => { onDragEnter(event) }}
                onDragLeave={(event) => { onDragLeave(event) }}
                onDragOver={(event) => { onDragOver(event) }}
                allowdrop="true"
                key={item.id} data-id={item.Id} data-type={"action"}
                actions={[
                    <Button type="link" onClick={() => { updateAction(item) }} key={`am-${item.Id}`}>
                        修改
            </Button>,
                    <Button type="link" onClick={() => { updateParameter(item) }} key={`ap-${item.Id}`}>
                        参数
                                        </Button>,
                    <Button type="link" danger onClick={() => { deleteAction(item) }} key={`ad-${item.Id}`}>
                        删除
            </Button>,
                ]}
            >
                <List.Item.Meta
                    title={(item.IsVisible) ? (<span>
                        {item.Name}
                        <Divider type="vertical" />
                        <EyeOutlined style={{ marginRight: 6 }} />
                    </span>) : (<span>
                        {item.Name}
                        <Divider type="vertical" />
                        <EyeInvisibleOutlined style={{ marginRight: 6 }} />
                    </span>)}
                    description={<span>
                        <TagOutlined style={{ marginRight: 6 }} />
                        {item.Text}
                        <Divider type="vertical" />
                        <CodeOutlined style={{ marginRight: 6 }} />
                        {item.Type}

                    </span>}
                />
            </List.Item>
        )}
    />
};
export default connect(({ actionList }) => ({
    actionList,
}))(ActionsList);