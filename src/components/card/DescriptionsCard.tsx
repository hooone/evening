import React, { ReactElement } from 'react';
import { Descriptions, Empty } from 'antd';
import { getLocaleText } from '@/util'
import { ICard, IStyle } from '@/interfaces';
import cardInfoConfig from '@/models/cardInfoConfig';

interface cardProps {
    card: ICard,
}
const DescriptionsCard = (props: cardProps) => {
    if (!props.card.data || props.card.data.length < 1) {
        return (<Empty />)
    }
    let Items: ReactElement[] = [];
    props.card.Fields.forEach(p => {
        if (p.IsVisible) {
            Items.push(<Descriptions.Item key={"desc_" + p.Id}
                label={getLocaleText(p.Locale)}>
                {props.card.data[0][p.Name]}
            </Descriptions.Item>)
        }
    })
    return (<Descriptions bordered size={'default'}>
        {Items}
    </Descriptions>)
};
export default DescriptionsCard;