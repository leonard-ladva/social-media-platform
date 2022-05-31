<template>
	<div id="sidebar">
		<div id="currentUser" v-if="$store.state.user">
			<p class="title">Logged in user</p>
			<UserCard :user="currentUser" :active="true" />
			<a class="logout" href="javascript:void(0)" @click="logout">Log out</a>
		</div>


		<div id="activeUsers" v-if="$store.state.activeUsers">
			<h3 v-if="$store.state.activeUsers.size != 0">Active Users</h3>
			<UserCard
			v-for="user of $store.state.activeUsers.values()" 
			:user="user"
			:key="user.id" 
			:active ="true" />
		</div>

		<div id="offlineUsers" v-if="$store.state.offlineUsers">
			<h3>Offline Users</h3>
			<UserCard
			v-for="user of $store.state.offlineUsers.values()" 
			:user="user"
			:key="user.id" />
		</div>

	</div>	
</template>

<script>
import { mapGetters } from 'vuex'
import axios from 'axios'
import UserCard from './UserCard.vue'

export default {
	name: 'SideBar',
	methods: {
		logout() {
			this.$router.push({name: 'login'})
			localStorage.removeItem('token')
		},
		sortUsersStatus(allUsers) {
			let activeUsers = new Map()	
			let offlineUsers = new Map()

			for (let user of allUsers.values()) {
				if (user.id == this.$store.state.user.id) continue
				if (user.active == true) {
					activeUsers.set(user.id, user)
					continue	
				}
				offlineUsers.set(user.id, user)
			}

			this.$store.dispatch('activeUsers', activeUsers)
			this.$store.dispatch('offlineUsers', offlineUsers)
		}
	},
	computed: {
		...mapGetters(['user', 'allUsers']),
		currentUser() {
			return this.$store.state.user
		},
	},
	async created() {
		let response = await axios.get('/users')

		let users = new Map()
		for (let user of response.data) {
			users.set(user.id, user)
		}
		this.$store.dispatch('allUsers', users)
		
		this.sortUsersStatus(users)
		
		await Notification.requestPermission();
		new Notification('Hi, How are you?',{
		body: 'Have a good day',
		});
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

	#activeUsers, #offlineUsers {
		margin-top: 2.3rem;
	}
	#activeUsers .userCard, #offlineUsers .userCard {
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