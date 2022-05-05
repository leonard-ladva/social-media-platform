<template>
  <div id="primaryPageWrapper">
    <router-view/>
  </div>
</template>

<script>
// import NavBar from './components/Nav.vue'
import axios from './plugins/axios'

export default {
	name: 'App',
	async created () {
		const response = await axios.get('user')

		this.$store.dispatch('user', response.data.User)
	}
}
</script>

<style scoped>
#primaryPageWrapper {
  height: 100%;
}
</style>
