
import { message } from 'antd';
import { getLocale } from 'umi';
import reqwest from 'reqwest'

export interface CommonResult {
    Success: boolean,
    Data: any,
    Message: string,
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