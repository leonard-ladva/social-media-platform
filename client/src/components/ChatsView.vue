<template>
		<div id="previousMessages">
			<div v-if="$store.state.allUsers" id="conversationEnd">
				<h3>This is the start of your conversation with {{ otherUser.nickname }}</h3>
			</div>

			<TriggerIntersect v-if="$store.state.user != null" id="trigger" @triggerIntersected="getMessages" />
			<ChatMessage
				v-for="msg in messages"	
				:key="msg.ID"

				:message="msg"
			/>
		</div>

		<div id="messageCreate" v-if="$store.state.allUsers">
			<form @submit.prevent="handleSubmit" id="messageForm">
				<textarea 
					v-model="message" 
					class="content" 
					:placeholder="inputPlaceHolder"
					@input="resizeTextArea()"	
					ref="content"
					resize="none"
					rows="1"
				></textarea>
				<button id="send">
					Send
				</button>
			</form>
		</div>

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
		},
		inputPlaceHolder() {
			return `Secret for ${this.otherUser.nickname}`
		}
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
		resizeTextArea() {
			const element = this.$refs.content
			element.style.height = "auto"
			element.style.height = (element.scrollHeight + "px")
		}
	}
}
</script>

<style>
#previousMessages {
	display: flex;
	flex-direction: column;
	align-items: center;
}

#conversationEnd {
	margin: 50px 0;
}

#messageCreate {
	position: sticky;
	bottom: 0;
	height: auto;
	background-color: var(--white);
	padding: 1rem 0 1.5rem 0;
}
#messageForm {
	display: flex;
	align-items: flex-end;
}
#messageForm .content {
	padding: 0.5rem 1.4rem;
	width: 100%;
	border-radius: 2rem;
	background-color: var(--extraLightGrey);
	border: none;
	resize: none;
}
#messageForm .content:focus {
	outline: none;
	box-shadow: none;
}
#messagesFeed {
	display: flex;
	flex-direction: column;
	height: 100%;
}

#send {
	background-color: var(--blue);
	border: none;
	border-radius: 2rem;
	padding: 0.4rem 0.7rem;
	font-size: 1.1rem;
	font-family: "Chirp Bold";
	color: var(--white);
	margin: 0 0.6rem 0 0.3rem;
}

</style>