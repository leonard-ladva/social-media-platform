<template>
	<form @submit.prevent="submitForm">
		<BaseInput
			v-model="post.title"
			label="Title"
			type="text"
		/>

		<BaseInput
			v-model="post.description"
			label="Description"
			type="text"
		/>

		<label>Post Type:</label>
		<select v-model="post.category">
			<option
				v-for="option in categories"
				:value="option"
				:key="option"
				:selected="option === post.category"
			>{{ option }}</option>
		</select>

		<input type="submit" value="Submit">
	</form>
</template>

<style scoped>
	form {
		display: flex;
		flex-direction: column;
		width: 500px;
		margin: auto;
	}
</style>

<script>
import BaseInput from '@/components/BaseInput.vue'
export default {
	components: {BaseInput},
	data() {
		return {
			categories: [
				'story',
				'question',
				'discussion'
			],
			post: {
				category: '',
				title: '',
				description: '',
			}

		}
	},
	methods: {
		submitForm: function() {
			console.log(this.post)
			// send postData to backend 

			// empty form
			this.post.category = ''
			this.post.title = ''
			this.post.description = ''
			
			// route back to feed
			this.$router.push('/feed')
		
		}
	}
}
</script>