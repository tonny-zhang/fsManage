import axios from 'axios';
//axios 相关配置，注意：axios请求头的 Content-Type 默认是 application/json，而postman默认的是 application/x-www-form-urlencoded
var instance = axios.create({
    //请求根目录
	baseURL: 'http://192.168.31.134:4000',  
    // headers: {
    //     "Access-Control-Allow-Origin": "*",
    // }   
	// withCredentials : true,
	//timeout: 1000,
	// headers: {'X-Custom-Header': 'foobar'},
});
//------------------- 一、请求拦截器 忽略
instance.interceptors.request.use(function (config) {
    return config;
}, function (error) {
    // 对请求错误做些什么
    
    return Promise.reject(error);
});
 
//----------------- 二、响应拦截器 忽略
instance.interceptors.response.use(function (response) {
    
    return response.data;
}, function (error) {
    // 对响应错误做点什么
    console.log('拦截器报错');
    return Promise.reject(error);
});

const params = {};
location.search.substr(1).split('&').forEach(v => {
    let arr = v.split('=');
    params[arr[0]] = arr[1];
})
export const getList = function(e) {
    return instance.get('/api/list', {
        params: {
            dir: e
        }
    })
}
export const deleteFile = function(e) {
    return instance.post('/api/remove', {
        file: e
    })
}
export const PARAMS = params;