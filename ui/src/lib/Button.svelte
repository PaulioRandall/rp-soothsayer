<script>
	import { getContext } from 'svelte'

	const gotoView = getContext('gotoView')

	let { children, ...rest } = $props()

	const gotoFileBrowserView = () => gotoView('FileBrowser')
</script>

<button onclick={gotoFileBrowserView} {...rest}>
	{@render children?.()}
</button>

<style>
	button {
		padding: 0.25rem 0.5rem;

		background: var(--color-grey-1);
		border: 2px solid var(--color-grey-3);
		border-radius: 8px;

		font-weight: 500;

		cursor: pointer;
	}

	button:hover {
		background: #dddddd;
	}
</style>
