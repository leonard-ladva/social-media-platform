<template>
	<div id="sidebar">
		<div id="currentUser">
			<p class="title">Logged in user</p>
			<UserCard :user="currentUser" />
			<a class="logout" href="javascript:void(0)" @click="handleClick">Log out</a>
		</div>

		<div id="activeUsers">
			<h3>Active Users</h3>
			<UserCard 
			v-for="activeUser in activeUsers" 
			:user="activeUser"
			:key="activeUser.id" />
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
			activeUsers: '',
			inactiveUsers: '',
		}
	},
	methods: {
		handleClick() {
			localStorage.removeItem('token')
			this.$store.dispatch('user', null)
			this.$router.push('/login')
		}
	},
	computed: {
		...mapGetters(['user']),
		currentUser() {
			return this.$store.state.user
		}
	},
	async created() {
		let response = await axios.get('/users')
		console.log(response.data)
		this.activeUsers = response.data
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
	.userCard {
		background-color: rgb(239, 243, 244);
		height: 55px;
		border-radius: 30px;
		display: flex;
		align-items: center;
	}
	.userCard .profilePicture {
		height: 45px;
		width: 45px;	
		margin: 5px;
	}
	#activeUsers .userCard {
		margin: 10px 0;
	}
	.userCard .nickname {
		width: calc(100% - 80px);
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		font-size: 20px;
		padding-bottom: 3px;
	}
	#currentUser .title {
		font-family: Chirp Bold;
		margin-bottom: 0px;
	}
	#currentUser .logout {
		font-size: 0.8rem;
		text-decoration: none;
	}
</style>