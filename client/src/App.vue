<template>
  <div id="primaryPageWrapper">
	<!-- <button v-on:click="sendMessage('hello')">Say Gekki!</button> -->
    <router-view/>
  </div>
</template>

<script>
// import NavBar from './components/Nav.vue'
import axios from './plugins/axios'

export default {
	name: 'App',
	data() {
		return {
			connection: null
		}
	},
	async created () {
		const response = await axios.get('user')
		if (response.status === 200) {
			this.webSocket()	
		}

		this.$store.dispatch('user', response.data.User)
	},
	methods: {
		webSocket() {
			console.log("Starting connection to WebSocket Server")
			this.connection = new WebSocket("ws://localhost:9100/ws")

			this.connection.onmessage = function(event) {
				console.log(event);
			}

			this.connection.onopen = function(event) {
				console.log(event)
				console.log("Successfully connected to the websocket")
			}
		},
		sendMessage(message) {
			console.log(this.connection)
			this.connection.send(message)
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
