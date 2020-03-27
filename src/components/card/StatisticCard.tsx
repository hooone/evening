import React, { ReactElement } from 'react';
import { Statistic, Empty, Row, Col } from 'antd';
import { getLocaleText } from '@/util'
import { ICard, IStyle } from '@/interfaces';
import cardInfoConfig from '@/models/cardInfoConfig';
import { CardContentProps } from '@/models/card';

const StatisticCard = (props: CardContentProps) => {
    if (!props.card.data || props.card.data.length < 1) {
        return (<Empty />)
    }
    let Items: ReactElement[] = [];
    let fields = props.card.Fields.filter(f => f.IsVisible)
    fields.forEach(p => {
        Items.push(<Col key={"stat_" + p.Id} span={Math.max(3, Math.floor(24 / fields.length))}>
            <Statistic title={getLocaleText(p.Locale)}
                value={props.card.data[0][p.Name]} />
        </Col>)
    })
    return (<Row gutter={16}>
        {Items}
    </Row>)
};
export default StatisticCard;