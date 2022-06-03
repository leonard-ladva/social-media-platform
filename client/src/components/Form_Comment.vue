<template>
	<form @submit.prevent="submitComment()" id="commentForm">
		<div class="form-group" id="middle">
			<textarea
				id="content"
				type="text"
				v-model="content"
				class="form-control"
				placeholder="What do you think about this?"

				@input="resizeTextArea()"
				ref="content"

				@blur="v$.content.$touch()"
				@focus="v$.content.$reset()"
			/>
			<button value="Comment" id="submit" :disabled="!readyToComment">Reply</button>
		</div>
		<div class="form-group" id="bottom">
			<span v-if="error" class="error badge bg-secondary">Something went wrong</span>
		</div>
	</form>
</template>

<script>
import useVuelidate from '@vuelidate/core'
import { required, maxLength, helpers } from '@vuelidate/validators'
import axios from '../plugins/axios'
import { directive as VueInputAutowidth } from "vue-input-autowidth"

const printableChars = helpers.regex(/[ -~]/)

export default {
	name: 'CreateComment',
	directives: { autowidth: VueInputAutowidth },
	setup: () => ({ v$: useVuelidate() }),
	data() {
		return {
			content: '',
			error: false,
		}
	},
	validations() {
		return {
			content: { required, maxLength: maxLength(500), printableChars },
		}
	},
	methods: {
		async submitComment() {
			const isFormCorrect = await this.v$.$validate()
			if (!isFormCorrect) return
			
			const data = {
				content: this.content,
				userId: this.$store.state.currentUser.id,
				postId: this.$route.params.postId,
			}
			axios.post('submitComment', data)
			.then(({status}) => {
				if (status === 200) {
					this.$emit('newComment')
					this.content = ''
				}
			})
			.catch(() => {
				this.error = true
			})
		},
		resizeTextArea() {
			const element = this.$refs.content
			element.style.height = "auto"
			element.style.height = (element.scrollHeight + "px")
		}
	},
	computed: {
		readyToComment() {
			const validContent = this.content.length > 0 && this.content.length < 501
			return validContent
		},
	},
}
</script>

<style>
	#commentForm {
		padding: 1rem 1.5rem;
	}
	#middle {
		display: flex;
		align-items: flex-start;
	}
	textarea#content {
		border: none;
		resize: none;
		font-size: 1.3rem;
	}
	textarea#content:focus, input.tag:focus {
		outline: none;
		box-shadow: none;
	}

	#submit {
		background-color: var(--blue);
		border: none;
		border-radius: 2rem;
		padding: 0.4rem 0.7rem;
		font-size: 1rem;
		font-family: "Chirp Bold";
		color: var(--white);
		margin-top: 0.3rem;
		height: auto;
	}
	#submit:disabled {
		background-color: var(--lightGrey)
	}

</style>