<template>
	<form @submit.prevent="submitPost" id="postForm">
		<div class="form-group">
			<textarea
				type="text"
				v-model="content"
				class="form-control"
				placeholder="What's on your mind?"

				@blur="v$.content.$touch()"
				@focus="v$.content.$reset()"
			/>
			<p
				v-for="error of v$.content.$errors"
				:key="error.$uid"
			>
				<strong>{{ error.$message }}</strong>
			</p>
		</div>

		<label for="validationTags">Choose a Tag</label>
		<select 
			class="form-select"
			id="tags" 
			name="tags" 
			v-model="tag" 
			data-allow-new="true"
			data-allow-clear="Hello"

			@blur="v$.tag.$touch()"
			@focus="v$.tag.$reset()"
			>
			<option disabled hidden value="">Choose a tag...</option>
            <option value="Just chilling" selected="selected" >Just chilling</option>
			<option v-for="tag of tags" :key="tag.ID" :value="tag.Title">{{ tag.Title }}</option>
        </select>
		<p 
			v-for="error of v$.tag.$errors"
			:key="error.$uid"
			>
			<strong>{{ error.$message }}</strong>		
		</p>
        <div class="invalid-feedback">Please select a valid tag.</div>

		<input type="submit" value="Submit">
			
	</form>
</template>

<script>
import useVuelidate from '@vuelidate/core'
import { required, maxLength, helpers } from '@vuelidate/validators'
import axios from '../plugins/axios'
import { mapGetters } from 'vuex'

const printableChars = helpers.regex(/[ -~]/)

export default {
	name: 'MakePost',
	setup: () => ({ v$: useVuelidate() }),
	data() {
		return {
			content: '',
			tag: 'Just chilling',
			userId: '',
		}
	},
	validations() {
		return {
			content: { required, printableChars },
			tag: { required, maxLength: maxLength(50), printableChars },
		}
	},
	methods: {
		async submitPost() {
			const isFormCorrect = await this.v$.$validate()
			if (!isFormCorrect) return
			
			// this.post.userId = this.$store.state.user.Id
			const data = {
				tag: this.tag,
				content: this.content,
				userId: this.$store.state.user.id,
			}
			const response = await axios.post('submitPost', data)
			console.log(response)
		},
	},
	computed: {
		...mapGetters(['tags'])
	},
	async created() {
		const response = await axios.get('tags')
		this.$store.dispatch('tags', response.data)
	}
}
</script>

<style>
	#postForm {
		border: 0.1rem solid rgba(220, 220, 220, 1);
		border-top: none;
	}
</style>