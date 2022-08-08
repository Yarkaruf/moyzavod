<script type="text/javascript">
	import axios from "axios";
	import Viewer from "./Viewer.svelte";
	import Menu from "./Menu.svelte";
	import Nav from "./Nav.svelte";
	import Select from "./Select.svelte";

	let service = {};

	let service_types = [
		{
			name: "3D печать",
			description: `Oт прототипа до готовых деталей с использованием технологий FDM, SLS, SLA, PolyJet и Multi Jet Fusion.`,
			files: [".STL", ".STEP"],
		},
		{
			name: "Обработка с ЧПУ",
			description: `Прецизионная обработка деталей: фрезеровочные, токарные работы по металлу.`,
			files: [".STEP"],
		},
		{
			name: "Листовые материалы",
			description:
				"Лазерная резка и фрезеровка как металлов, так и других материалов: фанера, оргстекло, алюминиевый композит и др.",
			files: [".DXF"],
		},
		{
			name: "Литье в силикон",
			description:
				"Изготовление мелкосерийных партий товаров (1-1000 шт.) из пластмассы.",
			files: [".STEP"],
		},
	];

	let service_additionals = [
		{
			name: "3D моделирование",
			description:
				"Профессиональное создание любой 3D модели по эскизам или чертежам.",
			files: [".PDF", ".DOCX"],
		},
		{
			name: "3D сканирование",
			description:
				"Оцифровка трехмерного объекта с сохранением всех параметров.",
			files: [".PDF", ".DOCX"],
		},
		{
			name: "Постобработка",
			description:
				"Механическая, химическая, покрасочные и другие виды работ с вашим изделием.",
			files: [".PDF", ".DOCX"],
		},
	];

	let new_order = false;

	function closeNewOrder() {
		new_order = false;
	}

	let file_input;
	let order_file;

	function loadFile(ev) {
		file_input.click(ev);
	}

	function onFileChange(ev) {
		ev.target.files[0];
		order_file = ev.target.files[0];
		new_order = true;
	}

	function dropHandler(ev) {
		console.log("File(s) dropped");

		// Prevent default behavior (Prevent file from being opened)
		ev.preventDefault();

		if (ev.dataTransfer.items) {
			// Use DataTransferItemList interface to access the file(s)

			for (var i = 0; i < ev.dataTransfer.items.length; i++) {
				// If dropped items aren't files, reject them
				if (ev.dataTransfer.items[i].kind === "file") {
					var file = ev.dataTransfer.items[i].getAsFile();

					console.log("... file[" + i + "].name = " + file.name);
				}
			}
			order_file = ev.dataTransfer.items[0].getAsFile();
			new_order = true;
		} else {
			// Use DataTransfer interface to access the file(s)
			for (var i = 0; i < ev.dataTransfer.files.length; i++) {
				console.log(
					"... file[" + i + "].name = " + ev.dataTransfer.files[i].name
				);
			}

			order_file = ev.dataTransfer.files[0];
			new_order = true;
		}
	}

	function dragOverHandler(ev) {
		console.log("File(s) in drop zone");

		// Prevent default behavior (Prevent file from being opened)
		ev.preventDefault();
	}

	function checkAuth() {
		axios
			.get("/api/user")
			.then(function (resp) {})
			.catch(function (resp) {
				console.log(resp);
				window.location.href = "/login";
			});
	}
	checkAuth();

	function deauth() {
		axios
			.get("/api/deauth")
			.then(function () {
				checkAuth();
			})
			.catch(console.log);
	}

	let info;
	let size_box;
	let area;
	let volume;
	let preview;

	$: if (info) {
		let formData = new FormData();
		formData.append("model", order_file);
		axios
			.post("/api/calculate-model", formData)
			.then(function (response) {
				const model = response.data;
				area = (model.area / 100).toFixed(2) + "cm²";
				volume = (model.volume / 1000).toFixed(2) + "cm³";
				size_box = [
					model.bounds.x.toFixed(0) + "mm",
					model.bounds.y.toFixed(0) + "mm",
					model.bounds.z.toFixed(0) + "mm",
				].join(" x ");
			})
			.catch(function (response) {
				console.log(response);
			});
	}

	let params;

	let technologies = [];
	let technology;
	let plastics = [];
	let plastic;

	let quality;
	let filling;

	let comment;
	let count = 1;
	let color;
	let dates = [
		new Date(Date.now() + 1 * 24 * 3600 * 1000),
		new Date(Date.now() + 3 * 24 * 3600 * 1000),
		new Date(Date.now() + 7 * 24 * 3600 * 1000),
		new Date(Date.now() + 14 * 24 * 3600 * 1000),
	];
	let date;

	function calc(info, technology, plastic) {
		if (!params || !info || !technology || !plastic) {
			return 0;
		}

		let t = params.technologies[technology];
		let p = params.plastics[plastic];
		if (!t || !p) {
			return 0;
		}

		let density = 1.2;
		return (
			(info.volume / 1000) *
				density *
				Number.parseFloat(t.cost) *
				Number.parseFloat(t.quality) +
			Number.parseFloat(t.service)
		);
	}

	let cost;
	$: cost = calc(info, technology, plastic);

	function makeOrder() {
		let formData = new FormData();

		formData.append("model", order_file);
		formData.append("preview", preview);
		formData.append(
			"order",
			JSON.stringify({
				count,
				comment,
				cost,
				technology,
				plastic,
				color,
				quality,
				filling,
				date,
			})
		);
		axios
			.post("/api/create", formData, {
				headers: {
					"Content-Type": "multipart/form-data",
				},
			})
			.then(function () {
				window.location.href = "/configuration";
			})
			.catch(console.log);
	}

	function getParams() {
		axios.get("/api/params").then(function (response) {
			params = JSON.parse(response.data.params);
			Object.entries(params.technologies).forEach(([k, v]) =>
				technologies.push(k)
			);
			Object.entries(params.plastics).forEach(([k, v]) => plastics.push(k));
		});
	}
	getParams();
