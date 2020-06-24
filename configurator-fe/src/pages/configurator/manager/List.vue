<template>
    <a-table :columns="columns" :dataSource="data">
        <a slot="name" slot-scope="text">{{text}}</a>
    </a-table>
</template>

<script>
    "use strict"
    import resource from "resource-axios";
    import axios from "axios";

    const Configurators = resource("/configurators", axios);
    const columns = [
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
            dataIndex: 'properties',
            key: 'properties',
            ellipsis: true,
        }
    ];

    export default {
        mounted() {
            this.getConfigurators();
        },
        data() {
            return {
                data: null,
                columns
            };
        },
        methods: {
            getConfigurators() {
                let _this = this;
                Configurators.get().then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.data = res.data.data;
                    }
                })
            }
        },
        name: "PagesConfiguratorManagerList",
        components: {}
    };
</script>

<style>
</style>
