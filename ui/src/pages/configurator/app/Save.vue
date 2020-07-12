<template>
    <div>
        <h1 v-if="form.id > 0">修改</h1>
        <h1 v-else>新建</h1>
        <hr/>
        <a-form-model ref="form" :model="form" :label-col="{ span: 4 }" :wrapper-col="{ span: 10 }" :rules="rules"
                      style="margin-top: 50px;">
            <a-form-model-item label="环境" prop="env">
                <a-select style="width: 120px" @change="envChange" v-model="defaultEnv">
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
    const Envs = resource('/envs', axios);
    const UserEnv = resource('/user_env', axios);

    export default {
        created() {
            let _this = this;
            let id = 0;
            if (_this.$route.query != null && _this.$route.query.id != null) {
                id = _this.$route.query.id;
            }
            if (id > 0) {
                App.get(id).then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.form.id = id;
                        _this.form.env = res.data.data.env;
                        _this.form.group = res.data.data.group;
                        _this.form.project = res.data.data.project;
                        _this.form.version = res.data.data.version;
                        _this.defaultEnv = _this.form.env;
                    }
                });
            }
            Envs.get().then(function (res1) {
                if (res1.status == 200 && res1.data.code == 200) {
                    _this.envs = {};
                    for (let i in res1.data.data) {
                        _this.envs[res1.data.data[i].id] = res1.data.data[i].name;
                    }
                    UserEnv.get().then(function (res2) {
                        if (res2.status == 200 && res2.data.code == 200) {
                            if (_this.userEnvs == null) {
                                _this.userEnvs = res2.data.data;
                            }
                        }
                    });
                }
            });
        },
        data() {
            return {
                defaultEnv: '点击选择',
                userEnvs: null,
                envs: null,
                form: {
                    id: 0,
                    env: '',
                    group: '',
                    project: '',
                    version: ''
                },
                rules: {
                    env: [{
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
            envChange(envId) {
                let _this = this;
                _this.form.env = _this.envs[envId];
            },
        },
        name: "PagesConfiguratorAppSave",
        components: {}
    };
</script>

<style>
</style>
