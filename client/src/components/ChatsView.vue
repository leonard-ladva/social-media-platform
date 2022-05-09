<template>
	<div v-if="$store.state.allUsers">
		<h3>Here you can chat with {{ otherUser.nickname }}</h3>
	</div>
</template>

<script>
import axios from 'axios'

export default {
	name: 'ChatsView',
	computed: {
		chatId() {
			return this.$route.params.id
		},
		otherUser() {
			let otherUser = this.$store.state.allUsers.get(this.$route.params.id)
			return otherUser
		},
	},
	async created() {
		if (this.$store.state.allUsers) { return }
		let response = await axios.get('/users')

		let users = new Map()
		for (let user of response.data) {
			users.set(user.id, user)
		}
		this.$store.dispatch('allUsers', users)
	},
}
</script>
