
import axios from 'axios';


const baseURL = '/'; // 替换为实际的 API 服务器地址

const instance = axios.create({
    baseURL,
    timeout: 10000, // 超时时间
});

export const post = (url, data) => {
    console.log("data:", JSON.stringify(data))

    const authority_password = localStorage.getItem('authority_password')
    if (authority_password) {
        data['authority_password'] = authority_password
    } else {
        data['authority_password'] = ""
    }
    return instance.post(url, JSON.stringify(data), {
        headers: {
            'Content-Type': 'application/json',
        }
    })
        .then(response => {
            return response.data;
        })
        .catch(error => {
            throw error;
        });
};