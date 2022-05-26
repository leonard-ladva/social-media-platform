<template>
	<form @submit.prevent="submitPost" id="postForm">
		<div class="form-group" id="top">
			<input class="tag" type="text" v-model="tag" v-autowidth="{comfortZone: '0.3rem', minWidth: '4rem'}" aria-disabled="true"/>
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
		<div class="form-group" id="bottom">
			<input type="submit" value="Post" id="submit">
		</div>
	</form>
</template>

<script>
import useVuelidate from '@vuelidate/core'
import { required, maxLength, helpers } from '@vuelidate/validators'
import axios from '../plugins/axios'
import { mapGetters } from 'vuex'
import { directive as VueInputAutowidth } from "vue-input-autowidth"


const printableChars = helpers.regex(/[ -~]/)

export default {
	name: 'MakePost',
	directives: { autowidth: VueInputAutowidth },
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
			
			const data = {
				tag: this.tag,
				content: this.content,
				userId: this.$store.state.user.id,
			}
			const response = await axios.post('submitPost', data)
			console.log(response)
		},
		resizeTextArea() {
			const element = this.$refs.content
			element.style.height = "auto"
			element.style.height = (element.scrollHeight + "px")
		}
	},
	computed: {
		...mapGetters(['tags'])
	},
}
</script>

<style>
	#postForm {
		border: 0.1rem solid var(--extraExtraLightGrey);
		border-top: none;
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

</style>