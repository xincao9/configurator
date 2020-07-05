<template>
    <div>
        <h1>新建</h1>
        <hr/>
        <a-form-model ref="form" :model="form" :label-col="{ span: 4 }" :wrapper-col="{ span: 10 }" :rules="rules"
                      style="margin-top: 50px;">
            <a-form-model-item label="用户" prop="username">
                <a-input v-model="form.username" placeholder="xincao9"/>
            </a-form-model-item>
            <a-form-model-item label="角色" prop="role">
                <a-select style="width: 120px" default-value="1" @change="roleChange">
                    <a-select-option value="1">
                        普通
                    </a-select-option>
                    <a-select-option value="2">
                        经理
                    </a-select-option>
                    <a-select-option value="3">
                        管理
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item label="环境">
                <a-checkbox-group @change="envChange">
                    <a-row>
                        <span v-for="env in envs" v-bind:key="env.id">
                            <a-checkbox :value="env.id">
                                {{ env.name }}
                            </a-checkbox>
                        </span>
                    </a-row>
                </a-checkbox-group>
            </a-form-model-item>
            <a-form-model-item label="密码" prop="password">
                <a-input v-model="form.password" placeholder="password" type="password"/>
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

    const User = resource('/user', axios);
    const Envs = resource('/envs', axios);
    const UserEnv = resource('/user_env', axios);

    export default {
        mounted() {
            this.getEnvs();
        },
        data() {
            return {
                form: {
                    username: '',
                    password: '',
                    role: 1,
                },
                rules: {
                    username: [{
                        required: true,
                        message: '用户'
                    }],
                    role: [{
                        required: true,
                        message: '角色'
                    }],
                    password: [{
                        required: true,
                        message: '密码'
                    }],
                },
                envs: null,
                changeEnvs: null,
            }
        },
        methods: {
            getEnvs() {
                let _this = this;
                Envs.get().then(function (res) {
                    if (res.status == 200 && res.data.code == 200) {
                        _this.envs = res.data.data;
                    }
                });
            },
            onSubmit(e) {
                e.preventDefault();
                let _this = this;
                _this.$refs['form'].validate(
                    valid => {
                        if (valid) {
                            User.post(_this.form).then(function (res) {
                                if (res.status == 200 && res.data.code == 200) {
                                    for (let id in _this.changeEnvs) {
                                        UserEnv.post({
                                            'user_id': res.data.data.id,
                                            'env_id': parseInt(_this.changeEnvs[id])
                                        }).then(function (re) {
                                            if (re.status == 200 && re.data.code == 200) {
                                                console.log(re)
                                            } else {
                                                console.log(re)
                                            }
                                        });
                                    }
                                    _this.$router.push({
                                        path: "/pages/user_setting/user/list"
                                    });
                                    _this.$emit('basicsync');
                                }
                            });
                        }
                    }
                );
            },
            roleChange(role) {
                this.form.role = parseInt(role);
            },
            envChange(envs) {
                this.changeEnvs = envs;
                console.log(this.changeEnvs);
            },
        },
        name: "PagesUserSettingUserSave",
        components: {}
    };
</script>

<style type="text/css">
</style>
