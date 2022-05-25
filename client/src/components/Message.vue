<template>
	<div class="message" :class="messageType">
		<span>{{message.content}}</span>	
		<!-- <span>{{message.userId}}</span>	
		<span>{{message.createdAt}}</span>	 -->
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
			return this.message.userId == this.$store.state.user.id ? "sentMessage" : "receivedMessage"
		}
	}
}

</script>

<style>
.message {
	/* background-color: lightblue; */
	display: flex;
	flex-direction: column;
	margin: 0.2rem;
	padding: 0.25rem 0.8rem;
	width: fit-content;
	border-radius: 1rem;
}
.sentMessage {
	background-color: pink;
	align-self: end;
}
.receivedMessage {
	background-color: lightblue;
	align-self: start;
}
</style>