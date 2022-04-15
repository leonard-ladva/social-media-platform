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
