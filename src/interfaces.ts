


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
    FolderLocale: ILocale,
    Rows: any[],
}
export interface ICard {
    Id: number,
    PageId: number,
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
    IsVisible: boolean,
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
    Default: string,
    Filter: string,
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
    IsVisible: boolean,
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
    "Default": string,
    "Name"?: string,
}

export interface renderStateProps {

}