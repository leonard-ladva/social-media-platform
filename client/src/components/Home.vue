<template>
	<div id="home">
		<div id="mainView">
			<div id="titlebar">
				<div v-if="this.$route.name != 'feed'">
					<router-link :to="{name: 'home'}">
						<i id="homeLink" class="iconoir-arrow-left"></i>
					</router-link>
				</div>
				<div>
					<h3>
						{{currentPath}}
					</h3>
				</div>
				<!-- <div>
					<h3 v-if="user">Hi, {{ user.firstName }} {{ user.lastName }}!</h3>
					<h3 v-if="!user">You are not logged in!</h3>	
				</div> -->
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
			...mapGetters(['user']),
			currentPath() {
				return this.$route.name == "feed" ? "Home" : this.$route.name.replace(/^\w/, (c) => c.toUpperCase());
			}
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
		height: 100%;
		margin: 0 auto;
	}	
	#mainView {
		width: 68%;
		height: 100%;
		overflow-y: auto;
	}
	#titlebar {
		display: flex;
		position: sticky;
		top: 0;
		height: 3rem;
		width: 100%;
		background-color: var(--white);
		opacity: 0.85;
	}
	#titlebar p {
		font-family: "Chirp Bold";
		font-size: 1.4rem;
	}
	#titlebar #homeLink {
		font-size: 2.1rem;
		margin-right: 1rem;
		color: var(--black);
	}
</style>