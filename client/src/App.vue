<template>
  <div id="primaryPageWrapper">
    <router-view/>
  </div>
</template>

<script>
// import NavBar from './components/Nav.vue'
import axios from './plugins/axios'
import { ws } from './plugins/websocket.js'


export default {
	name: 'App',
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
  height: 100vh;
  padding: 0 3rem;
}
</style>
