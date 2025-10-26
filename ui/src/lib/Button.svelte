<script>
	import { getContext } from 'svelte'

	const gotoView = getContext('gotoView')

	let { children, ...rest } = $props()

	const gotoFileBrowserView = () => gotoView("FileBrowser")
</script>

<button onclick={gotoFileBrowserView} {...rest}>
	{@render children?.()}
</button>

<style>
	button {
		padding: 1rem;

		background: var(--color-grey-2);
		border: 2px solid var(--color-grey-3);
		border-radius: 8px;

		cursor: pointer;
	}
</style>
