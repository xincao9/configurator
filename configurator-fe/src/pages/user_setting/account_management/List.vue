<template>
	<a-table :columns="columns" :dataSource="data">
		<a slot="name" slot-scope="text">{{text}}</a>
	</a-table>
</template>

<script>
	"use strict"
	import resource from "resource-axios";
	import axios from "axios";

	const Accounts = resource("/accounts", axios);
	// const dict = {
	// 	'0': '正常',
	// 	'1': '封禁',
	// 	'2': '删除',
	// 	'component': '普通账户',
	// 	'manager': '经理',
	// 	'admin': '超级管理'
	// };
	const columns = [{
			title: '账号',
			dataIndex: 'share_id',
			key: 'share_id',
			ellipsis: true,
		},
		{
			title: '角色',
			dataIndex: 'role',
			key: 'role',
			ellipsis: true,
		},
		{
			title: '邮箱',
			dataIndex: 'email',
			key: 'email',
			ellipsis: true,
		},
		{
			title: '组',
			dataIndex: 'group',
			key: 'group',
			ellipsis: true,
		},
		{
			title: '状态',
			dataIndex: 'status',
			key: 'status',
			ellipsis: true,
		}
	];

	export default {
		mounted() {
			this.getAccounts();
		},
		data() {
			return {
				data: null,
				columns
			};
		},
		methods: {
			getAccounts() {
				let _this = this;
				Accounts.get().then(function(res) {
					if (res.status == 200 && res.data.code == 200) {
						_this.data = res.data.data;
					}
				})
			}
		},
		name: "PagesUserSettingAccountManagementList",
		components: {}
	};
</script>

<style>
</style>
