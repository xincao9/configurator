<template>
    <a-table :columns="columns" :dataSource="data">
        <a slot="properties" slot-scope="record" @click="showProperties (record.id)">查看</a>
    </a-table>
</template>

<script>
    "use strict"
    import resource from "resource-axios";
    import axios from "axios";

    const Apps = resource("/apps", axios);
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
            };
        },
        methods: {
            getApps() {
                let _this = this;
                Apps.get().then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.data = res.data.data;
                    }
                })
            },
            showProperties(data) {
                console.log(data);
            },
        },
        name: "PagesConfiguratorManagerList",
        components: {}
    };
</script>

<style>
</style>
