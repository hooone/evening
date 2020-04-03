import React, { InputHTMLAttributes, ReactElement } from 'react';
import { List, Divider, Button, Modal } from 'antd';
import { DispatchProp } from 'react-redux';
import { connect, useIntl } from 'umi';
import { EyeOutlined, EyeInvisibleOutlined, TagOutlined, CodeOutlined, ExclamationCircleOutlined } from '@ant-design/icons';
import { ICard } from '@/interfaces';
import { findAttribute } from '@/util';
import { IStore } from '@/store'

const { confirm } = Modal;

interface CardListProps extends DispatchProp {
    cardList: ICard[],
}
const CardsList = (props: CardListProps) => {
    const intl = useIntl();
    function onDragStart(e: React.DragEvent) {
        let real = e.target as HTMLElement;
        if (real.tagName === "LI" && real.classList.contains("ant-list-item")) {
            e.dataTransfer.setData("id", findAttribute(real, "data-id"));
            e.dataTransfer.setData("type", findAttribute(real, "data-type"));
        }
    }
    function getRealDOM(s: HTMLElement) {
        let real = s;
        if (real.tagName === "path") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "svg") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "EM") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "A") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "I") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "LI" && !real.classList.contains("ant-list-item")) {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "UL") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "H4") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "DIV") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "DIV") {
            real = real.parentElement as HTMLElement;
        }
        if (real.tagName === "DIV") {
            real = real.parentElement as HTMLElement;
        }
        return real
    }
    function onDrop(e: React.DragEvent) {
        e.preventDefault()
        let formdata = {
            Source: 0,
            Target: 0,
            Position: 0,
            PageId:props.cardList[0].PageId,
        }
        let real = getRealDOM(e.target as HTMLElement)
        real.classList.add("dragEnterList")
        if (real.tagName === "LI" && real.classList.contains("ant-list-item")) {
            real.classList.remove("dragEnterList")
            formdata.Source = parseInt(e.dataTransfer.getData("id"));
            formdata.Target = parseInt(findAttribute(real, "data-id"));
            let sourceType = e.dataTransfer.getData("type");
            if (e.clientY - 71 - real.offsetTop < real.clientHeight / 2) {
                formdata.Position = 1;
            }
            else {
                formdata.Position = 2;
            }
            props.dispatch({
                type: 'cardList/updateSeq',
                move: formdata,
            });
        }

    }
    function onDragOver(e: React.DragEvent) {
        let real = getRealDOM(e.target as HTMLElement);
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
    function onDragEnter(e: React.DragEvent) {
        let real = getRealDOM(e.target as HTMLElement);
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            real.classList.add("dragEnterList")
        }
        e.preventDefault()
    }
    function onDragLeave(e: React.DragEvent) {
        let real = e.target as HTMLElement;
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            real.classList.remove("dragEnterList")
        }
        e.preventDefault()
    }

    return <List
        dataSource={props.cardList}
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
export default connect((state: IStore) => {
    return {
        cardList: state.cardList,
    }
})(CardsList);