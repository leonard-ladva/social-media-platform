<template>
	<div id="commentsView">
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
	import TriggerIntersect from './Trigger.vue'
	import Form_Comment from './Form_Comment.vue'

	export default {
		name: 'CommentsView',
		data() {
			return {
				comments: [],
				lastEarliestComment: '-1',
				outOfComments: false,
			}
		},
		components: {
			Template_Comment,
			TriggerIntersect,
			Form_Comment,
		},
		methods: {
			async getComments() {
				const response = await axios.get('latestComments', {params: {lastEarliestComment: this.lastEarliestComment}})
				this.comments = this.comments.concat(response.data)
				// get the createdAt time of the last comment gotten, last comment is commented earliest
				if (response.data.length !== 0) {
					this.lastEarliestComment = [...response.data].pop().comment.createdAt
				} 

				if (response.data.length < 5) { this.outOfComments = true } 
			},
			newComment() {
				this.lastEarliestComment = -1
				this.comments = []
				this.outOfComments = false
			}
		},
	}
</script>

<style>
	#commentsView {
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