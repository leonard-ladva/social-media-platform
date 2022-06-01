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
				const response = await axios.post('login', {
					nickname: this.nickname,
					passwordPlain: this.password,
				})
				if (response.status === 200) {
					localStorage.setItem('token', response.data.Token)
					this.$store.dispatch('logInUser', response.data.User)
					this.$router.push({name: 'feed'})
				} else {
					console.log(`ERROR: Login. Status Code: ${response.status}`)
				}
			},
		},
		created() {
			localStorage.removeItem('token')
			this.$store.dispatch('logOutUser')
		}
	}
</script>
