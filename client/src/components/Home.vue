<template>
	<div id="home">
		<div id="mainView">
			<TitleBar/>	
			<router-view :key="$route.params.receiverId"/>
		</div>
		<SideBar/>
	</div>
	<MessageNotification v-for="notification in notifications" :key="notification.id" :message="notification"/>
</template>

<script>
import SideBar from './SideBar.vue'
import TitleBar from './Titlebar.vue'
import MessageNotification from './Notification.vue'

export default {
	name: 'HomePage',
	components: {
		SideBar,
		TitleBar,
		MessageNotification,
	},
	computed: {
		notifications() {
			return this.$store.state.notifications
		},
	},
	async created() {
		await this.$store.dispatch('getUsers')
	},
}
</script>

<style>
	#home {
		display: flex;
		max-width: 1000px;
		height: 100%;
		margin: auto;
	}	
	#mainView {
		width: 65%;
		height: 100%;
		overflow-y: auto;
		display: flex;
		flex-flow: column nowrap;
		border-left: 1px solid var(--extraLightGrey);
		border-right: 1px solid var(--extraLightGrey);
	}
</style>