<script>
	import { setContext } from 'svelte'
	import views from './views'
	import './global.js'

	let view = $state(views['landing'])

	setContext('gotoView', (viewName) => {
		if (typeof viewName !== 'string') {
			throw new Error('View name must be a string')
		}

		const name = viewName.toLowerCase()
		const v = views[name]

		if (!v) {
			throw new Error(`Cannot goto unknown view '${viewName}'`)
		}

		view = v
	})
</script>

<div class="app">
	{#if view}
		{@render view()}
	{/if}
</div>

<style>
	.app {
		margin: 0;
		padding: 0;

		width: 100vw;
		height: 100vh;
		min-width: 100vw;
		min-height: 100vh;
		max-width: 100vw;
		max-height: 100vh;

		overflow: hidden;
	}
</style>
