<template>
	<div id="sidebar">
		<div id="currentUser" v-if="$store.state.loggedIn">
			<p class="title">Logged in user</p>
			<UserCard :user="$store.state.currentUser" :active="true" />
			<a class="logout" href="javascript:void(0)" @click="logout">Log out</a>
		</div>

		<div id="allUsers">
			<h4>Chats</h4>
			<UserCard
			v-for="user of activeWithoutCurrent" 
			:user="user"
			:key="user.id" 
			:active ="true" 
			/>

			<UserCard
			v-for="user of offlineWithoutCurrent" 
			:user="user"
			:key="user.id" 
			/>
		</div>

	</div>	
</template>

<script>
import UserCard from './UserCard.vue'
import { ws } from '../plugins/websocket.js'

export default {
	name: 'SideBar',
	methods: {
		logout() {
			this.$router.push({name: 'login'})
			localStorage.removeItem('token')
			ws.disconnect()
		},
	},
	computed: {
		activeWithoutCurrent() {
			let activeUsers = Array.from(this.$store.getters.activeUsers)
			activeUsers = activeUsers.filter(user => user[0] !== this.$store.state.currentUser.id)
			return activeUsers.map(user => user[1])
		},
		offlineWithoutCurrent() {
			let offlineUsers = Array.from(this.$store.getters.offlineUsers)
			offlineUsers = offlineUsers.filter((user) => user[0] !== this.$store.state.currentUser.id)
			return offlineUsers.map(user => user[1])
		}
	},
	components: {
		UserCard,
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