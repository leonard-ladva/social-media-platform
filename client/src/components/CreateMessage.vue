<template>
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
			<button id="send">Send</button>
		</form>
	</div>
</template>
<script>
import { ws } from '../plugins/websocket.js'

export default {
	name: 'CreateMessage',
	data() {
		return {
			message: null,
		}
	},
	methods: {
		handleSubmit() {
			let msg = {
				userId: this.$store.state.user.id, 
				receiverId: this.receiverID, 
				content: this.message,
			}
			console.log('yes')
			ws.sendMessage(msg)
		},
		resizeTextArea() {
			const element = this.$refs.content
			element.style.height = "auto"
			element.style.height = (element.scrollHeight + "px")
		},
	},
	computed: {
		receiverID() {
			return this.$route.params.receiverId
		},
		inputPlaceHolder() {
			return `Secret for ${this.otherUser.nickname}`
		},
		otherUser() {
			return this.$store.state.allUsers.get(this.$route.params.receiverId)
		},
	}
}
</script>

<style>
#messageCreate {
	height: auto;
	background-color: var(--white);
	padding: 1rem 0 1.5rem 1rem;
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