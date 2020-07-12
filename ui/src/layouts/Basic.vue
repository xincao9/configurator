<template>
    <a-layout id="layout_basic">
        <a-layout-header class="header">
            <a-row>
                <a-col :span="6">
                    <div class="logo"/>
                </a-col>
                <a-col :span="12">
                    <a-menu theme="dark" mode="horizontal" :defaultSelectedKeys="[selectedKeys]"
                            :style="{ lineHeight: '64px' }"
                            @click="click">
                        <a-menu-item key="configurator">配置中心</a-menu-item>
                        <a-menu-item key="user_setting">用户设置</a-menu-item>
                    </a-menu>
                </a-col>
                <a-col :span="4">
                </a-col>
                <a-col :span="1">
                    <a-avatar>
                        {{ user.username }}
                    </a-avatar>
                </a-col>
                <a-col :span="1">
                    <a @click="deleteSession(user.id)"><strong>注销</strong></a>
                </a-col>
            </a-row>
        </a-layout-header>
        <a-layout>
            <a-layout-sider width="200" style="background: #fff">
                <a-menu v-if="selectedKeys == 'configurator'" mode="inline" :defaultSelectedKeys="[siderSelectedKeys]"
                        :selectedKeys="[siderSelectedKeys]"
                        :defaultOpenKeys="[openKeys]" :style="{ height: '100%', borderRight: 0 }" @click="siderClick">
                    <a-sub-menu key="app">
						<span slot="title">
							<a-icon type="file"/>应用</span>
                        <a-menu-item key=":pages:configurator:app:list">
                            <a-icon type="unordered-list"/>
                            列表
                        </a-menu-item>
                        <a-menu-item key=":pages:configurator:app:save">
                            <a-icon type="save"/>
                            保存
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
                    <a-sub-menu key="user" v-if="user != null && user.role != 1">
						<span slot="title">
							<a-icon type="user"/>账号管理</span>
                        <a-menu-item key=":pages:user_setting:user:list">
                            <a-icon type="unordered-list"/>
                            列表
                        </a-menu-item>
                        <a-menu-item key=":pages:user_setting:user:save">
                            <a-icon type="save"/>
                            保存
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
                <a-layout-footer style="text-align: center">
                    configurator ©2020 开源 <a href="https://github.com/xincao9/configurator">github.com/xincao9/configurator</a>
                </a-layout-footer>
            </a-layout>
        </a-layout>
    </a-layout>
</template>
<script>
    "use strict"
    import resource from "resource-axios";
    import axios from "axios";

    const User = resource("/user", axios);
    const Session = resource("/session", axios);

    const dict = {
        ['configurator']: '配置中心',
        ['app']: '应用',
        [':pages:configurator:app:list']: '列表',
        [':pages:configurator:app:save']: '保存',
        ['user_setting']: "用户设置",
        ['notification']: "通知管理",
        [':pages:user_setting:notification:message_box']: "消息箱",
        [':pages:user_setting:notification:operation_log']: "操作日志",
        ['user']: "账号管理",
        [':pages:user_setting:user:list']: "列表",
        [':pages:user_setting:user:save']: "保存",
    }
    export default {
        created() {
            let _this = this;
            this.onSync();
            User.get().then(function (res) {
                if (res.status == 200 && res.data.code == 200) {
                    _this.user = res.data.data;
                }
            });
        },
        data() {
            return {
                user: null,
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
            },
            deleteSession: function (userId) {
                if (userId <= 0) {
                    return;
                }
                let _this = this;
                _this.$confirm({
                    title: '提醒',
                    content: '确认注销吗?',
                    okText: '确认',
                    cancelText: '取消',
                    onOk: function () {
                        Session.delete(userId).then(function (res) {
                            if (res.status == 200 && res.data.code == 200) {
                                _this.$router.push({
                                    path: '/'
                                });
                            }
                        });
                    }
                });
            }
        }
    };
</script>

<style>
    #layout_basic .logo {
        width: 240px;
        height: 31px;
        background: rgba(255, 255, 255, 0.2);
        margin: 16px 28px 16px 0;
        float: left;
    }
</style>
