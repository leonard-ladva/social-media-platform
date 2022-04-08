<template>
		<nav class="navbar navbar-expand navbar-light fixed-top">
		<div class="container">
			<a href="#" class="navbar-brand">Home</a>
			<div class="collapse navbar-collapse">
				<ul class="navbar-nav ml-auto">
					<li class="nav-item">
						<router-link to="login" class="nav-link">Login</router-link>
					</li>
					<li class="nav-item">
						<router-link to="register" class="nav-link">Sign up</router-link>
					</li>
				</ul>
			</div>
		</div>	
	</nav>
</template>

<script>
	export default {
		name: 'NavBar',
	}
</script>
