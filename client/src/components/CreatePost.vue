<template>
	<form @submit.prevent="submitPost" id="postForm">
		<div class="form-group" id="top">
			<input 
				class="tag" 
				type="text" 
				v-model="tag" 
				v-autowidth="{comfortZone: '0.5rem', minWidth: '4rem'}" 
				aria-disabled="true"
				:style="tagLengthValid"
			/>
		</div>
		<div class="form-group" id="middle">
			<textarea
				id="content"
				type="text"
				v-model="content"
				class="form-control"
				placeholder="What's in your noggin?"

				@input="resizeTextArea()"
				ref="content"

				@blur="v$.content.$touch()"
				@focus="v$.content.$reset()"
			/>
			<hr>
		</div>
		<span v-if="error" class="error badge bg-secondary">Something went wrong</span>
		<div class="form-group" id="bottom">
			<button value="Post" id="submit" :disabled="!readyToPost">Post</button>
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
	name: 'CreatePost',
	directives: { autowidth: VueInputAutowidth },
	setup: () => ({ v$: useVuelidate() }),
	data() {
		return {
			content: '',
			tag: '#NoCategory',
			userId: '',
			error: false,
		}
	},
	validations() {
		return {
			content: { required, maxLength: maxLength(500), printableChars },
			tag: { required, maxLength: maxLength(50), printableChars },
		}
	},
	methods: {
		async submitPost() {
			const isFormCorrect = await this.v$.$validate()
			if (!isFormCorrect) return
			
			const data = {
				tag: this.tag,
				content: this.content,
				userId: this.$store.state.currentUser.id,
			}
			axios.post('submitPost', data)
			.then(({status}) => {
				if (status === 200) {
					this.$emit('newPost')
					this.content = ''
					this.tag = '#NoCategory'
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
		readyToPost() {
			const validTag = this.tag.length > 0 && this.tag.length < 51
			const validContent = this.content.length > 0 && this.content.length < 501
			return validTag && validContent
		},
		tagLengthValid() {
			return `background-color: ${this.tag.length > 0 && this.tag.length < 51 ? 'var(--blue)' : 'red'}`
		}
	},
}
</script>

<style>
	#postForm {
		padding: 1rem 1.5rem;
	}

	#top, #bottom {
		text-align: right;		
	}
	textarea#content {
		border: none;
		resize: none;
		font-size: 1.5rem;
	}
	textarea#content:focus, input.tag:focus {
		outline: none;
		box-shadow: none;
	}

	input.tag {
		font-family: "Chirp Heavy";
		border: none;
		display: inline-block;
		padding: .35em .65em;
		font-size: .75em;
		font-weight: 700;
		line-height: 1;
		color: var(--white);
		text-align: center;
		white-space: nowrap;
		vertical-align: baseline;
		border-radius: .25rem;
	}
	#submit {
		background-color: var(--blue);
		border: none;
		border-radius: 2rem;
		padding: 0.4rem 0.7rem;
		font-size: 1rem;
		font-family: "Chirp Bold";
		color: var(--white);
		margin: 0 0.6rem 0 0.3rem;
	}
	#submit:disabled {
		background-color: var(--lightGrey)
	}

</style>