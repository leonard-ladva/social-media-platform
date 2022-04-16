<template>
	<div class="root">
	<form @submit.prevent="handleSubmit">
		<h3>Sign up</h3>

		<div class="form-group">
			<label>First Name</label>
			<input class="form-control"
			type="text"
			v-model.lazy="firstName"
			@blur="v$.firstName.$touch"
			@focus="v$.firstName.$reset()"
			autocomplete="given-name"
			placeholder="First Name"/>
			<p
				v-for="error of v$.firstName.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<div class="form-group">
			<label>Last Name</label>
			<input class="form-control"
			type="text"
			v-model.lazy="lastName"
			@blur="v$.lastName.$touch"
			@focus="v$.lastName.$reset()"
			autocomplete="family-name"
			placeholder="Last Name"/>
			<p
				v-for="error of v$.lastName.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<div class="form-group">
			<label>Email</label>
			<input class="form-control"
			type="email"
			v-model.lazy="email"
			@blur="v$.email.$touch"
			@focus="v$.email.$reset()"
			autocomplete="email"
			placeholder="Email"
			novalidate/>
			<p
				v-for="error of v$.email.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<div class="form-group">
			<label>Password</label>
			<input class="form-control"
			type="password"
			v-model.lazy="password"
			@blur="v$.password.$touch"
			@focus="v$.password.$reset()"
			autocomplete="new-password"
			placeholder="Password"/>
			<p
				v-for="error of v$.password.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<div class="form-group">
			<label>Confirm Password</label>
			<input class="form-control"
			type="password"
			v-model.lazy="passwordConfirm"
			@blur="v$.passwordConfirm.$touch"
			@focus="v$.passwordConfirm.$reset()"
			autocomplete="new-password"
			placeholder="Confirm Password"/>
			<p
				v-for="error of v$.passwordConfirm.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<div class="form-group">
			<label>Nickname</label>
			<input class="form-control"
			type="text"
			v-model.lazy="nickname"
			@blur="v$.nickname.$touch"
			@focus="v$.nickname.$reset()"
			autocomplete="nickname"
			placeholder="Nickname"/>
			<p
				v-for="error of v$.nickname.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<div class="form-group">
			<label>Gender</label>
			<input class="form-control"
			type="text"
			v-model.lazy="gender"
			@blur="v$.gender.$touch"
			@focus="v$.gender.$reset()"
			autocomplete="sex"
			placeholder="Gender"/>
			<p
				v-for="error of v$.gender.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<div class="form-group">
			<label>Age</label>
			<input class="form-control"
			type="number"
			v-model="age"
			@blur="v$.age.$touch"
			@focus="v$.age.$reset()"
			placeholder="Age"/>
			<p
				v-for="error of v$.age.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<button class="btn btn-primary btn-block">Sign Up</button>
	</form>
	
	</div>
</template>
 
<script>
import useVuelidate from '@vuelidate/core'
import { required, email, minLength, maxLength, between, sameAs, alpha, alphaNum, integer, helpers } from '@vuelidate/validators'
import axios from '../plugins/axios'
const { withAsync, withMessage } = helpers

export default {
	name: "RegisterForm",
	setup: () => ({ v$: useVuelidate() }),
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
	validations() {
		return {
			firstName:		{ required, maxLength: maxLength(50), alpha },
			lastName:		{ required, maxLength: maxLength(50), alpha },
			email:			{ required, email,
				isUnique: withAsync(withMessage('Oops, Email already in use', async (value) => {
					if (value === '') return true
					const resp = await axios.post('isUnique', {Email: value})
					return resp.data 
				})) },
			password:		{ required, minLength: minLength(8), maxLength: maxLength(50) },
			passwordConfirm:{ required, sameAs: sameAs(this.password) },
			nickname:		{ required, minLength: minLength(3), maxLength: maxLength(20), alphaNum, 
				isUnique: withAsync(withMessage('Oops, Nickname already taken', async (value) => {
					if (value === '') return true
					const resp = await axios.post('isUnique', {Nickname: value})
					return resp.data 
				})) },
			gender:			{ required, maxLength: maxLength(50), alpha },
			age:			{ required, between: between(0, 120), integer },
		}
	},
	methods: {
		async handleSubmit() {
			const ifFormCorrect = await this.v$.$validate()
			if (!ifFormCorrect) return
		
			await axios.post('register', {
				firstName: this.firstName,
				lastName: this.lastName,
				email: this.email,
				passwordPlain: this.password,
				passwordConfirm: this.passwordConfirm,
				nickname: this.nickname,
				gender: this.gender,
				age: this.age,
			})
			this.$router.push('/login')
		},
		// async isUnique(field, value) {
		// 	console.log(field, value)
		// 	if (value === '') return

		// 	// const data = `{${field}: ${value}}`
		// 	if (field === 'Email') {
		// 		const resp = await axios.post('isUnique', {Email: value})
		// 		console.log(resp.data)
		// 		return resp.data
		// 	} 
		// 	const resp = await axios.post('isUnique', {Nickname: value})
		// 	console.log(resp.data, typeof resp.data)
		// 	return resp.data
		// }
	},
}
</script>