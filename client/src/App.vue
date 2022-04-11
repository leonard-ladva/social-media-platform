<template>
  <div id="primaryPageWrapper">
    <NavBar/>
    <div class="auth-wrapper">
      <div class="auth-inner">
        <router-view/>
      </div>
    </div>
  </div>
  <!-- <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link> |
    <router-link to="/feed">Feed</router-link> |
    <router-link to="/makepost">New Post</router-link>
  </nav> -->
</template>

<script>
import NavBar from './components/Nav.vue'

export default {
  name: 'App',
  components: {
    NavBar,
  },
}
</script>

<style scoped>
#primaryPageWrapper {
  height: 100%;
}
</style>
