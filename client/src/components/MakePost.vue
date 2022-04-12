<template>
	<form @submit.prevent="submitPost">
		<div class="form-group">
			<textarea type="text" v-model="post.content" class="form-control" placeholder="What's on your mind?"/>
		</div>

		<label>Tags:</label>
		<select v-model="post.tag">
			<option
				v-for="option in tags"
				:value="option"
				:key="option"
				:selected="option === post.tag"
			>{{ option }}</option>
		</select>

		<input type="submit" value="Submit">
	</form>
</template>

<script>
import axios from 'axios'
export default {
	name: 'MakePost',
	data() {
		return {
			tags: [
				'f1',
				'drawing',
				'politics',
				'coding',
				'entrepreneurship',
			],
			post: {
				tag: '',
				content: '',
			}

		}
	},
	methods: {
		submitPost: function() {
			const data = this.post
			console.log(data)

			axios.post('http://localhost:9100/submitpost', data)
			.then(res => {
				console.log(res)
			})
			.catch(err => {
				console.log(err)
			})
		}
	}
}
</script>