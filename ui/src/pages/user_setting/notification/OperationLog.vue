<template>
    <div>
        <a-table :locale="{emptyText: '暂无数据'}" :columns="columns" :dataSource="data" :rowKey="getRowKey"/>
    </div>
</template>

<script type="text/javascript">
    "use strict"
    import resource from "resource-axios";
    import axios from "axios";

    const OperationLogs = resource("/operation_logs", axios);
    const columns = [
        {
            title: '主键',
            dataIndex: 'id',
            key: 'id',
            ellipsis: true,
        },
        {
            title: '用户',
            dataIndex: 'username',
            key: 'username',
            ellipsis: true,
        },
        {
            title: '消息',
            dataIndex: 'message',
            key: 'message',
            ellipsis: true,
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
            ellipsis: true,
        }
    ];

    export default {
        mounted() {
            this.getOperationLogs();
        },
        data() {
            return {
                data: null,
                columns,
            };
        },
        methods: {
            getRowKey(record) {
                return record.id;
            },
            getOperationLogs() {
                let _this = this;
                OperationLogs.get().then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.data = res.data.data;
                    }
                })
            },
        },
        name: "PagesUserSettingNotificationOperationLog"
    }
</script>
<style>
</style>
