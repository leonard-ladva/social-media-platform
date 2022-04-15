<template>
	<div class="form-group">
		<label>{{ label }}</label>
		<input class="form-control"
		:type="type"
		:v-model="something"
		:@blur="blurValue"
		placeholder="{{ label }}"/>
		<p
			:v-for="errorLoop"
			:key="error.$uid"
		>
			<strong>{{ error.$message }}</strong>
		</p>
	</div>	
</template>

<script>
export default {
	name: "InputField",
	props: {
		label: {
			type: String,
			required: true,
		},
		field: {
			type: String,
			required: true,
		},
		type: {
			type: String,
			required: true,
		}
	},
	computed: {
		blurValue () {
			return "v$." + this.field + ".$touch"
		},
		errorLoop () {
			return "error of v$." + this.field + ".$errors"
		}
	}
}

</script>