<template>
	<div id="sidebar">
		<!-- Logged in user -->
		<!-- Log out button -->
		<!-- List of chats orderded by last message time or else alphabetically-->
			<!-- Active users -->
			<!-- Other users -->
		<div id="currentUser">
			<p class="title">Logged in user</p>
			<div id="userInfo">
				<span class="profilePicture" :style="style"></span>
				<span class="nickname">{{ user.nickname }}</span>
			</div>
			<a class="logout" href="javascript:void(0)" @click="handleClick">Log out</a>
		</div>
		<div id="activeUsers">
			<h3>Active Users</h3>

		</div>
		<div id="unactiveUsers">
			<h3>Offline Users</h3>

		</div>
	</div>	
</template>

<script>
import { mapGetters } from 'vuex'

export default {
	name: 'SideBar',
	methods: {
		handleClick() {
			localStorage.removeItem('token')
			this.$store.dispatch('user', null)
			this.$router.push('/login')
		}
	},
	computed: {
		...mapGetters(['user']),
		style() {
			return 'background-color: ' + this.$store.state.user.color
		},
	},
}
</script>

<style>
	#sidebar {
		width: 30%;
		border-color: rgb(247, 249, 249);
		border-radius: 20px;
	}
	#userInfo {
		background-color: rgb(239, 243, 244);
		height: 55px;
		border-radius: 30px;
		display: flex;
		align-items: center;
	}
	#userInfo .profilePicture {
		height: 45px;
		width: 45px;	
		margin: 5px;
	}

	#userInfo .nickname {
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