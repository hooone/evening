import { DefaultRootState } from 'react-redux';
import { modalStateProps } from './models/modal';
import { globalStateProps } from './models/global';
import { navStateProps } from './models/navigation';
import { pageStateProps } from './models/page';
import { drawStateProps } from './models/draw';
import { rectChartStateProps } from './models/rectChartConfig';
import { pointChartStateProps } from './models/pointChartConfig';
import { actionInfoStateProps } from './models/actionInfoConfig';
import { cardInfoStateProps } from './models/cardInfoConfig';
import { actionListStateProps } from './models/actionList';
import { fieldInfoStateProps } from './models/fieldInfoConfig';
import { fieldListStateProps } from './models/fieldList';
import { parameterStateProps } from './models/parameterInfoConfig';
import { contextMenuStateProps } from './models/contextMenu';
import { ICard } from './interfaces';

export interface IStore extends DefaultRootState {
    props: any,
    modal: modalStateProps,
    global: globalStateProps,
    nav: navStateProps,
    page: pageStateProps,
    draw: drawStateProps,
    contextMenu: contextMenuStateProps,
    cardList: ICard[],
    fieldInfoConfig: fieldInfoStateProps,
    fieldList: fieldListStateProps,
    actionList: actionListStateProps,
    actionInfoConfig: actionInfoStateProps,
    parameterInfoConfig: parameterStateProps,
    cardInfoConfig: cardInfoStateProps,
    rectChartConfig: rectChartStateProps,
    pointChartConfig: pointChartStateProps,
}