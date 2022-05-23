<template>
  <div id="primaryPageWrapper">
    <router-view/>
  </div>
</template>

<script>
// import NavBar from './components/Nav.vue'
import axios from './plugins/axios'
import { ws } from './assets/js/websocket.js'


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
  height: 100%;
  margin: 1rem 2rem;
}
</style>
