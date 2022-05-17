<template>
  <div id="primaryPageWrapper">
	<button v-on:click="sendMessage('hello')">Say Gekki!</button>
    <router-view/>
  </div>
</template>

<script>
// import NavBar from './components/Nav.vue'
import axios from './plugins/axios'
import { ws } from './assets/js/websocket.js'


export default {
	name: 'App',
	// data() {
	// 	return {
	// 		// connection: null
	// 	}
	// },
	async created () {
		if (localStorage.getItem('token')) {
			await this.getCurrentUser()	
			ws.connect(this.$store.state.user)
		}
	},
	methods: {
		async getCurrentUser() {
			try {
				const response = await axios.get('user')
				this.$store.dispatch('user', response.data.User)
			} catch (err) {
				console.log("User not logged in")
				this.$store.dispatch('user', null)
			}
		}
	}
}
</script>

<style scoped>
#primaryPageWrapper {
  height: 100%;
  margin: 0 2rem;
}
</style>
