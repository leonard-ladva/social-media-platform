<template>
	<div>
		<router-link :to="{name: 'home'}">Back Home</router-link>
	</div>
	<div v-if="$store.state.allUsers">
		<h3>This is the start of your conversation with {{ otherUser.nickname }}</h3>
	</div>

	<div>
		<form @submit.prevent="handleSubmit">
			<input v-model="message" type="text">
			<input type="submit" value="send">
		</form>
	</div>
	<div id="messagesFeed">
		<ChatMessage
			v-for="msg in messages"	
			:key="msg.ID"

			:message="msg"
		/>
	</div>
	<TriggerIntersect v-if="$store.state.user != null" id="trigger" @triggerIntersected="getMessages" />
</template>

<script>
import axios from 'axios'
import { ws } from '../assets/js/websocket.js'
import ChatMessage from './Message.vue'
import TriggerIntersect from './Trigger.vue'

export default {
	name: 'ChatsView',
	data() {
		return {
			message: null,
			messages: [],
			lastEarliest: null,
		}
	},
	components: {
		ChatMessage,
		TriggerIntersect,
	},
	computed: {
		receiverID() {
			return this.$route.params.receiverId
		},
		otherUser() {
			return this.$store.state.allUsers.get(this.$route.params.receiverId)
		},
		// chatID returns receiver and sender IDs concatenated in alphabetical order
		chatID() {
			return this.userID < this.receiverID ? this.userID + this.receiverID : this.receiverID + this.userID
		},
		userID() {
			return this.$store.state.user.id
		},
		currentTime() {
			return (new Date).getTime()
		}
	},
	// async created() {
	// 	if (this.$store.state.allUsers) { return }
	// 	let response = await axios.get('/users')

	// 	let users = new Map()
	// 	for (let user of response.data) {
	// 		users.set(user.id, user)
	// 	}
	// 	this.$store.dispatch('allUsers', users)
	// },
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
		},
		async getMessages() {
			const response = await axios.get('/latestMessages', {params: {lastEarliest: this.lastEarliest ? this.lastEarliest : this.currentTime, chatID: this.chatID}})
			this.messages = this.messages.concat(response.data)
			// get the createdAt time of the last post gotten, last post is posted earliest
			if (response.data.length !== 0) {
				this.lastEarliest = [...response.data].pop().createdAt
			} 
			if (response.data.length < 10) {
				const trigger = document.querySelector('#trigger')
				trigger.remove()
			} 
		},
	}
}
</script>

<style>
#messagesFeed {
	display: flex;
	flex-direction: column;
	align-items: center;
}
</style>