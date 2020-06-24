<template>
	<div>
		<h1>新建</h1>
		<hr />
		<a-form-model :model="form" :label-col="{ span: 4 }" :wrapper-col="{ span: 14 }" :rules="rules" style="margin-top: 50px;">
			<a-form-model-item label="角色" prop="name">
				<a-select v-model="form.name" placeholder="请选择角色">
					<a-select-option value="">
						请选择
					</a-select-option>
					<a-select-option value="common">
						普通
					</a-select-option>
					<a-select-option value="manager">
						经理
					</a-select-option>
					<a-select-option value="admin">
						超级管理
					</a-select-option>
				</a-select>
			</a-form-model-item>
			<a-form-model-item label="标签" prop="label">
				<a-input v-model="form.label" placeholder="权限字符串"/>
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

	const AccountRole = resource("/account/role", axios);

	export default {
		data() {
			return {
				form: {
					name: '',
					label: ''
				},
				rules: {
					name: [{
						required: true,
						message: '请选择角色!'
					}],
					label: [{
						required: true,
						message: '请输入标签，标签是根据请求体协议遵循RESTful风格生成!'
					}],
				}
			};
		},
		methods: {
			onSubmit(e) {
				e.preventDefault();
				let _this = this;
				AccountRole.post(_this.form).then(function(res) {
					if (res.status == 200 && res.data.code == 200) {
						_this.$router.push({
							path: "/pages/user_setting/role_management/list"
						});
						_this.$emit('basicsync');
					}
				})
			},
		},
		name: "PagesUserSettingRoleManagementNew",
		components: {}
	};
</script>

<style>
</style>
