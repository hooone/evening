
import { ILocale } from '@/interfaces'
import { getLocale } from 'umi';

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
    let p1 = document.getElementById(id);
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