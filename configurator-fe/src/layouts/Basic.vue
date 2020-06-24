<template>
	<a-layout id="layout_basic">
		<a-layout-header class="header">
			<div class="logo" />
			<a-menu theme="dark" mode="horizontal" :defaultSelectedKeys="[selectedKeys]" :style="{ lineHeight: '64px' }" @click="click">
				<a-menu-item key="ticket_system">工单系统</a-menu-item>
				<a-menu-item key="cluster_management">集群管理</a-menu-item>
				<a-menu-item key="user_setting">用户设置</a-menu-item>
			</a-menu>
		</a-layout-header>
		<a-layout>
			<a-layout-sider width="200" style="background: #fff">
				<a-menu v-if="selectedKeys == 'ticket_system'" mode="inline" :defaultSelectedKeys="[siderSelectedKeys]" :selectedKeys="[siderSelectedKeys]" :defaultOpenKeys="[openKeys]" :style="{ height: '100%', borderRight: 0 }" @click="siderClick">
					<a-sub-menu key="my_ticket">
						<span slot="title">
							<a-icon type="file" />我的工单</span>
						<a-menu-item key=":pages:ticket_system:my_ticket:list">列表</a-menu-item>
						<a-menu-item key=":pages:ticket_system:my_ticket:new">新建</a-menu-item>
					</a-sub-menu>
				</a-menu>
				<a-menu v-if="selectedKeys == 'cluster_management'" mode="inline" :defaultSelectedKeys="[siderSelectedKeys]" :selectedKeys="[siderSelectedKeys]" :defaultOpenKeys="[openKeys]" :style="{ height: '100%', borderRight: 0 }" @click="siderClick">
					<a-sub-menu key="segment_management">
						<span slot="title">
							<a-icon type="cloud" />segment管理</span>
						<a-menu-item key=":pages:cluster_management:segment_management:list">列表</a-menu-item>
					</a-sub-menu>
				</a-menu>
				<a-menu v-if="selectedKeys == 'user_setting'" mode="inline" :defaultSelectedKeys="[siderSelectedKeys]" :selectedKeys="[siderSelectedKeys]" :defaultOpenKeys="[openKeys]" :style="{ height: '100%', borderRight: 0 }" @click="siderClick">
					<a-sub-menu key="notification_management">
						<span slot="title">
							<a-icon type="notification" />通知管理</span>
						<a-menu-item key=":pages:user_setting:notification_management:message_box">消息箱</a-menu-item>
					</a-sub-menu>
					<a-sub-menu key="account_management">
						<span slot="title">
							<a-icon type="user" />账号管理</span>
						<a-menu-item key=":pages:user_setting:account_management:list">列表</a-menu-item>
						<a-menu-item key=":pages:user_setting:account_management:update">修改</a-menu-item>
					</a-sub-menu>
					<a-sub-menu key="role_management">
						<span slot="title">
							<a-icon type="setting" />角色管理</span>
						<a-menu-item key=":pages:user_setting:role_management:list">列表</a-menu-item>
						<a-menu-item key=":pages:user_setting:role_management:new">新建</a-menu-item>
						<a-menu-item key=":pages:user_setting:role_management:update">修改</a-menu-item>
					</a-sub-menu>
				</a-menu>
			</a-layout-sider>
			<a-layout style="padding: 0 24px 24px">
				<a-breadcrumb style="margin: 16px 0">
					<a-breadcrumb-item>{{ dict[selectedKeys] }}</a-breadcrumb-item>
					<a-breadcrumb-item>{{ dict[openKeys] }}</a-breadcrumb-item>
					<a-breadcrumb-item>{{ dict[siderSelectedKeys] }}</a-breadcrumb-item>
				</a-breadcrumb>
				<a-layout-content :style="{ background: '#fff', padding: '24px', margin: 0, minHeight: '280px' }">
					<router-view @basicsync="onSync"></router-view>
				</a-layout-content>
			</a-layout>
		</a-layout>
	</a-layout>
</template>
<script>
	const dict = {
		['user_setting']: '用户设置',
		['account_management']: '账号管理',
		[':pages:user_setting:account_management:list']: '列表',
		[':pages:user_setting:account_management:update']: '修改',
		['role_management']: '角色管理',
		[':pages:user_setting:role_management:list']: '列表',
		[':pages:user_setting:role_management:new']: '新建',
		[':pages:user_setting:role_management:update']: '修改',
		['notification_management']: '通知管理',
		[':pages:user_setting:notification_management:message_box']: '消息箱',
		['cluster_management']: '集群管理',
		['segment_management']: 'segment管理',
		[':pages:cluster_management:segment_management:list']: '列表',
		['ticket_system']: '工单系统',
		['my_ticket']: '我的工单',
		[':pages:ticket_system:my_ticket:list']: '列表',
		[':pages:ticket_system:my_ticket:new']: '新建',
	}
	export default {
		created() {
			this.onSync();
		},
		data() {
			return {
				selectedKeys: null,
				openKeys: null,
				siderSelectedKeys: null,
				dict
			};
		},
		methods: {
			click: function(event) {
				let key = event.key;
				if (key == 'user_setting') {
					this.selectedKeys = 'user_setting';
					this.openKeys = 'notification_management';
					this.siderSelectedKeys = ':pages:user_setting:notification_management:message_box';
					this.$router.push({
						path: '/pages/user_setting/notification_management/message_box'
					});
				} else if (key == 'cluster_management') {
					this.selectedKeys = 'cluster_management';
					this.openKeys = 'segment_management';
					this.siderSelectedKeys = ':pages:cluster_management:segment_management:list';
					this.$router.push({
						path: '/pages/cluster_management/segment_management/list'
					});
				} else if (key == 'ticket_system') {
					this.selectedKeys = 'ticket_system';
					this.openKeys = 'my_ticket';
					this.siderSelectedKeys = ':pages:ticket_system:my_ticket:list';
					this.$router.push({
						path: '/pages/ticket_system/my_ticket/list'
					});
				}
			},
			siderClick: function(event) {
				let key = event.key;
				if (key == null || key.length <= 0) {
					return;
				}
				let ss = key.split(':');
				if (ss != null && ss.length > 2) {
					this.selectedKeys = ss[2];
					this.openKeys = ss[3];
					this.siderSelectedKeys = key;
				}
				let newUrl = ss.join('/');
				if (newUrl == this.$route.path) {
					return;
				}
				this.$router.push({
					path: ss.join('/')
				});
			},
			onSync: function() {
				let path = this.$route.path;
				if (path != null && path.length > 0) {
					let ss = path.split('/');
					if (ss != null && ss.length > 2) {
						this.selectedKeys = ss[2];
						this.openKeys = ss[3];
						this.siderSelectedKeys = ss.join(':');
					}
				}
			}
		}
	};
</script>

<style>
	#layout_basic .logo {
		width: 120px;
		height: 31px;
		background: rgba(255, 255, 255, 0.2);
		margin: 16px 28px 16px 0;
		float: left;
	}
</style>
