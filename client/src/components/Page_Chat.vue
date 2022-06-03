<template>
	<div id="chatsView" :key="this.lastEarliest" v-if="$store.getters.loaded">
		<CreateMessage ref="createMessage"/>
		<div id="previousMessages">
			<ChatMessage
				v-for="msg in messages"	
				:key="msg.id"

				:message="msg"
			/>
			<TriggerIntersect v-if="!outOfMessages" @triggerIntersected="getMessages" />
			<div id="conversationEnd">
				<h3>This is the start of your conversation with {{ otherUser.nickname }}</h3>
			</div>
		</div>

	</div>
</template>

<script>
import axios from 'axios'
import ChatMessage from './Message.vue'
import TriggerIntersect from './Trigger.vue'
import CreateMessage from './CreateMessage.vue'

export default {
	name: 'ChatsView',
	data() {
		return {
			messages: [],
			lastEarliest: null,
			outOfMessages: false,
		}
	},
	components: {
		ChatMessage,
		TriggerIntersect,
		CreateMessage,
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
			return this.$store.state.currentUser.id
		},
		currentTime() {
			return (new Date).getTime()
		},
		lastEarliestMessage() {
			return !this.lastEarliest ? this.currentTime : this.lastEarliest
		},
		newMessages() {
			return this.$store.state.newMessages
		},

	},
	methods: {
		async getMessages() {
			const response = await axios.get('/latestMessages', {
				params: {
					lastEarliest: this.lastEarliestMessage,
					chatID: this.chatID
				}
			})
			if (response.status === 200) {
				this.messages = this.messages.concat(response.data)

				// get the createdAt time of the last post gotten, last post is the one posted earliest
				if (response.data.length !== 0) {
					this.lastEarliest = [...response.data].pop().createdAt
				} 
				if (response.data.length < 10) {
					this.outOfMessages = true
				} 
			}
		},
	},
	watch: {
		newMessages: {
			deep: true,
			handler() {
				if(this.newMessages.has(this.chatID)) {
					this.messages.unshift(this.newMessages.get(this.chatID))
				}
				this.$refs.createMessage.messageSentSuccess()
			}
		}
	}
}
</script>

<style>
#chatsView {
	height: 100%;
}
#previousMessages {
	flex-flow: column nowrap;
	display: flex;
	margin: 0 0.7rem;
}
#conversationEnd {
	padding: 50px 0;
	/* margin: -1px 0; */
}
</style>