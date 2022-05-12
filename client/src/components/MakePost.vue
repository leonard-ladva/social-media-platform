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

				@blur="v$.content.$touch()"
				@focus="v$.content.$reset()"
			/>
			<hr>
		</div>
		<div class="form-group" id="bottom">
			<input type="submit" value="Post">
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
	},
	computed: {
		...mapGetters(['tags'])
	},
}
</script>

<style>
	#postForm {
		border: 0.1rem solid rgb(239, 243, 244);
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
		font-family: Chirp Heavy;
		border: none;
		display: inline-block;
		padding: .35em .65em;
		font-size: .75em;
		font-weight: 700;
		line-height: 1;
		color: #fff;
		text-align: center;
		white-space: nowrap;
		vertical-align: baseline;
		border-radius: .25rem;
	}
</style>