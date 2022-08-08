<script type="text/javascript">
	import Nav from "./Nav.svelte";
	import Menu from "./Menu.svelte";
	import axios from "axios";
	import { onMount } from "svelte";
	let technologies = {
		FDM: {
			plastics: ["ABS", "PETG"],
			quality: [0.1, 0.2],
			cost: [1.5],
			filling: [0.1, 0.2, 0.5, 1],
			service: [100],
		},
		SLA: {
			plastics: ["HIPS", "ABS"],
			quality: [0.1, 0.2],
			filling: [0.1, 0.2, 0.5, 1],
			cost: [1.5],
			service: [100],
		},
	};

	let plastics = {
		ABS: {
			colors: ["Желтый", "Белый", "Черный"],
			cost: [1.2],
		},
		PETG: {
			colors: ["Желтый", "Белый", "Черный"],
			cost: [1.2],
		},
		HIPS: {
			colors: ["Желтый", "Белый", "Черный"],
			cost: [1.2],
		},
	};

	let params = {
		plastics,
		technologies,
	};

	function getParams() {
		axios.get("/api/params").then(function (response) {
			params = JSON.parse(response.data);
		});
	}
	getParams();

	function save() {
		axios
			.post("/api/params", {
				params: JSON.stringify(params),
			})
			.then(console.log)
			.catch(console.log);
	}
</script>

<div class="wrapper">
	<Nav />
	<div class="general">
		<Menu />
		<main class="main">
			<h1>3D печать</h1>
			<table>
				<tbody>
					<tr>
						<td><h2>Технология печати</h2></td>
					</tr>
					{#each Object.entries(params.technologies) as [key, value]}
						<tr>
							<td><h3>{key}</h3></td>
						</tr>
						{#each Object.entries(value) as [k, v]}
							{#if Array.isArray(v)}
								<tr>
									<td><h4>{k}</h4></td>
								</tr>
								{#each v as item, i}
									<tr>
										<td
											><input
												on:change={() => {
													console.log(plastics);
												}}
												bind:value={item}
											/><button
												on:click={() => {
													v.splice(i, 1);
													plastics = { ...plastics };
												}}>-</button
											></td
										>
									</tr>
								{/each}
							{/if}
							<tr>
								<td
									><input type="text" bind:value={v.new_item} /><button
										on:click={() => {
											v.push(v.new_item);
											plastics = { ...plastics };
										}}>+</button
									></td
								>
							</tr>
						{/each}
					{/each}
				</tbody>
			</table>

			<table>
				<tbody>
					<tr>
						<td><h2>Пластик</h2></td>
					</tr>
					{#each Object.entries(params.plastics) as [key, value]}
						<tr>
							<td><h3>{key}</h3></td>
						</tr>
						{#each Object.entries(value) as [k, v]}
							{#if Array.isArray(v)}
								<tr>
									<td><h4>{k}</h4></td>
								</tr>
								{#each v as item, i}
									<tr>
										<td
											><input
												on:change={() => {
													console.log(plastics);
												}}
												bind:value={item}
											/><button
												on:click={() => {
													v.splice(i, 1);
													plastics = { ...plastics };
												}}>-</button
											></td
										>
									</tr>
								{/each}
							{/if}
							<tr>
								<td
									><input type="text" bind:value={v.new_item} /><button
										on:click={() => {
											v.push(v.new_item);
											plastics = { ...plastics };
										}}>+</button
									></td
								>
							</tr>
						{/each}
					{/each}
				</tbody>
			</table>
			<button on:click={save}>Save</button>
			Формула цена volume*technology_filling*technology_cost/quality (время работы
			принтера) + volume*technology_filling*plastic_cost (цена материала) + service_cost
		</main>
	</div>
</div>
