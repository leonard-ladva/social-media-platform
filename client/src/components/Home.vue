<template>
	<h3>Helloss</h3>	
</template>

<script>
	import axios from '../plugins/axios'

	export default {
		name: 'HomePage',
		async created() {
			const resp = await axios.get('user')

			console.log(resp)
		}
	}
</script>