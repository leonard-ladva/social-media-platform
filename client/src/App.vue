<template>
  <div id="primaryPageWrapper">
    <NavBar/>
    <div class="auth-wrapper">
      <div class="auth-inner">
        <router-view/>
      </div>
    </div>
  </div>
</template>

<script>
import NavBar from './components/Nav.vue'
import axios from './plugins/axios'

export default {
	name: 'App',
	components: {
		NavBar,
	},
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
