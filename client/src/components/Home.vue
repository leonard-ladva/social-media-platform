<template>
	<div id="home">
		<div id="mainView">
			<TitleBar/>	
			<router-view :key="$route.params.receiverId"/>
		</div>
		<SideBar/>
	</div>
	<MessageNotification 
	v-for="message in $store.state.messages.values()" 
	:key="message.chatId"
	:message="message"
	/>
</template>

<script>
import { mapGetters } from 'vuex'
import SideBar from './SideBar.vue'
import TitleBar from './Titlebar.vue'
import MessageNotification from './Notification.vue'

export default {
	name: 'HomePage',
	computed: {
		...mapGetters(['user']),
	},
	components: {
		SideBar,
		TitleBar,
		MessageNotification,
	},
	// created() {
	// 	if (!localStorage.getItem('token')) {
	// 		this.$router.push({name: 'login'})
	// 	}
	// }
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