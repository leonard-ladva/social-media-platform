<template>
		<nav class="navbar navbar-expand navbar-light fixed-top">
		<div class="container">
			<router-link to="/" class="navbar-brand">Home</router-link>
			<div class="collapse navbar-collapse">
				<ul class="navbar-nav ml-auto" v-if="!user">
					<li class="nav-item">
					</li>
					<li class="nav-item">
					</li>
				</ul>
				<ul class="navbar-nav ml-auto" v-if="user">
					<li class="nav-item">
					</li>
				</ul>
			</div>
		</div>	
	</nav>
</template>

<script>
	import { mapGetters } from 'vuex'

	export default {
		name: 'NavBar',
		methods: {
			handleClick() {
				localStorage.removeItem('token')
				this.$store.dispatch('user', null)
				this.$router.push('/')
			}
		},
		computed: {
			...mapGetters(['user'])
		}
	}
</script>