</script>

<div class="wrapper">
	<Nav />
	<div class="general">
		<Menu />
		<main class="main">
			<div class="block">
				<div class="container">
					<div class="unit">
						<div class="unit__header">
							<h2 class="block-title unit__title">Выберите услугу</h2>
						</div>
						<div class="unit__body">
							<div class="checklist">
								{#each service_types as service_type}
									<label class="custom-check checklist__item">
										<input
											type="radio"
											on:change={() => {
												service.service_type = service_type;
											}}
											class="radio radio--hidden"
											name="one-group"
										/>
										<div class="custom-check__wrap">
											<div class="custom-check__header">
												<span class="custom-check__select" />
												<div class="custom-check__info">
													<span class="custom-check__title"
														>{service_type.name}</span
													>
												</div>
											</div>
											<p class="custom-check__desc">
												{service_type.description}
											</p>
										</div>
									</label>
								{/each}
							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="content">
				<div class="container">
					<div class="content__wrap">
						<div class="content__block">
							<div class="unit">
								<div class="unit__header">
									<h2 class="block-title unit__title">Загрузите детали</h2>
									{#if service.service_type}<a
											href="#"
											class="button button--no-active"
											>{service.service_type.name}</a
										>{/if}
								</div>
								<div class="unit__body">
									<div
										on:drop={dropHandler}
										on:dragover={dragOverHandler}
										class="loading"
									>
										<div class="loading__info">
											<div class="loading__info-header">
												<h2 class="block-title unit__title">
													Перетащите модели и чертежи САПР или
												</h2>
												<a
													on:click={loadFile}
													class="button button--no-uppercase-default"
													>Выбрать файлы</a
												>
												<input
													on:change={onFileChange}
													style="display: none;"
													bind:this={file_input}
													multiple
													type="file"
													accept=".iges,.step,.igs,.stp,.stl,.docx,.pdf"
												/>
											</div>
											<div class="loading__info-body">
												<div class="loading__info-list">
													{#if service.service_type}
														{#each service.service_type.files as file}
															<div class="loading__info-item file__type">
																<div class="file__type-name">{file}</div>
																<svg
																	width="69"
																	height="84"
																	viewBox="0 0 69 84"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<g filter="url(#filter0_i_202_198)">
																		<path
																			d="M1 0H45.3478L69 23.5686V84H1V0Z"
																			fill="black"
																			fill-opacity="0.09"
																		/>
																		<path
																			d="M49 24H69L45 0V20C45 22.2091 46.7909 24 49 24Z"
																			fill="black"
																			fill-opacity="0.15"
																		/>
																	</g>
																	<defs>
																		<filter
																			id="filter0_i_202_198"
																			x="1"
																			y="-0.941043"
																			width="68"
																			height="84.941"
																			filterUnits="userSpaceOnUse"
																			color-interpolation-filters="sRGB"
																		>
																			<feFlood
																				flood-opacity="0"
																				result="BackgroundImageFix"
																			/>
																			<feBlend
																				mode="normal"
																				in="SourceGraphic"
																				in2="BackgroundImageFix"
																				result="shape"
																			/>
																			<feColorMatrix
																				in="SourceAlpha"
																				type="matrix"
																				values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
																				result="hardAlpha"
																			/>
																			<feOffset dy="-0.941043" />
																			<feGaussianBlur stdDeviation="2" />
																			<feComposite
																				in2="hardAlpha"
																				operator="arithmetic"
																				k2="-1"
																				k3="1"
																			/>
																			<feColorMatrix
																				type="matrix"
																				values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.15 0"
																			/>
																			<feBlend
																				mode="normal"
																				in2="shape"
																				result="effect1_innerShadow_202_198"
																			/>
																		</filter>
																	</defs>
																</svg>
															</div>
														{/each}
													{/if}
													{#if service.service_additional}
														{#each service.service_additional.files as file}
															<div class="loading__info-item file__type">
																<div class="file__type-name">{file}</div>
																<svg
																	width="69"
																	height="84"
																	viewBox="0 0 69 84"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<g filter="url(#filter0_i_202_198)">
																		<path
																			d="M1 0H45.3478L69 23.5686V84H1V0Z"
																			fill="black"
																			fill-opacity="0.09"
																		/>
																		<path
																			d="M49 24H69L45 0V20C45 22.2091 46.7909 24 49 24Z"
																			fill="black"
																			fill-opacity="0.15"
																		/>
																	</g>
																	<defs>
																		<filter
																			id="filter0_i_202_198"
																			x="1"
																			y="-0.941043"
																			width="68"
																			height="84.941"
																			filterUnits="userSpaceOnUse"
																			color-interpolation-filters="sRGB"
																		>
																			<feFlood
																				flood-opacity="0"
																				result="BackgroundImageFix"
																			/>
																			<feBlend
																				mode="normal"
																				in="SourceGraphic"
																				in2="BackgroundImageFix"
																				result="shape"
																			/>
																			<feColorMatrix
																				in="SourceAlpha"
																				type="matrix"
																				values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
																				result="hardAlpha"
																			/>
																			<feOffset dy="-0.941043" />
																			<feGaussianBlur stdDeviation="2" />
																			<feComposite
																				in2="hardAlpha"
																				operator="arithmetic"
																				k2="-1"
																				k3="1"
																			/>
																			<feColorMatrix
																				type="matrix"
																				values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.15 0"
																			/>
																			<feBlend
																				mode="normal"
																				in2="shape"
																				result="effect1_innerShadow_202_198"
																			/>
																		</filter>
																	</defs>
																</svg>
															</div>
														{/each}
													{/if}
												</div>
												<a href="#" class="button-link"
													>Показать больше поддерживаемых типов файлов</a
												>
											</div>
											<div class="afterword">
												<p class="afterword__text">
													Загруженные файлы находятся в безопасности и строго
													защищены,
													<a href="#" class="button-link afterword__link"
														>политика безопасности Мой Завод</a
													>
												</p>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
						<div class="content__block">
							<div class="unit">
								<div class="unit__header">
									<h2 class="block-title unit__title">Дополнительные услуги</h2>
								</div>
								<div class="unit__body">
									<div class="checklist checklist--column">
										{#each service_additionals as item}
											<label class="custom-check checklist__item">
												<input
													type="radio"
													on:change={() => {
														service.service_additional = item;
													}}
													class="radio radio--hidden"
													name="two-group"
												/>
												<div class="custom-check__wrap">
													<div class="custom-check__header">
														<span class="custom-check__select" />
														<div class="custom-check__info">
															<span class="custom-check__title"
																>{item.name}</span
															>
														</div>
													</div>
													<p class="custom-check__desc">
														{item.description}
													</p>
												</div>
											</label>
										{/each}
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</main>
		{#if new_order}
			<div id="popup" class="popup open">
				<div class="popup__body">
					<div class="popup__content">
						<div class="config">
							<div class="config__header">
								<div class="config__header-top">
									<span class="config__name">{order_file.name}</span>
									<div class="config__panel">
										<button class="config__button">
											<svg
												width="22"
												height="20"
												viewBox="0 0 22 20"
												fill="none"
												xmlns="http://www.w3.org/2000/svg"
											>
												<path
													d="M21 14.603C21 14.3409 20.7875 14.1284 20.5255 14.1284C20.2634 14.1284 20.0509 14.3409 20.0509 14.603V17.662C20.0509 17.9381 19.8271 18.162 19.5509 18.162H1.94907C1.67293 18.162 1.44907 17.9381 1.44907 17.662V14.603C1.44907 14.3409 1.23662 14.1284 0.974537 14.1284C0.712458 14.1284 0.5 14.3409 0.5 14.603V18.6111C0.5 18.8872 0.723858 19.1111 1 19.1111H20.5C20.7761 19.1111 21 18.8872 21 18.6111V14.603Z"
													fill="#333333"
													stroke="#333333"
													stroke-width="0.4"
												/>
												<path
													d="M9.4216 11.7563C9.73658 12.0713 10.2752 11.8482 10.2752 11.4027V1.36332C10.2752 1.10124 10.4876 0.88878 10.7497 0.88878C11.0118 0.88878 11.2242 1.10124 11.2242 1.36332V11.4027C11.2242 11.8482 11.7628 12.0713 12.0778 11.7563L14.9731 8.861C15.1565 8.67755 15.454 8.67755 15.6374 8.861C15.8209 9.04446 15.8209 9.3419 15.6374 9.52535L11.1032 14.0595C10.908 14.2548 10.5914 14.2548 10.3961 14.0595L5.86196 9.52535C5.67851 9.3419 5.67851 9.04446 5.86196 8.861C6.04542 8.67755 6.34286 8.67755 6.52631 8.861L9.4216 11.7563Z"
													fill="#333333"
													stroke="#333333"
													stroke-width="0.4"
												/>
											</svg>
										</button>
										<button on:click={closeNewOrder} class="config__button">
											<svg
												width="14"
												height="14"
												viewBox="0 0 14 14"
												fill="none"
												xmlns="http://www.w3.org/2000/svg"
											>
												<path
													d="M13.7704 12.6714C13.8431 12.7433 13.9009 12.829 13.9403 12.9235C13.9797 13.018 14 13.1194 14 13.2218C14 13.3242 13.9797 13.4256 13.9403 13.5201C13.9009 13.6145 13.8431 13.7002 13.7704 13.7722C13.6246 13.9181 13.427 14 13.2209 14C13.0148 14 12.8172 13.9181 12.6714 13.7722L7 8.09147L1.32861 13.7722C1.18281 13.9181 0.985164 14 0.779096 14C0.573028 14 0.375386 13.9181 0.229583 13.7722C0.156855 13.7002 0.0991129 13.6145 0.059703 13.5201C0.0202932 13.4256 0 13.3242 0 13.2218C0 13.1194 0.0202932 13.018 0.059703 12.9235C0.0991129 12.829 0.156855 12.7433 0.229583 12.6714L5.90097 6.99063L0.229583 1.30989C0.0922254 1.16223 0.0174469 0.966942 0.0210013 0.765155C0.0245558 0.563368 0.106166 0.37084 0.248638 0.228133C0.39111 0.0854261 0.58332 0.00368164 0.784775 0.000121349C0.98623 -0.00343895 1.1812 0.0714629 1.32861 0.209047L7 5.88979L12.6714 0.209047C12.8188 0.0714629 13.0138 -0.00343895 13.2152 0.000121349C13.4167 0.00368164 13.6089 0.0854261 13.7514 0.228133C13.8938 0.37084 13.9754 0.563368 13.979 0.765155C13.9826 0.966942 13.9078 1.16223 13.7704 1.30989L8.09903 6.99063L13.7704 12.6714Z"
													fill="#333333"
												/>
											</svg>
										</button>
									</div>
								</div>
								<div class="config__header-bottom">
									<ul class="config__list">
										<li class="config-item config-item--active">
											<button class="config-item__button">Конфигурация</button>
										</li>
										<li class="config-item">
											<button class="config-item__button">Анализ модели</button>
										</li>
									</ul>
								</div>
							</div>
							<div class="config__body">
								<div class="config__flex">
									<div class="config__work">
										<Viewer bind:file={order_file} bind:info bind:preview />
										<div class="config-info">
											<div class="config-info__list">
												<div class="config-info__item">
													<span class="config-info__category">Объем модели</span
													>
													<input
														type="text"
														class="config-info__input"
														bind:value={volume}
													/>
												</div>
												<div class="config-info__item">
													<span class="config-info__category"
														>Площадь поверхности</span
													>
													<input
														type="text"
														class="config-info__input"
														bind:value={area}
													/>
												</div>
												<div class="config-info__item">
													<span class="config-info__category">Габариты</span>
													<input
														type="text"
														class="config-info__input"
														bind:value={size_box}
													/>
												</div>
											</div>
										</div>
									</div>
									<div class="config__right">
										<div class="config-options">
											<form
												action=""
												method="post"
												class="config-options__form"
											>
												<div class="config-options__group">
													<span
														class="config-options__title config-options__title--no-border"
														>Процесс</span
													>
													<Select
														list={technologies}
														bind:current={technology}
													/>
												</div>
												<div class="config-options__group">
													<span class="config-options__title">Процесс</span>
													<div class="config-options__row">
														<span class="config-options__category"
															>Материал</span
														>
														{#if technology}<Select
																list={params.technologies[technology].plastics}
																bind:current={plastic}
															/>{/if}
													</div>
													<div class="config-options__row">
														<span class="config-options__category">Цвет</span>
														{#if plastic}<Select
																list={params.plastics[plastic].colors}
																bind:current={color}
															/>{/if}
													</div>
													<div class="config-options__row">
														<span class="config-options__category"
															>Качество</span
														>
														{#if technology}<Select
																list={params.technologies[technology].quality}
																bind:current={quality}
															/>{/if}
													</div>
													<div class="config-options__row">
														<span class="config-options__category"
															>Заполнение</span
														>
														{#if technology}<Select
																list={params.technologies[technology].filling}
																bind:current={filling}
															/>{/if}
													</div>
												</div>
												<div class="config-options__group">
													<span class="config-options__title"
														>Дополнительные услуги</span
													>
													<div class="config-options__row">
														<span class="config-options__category"
															>Шлифовка</span
														>
														<div class="select">
															<div class="select__header">
																<span class="select__current">Выбор опции</span>
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
																<div class="select__item">Value 1</div>
																<div class="select__item">Value 2</div>
																<div class="select__item">Value 3</div>
															</div>
														</div>
													</div>
													<div class="config-options__row">
														<span class="config-options__category"
															>Грунтовка</span
														>
														<div class="select">
															<div class="select__header">
																<span class="select__current">Выбор опции</span>
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
																<div class="select__item">Value 1</div>
																<div class="select__item">Value 2</div>
																<div class="select__item">Value 3</div>
															</div>
														</div>
													</div>
													<div class="config-options__row">
														<span class="config-options__category"
															>Покраска</span
														>
														<div class="select">
															<div class="select__header">
																<span class="select__current">Выбор опции</span>
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
																<div class="select__item">Value 1</div>
																<div class="select__item">Value 2</div>
																<div class="select__item">Value 3</div>
															</div>
														</div>
													</div>
												</div>
												<div class="config-options__group">
													<span class="config-options__comment"
														>Комментарий к заказу</span
													>
													<textarea
														class="config-options__textarea"
														name=""
														id=""
														cols="30"
														rows="10"
														bind:value={comment}
													/>
												</div>
												<div class="config-options__buttons">
													<a
														href="#"
														class="button button--no-uppercase-default"
														on:click={makeOrder}>Сохранить</a
													>
													<a
														href="#"
														class="button button--border button--no-uppercase-border"
														>Отмена</a
													>
												</div>
											</form>
										</div>
										<div class="config-price">
											<div class="config-price__row">
												<div class="config-price__column">
													<div class="config-price__item">
														<span class="config-price__category"
															>Срок изготовления</span
														>
														<Select list={dates} date bind:current={date} />
													</div>
												</div>
												<div class="config-price__column">
													<div class="config-price__item">
														<span class="config-price__category"
															>Количество</span
														>
														<div class="counter">
															<button
																class="counter__button counter__button--left"
																on:click={() => {
																	count -= 1;
																	count = count < 1 ? 1 : count;
																}}
															>
																<svg
																	width="10"
																	height="2"
																	viewBox="0 0 10 2"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<path
																		d="M4.61538 0.5H0.416667C0.186548 0.5 0 0.679087 0 0.9C0 1.12091 0.186548 1.3 0.416667 1.3H4.61538H5.44872H9.58333C9.81345 1.3 10 1.12091 10 0.9C10 0.679087 9.81345 0.5 9.58333 0.5H5.44872H4.61538Z"
																		fill="#5C5C5C"
																	/>
																</svg>
															</button>
															<input
																type="number"
																class="counter__number"
																bind:value={count}
																min="0"
																max="10000"
															/>
															<button
																class="counter__button counter__button--right"
																on:click={() => {
																	count += 1;
																}}
															>
																<svg
																	width="10"
																	height="11"
																	viewBox="0 0 10 11"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<path
																		d="M5.44872 0.916667C5.44872 0.686548 5.26217 0.5 5.03205 0.5C4.80193 0.5 4.61538 0.686548 4.61538 0.916667V5.05128H0.416667C0.186548 5.05128 0 5.23783 0 5.46795C0 5.69807 0.186548 5.88461 0.416667 5.88461H4.61538V10.0833C4.61538 10.3135 4.80193 10.5 5.03205 10.5C5.26217 10.5 5.44872 10.3135 5.44872 10.0833V5.88461H9.58333C9.81345 5.88461 10 5.69807 10 5.46795C10 5.23783 9.81345 5.05128 9.58333 5.05128H5.44872V0.916667Z"
																		fill="#5C5C5C"
																	/>
																</svg>
															</button>
														</div>
													</div>
												</div>
											</div>
											<div class="config-price__row">
												<div class="config-price__column">
													<div class="config-price__item">
														<span class="config-price__category"
															>Дата изготовления</span
														>
														<span class="config-price__date">21.11.2021</span>
													</div>
												</div>
												<div class="config-price__column">
													<div class="config-price__item">
														<span class="config-price__category">Цена</span>
														<div class="price">
															<span class="price__small"
																>1 шт: {cost.toFixed()} ₽</span
															>
															<span class="price__total"
																>Итого: {(cost * count).toFixed()} ₽</span
															>
														</div>
													</div>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>

<style type="text/css">
	.file__type {
		position: relative;
	}

	.file__type-name {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
	}
</style>
