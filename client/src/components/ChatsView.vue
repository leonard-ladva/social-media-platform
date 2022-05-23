<template>
	<div>
		<router-link :to="{name: 'home'}">Back Home</router-link>
	</div>
	<div v-if="$store.state.allUsers">
		<h3>Here you can chat with {{ otherUser.nickname }}</h3>
	</div>

	<div>
		<form @submit.prevent="handleSubmit">
			<input v-model="message" type="text">
			<input type="submit" value="send">
		</form>	
	</div>
</template>

<script>
import axios from 'axios'
import { ws } from '../assets/js/websocket.js'

export default {
	name: 'ChatsView',
	data() {
		return {
			message: null,
		}
	},
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
	methods: {
		handleSubmit() {
			let wsMsg = {
				type: 'message',
				user: this.$store.state.user, 
				to: this.otherUser.id, 
				message: this.message
			}
			console.log(wsMsg)
			ws.sendMessage(wsMsg)
		}
	}
}
</script>
