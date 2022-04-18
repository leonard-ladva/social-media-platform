<template>
<div>
	<h3 v-if="user">Hi, {{ user.FirstName }} {{ user.LastName }}!</h3>
	<h3 v-if="!user">You are not logged in!</h3>	
</div>
</template>

<script>
	import { mapGetters } from 'vuex'

	export default {
		name: 'HomePage',
		computed: {
			...mapGetters(['user'])
		}
	}
</script>