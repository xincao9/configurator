<template>
    <div>
        <a-modal v-model="visible" title="配置属性" @ok="handleOk" ok-text="确认" cancel-text="取消">
            <a-textarea
                v-model="properties"
                placeholder="目前仅支持JSON格式"
                :auto-size="{ minRows: 10, maxRows: 50 }"
            />
        </a-modal>
        <a-table :columns="columns" :dataSource="data" :rowKey="getRowKey">
            <span slot="properties" slot-scope="record">
                <a @click="showProperties(record)">编辑</a>
                <a-divider type="vertical"/>
                <a @click="deleteApp(record)">删除</a>
            </span>
        </a-table>
    </div>
</template>

<script>
    "use strict"
    import resource from "resource-axios";
    import axios from "axios";

    const Apps = resource("/apps", axios);
    const App = resource("/app", {
        getPropertiesById: (id) => axios.get(`/app/${id}/properties`),
        savePropertiesById: (id, props) => axios.put(`/app/${id}/properties`, props),
    }, axios);

    const columns = [
        {
            title: '主键',
            dataIndex: 'id',
            key: 'id',
            ellipsis: true,
        },
        {
            title: '环境',
            dataIndex: 'env',
            key: 'env',
            ellipsis: true,
        },
        {
            title: '组',
            dataIndex: 'group',
            key: 'group',
            ellipsis: true,
        },
        {
            title: '项目',
            dataIndex: 'project',
            key: 'project',
            ellipsis: true,
        },
        {
            title: '版本',
            dataIndex: 'version',
            key: 'version',
            ellipsis: true,
        },
        {
            title: '属性',
            key: 'properties',
            ellipsis: true,
            scopedSlots: {customRender: 'properties'},
        }
    ];

    export default {
        mounted() {
            this.getApps();
        },
        data() {
            return {
                data: null,
                columns,
                visible: false,
                properties: null,
                appId: 0,
            };
        },
        methods: {
            getRowKey(record) {
                return record.id;
            },
            getApps() {
                let _this = this;
                Apps.get().then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.data = res.data.data;
                    }
                })
            },
            showProperties(app) {
                if (app == null) {
                    return
                }
                let _this = this;
                _this.appId = app.id;
                App.getPropertiesById(app.id).then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.properties = res.data.data;
                    }
                });
                this.visible = true;
            },
            handleOk() {
                let _this = this;
                if (this.properties != null) {
                    App.savePropertiesById(_this.appId, _this.properties).then(function (res) {
                        if (res.status == 200 && res.data.code == 200) {
                            _this.visible = false;
                            _this.properties = null;
                            _this.appId = 0;
                        }
                    });
                }
            },
            deleteApp(app) {
                if (app == null) {
                    return
                }
                let _this = this;
                _this.$confirm({
                    title: '提醒',
                    content: '确认删除吗?',
                    okText: '确认',
                    cancelText: '取消',
                    onOk: function () {
                        App.delete(app.id).then(function (res) {
                            if (res.status == 200 && res.data.code == 200) {
                                _this.getApps();
                            }
                        });
                    }
                });
            }
        },
        name: "PagesConfiguratorManagerList",
        components: {}
    };
</script>

<style>
</style>
