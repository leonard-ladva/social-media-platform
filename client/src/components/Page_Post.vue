<template>
	<div id="postPage">
		<Template_Post v-if="currentPost !== null"
			:post="currentPost.post"
			:user="currentPost.user"
			:tag="currentPost.tag"
		/>
		<hr class="divider">
		<Form_Comment @newComment="newComment()"/>
		<div id="comments">
			<Template_Comment
				v-for="comment in comments"	
				:key="comment.comment.id"

				:comment="comment.comment"
				:user="comment.user"
			/>
			<TriggerIntersect v-if="!outOfComments" @triggerIntersected="getComments()" />
			<h4 id="feedEnd">
				You've reached the end
			</h4>
		</div>
	</div>
</template>
	
<script>
	import axios from '../plugins/axios'
	import Template_Comment from './Template_Comment.vue' 
	import Template_Post from './Template_Post.vue' 
	import TriggerIntersect from './Trigger.vue'
	import Form_Comment from './Form_Comment.vue'

	export default {
		name: 'CommentsView',
		data() {
			return {
				comments: [],
				lastEarliestComment: '-1',
				outOfComments: false,
				currentPost: null,
			}
		},
		components: {
			Template_Comment,
			TriggerIntersect,
			Form_Comment,
			Template_Post,
		},
		methods: {
			async getComments() {
				const response = await axios.get('latestComments', {params: {lastEarliestComment: this.lastEarliestComment}})
				this.comments = this.comments.concat(response.data)
				// get the createdAt time of the last comment gotten, last comment is commented earliest
				if (response.data.length !== 0) {
					this.lastEarliestComment = [...response.data].pop().comment.createdAt
				} 

				if (response.data.length < 10) { this.outOfComments = true } 
			},
			async newComment() {
				this.lastEarliestComment = -1
				this.comments = []
				await this.getComments()
				this.outOfComments = false
			}
		},
		async created() {
			axios.get('/post', {params: {postId: this.$route.params.postId}})
			.then(({data, status}) => {
				if (status === 200) {
					this.currentPost = data
				}
			})
			.catch((error) => {
				console.log(error)
			})
		}
	}
</script>

<style>
	#postPage {
		height: 100%;
	}

	#feedEnd {
		padding: 50px 0;
		border: 1px solid var(--extraLightGrey);
		margin: -1px 0 0 -1px;
	}
	.divider {
		margin-left: 8%;
		width: 84%;
	}
	#postPage .post {
		margin: -1px 0 0 -1px ;
		border: 1px solid var(--extraLightGrey);
		border-bottom: none;
		padding: 1.2rem;
	}
	#postPage .post .tag {
		font-size: 0.7rem;
		background-color: var(--blue);
	}
	#postPage .post .body {
		text-align: left;
		margin: -1rem 0 0 calc(50px + 1rem);
		font-size: 1.5rem;
	}
	#postPage .post:hover {
		background-color: var(--white);
	}
</style>