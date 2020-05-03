
import { ILocale, CommonResult } from '@/interfaces'
import { message } from 'antd';
import { getLocale } from 'umi';
import reqwest from 'reqwest'

export function getLocaleText(locale: ILocale): string {
    let lang = getLocale();
    if (lang === "zh-CN" && locale["zh-CN"] !== undefined && locale["zh-CN"] !== "") {
        return locale["zh-CN"]
    }
    else if (lang === "en-US" && locale["en-US"] !== undefined && locale["en-US"] !== "") {
        return locale["en-US"]
    }
    else {
        return locale.Default
    }
}
//用原生JS进行表单操作，根据元素id获取输入值
export function getInputValue(id: string): string {
    let result: string = '';
    let p1 = document.getElementsByClassName(id)[0];
    if (!p1) {
        return '';
    }
    if (p1.nodeName === 'INPUT') {
        let p1Input = p1 as HTMLInputElement
        result = p1Input.value;
    }
    else if (typeof (p1.getAttribute('data-value')) !== 'undefined' && p1.getAttribute('data-value') != null) {
        result = p1.getAttribute('data-value') ?? ''
    }
    else {
        p1.childNodes.forEach(
            p1Child => {
                let p2 = p1Child as (HTMLElement);
                if (p2.nodeName === 'INPUT') {
                    let p2Input = p2 as HTMLInputElement
                    result = p2Input.value;
                }
                else if (typeof (p2.getAttribute('data-value')) !== 'undefined' && p2.getAttribute('data-value') != null) {
                    result = p2.getAttribute('data-value') ?? ""
                }
                else {
                    p2.childNodes.forEach(
                        p2Child => {
                            let p3 = p2Child as (HTMLElement);
                            if (p3.nodeName === 'INPUT') {
                                let p3Input = p3 as HTMLInputElement
                                result = p3Input.value;
                            }
                            else if (typeof (p3.getAttribute('data-value')) !== 'undefined' && p3.getAttribute('data-value') != null) {
                                result = p3.getAttribute('data-value') ?? ""
                            }
                            else {
                                p3.childNodes.forEach(
                                    p3Child => {
                                        let p4 = p3Child as (HTMLElement);
                                        if (p4.nodeName === 'INPUT') {
                                            let p4Input = p4 as HTMLInputElement
                                            result = p4Input.value;
                                        }
                                        else if (typeof (p4.getAttribute('data-value')) !== 'undefined' && p4.getAttribute('data-value') != null) {
                                            result = p4.getAttribute('data-value') ?? ""
                                        }
                                    }
                                )
                            }
                        }
                    )
                }
            }
        )
    }
    return result;
}

//获取元素的attribute
export function findAttribute(ele: HTMLElement, attr: string): string {
    if (!ele)
        return '';
    while (!ele.getAttribute(attr)) {
        if (ele.id === 'root')
            return '';
        else {
            let parent = ele.parentElement;
            if (parent) {
                ele = parent
            }
            else {
                return '';
            }
        }
    }
    return ele.getAttribute(attr) ?? "";
}

function dateFormat(date: Date, fmt: string) {
    var o = {
        "M+": date.getMonth() + 1,
        "d+": date.getDate(),
        "h+": date.getHours(),
        "m+": date.getMinutes(),
        "s+": date.getSeconds(),
        "q+": Math.floor((date.getMonth() + 3) / 3),
        "S": date.getMilliseconds()
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}

export function Offset2Datetime(offset: any) {
    let timeOffset = parseInt(offset)
    if (typeof (timeOffset) !== 'undefined' && (!isNaN(timeOffset)) && (offset.split(':').length == 1) && (offset.split('/').length == 1)) {
        let now = new Date().getTime();
        let target = new Date(now - timeOffset);
        return dateFormat(target, "yyyy-MM-dd hh:mm:ss");
    }
    else if (typeof (offset) === "string") {
        let timeA = new Date(offset.replace("-", "/"));
        if (!isNaN(timeA.getTime())) {
            return dateFormat(timeA, "yyyy-MM-dd hh:mm:ss");
        }
        else {
            return dateFormat(new Date(), "yyyy-MM-dd hh:mm:ss");
        }
    }
    else if (typeof (offset) === "object") {
        if (typeof (offset.getTime) !== "undefined" && (!isNaN(offset.getTime())) && (offset.split(':').length == 1) && (offset.split('/').length == 1)) {
            return dateFormat(offset, "yyyy-MM-dd hh:mm:ss");
        }
        else {
            return dateFormat(new Date(), "yyyy-MM-dd hh:mm:ss");
        }
    }
    else {
        return dateFormat(new Date(), "yyyy-MM-dd hh:mm:ss");
    }
}

export function Datetime2Offset(time: any) {
    if (typeof (time) === "string") {
        time = time.replace("-", "/");
    }
    let timeformat = new Date(time);
    if (typeof (timeformat) !== "undefined") {
        return new Date().getTime() - timeformat.getTime();
    }
    else {
        return 0;
    }
}

export function Offset2Date(offset: any) {
    let timeOffset = parseInt(offset)
    if (typeof (timeOffset) !== 'undefined' && (!isNaN(timeOffset))) {
        let today = new Date()
        let now = new Date(today.getFullYear(), today.getMonth(), today.getDate()).getTime();
        let target = new Date(now - timeOffset);
        return dateFormat(target, "yyyy-MM-dd");
    }
    else if (typeof (offset) === "string") {
        let timeA = new Date(offset.replace("-", "/"));
        if (!isNaN(timeA.getTime())) {
            return dateFormat(timeA, "yyyy-MM-dd");
        }
        else {
            return dateFormat(new Date(), "yyyy-MM-dd");
        }
    }
    else if (typeof (offset) === "object") {
        if (typeof (offset.getTime) !== "undefined" && (!isNaN(offset.getTime()))) {
            return dateFormat(offset, "yyyy-MM-dd");
        }
        else {
            return dateFormat(new Date(), "yyyy-MM-dd");
        }
    }
    else {
        return dateFormat(new Date(), "yyyy-MM-dd");
    }
}

export function Date2Offset(time: any) {
    if (typeof (time) === "string") {
        time = time.replace("-", "/");
    }
    let timeformat = new Date(time);
    if (typeof (timeformat) !== "undefined") {
        let today = new Date()
        let now = new Date(today.getFullYear(), today.getMonth(), today.getDate()).getTime();
        let inputtime = new Date(timeformat.getFullYear(), timeformat.getMonth(), timeformat.getDate()).getTime();
        return now - inputtime;
    }
    else {
        return 0;
    }
}

export function AJAX(url: string, data: any) {
    let promise = reqwest({
        url: url,
        type: 'json',
        method: 'post',
        data: { lang: getLocale(), ...data },
    })
    promise.then(function (resp: CommonResult) {
        if (!resp.Success) {
            if (resp.Message) {
                message.error(resp.Message);
            }
            else {
                message.error('request fail');
            }
        }
        return resp;
    }, function (err: any) {
        message.error(err.responseText);
        return {
            Success: false,
            Data: [],
            Message: err.responseText
        };
    })
    return promise;
}