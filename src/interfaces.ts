


export interface CommonResult {
    Success: boolean,
    Data: any,
    Message: string,
}

export interface IFolder {
    Id: number,
    Name: string,
    Text: string,
    IsFolder: boolean,
    Locale: ILocale,
    Pages: IPage[],
}
export interface IPage {
    Id: number,
    Name: string,
    Text: string,
    Cards: ICard[],
    Locale: ILocale,
}
export interface ICard {
    Id: number,
    Name: string,
    Text: string,
    Width: number,
    Pos: number,
    Style: string,
    Actions: IViewAction[],
    Reader: IViewAction,
    Locale: ILocale,
    Fields: IField[],
    data: any[],
    Styles: IStyle[],
}

export interface IModal {
    visible: boolean,
    action: IViewAction,
    record: IRecord,
    confirmLoading: boolean,
}
export interface IField {
    Id: number,
    CardId: number,
    Name: string,
    Text: string,
    Type: string,
    IsVisible: boolean,
    Locale: ILocale,
}
export interface IStyle {
    Id: number,
    CardId: number,
    Type: string,

    Property: string,
    FieldId: number,
    Field: IField,
    Value: string,
}
export interface IViewAction {
    Id: number,
    CardId: number,
    Name: string,
    Text: string,
    Type: string,
    URL: string,
    DoubleCheck: boolean,
    Locale: ILocale,
    Parameters: IParameter[],
}
export interface IParameter {
    Id: number,
    IsEditable: boolean,
    IsVisible: boolean,
    Field: IField,
    Default: string,
    Compare: string,
}
export interface IRecord {
    __Key: string,
}
export interface IValueChange {
    type: string,
    id: number,
    name: string,
    value: string,
}
export interface IMove {
    Source: number,
    Target: number,
    Position: number,
}
export interface navMoveProps {

}

export interface ILocale {
    "zh-CN": string,
    "en-US": string,
    Default: string,
}

export interface drawStateProps {
    cardId: number,
    cardChanged: boolean,
    visible: boolean,
    subVisible: boolean,
    styleVisible: boolean,
    title: string,
    subTitle: string,
}
export interface contextMenuStateProps {
    visible: boolean,
    menu: string,
    left: number,
    top: number,
    record: any,
}
export interface actionListStateProps {
    cardId: number,
    actions: IViewAction[],
}
export interface renderStateProps {

}
export interface fieldListStateProps {
    cardId: number,
    fields: IField[],
}
export interface fieldInfoStateProps extends IField {
    visible: boolean,
    dirty: boolean,
}
export interface parameterStateProps {
    cardId: number,
    visible: boolean,
    dirty: boolean,
    parameters: IParameter[],
}