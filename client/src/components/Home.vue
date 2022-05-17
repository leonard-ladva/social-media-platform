<template>
	<div id="home">
		<div id="mainView">
			<div>
				<h3 v-if="user">Hi, {{ user.firstName }} {{ user.lastName }}!</h3>
				<h3 v-if="!user">You are not logged in!</h3>	
			</div>

			<router-view />
		</div>
		<SideBar/>
	</div>
</template>

<script>
	import { mapGetters } from 'vuex'
	import SideBar from './SideBar.vue'

	export default {
		name: 'HomePage',
		computed: {
			...mapGetters(['user'])
		},
		components: {
			SideBar,
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
		justify-content: space-between;
		max-width: 900px;
		margin: 0 auto;
	}	
	#mainView {
		margin-top: 3rem;
		width: 65%;
	}
	#titleBar {
		position: fixed;
		background-color: #bbbbbb;
		height: 3rem;
		opacity: 95%;
		width: 100%;

	}
	#titleBar p {
		font-family: Chirp Bold;
		font-size: 1.4rem;

	}
</style>