<script type="text/javascript">
	let opened = false;
	export let list;
	export let date = false;
	export let blue = false;
	export let current;
	$: current = list.length > 0 ? list[0] : "";

	function pre(v) {
		if (date) {
			let value = new Date(v);
			return value.toLocaleDateString("ru");
		}
		return v;
	}
</script>

<div
	class="select"
	class:select--opened={opened}
	class:select--blue={blue}
	on:click={() => {
		opened = !opened;
	}}
>
	<div class="select__header">
		<span class="select__current">{pre(current)}</span>
		<svg
			width="8"
			height="5"
			viewBox="0 0 8 5"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
		>
			<path
				d="M4.02598 4.49891C4.10696 4.49387 4.18393 4.46189 4.24463 4.40804L7.87929 1.13685C7.91507 1.10483 7.94418 1.06606 7.96495 1.02276C7.98571 0.979471 7.99773 0.93251 8.00032 0.884562C8.0029 0.836615 7.99599 0.788628 7.98 0.743354C7.964 0.698081 7.93923 0.656407 7.90709 0.620725C7.87496 0.585044 7.8361 0.556057 7.79275 0.535422C7.74939 0.514786 7.70239 0.502907 7.65443 0.50047C7.60648 0.498034 7.55851 0.505083 7.51329 0.521218C7.46806 0.537353 7.42647 0.562258 7.39088 0.594498L4.00042 3.64704L0.609962 0.594498C0.574378 0.562257 0.532782 0.537353 0.487558 0.521218C0.442333 0.505083 0.394368 0.498034 0.346413 0.50047C0.298458 0.502907 0.251455 0.51478 0.208098 0.535416C0.164742 0.556052 0.125884 0.585044 0.0937514 0.620725C0.0616192 0.656407 0.0368447 0.69808 0.0208476 0.743354C0.00485054 0.788628 -0.00205448 0.836609 0.000528325 0.884557C0.00311113 0.932504 0.015131 0.979471 0.0358989 1.02276C0.0566668 1.06606 0.0857746 1.10483 0.121554 1.13685L3.75622 4.40804C3.79276 4.44056 3.83548 4.46537 3.88183 4.48099C3.92819 4.4966 3.97722 4.50269 4.02598 4.49891Z"
				fill="#5C5C5C"
			/>
		</svg>
	</div>
	<div class="select__body">
		{#each list as item}
			<div
				class="select__item"
				on:click={() => {
					current = item;
				}}
			>
				{pre(item)}
			</div>
		{/each}
	</div>
</div>
