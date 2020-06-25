<template>
    <a-layout id="layout_basic">
        <a-layout-header class="header">
            <div class="logo"/>
            <a-menu theme="dark" mode="horizontal" :defaultSelectedKeys="[selectedKeys]" :style="{ lineHeight: '64px' }"
                    @click="click">
                <a-menu-item key="configurator">配置</a-menu-item>
            </a-menu>
        </a-layout-header>
        <a-layout>
            <a-layout-sider width="200" style="background: #fff">
                <a-menu v-if="selectedKeys == 'configurator'" mode="inline" :defaultSelectedKeys="[siderSelectedKeys]"
                        :selectedKeys="[siderSelectedKeys]"
                        :defaultOpenKeys="[openKeys]" :style="{ height: '100%', borderRight: 0 }" @click="siderClick">
                    <a-sub-menu key="manager">
						<span slot="title">
							<a-icon type="user"/>管理</span>
                        <a-menu-item key=":pages:configurator:manager:list"><a-icon type="unordered-list"/>列表</a-menu-item>
						<a-menu-item key=":pages:configurator:manager:save"><a-icon type="save"/>保存</a-menu-item>
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
        ['configurator']: '配置',
        ['manager']: '管理',
        [':pages:configurator:manager:list']: '列表',
		[':pages:configurator:manager:save']: '保存',
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
                    this.openKeys = 'manager';
                    this.siderSelectedKeys = ':pages:configurator:manager:list';
                    this.$router.push({
                        path: '/pages/configurator/manager/list'
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
