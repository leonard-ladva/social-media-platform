<template>
	<div id="posts">
		<PostTemplate
			v-for="post in posts"	
			:key="post.ID"

			:post="post.post"
			:user="post.user"
			:tag="post.tag"
		/>
		<TriggerIntersect id="trigger" @triggerIntersected="getPosts" />
		<h4 id="feedEnd">
			You've reached the end
		</h4>
	</div>
</template>
	
<script>
	import axios from 'axios'
	import PostTemplate from './PostTemplate.vue' 
	import TriggerIntersect from './Trigger.vue'

	export default {
		data() {
			return {
				posts: [],
				lastEarliestPost: '-1',
			}
		},
		components: {
			PostTemplate,
			TriggerIntersect,
		},
		methods: {
			async getPosts() {
				const response = await axios.get('latestPosts', {params: {lastEarliestPost: this.lastEarliestPost}})
				console.log(response.data)
				this.posts = this.posts.concat(response.data)
				// get the createdAt time of the last post gotten, last post is posted earliest
				if (response.data.length !== 0) {
					this.lastEarliestPost = [...response.data].pop().post.createdAt
				} 
				if (response.data.length < 5) {
					const trigger = document.querySelector('#trigger')
					trigger.remove()
				} 
			},
		},
	}
</script>

<style>
	.post {
		border: 0.1rem solid rgb(239, 243, 244);
	}

	.post .nickname {
		font-family: Chirp Bold;
	}
	.post:hover {
		background-color: rgb(245 248 250);
	}
	/* .post .header {
		display: flex;
		justify-content: space-around;
	}
	.post .body {
		display: flex;
		justify-content: space-around;
	} */
	#feedEnd {
		margin: 50px 0;
	}
</style>