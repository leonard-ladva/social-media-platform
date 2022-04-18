<template>
	<form @submit.prevent="handleSubmit">
		<h3>Login</h3>	
	
		<div class="form-group">
			<label>Nickname or Email</label>
			<input class="form-control"
			v-model="nickname"
			autocomplete="username"
			placeholder="Nickname or Email"/>
		</div>

		<div class="form-group">
			<label>Password</label>
			<input class="form-control"
			type="password"
			v-model="password"
			autocomplete="password"
			placeholder="Password"/>
		</div>

		<button class="btn btn-primary btn-block">Login</button>
	</form>	
</template>

<script>
	import axios from '../plugins/axios'

	export default {
		name: 'LoginForm',
		data() {
			return {
				nickname: '',
				password: '',
			}
		},
		methods: {
			async handleSubmit() {
				const resp = await axios.post('login', {
					nickname: this.nickname,
					passwordPlain: this.password,
				})
				localStorage.setItem('token', resp.data.Token)
				this.$store.dispatch('user', resp.data.User)
				this.$router.push('/')
			}
		}
	}
</script>
