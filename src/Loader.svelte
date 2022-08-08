<div style="position: relative;">
	<div class="loader" />
	<div class="loader loader-blue" />
</div>

<style type="text/css">
	.loader {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		width: 2.5em;
		height: 2.5em;
		animation-duration: 2s;
		border: 0.3em solid grey;
		animation-name: spin;
		animation-iteration-count: infinite;
	}

	.loader-blue {
		border: 0.3em solid #6e96f8;
		animation-duration: 2.5s;
	}

	.loader-pink {
		border: 0.3em solid #ecd4f9;
		animation-duration: 3s;
	}
	@keyframes spin {
		from {
			transform: rotate(0deg);
		}

		to {
			transform: rotate(360deg);
		}
	}
</style>
