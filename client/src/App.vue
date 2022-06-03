<template>
  <div id="primaryPageWrapper">
    <router-view/>
  </div>
</template>

<script>
export default {
	name: 'App',
}
</script>

<style scoped>
#primaryPageWrapper {
  height: 100vh;
  padding: 0 3rem;
}
</style>
