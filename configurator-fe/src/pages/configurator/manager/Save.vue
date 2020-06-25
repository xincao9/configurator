<template>
    <div>
        <h1>新建</h1>
        <hr/>
        <a-form-model :model="form" :label-col="{ span: 4 }" :wrapper-col="{ span: 14 }" :rules="rules"
                      style="margin-top: 50px;">
            <a-form-model-item label="环境" prop="label">
                <a-input v-model="form.env" placeholder="test"/>
            </a-form-model-item>
            <a-form-model-item label="组" prop="label">
                <a-input v-model="form.group" placeholder="cbs"/>
            </a-form-model-item>
            <a-form-model-item label="项目" prop="label">
                <a-input v-model="form.project" placeholder="user-service"/>
            </a-form-model-item>
            <a-form-model-item label="版本" prop="label">
                <a-input v-model="form.version" placeholder="1.0"/>
            </a-form-model-item>
            <a-form-model-item :wrapper-col="{ span: 14, offset: 4 }">
                <a-button type="primary" @click="onSubmit">
                    提交
                </a-button>
            </a-form-model-item>
        </a-form-model>
    </div>
</template>

<script type="text/javascript">
    "use strict"
    import resource from "resource-axios";
    import axios from "axios";

    const App = resource("/app", axios);

    export default {
        data() {
            return {
                form: {
                    env: '',
                    group: '',
                    project: '',
                    version: ''
                },
                rules: {
                    env: [{
                        required: true,
                        message: '环境!'
                    }],
                    group: [{
                        required: true,
                        message: '组'
                    }],
                    project: [{
                        required: true,
                        message: '项目'
                    }],
                    version: [{
                        required: true,
                        message: '版本'
                    }]
                }
            }
        },
        methods: {
            onSubmit(e) {
                e.preventDefault();
                let _this = this;
                App.post(_this.form).then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.$router.push({
                            path: "/pages/configurator/manager/list"
                        });
                        _this.$emit('basicsync');
                    }
                })
            },
        },
        name: "PagesConfiguratorManagerSave",
        components: {}
    };
</script>

<style>
</style>
