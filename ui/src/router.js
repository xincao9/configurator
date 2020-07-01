"use strict"

import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from './pages/Login.vue';
import PagesConfiguratorManagerList from './pages/configurator/manager/List.vue';
import PagesConfiguratorManagerSave from './pages/configurator/manager/Save.vue';
import PagesUserSettingAccountList from './pages/user_setting/account/List.vue';
import PagesUserSettingAccountSave from './pages/user_setting/account/Save';
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
            path: '/pages/configurator/manager/list',
            component: PagesConfiguratorManagerList
        }, {
            path: '/pages/configurator/manager/save',
            component: PagesConfiguratorManagerSave
        }, {
            path: '/pages/user_setting/account/list',
            component: PagesUserSettingAccountList
        }, {
            path: '/pages/user_setting/account/save',
            component: PagesUserSettingAccountSave
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
