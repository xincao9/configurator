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
                    <a-select-option value="2" v-if="user != null && user.role == 3">
                        经理
                    </a-select-option>
                    <a-select-option value="3" v-if="user != null && user.role == 3">
                        管理
                    </a-select-option>
                </a-select>
            </a-form-model-item>
            <a-form-model-item label="环境">
                <a-checkbox-group @change="envChange">
                    <a-row>
                        <span v-for="userEnv in userEnvs" v-bind:key="userEnv.id">
                            <a-checkbox :value="userEnv.env_id">
                                {{ envs[userEnv.env_id] }}
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
        created() {
            let _this = this;
            User.get().then(function (res1) {
                if (res1.status == 200 && res1.data.code == 200) {
                    if (_this.user == null) {
                        _this.user = res1.data.data;
                    }
                    Envs.get().then(function (res2) {
                        if (res2.status == 200 && res2.data.code == 200) {
                            if (_this.envs == null) {
                                _this.envs = {};
                                for (let i in res2.data.data) {
                                    _this.envs[res2.data.data[i].id] = res2.data.data[i].name;
                                }
                            }
                            UserEnv.get().then(function (res3) {
                                if (res3.status == 200 && res3.data.code == 200) {
                                    if (_this.userEnvs == null) {
                                        _this.userEnvs = res3.data.data;
                                    }
                                }
                            });
                        }
                    });
                }
            });
        },
        data() {
            return {
                user: null,
                envs: null,
                userEnvs: null,
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
                changeEnvs: null,
            }
        },
        methods: {
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
            },
        },
        name: "PagesUserSettingUserSave",
        components: {}
    };
</script>

<style type="text/css">
</style>
