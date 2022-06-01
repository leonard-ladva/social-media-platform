<template>
	<div class="message" :class="messageType">
		<p class="content">{{message.content}}</p>	
		<span class="time">{{messageTime}}</span>	
	</div>	
</template>

<script>

export default {
	name: 'ChatMessage',
	props: {
		message: Object,
	},
	computed: {
		messageType() {
			return this.message.userId == this.$store.state.currentUser.id ? "sentMessage" : "receivedMessage"
		},
		messageTime() {
			let date = new Date(this.message.createdAt)
			// let month = date.toLocaleString("default", {month: "short"})
			let hours = date.getHours()
			let minutes = date.getMinutes()
			minutes = minutes < 10 ? `0${minutes}` : minutes

			return `${hours}:${minutes}`
		}
	}
}

</script>

<style>
.message {
	display: flex;
	margin: 0.3rem 0 0.1rem 0;
	padding: 0.5rem 0.8rem;
	width: fit-content;
	border-radius: 1.2rem;
	max-width: 85%;
	line-height: 1.2rem;
}
.sentMessage {
	background-color: pink;
	align-self: flex-end;
}
.receivedMessage {
	background-color: lightblue;
	align-self: flex-start;
}
.message .time {
	font-size: 0.65rem;
	color: var(--darkGrey);
	font-family: "Chirp Regular";
	margin-left: 0.45rem;
	align-self: flex-end;
	height: fit-content;
}
.message .content {
	font-family: "Chirp Heavy";
	margin-bottom: 0.25rem;
	text-align: left;
}
</style>