import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Table, Popconfirm, Button, Row, Col, Layout } from 'antd';
import { connect, getDvaApp } from 'umi';
import { DispatchProp } from 'react-redux';
import DataCard from '@/components/DataCard';
import { Component } from 'react';
import { IPage, ICard } from '@/interfaces'
import { IStore } from '@/store'

const { Content, Header, Sider } = Layout;


interface PageProps extends DispatchProp {
    page: IPage,
}
interface cardCfg {
    Id: number,
    Pos: number,
    Width: number,
    Children: React.ReactElement[],
}
interface rowCfg {
    Id: number,
    Cols: cardCfg[],
}
interface rootCfg {
    Rows: rowCfg[]
}

const PageContent = (props: PageProps) => {
    let root: rootCfg = {
        Rows: [{
            Id: 0,
            Cols: []
        }]
    };
    props.page.Cards.forEach(
        card => {
            let cols: cardCfg[] = root.Rows[root.Rows.length - 1].Cols
            let oldCol = false;
            let newCol = false;
            cols.forEach(col => {
                if (col.Pos === card.Pos && col.Width === card.Width) {
                    oldCol = true
                    col.Id = card.Id
                    col.Children.push(<DataCard key={"card_" + card.Id}
                        cardInfo={card}
                    > </ DataCard>)
                }
            });
            if (!oldCol) {
                newCol = true
                cols.forEach(col => {
                    if ((col.Pos <= card.Pos && (col.Pos + col.Width) > card.Pos) ||
                        (col.Pos < (card.Pos + card.Width) && (col.Pos + col.Width) >= (card.Pos + card.Width)) ||
                        (col.Pos >= card.Pos && (col.Pos + col.Width) <= (card.Pos + card.Width))) {
                        newCol = false
                    }
                });
                if (newCol) {
                    cols.push({
                        Id: card.Id,
                        Pos: card.Pos,
                        Width: card.Width,
                        Children: [<DataCard key={"card_" + card.Id}
                            cardInfo={card}
                            data-contextmenu="treepage"
                        > </ DataCard>]
                    })
                } else {
                    root.Rows.push({
                        Id: card.Id,
                        Cols: [{
                            Id: card.Id,
                            Pos: card.Pos,
                            Width: card.Width,
                            Children: [<DataCard key={"card_" + card.Id}
                                cardInfo={card}
                                data-contextmenu="treepage"
                            > </ DataCard>]
                        }]
                    })
                }
            }
        }
    )
    let contents: React.ReactElement[] = []
    root.Rows.forEach(
        row => {
            let cols: React.ReactElement[] = []
            let offset = 0
            row.Cols.forEach(col => {
                cols.push(<Col key={"col_" + col.Id} span={col.Width} offset={col.Pos - offset}> {col.Children}</Col>)
                offset = col.Pos + col.Width
            });
            contents.push(<Row key={"row_" + row.Id} >{cols}</Row>)
        }
    )
    return <Content style={{ margin: '0 16px', paddingBottom: '100px' }}
        data-contextmenu="content" data-pageid={props.page.Id} data-pagename={props.page.Name} >
        {contents}
    </Content>
};

export default connect((state: IStore) => {
    return {
        page: state.page,
    }
})(PageContent);