<script>
	let { value = $bindable(''), ...rest } = $props()
</script>

<input type="text" bind:value {...rest} />

<style>
	input {
		font-size: 120%;

		margin: 0;
		padding: 0.25rem;

		border: 1px solid black;
		border-radius: 4px;
	}
</style>
