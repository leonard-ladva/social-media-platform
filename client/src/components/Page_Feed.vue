<template>
	<div id="postsView">
		<Form_Post @newPost="newPost()"/>
		<div id="posts">
			<Template_Post
				v-for="post in posts"	
				:key="post.id"

				:post="post.post"
				:user="post.user"
				:tag="post.tag"
			/>
			<TriggerIntersect v-if="!outOfPosts" @triggerIntersected="getPosts" />
			<h4 id="feedEnd">
				You've reached the end
			</h4>
		</div>
	</div>
</template>
	
<script>
	import axios from '../plugins/axios'
	import Template_Post from './Template_Post.vue' 
	import TriggerIntersect from './Trigger.vue'
	import Form_Post from './Form_Post.vue'

	export default {
		name: 'PostsView',
		data() {
			return {
				posts: [],
				lastEarliestPost: '-1',
				outOfPosts: false,
			}
		},
		components: {
			Template_Post,
			TriggerIntersect,
			Form_Post,	
		},
		methods: {
			async getPosts() {
				const response = await axios.get('latestPosts', {params: {lastEarliestPost: this.lastEarliestPost}})
				this.posts = this.posts.concat(response.data)
				// get the createdAt time of the last post gotten, last post is posted earliest
				if (response.data.length !== 0) {
					this.lastEarliestPost = [...response.data].pop().post.createdAt
				} 

				if (response.data.length < 5) { this.outOfPosts = true } 
			},
			newPost() {
				this.lastEarliestPost = -1
				this.posts = []
			}
		},
	}
</script>

<style>
	#postsView {
		height: 100%;
	}

	#feedEnd {
		padding: 50px 0;
		border: 1px solid var(--extraLightGrey);
		margin: -1px 0 0 -1px;
	}
	.tag {
		background-color: var(--blue);
	}
</style>