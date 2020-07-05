<template>
    <div id="login">
        <div id="login_form_box">
            <h2>账户密码登录</h2>
            <a-form id="login_form" :form="form" class="login-form" @submit="onSubmit">
                <a-form-item>
                    <a-input v-decorator="[
            'username',
            { rules: [{ required: true, message: '请输入账号!' }] },
          ]"
                             placeholder="账号">
                        <a-icon slot="prefix" type="user" style="color: rgba(0,0,0,.25)"/>
                    </a-input>
                </a-form-item>
                <a-form-item>
                    <a-input v-decorator="[
            'password',
            { rules: [{ required: true, message: '请输入密码!' }] },
          ]"
                             type="password" placeholder="密码">
                        <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)"/>
                    </a-input>
                </a-form-item>
                <a-form-item>
                    <a-checkbox v-decorator="[
            'remember',
            {
              valuePropName: 'checked',
              initialValue: true,
            },
          ]">自动登录
                    </a-checkbox>
                    <a class="login-form-forgot" href>忘记密码</a>
                    <a-button type="primary" html-type="submit" class="login-form-button">登录</a-button>
                    或者
                    <a href>注册</a>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<script>
    import router from '../router.js';
    import resource from "resource-axios";
    import axios from "axios";

    const Session = resource("/session", axios);

    export default {
        beforeCreate() {
            this.form = this.$form.createForm(this, {
                name: "login_form_box"
            });
        },
        methods: {
            onSubmit(e) {
                e.preventDefault();
                this.form.validateFields((err, values) => {
                    if (!err) {
                        Session.post(values).then(function (res) {
                            if (res.status == 200 && res.data.code == 200) {
                                router.push({
                                    path: "/pages/configurator/app/list"
                                });
                            }
                        });
                    }
                });
            }
        },
        name: "Login",
    };
</script>

<style>
    #login {
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        -webkit-box-align: center;
        -webkit-align-items: center;
        -ms-flex-align: center;
        align-items: center;
        height: 100vh;
        background-image: url(../assets/login_background.png);
        background-repeat: no-repeat;
        background-size: cover;
    }

    #login_form_box {
        color: #333;
        width: 356px;
        margin: 0 auto;
    }

    #login_form .login-form {
        max-width: 300px;
    }

    #login_form .login-form-forgot {
        float: right;
    }

    #login_form .login-form-button {
        width: 100%;
    }
</style>
