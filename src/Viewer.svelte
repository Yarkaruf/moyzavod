<script type="text/javascript">
	import { onMount, onDestroy } from "svelte";
	import Loader from "./Loader.svelte";
	export let file;
	export let info;
	export let preview;
	let info_input;
	let viewport;
	let file_input;
	onMount(function () {
		let interval = setInterval(function () {
			if (file_input.ModelViewer) {
				file_input.ModelViewer(file);
				clearInterval(interval);
			}
		}, 200);

		let intervalResult = setInterval(function () {
			if (info_input.info) {
				info = info_input.info;
				clearInterval(intervalResult);
				setTimeout(function () {
					viewport.children[0].toBlob(function (blob) {
						preview = new File([blob], "preview.png");
					}, "image/png");
				}, 1000);
			}
		}, 200);
	});

	onDestroy(function () {
		info = null;
	});
</script>

<div style="height: 100%; position: relative;">
	<div
		style="position: absolute; top:50%; left: 50%; transform: translate(-50%,-50%);"
	/>
	<script src="/js/viewer5.js"></script>
	<div id="menu" style="display: none;">
		<input
			bind:this={file_input}
			id="step-file"
			name="step-file"
			type="file"
			accept=".iges,.step,.igs,.stp,.stl"
		/>
		<div><input id="support" type="checkbox" name="support" />support</div>
		<div>
			<input
				id="angle"
				type="number"
				name="angle"
				min="0"
				max="90"
				style="width: 50px;"
			/>
			angle <button id="recalculate">recalculate</button>
		</div>
		<div bind:this={info_input} id="info" style="font-size: 0.75em;" />
	</div>
	<div style=" position: relative; height: 100%;">
		<div
			id="viewport"
			bind:this={viewport}
			style="height: 100%;"
			align="centers"
		/>
		<div
			style="position: absolute; top:50%; left: 50%; transform: translate(-50%;-50%);"
		>
			{#if !info}<Loader />{/if}
		</div>
	</div>
	<div
		style="position: absolute; top:100px; right:50px; width: 200px; height: 200px;"
		id="helper"
	>
		<button class="arrow button--up" id="upView">❮</button>
		<button class="arrow button--down" id="downView">❮</button>
		<button class="arrow button--left" id="leftView">❮</button>
		<button class="arrow button--right" id="rightView">❮</button>
	</div>
</div>

<style>
	.noselect {
		-webkit-touch-callout: none; /* iOS Safari */
		-webkit-user-select: none; /* Safari */
		-khtml-user-select: none; /* Konqueror HTML */
		-moz-user-select: none; /* Firefox */
		-ms-user-select: none; /* Internet Explorer/Edge */
		user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
	}
	body {
		margin: 0;
		overflow: hidden;
		height: 100vh;
		display: flex;
		flex-direction: column;
	}
	#menu {
		flex-basis: 0;
		display: flex;
		flex-direction: row;
		justify-content: space-evenly;
	}
	#viewport {
		flex-basis: 0;
		flex-grow: 1;
	}

	.arrow {
		position: absolute;
		font-size: 30px;
		padding: 0;
		line-height: 30px;
		margin: 0;
		width: 50px;
		height: 50px;
		background: transparent;
		border: none;
	}
	.button--up {
		top: 0%;
		left: 50%;
		transform: translateX(-50%) translateY(-50%) rotate(90deg);
	}
	.button--down {
		top: 100%;
		left: 50%;
		transform: translateX(-50%) translateY(-50%) rotate(-90deg);
	}
	.button--left {
		top: 50%;
		left: 0%;
		transform: translateX(-50%) translateY(-50%) rotate(0deg);
	}

	.button--right {
		top: 50%;
		left: 100%;
		transform: translateX(-50%) translateY(-50%) rotate(180deg);
	}

	label {
		margin-right: 10px;
	}
</style>
