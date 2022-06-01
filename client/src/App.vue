<template>
  <div id="primaryPageWrapper">
    <router-view/>
  </div>
</template>

<script>
import { ws } from './plugins/websocket.js'

export default {
	name: 'App',
	async created () {
		if (localStorage.getItem('token')) {
			await this.$store.dispatch('getCurrentUser')	
			ws.connect(this.$store.state.currentUser)
		}
	},
}
</script>

<style scoped>
#primaryPageWrapper {
  height: 100vh;
  padding: 0 3rem;
}
</style>
