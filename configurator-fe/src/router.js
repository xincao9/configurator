"use strict"

import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from './pages/Login.vue';
import PagesUserSettingAccountManagementList from './pages/user_setting/account_management/List.vue';
import PagesUserSettingAccountManagementUpdate from './pages/user_setting/account_management/Update.vue';
import PagesUserSettingRoleManagementList from './pages/user_setting/role_management/List.vue';
import PagesUserSettingRoleManagementNew from './pages/user_setting/role_management/New.vue'
import PagesUserSettingRoleManagementUpdate from './pages/user_setting/role_management/Update.vue';
import PagesUserSettingNotificationManagementMessageBoxfrom from './pages/user_setting/notification_management/MessageBox.vue';
import PagesTicketSystemMyTicketList from './pages/ticket_system/my_ticket/List.vue';
import PagesTicketSystemMyTicketNew from './pages/ticket_system/my_ticket/New.vue';
import PagesClusterManagementSegmentManagementList from './pages/cluster_management/segment_management/List.vue';

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
				path: 'user_setting/account_management/list',
				component: PagesUserSettingAccountManagementList
			},
			{
				path: 'user_setting/account_management/update',
				component: PagesUserSettingAccountManagementUpdate
			}, {
				path: 'user_setting/role_management/list',
				component: PagesUserSettingRoleManagementList
			},
			{
				path: 'user_setting/role_management/new',
				component: PagesUserSettingRoleManagementNew
			},
			{
				path: 'user_setting/role_management/update',
				component: PagesUserSettingRoleManagementUpdate
			},
			{
				path: 'user_setting/notification_management/message_box',
				component: PagesUserSettingNotificationManagementMessageBoxfrom
			},
			{
				path: 'ticket_system/my_ticket/list',
				component: PagesTicketSystemMyTicketList
			},
			{
				path: 'ticket_system/my_ticket/new',
				component: PagesTicketSystemMyTicketNew
			},
			{
				path: 'cluster_management/segment_management/list',
				component: PagesClusterManagementSegmentManagementList
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
