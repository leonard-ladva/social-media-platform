<template>
	<MakePost />
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
	import axios from '../plugins/axios'
	import PostTemplate from './PostTemplate.vue' 
	import TriggerIntersect from './Trigger.vue'
	import MakePost from './MakePost.vue'

	export default {
		name: 'PostsView',
		data() {
			return {
				posts: [],
				lastEarliestPost: '-1',
			}
		},
		components: {
			PostTemplate,
			TriggerIntersect,
			MakePost
		},
		methods: {
			async getPosts() {
				const response = await axios.get('latestPosts', {params: {lastEarliestPost: this.lastEarliestPost}})
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
		border: 0.1rem solid var(--extraExtraLightGrey);
	}

	.post .nickname {
		font-family: "Chirp Bold";
	}
	.post:hover {
		background-color: var(--extraExtraLightGrey);
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
	.tag {
		background-color: var(--blue);
	}
</style>