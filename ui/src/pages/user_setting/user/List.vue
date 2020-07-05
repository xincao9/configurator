<template>
    <div>
        <a-table :columns="columns" :dataSource="data" :rowKey="getRowKey">
            <span slot="role" slot-scope="record">
               {{ roles[record.role] }}
            </span>
        </a-table>
    </div>
</template>

<script type="text/javascript">
    "use strict"
    import resource from "resource-axios";
    import axios from "axios";

    const Users = resource("/users", axios);
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
            title: '角色',
            key: 'role',
            ellipsis: true,
            scopedSlots: {customRender: 'role'},
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
            this.getUsers();
        },
        data() {
            let roles = {
                1: "普通",
                2: "经理",
                3: "管理",
            };
            return {
                roles,
                data: null,
                columns,
            };
        },
        methods: {
            getRowKey(record) {
                return record.id;
            },
            getUsers() {
                let _this = this;
                Users.get().then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.data = res.data.data;
                    }
                })
            },
        },
        name: "PagesUserSettingUserList"
    }
</script>
<style>
</style>
