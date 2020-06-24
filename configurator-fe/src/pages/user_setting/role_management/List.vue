<template>
	<a-table :columns="columns" :dataSource="data">
		<a slot="name" slot-scope="text">{{text}}</a>
	</a-table>
</template>

<script>
	import resource from "resource-axios";
	import axios from "axios";

	const AccountRoles = resource("/account/roles", axios);

	const columns = [{
			title: '角色',
			dataIndex: 'name',
			key: 'name',
			ellipsis: true,
		},
		{
			title: '标签',
			dataIndex: 'label',
			key: 'label',
			ellipsis: true,
		}
	];
	export default {
		mounted() {
			this.getAccountRoles();
		},
		data() {
			return {
				data: null,
				columns
			};
		},
		methods: {
			getAccountRoles() {
				let _this = this;
				AccountRoles.get().then(function(res) {
					if (res.status == 200 && res.data.code == 200) {
						_this.data = res.data.data;
					}
				})
			}
		},
		name: "PagesUserSettingRoleManagementList",
		components: {}
	};
</script>

<style>
</style>
