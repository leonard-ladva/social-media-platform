<template>
	<div id="sidebar" v-if="$store.state.loggedIn">
		<div id="currentUser">
			<p class="title">Logged in user</p>
			<Template_User :user="$store.state.currentUser" :active="true" />
			<a class="logout" href="javascript:void(0)" @click="logout">Log out</a>
		</div>

		<div id="allUsers">
			<h4>Chats</h4>
			<Template_User
			v-for="user of activeWithoutCurrent" 
			:user="user"
			:key="user.id" 
			:active ="true" 
			/>

			<Template_User
			v-for="user of offlineWithoutCurrent" 
			:user="user"
			:key="user.id" 
			/>
		</div>

	</div>	
</template>

<script>
import Template_User from './Template_User.vue'

export default {
	name: 'SideBar',
	methods: {
		logout() {
			this.$router.push({name: 'login'})
			this.$store.dispatch('logOutUser')
		},
	},
	computed: {
		activeWithoutCurrent() {
			return this.$store.getters.activeUsers
		},
		offlineWithoutCurrent() {
			return this.$store.getters.offlineUsers
		},

	},
	components: {
		Template_User,
	},
}
</script>

<style>
	#sidebar {
		width: 35%;
		max-width: 250px;
		margin-left: 1rem;
	}
	#allUsers {
		margin-top: 2.3rem;
	}
	#allUsers .userCard {
		margin: 0.7rem 0;
	}
	#currentUser .title {
		font-family: "Chirp Bold";
		margin-bottom: 0px;
	}
	#currentUser .logout {
		font-size: 0.8rem;
	}
</style>