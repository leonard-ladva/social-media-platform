<template>
	<div class="auth-wrapper">
		<div class="auth-inner">
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
			<router-link to="register" class="nav-link">Sign up</router-link>
		</div>
	</div>	
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
				let data = await this.loginRequest()
				if (data.User) {
					this.$router.push({name: 'feed'})
				}

			},
			async loginRequest() {
				try {
					const response = await axios.post('login', {
						nickname: this.nickname,
						passwordPlain: this.password,
					})
					localStorage.setItem('token', response.data.Token)
					this.$store.dispatch('user', response.data.User)
					return response.data 
				} catch(err) {
					console.log(err)
					return err
					// show error on login form
				}
			},
		},
	}
</script>
