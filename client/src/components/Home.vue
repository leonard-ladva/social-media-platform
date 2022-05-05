<template>
<div id="homeWrapper">
	<!-- <div id="titleBar">
		<p>Home</p>
	</div> -->
	<div id="feed">

		<div>
			<h3 v-if="user">Hi, {{ user.firstName }} {{ user.lastName }}!</h3>
			<h3 v-if="!user">You are not logged in!</h3>	
		</div>

		<MakePost/>
		<PostsFeed/>
	</div>
	<SideBar/>
</div>
</template>

<script>
	import { mapGetters } from 'vuex'
	import MakePost from './MakePost.vue'
	import PostsFeed from './PostsFeed.vue'
	import SideBar from './SideBar.vue'

	export default {
		name: 'HomePage',
		computed: {
			...mapGetters(['user'])
		},
		components: {
			MakePost,
			PostsFeed,
			SideBar,
		},
	}
</script>

<style>
	#homeWrapper {
		display: flex;
		justify-content: space-between;
		max-width: 900px;
		margin: 0 auto;
	}	
	#feed {
		margin-top: 3rem;
		width: 65%;
	}
	#titleBar {
		position: fixed;
		background-color: #bbbbbb;
		height: 3rem;
		opacity: 95%;
		width: 100%;

	}
	#titleBar p {
		font-family: Chirp Bold;
		font-size: 1.4rem;

	}
</style>