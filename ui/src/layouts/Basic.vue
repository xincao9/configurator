<template>
    <a-layout id="layout_basic">
        <a-layout-header class="header">
            <div class="logo"/>
            <a-menu theme="dark" mode="horizontal" :defaultSelectedKeys="[selectedKeys]" :style="{ lineHeight: '64px' }"
                    @click="click">
                <a-menu-item key="configurator">配置中心</a-menu-item>
                <a-menu-item key="user_setting">用户设置</a-menu-item>
            </a-menu>
        </a-layout-header>
        <a-layout>
            <a-layout-sider width="200" style="background: #fff">
                <a-menu v-if="selectedKeys == 'configurator'" mode="inline" :defaultSelectedKeys="[siderSelectedKeys]"
                        :selectedKeys="[siderSelectedKeys]"
                        :defaultOpenKeys="[openKeys]" :style="{ height: '100%', borderRight: 0 }" @click="siderClick">
                    <a-sub-menu key="manager">
						<span slot="title">
							<a-icon type="file"/>应用</span>
                        <a-menu-item key=":pages:configurator:app:list">
                            <a-icon type="unordered-list"/>
                            列表
                        </a-menu-item>
                        <a-menu-item key=":pages:configurator:app:save">
                            <a-icon type="save"/>
                            新建
                        </a-menu-item>
                    </a-sub-menu>
                </a-menu>
                <a-menu v-if="selectedKeys == 'user_setting'" mode="inline" :defaultSelectedKeys="[siderSelectedKeys]"
                        :selectedKeys="[siderSelectedKeys]"
                        :defaultOpenKeys="[openKeys]" :style="{ height: '100%', borderRight: 0 }" @click="siderClick">
                    <a-sub-menu key="notification">
						<span slot="title">
							<a-icon type="message"/>通知管理</span>
                        <a-menu-item key=":pages:user_setting:notification:message_box">
                            <a-icon type="inbox"/>
                            消息箱
                        </a-menu-item>
                        <a-menu-item key=":pages:user_setting:notification:operation_log">
                            <a-icon type="unordered-list"/>
                            操作日志
                        </a-menu-item>
                    </a-sub-menu>
                    <a-sub-menu key="user">
						<span slot="title">
							<a-icon type="user"/>账号管理</span>
                        <a-menu-item key=":pages:user_setting:user:list">
                            <a-icon type="unordered-list"/>
                            列表
                        </a-menu-item>
                        <a-menu-item key=":pages:user_setting:user:save">
                            <a-icon type="save"/>
                            新建
                        </a-menu-item>
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
        ['configurator']: '配置中心',
        ['app']: '应用',
        [':pages:configurator:app:list']: '列表',
        [':pages:configurator:app:save']: '新建',
        ['user_setting']: "用户设置",
        ['notification']: "通知管理",
        [':pages:user_setting:notification:message_box']: "消息箱",
        [':pages:user_setting:notification:operation_log']: "操作日志",
        ['user']: "账号管理",
        [':pages:user_setting:user:list']: "列表",
        [':pages:user_setting:user:save']: "新建",
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
            click: function (event) {
                let key = event.key;
                if (key == 'configurator') {
                    this.selectedKeys = 'configurator';
                    this.openKeys = 'app';
                    this.siderSelectedKeys = ':pages:configurator:app:list';
                    this.$router.push({
                        path: '/pages/configurator/app/list'
                    });
                } else if (key == 'user_setting') {
                    this.selectedKeys = 'user_setting';
                    this.openKeys = 'notification';
                    this.siderSelectedKeys = ':pages:user_setting:notification:message_box';
                    this.$router.push({
                        path: '/pages/user_setting/notification/message_box'
                    });
                }
            },
            siderClick: function (event) {
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
            onSync: function () {
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
