<template>
    <div>
        <h1>新建</h1>
        <hr/>
        <a-form-model ref="form" :model="form" :label-col="{ span: 4 }" :wrapper-col="{ span: 10 }" :rules="rules"
                      style="margin-top: 50px;">
            <a-form-model-item label="用户" prop="username">
                <a-input v-model="form.username" placeholder="xincao9"/>
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

    const Account = resource('/account', axios);

    export default {
        data() {
            return {
                form: {
                    username: '',
                    password: '',
                },
                rules: {
                    username: [{
                        required: true,
                        message: '用户名'
                    }],
                    password: [{
                        required: true,
                        message: '密码'
                    }],
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
                            Account.post(_this.form).then(function (res) {
                                if (res.status == 200 && res.data.code == 200) {
                                    _this.$router.push({
                                        path: "/pages/user_setting/account/list"
                                    });
                                    _this.$emit('basicsync');
                                }
                            });
                        }
                    }
                );
            },
        },
        name: "PagesUserSettingAccountSave",
        components: {}
    };
</script>

<style type="text/css">
</style>
