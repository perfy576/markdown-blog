import {Base64} from 'js-base64'
export function base64_encode(data) {
    return Base64.encode(data);
}

export function base64_decode(data) {
    return Base64.decode(data);
}

export function base64_url_encode(data) {
    return Base64.encodeURL(data);
}