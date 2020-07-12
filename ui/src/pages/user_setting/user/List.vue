<template>
    <div>
        <a-table :locale="{emptyText: '暂无数据'}" :columns="columns" :dataSource="data" :rowKey="getRowKey">
            <span slot="role" slot-scope="user">
                {{ roles[user.role] }}
            </span>
            <span slot="envs" slot-scope="user">
                <a-avatar v-for="env in user.envs" v-bind:key="env">
                    {{ env }}
                </a-avatar>
            </span>
            <span slot="option" slot-scope="user">
                <a @click="edit(user)">编辑</a>
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
            title: '环境',
            key: 'envs',
            ellipsis: true,
            scopedSlots: {customRender: 'envs'},
        },
        {
            title: '操作',
            key: 'option',
            ellipsis: true,
            scopedSlots: {customRender: 'option'},
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
            edit(user) {
                let _this = this;
                if (user == null) {
                    return;
                }
                let id = user.id;
                _this.$router.push({
                    path: "/pages/user_setting/user/save",
                    query: {
                        "id": id
                    }
                });
                _this.$emit('basicsync');
            }
        },
        name: "PagesUserSettingUserList"
    }
</script>
<style>
</style>
