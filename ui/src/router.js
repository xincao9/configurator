"use strict"

import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from './pages/Login.vue';
import PagesConfiguratorAppList from './pages/configurator/app/List.vue';
import PagesConfiguratorAppSave from './pages/configurator/app/Save.vue';
import PagesUserSettingUserList from './pages/user_setting/user/List.vue';
import PagesUserSettingUserSave from './pages/user_setting/user/Save';
import PagesUserSettingNotificationMessageBox from './pages/user_setting/notification/MessageBox.vue';
import PagesUserSettingNotificationOperationLog from './pages/user_setting/notification/OperationLog.vue';

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
            path: '/pages/configurator/app/list',
            component: PagesConfiguratorAppList
        }, {
            path: '/pages/configurator/app/save',
            component: PagesConfiguratorAppSave
        }, {
            path: '/pages/user_setting/user/list',
            component: PagesUserSettingUserList
        }, {
            path: '/pages/user_setting/user/save',
            component: PagesUserSettingUserSave
        }, {
            path: '/pages/user_setting/notification/message_box',
            component: PagesUserSettingNotificationMessageBox
        }, {
            path: '/pages/user_setting/notification/operation_log',
            component: PagesUserSettingNotificationOperationLog
        }
        ]
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
