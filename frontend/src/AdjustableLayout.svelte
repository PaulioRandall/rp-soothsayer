<script>
	let { left, right } = $props()
</script>

<div class="adjustable-layout">
	<div class="left-side">
		{@render left?.()}
	</div>
	<div class="divider-line"></div>
	<div class="right-side">
		{@render right?.()}
	</div>
</div>

<style>
	.adjustable-layout {
		display: flex;
		height: 100%;
		width: 100%;
	}

	.left-side {
		flex-grow: 0;
		flex-shrink: 0;
		width: 60%;
	}

	.divider-line {
		flex: 0 0 10px;
		height: 100%;
		background: dimgrey;
		cursor: col-resize;
	}

	.right-side {
		flex-grow: 1;
		flex-shrink: 0;
	}
</style>
