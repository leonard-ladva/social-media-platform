<template>
	<div class="notification">
		<router-link :to="link">
			<div class="body">
				<h5>New Message</h5>
				<span>{{sender.nickname}}</span>
				<span>Message</span>
			</div>	
		</router-link>
	</div>
</template>

<script>
export default {
	name: 'MessageNotifcation',
	props: {
		message: null
	},
	computed: {
		sender() {
			return this.$store.state.allUsers.get(this.message.userId)
		},
		link() {
			return `/chat/${this.sender.id}`
		}
	}
}
</script>
<style scoped>
	.notification {
		text-decoration: none;
		color: var(--black);
		background-color: var(--extraLightGrey);
		position: absolute;
		bottom: 2rem;
		right: 2rem;
		border-radius: 1rem;
		height: 4rem;
		width: 18rem;
	}


</style>
