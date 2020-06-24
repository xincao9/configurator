"use strict"

import Vue from 'vue';
import Antd from 'ant-design-vue';
import App from './App.vue';
import 'ant-design-vue/dist/antd.css';
import router from './router';
import axios from 'axios';

Vue.config.productionTip = false;

Vue.use(Antd);

axios.defaults.timeout = 5000;
axios.defaults.baseURL = process.env.API_URL || 'http://localhost:8080';
axios.defaults.withCredentials = true;

axios.interceptors.response.use(function (response) {
    if (response.status != 200 || response.data.code == 401) {
        router.push({
            path: "/"
        });
        return Promise.reject(response);
    }
    return response;
}, function (error) {
    return Promise.reject(error);
});

new Vue({
    router,
    render: h => h(App)
}).$mount('#app');
