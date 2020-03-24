import { DefaultRootState } from 'react-redux';
import { modalStateProps } from './models/modal';
import { globalStateProps } from './models/global';
import { navStateProps } from './models/navigation';
import { pageStateProps } from './models/page';
import { rectChartStateProps } from './models/rectChartConfig';
import { pointChartStateProps } from './models/pointChartConfig';
import { actionInfoStateProps } from './models/actionInfoConfig';
import { cardInfoStateProps } from 'umi';

export interface IStore extends DefaultRootState {
    props: any,
    modal: modalStateProps,
    global: globalStateProps,
    nav: navStateProps,
    page: pageStateProps,
    actionInfoConfig: actionInfoStateProps,
    cardInfoConfig: cardInfoStateProps,
    rectChartConfig: rectChartStateProps,
    pointChartConfig: pointChartStateProps,
}