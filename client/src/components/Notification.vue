<template>
	<div class="notification">
		<router-link :to="link" class="chatLink">
			<div class="body">
				<h5>New Message</h5>
				<p>{{sender.nickname}}</p>
				<span>{{message.content}}</span>
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
	},
	methods: {
		wait() {
			return new Promise((resolve) => {
				setTimeout(() => {
					resolve()
				}, 10000)
			})
		}
	},
	async created() {
		await this.wait()
		this.$store.dispatch('removeFirstNotification')
	}
}
</script>
<style scoped>
	.notification {
		background-color: var(--extraLightGrey);
		position: absolute;
		bottom: 2rem;
		right: 2rem;
		border-radius: 1rem;
		height: 4rem;
		width: 18rem;
	}
	.notification .chatLink {
		text-decoration: none;
	}
	.notification .body {
		color: var(--black);
	}


</style>
