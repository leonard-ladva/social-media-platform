<template>
	<form method="post">
		<h3>Login</h3>	

		<div class="form-group">
			<!-- <BaseInput 
			v-model="login.email"	
			label="Email"
			type="email"
			/> -->
			<label>Email
        <input type="email" class="form-control" placeholder="Email"/>
      </label>
    </div>

		<div class="form-group">
			<!-- <BaseInput
			v-model="login.password"
			label="Password"
			type="password"
			/> -->
			<label>Password
        <input type="password" class="form-control" placeholder="Password"/>
      </label>
    </div>

		<button class="btn btn-primary btn-block">Login</button>
	</form>	
</template>

<script>
	// import BaseInput from "./BaseInput.vue";
	export default {
		name: 'LoginForm',
		components: {
			// BaseInput,
		}
	}
</script>
