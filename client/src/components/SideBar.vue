<template>
	<div id="sidebar">
		<div id="currentUser" v-if="$store.state.user">
			<p class="title">Logged in user</p>
			<UserCard :user="currentUser" />
			<a class="logout" href="javascript:void(0)" @click="logout">Log out</a>
		</div>

		<div id="activeUsers" v-if="$store.state.allUsers">
			<h3>Active Users</h3>
			<UserCard
			v-for="user of $store.state.allUsers.values()" 
			:user="user"
			:key="user.id" />
		</div>

		<div id="unactiveUsers">
			<h3>Offline Users</h3>
		</div>
	</div>	
</template>

<script>
import { mapGetters } from 'vuex'
import axios from 'axios'
import UserCard from './UserCard.vue'

export default {
	name: 'SideBar',
	data() {
		return {
			inactiveUsers: '',
		}
	},
	methods: {
		logout() {
			this.$router.push({name: 'login'})
			localStorage.removeItem('token')
		}
	},
	computed: {
		...mapGetters(['user', 'allUsers']),
		currentUser() {
			return this.$store.state.user
		}
	},
	async created() {
		let response = await axios.get('/users')

		let users = new Map()
		for (let user of response.data) {
			users.set(user.id, user)
		}
		this.$store.dispatch('allUsers', users)
	},
	components: {
		UserCard,
	}
}
</script>

<style>
	#sidebar {
		width: 30%;
		border-color: rgb(247, 249, 249);
		border-radius: 20px;
	}

	#activeUsers .userCard {
		margin: 10px 0;
	}
	#currentUser .title {
		font-family: Chirp Bold;
		margin-bottom: 0px;
	}
	#currentUser .logout {
		font-size: 0.8rem;
	}
</style>