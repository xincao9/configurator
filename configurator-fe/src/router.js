"use strict"

import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from './pages/Login.vue';
import PagesConfiguratorManagerList from './pages/configurator/manager/List.vue';

import Basic from './layouts/Basic.vue';

Vue.use(VueRouter);

const routes = [{
    path: '/login',
    component: Login
},
    {
        path: '/pages',
        component: Basic,
        children: [{
            path: '/pages/configurator/manager/list',
            component: PagesConfiguratorManagerList
        }]
    },
    {
        path: '/',
        redirect: '/login'
    }
];

const router = new VueRouter({
    routes
});

export default router;