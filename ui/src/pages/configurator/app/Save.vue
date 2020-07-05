<template>
    <div>
        <h1>新建</h1>
        <hr/>
        <a-form-model ref="form" :model="form" :label-col="{ span: 4 }" :wrapper-col="{ span: 10 }" :rules="rules"
                      style="margin-top: 50px;">
            <a-form-model-item label="环境" prop="env_id">
                <a-select style="width: 120px" @change="envIdChange">
                    <a-select-option v-for="ue in userEnvs" v-bind:key="ue.env_id" :value="ue.env_id">
                        {{ envs[ue.env_id] }}
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item label="组" prop="group">
                <a-input v-model="form.group" placeholder="例如: BASE"/>
            </a-form-model-item>
            <a-form-model-item label="项目" prop="project">
                <a-input v-model="form.project" placeholder="例如: USER-SERVICE"/>
            </a-form-model-item>
            <a-form-model-item label="版本" prop="version">
                <a-input v-model="form.version" placeholder="v1.0"/>
            </a-form-model-item>
            <a-form-model-item :wrapper-col="{ offset: 8 }">
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

    const App = resource('/app', axios);
    const UserEnv = resource('/user_env', axios);
    const Envs = resource('/envs', axios);

    export default {
        mounted() {
            let _this = this;
            UserEnv.get().then(function (res) {
                if (res.status == 200 && res.data.code == 200) {
                    _this.userEnvs = res.data.data;
                }
            });
            Envs.get().then(function (res) {
                if (res.status == 200 && res.data.code == 200) {
                    _this.envs = {};
                    for (let i in res.data.data) {
                        _this.envs[res.data.data[i].id] = res.data.data[i].name;
                    }
                }
            });
        },
        data() {
            return {
                envs: null,
                userEnvs: null,
                form: {
                    env_id: 0,
                    group: '',
                    project: '',
                    version: ''
                },
                rules: {
                    env_id: [{
                        required: true,
                        message: '运行环境'
                    }],
                    group: [{
                        required: true,
                        message: '业务组'
                    }],
                    project: [{
                        required: true,
                        message: '项目名称'
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
                _this.$refs['form'].validate(
                    valid => {
                        if (valid) {
                            App.post(_this.form).then(function (res) {
                                if (res.status == 200 && res.data.code == 200) {
                                    _this.$router.push({
                                        path: "/pages/configurator/app/list"
                                    });
                                    _this.$emit('basicsync');
                                }
                            });
                        }
                    }
                );
            },
            envIdChange(envId) {
                this.form.env_id = parseInt(envId);
            },
        },
        name: "PagesConfiguratorAppSave",
        components: {}
    };
</script>

<style>
</style>
