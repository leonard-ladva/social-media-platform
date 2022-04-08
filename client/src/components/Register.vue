<template>
	<form @submit.prevent="handleSubmit">
		<h3>Sign up</h3>

		<div class="form-group">
			<label>First Name</label>
			<input type="text" v-model="firstName" class="form-control" placeholder="First Name"/>
		</div>

		<div class="form-group">
			<label>Last Name</label>
			<input type="text" v-model="lastName" class="form-control" placeholder="Last Name"/>
		</div>

		<div class="form-group">
			<label>Email</label>
			<input type="email" v-model="email" class="form-control" placeholder="Email"/>
		</div>

		<div class="form-group">
			<label>Password</label>
			<input type="password" v-model="password" class="form-control" placeholder="Password"/>
		</div>

		<div class="form-group">
			<label>Confirm Password</label>
			<input type="password" v-model="passwordConfirm" class="form-control" placeholder="Confirm Password"/>
		</div>

		<div class="form-group">
			<label>Nickname</label>
			<input type="text" v-model="nickname" class="form-control" placeholder="Nickname"/>
		</div>

		<div class="form-group">
			<label>Gender</label>
			<input type="text" v-model="gender" class="form-control" placeholder="Gender"/>
		</div>

		<div class="form-group">
			<label>Age</label>
			<input type="number" v-model="age" class="form-control" placeholder="Age"/>
		</div>

		<button class="btn btn-primary btn-block">Sign Up</button>
	</form>
</template>
 
<script>
	import axios from 'axios'
	export default {
    name: "RegisterForm",
    data() {
        return {
				firstName: "",
				lastName: "",
				email: "",
				password: "",
				passwordConfirm: "",
				nickname: "",
				gender: "",
				age: "",
        };
    },
    methods: {
        handleSubmit() {
			const data = {
				firstName: this.firstName,
				lastName: this.lastName,
				email: this.email,
				passwordPlain: this.password,
				passwordConfirm: this.passwordConfirm,
				nickname: this.nickname,
				gender: this.gender,
				age: this.age,
			}
			axios.post('http://localhost:9100/register', data)
			.then(
				res => {
					console.log(res)
				}
			).catch(
				err => {
					console.log(err)
				}
			)
        },
    },
}

</script>