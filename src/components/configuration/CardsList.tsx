import { List, Divider, Button, Modal } from 'antd';
import { connect, useIntl } from 'umi';
import { EyeOutlined, EyeInvisibleOutlined, TagOutlined, CodeOutlined, ExclamationCircleOutlined } from '@ant-design/icons';

const { confirm } = Modal;

const CardsList = ({ cardList, dispatch }) => {
    const intl = useIntl();
    console.log(cardList)
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
                type: 'cardList/updateSeq',
                move: formdata,
                cardId: cardList.cardId
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
        dataSource={cardList}
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
                key={item.Id} data-id={item.Id} data-type={"card"}
            >
                <List.Item.Meta
                    title={(item.IsVisible) ? (<span>
                        {item.Name}
                    </span>) : (<span>
                        {item.Name}
                        <Divider type="vertical" />
                        {intl.formatMessage(
                            {
                                id: item.Style,
                            }
                        )}
                    </span>)}
                    description={<span>
                        <TagOutlined style={{ marginRight: 6 }} />
                        {item.Text}
                    </span>}
                />
            </List.Item>
        )}
    />
};
export default connect(({ cardList }) => ({
    cardList,
}))(CardsList);