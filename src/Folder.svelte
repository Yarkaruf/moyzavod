<script>
	import Menu from "./Menu.svelte";
	import Nav from "./Nav.svelte";
	import { Link } from "svelte-routing";
	import Util from "./util.js";

	let flag = false;
	let display = true;
	let popup = false;
	let popupCreate = false;
	let popupRename = false;

	Util.dropdownToggle();

	Util.select();
</script>

<div class="wrapper">
	<Nav />
	<div class="general">
		<Menu />
		<main class="main">
			<div class="page page--files">
				<div class="container">
					<div class="page__wrap">
						<div class="page__header">
							<h3 class="page-title">Мои файлы</h3>
							<div class="search search--transparent">
								<div class="field">
									<input
										type="text"
										name=""
										id=""
										class="field__input search__input"
										placeholder="Введите название файла или папки"
									/>
									<div class="search__icon">
										<svg
											width="16"
											height="16"
											viewBox="0 0 16 16"
											fill="none"
											xmlns="http://www.w3.org/2000/svg"
										>
											<path
												fill-rule="evenodd"
												clip-rule="evenodd"
												d="M11.1751 13.0451C10.0822 13.7837 8.76614 14.2147 7.34932 14.2147C3.56947 14.2147 0.5 11.1417 0.5 7.35737C0.5 3.57308 3.56947 0.5 7.34932 0.5C11.1292 0.5 14.1986 3.57308 14.1986 7.35737C14.1986 8.76607 13.773 10.0768 13.044 11.1671L15.1125 13.2371C15.6292 13.7543 15.6292 14.5948 15.1125 15.1121C14.5959 15.6293 13.7564 15.6293 13.2397 15.1121L11.1751 13.0451V13.0451ZM7.34931 2.92653C9.79256 2.92653 11.775 4.91125 11.775 7.35737C11.775 9.80349 9.79256 11.7882 7.34931 11.7882C4.90607 11.7882 2.92368 9.80349 2.92368 7.35737C2.92368 4.91125 4.90607 2.92653 7.34931 2.92653Z"
												fill="#ADADAD"
											/>
										</svg>
									</div>
								</div>
							</div>
						</div>
						<div class="page__body">
							<div class="files files--modal">
								<div class="files__overlay" class:is-active={popup}>
									<div class="files-modal" class:is-active={popupRename}>
										<span class="files-modal__title">Переименовать</span>
										<div class="field files-modal__field">
											<input
												class="field__input"
												type="text"
												placeholder="Название папки"
											/>
										</div>
										<div class="files-modal__buttons">
											<a
												href="#"
												class="button button--no-uppercase-default"
												on:click={() => (popupRename = false)}
												on:click={() => (popup = false)}>Сохранить</a
											>

											<a
												href="#"
												class="button button--border button--no-uppercase-border"
												on:click={() => (popup = false)}
												on:click={() => (popupRename = false)}>Отменить</a
											>
										</div>
									</div>
									<div class="files-modal" class:is-active={popupCreate}>
										<span class="files-modal__title"
											>Укажите название папки</span
										>
										<div class="field files-modal__field">
											<input
												class="field__input"
												type="text"
												placeholder="Название папки"
											/>
										</div>
										<div class="files-modal__buttons">
											<a
												href="#"
												class="button button--no-uppercase-default"
												on:click={() => (popup = false)}
												on:click={() => (popupCreate = false)}
												>Создать
											</a>

											<a
												href="#"
												class="button button--border button--no-uppercase-border"
												on:click={() => (popup = false)}
												on:click={() => (popupCreate = false)}
											>
												Отменить</a
											>
										</div>
									</div>
								</div>
								<div class="files__header flag-1" class:is-active={!flag}>
									<Link to="/files">
										<a href="#" class="files-back">
											<svg
												width="11"
												height="9"
												viewBox="0 0 11 9"
												fill="none"
												xmlns="http://www.w3.org/2000/svg"
											>
												<path
													fill-rule="evenodd"
													clip-rule="evenodd"
													d="M0.304 3.78072L3.344 0.741719C3.734 0.351719 4.367 0.351719 4.758 0.741719C5.148 1.13172 5.148 1.76572 4.758 2.15572L3.414 3.49972H9.051C9.603 3.49972 10.051 3.94772 10.051 4.49972C10.051 5.05172 9.603 5.49972 9.051 5.49972H3.414L4.758 6.84372C5.148 7.23372 5.148 7.86772 4.758 8.25772C4.367 8.64772 3.734 8.64772 3.344 8.25772L0.304 5.21872C0.301 5.21472 0.297 5.21072 0.293 5.20672C0.098 5.01172 0 4.75572 0 4.49972C0 4.24372 0.098 3.98772 0.293 3.79272C0.297 3.78872 0.301 3.78472 0.304 3.78072Z"
													fill="#5C5C5C"
												/>
											</svg>
											<span class="files-back__text"
												>Название папки-родителя</span
											>
										</a>
									</Link>
									<div class="files__header-wrap">
										<div class="files__header-left">
											<div
												class="select select--transparent files__header-sort"
											>
												<div class="select__header" on:click={Util.select()}>
													<svg
														width="18"
														height="16"
														viewBox="0 0 18 16"
														fill="none"
														xmlns="http://www.w3.org/2000/svg"
													>
														<rect
															y="6.5"
															width="15"
															height="3"
															fill="#5C5C5C"
														/>
														<rect width="12" height="3" fill="#5C5C5C" />
														<rect y="13" width="18" height="3" fill="#5C5C5C" />
													</svg>
													<span
														>По <span class="select__current">названию</span
														></span
													>
													<svg
														width="8"
														height="5"
														class="select__icon"
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
													<div class="select__item">Названию</div>
													<div class="select__item">Размеру</div>
													<div class="select__item">Типу</div>
													<div class="select__item">Дате изменения</div>
												</div>
											</div>
											<div class="files__header-buttons">
												<button
													class="files__header-button"
													on:click={() => (display = false)}
												>
													<svg
														width="18"
														height="18"
														viewBox="0 0 18 18"
														fill="none"
														xmlns="http://www.w3.org/2000/svg"
													>
														<rect width="8" height="8" fill="#D6D6D6" />
														<rect y="10" width="8" height="8" fill="#D6D6D6" />
														<rect x="10" width="8" height="8" fill="#D6D6D6" />
														<rect
															x="10"
															y="10"
															width="8"
															height="8"
															fill="#D6D6D6"
														/>
													</svg>
												</button>
												<button
													class="files__header-button"
													on:click={() => (display = true)}
												>
													<svg
														width="18"
														height="16"
														viewBox="0 0 18 16"
														fill="none"
														xmlns="http://www.w3.org/2000/svg"
													>
														<rect width="5" height="3" fill="#5C5C5C" />
														<rect y="6.5" width="5" height="3" fill="#5C5C5C" />
														<rect
															x="6"
															y="6.5"
															width="12"
															height="3"
															fill="#5C5C5C"
														/>
														<rect x="6" width="12" height="3" fill="#5C5C5C" />
														<rect
															x="6"
															y="13"
															width="12"
															height="3"
															fill="#5C5C5C"
														/>
														<rect y="13" width="5" height="3" fill="#5C5C5C" />
													</svg>
												</button>
											</div>
										</div>
										<div class="files__header-right">
											<a
												href="#"
												class="button button--border button--no-uppercase-border"
												on:click={() => ((popup = true), (popupCreate = true))}
												>Новая папка</a
											>

											<a
												href="#"
												class="button button--no-uppercase-default"
												on:click={() => (popupRename = true)}
												on:click={() => (popup = true)}>Загрузить файл</a
											>
										</div>
									</div>
								</div>
								<div class="files__header flag-2" class:is-active={flag}>
									<div class="files__header-wrap">
										<a href="#" class="button button--no-uppercase-default"
											>Отправить выбранные файлы на расчет</a
										>
									</div>
								</div>
								<div class="files__body display-1" class:is-active={!display}>
									<ul class="files-list">
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
										<li class="files-list__item">
											<div class="files-list__name">
												<div class="files-list__name-icon">
													<img
														src="./images/content/files-list-img1.png"
														alt=""
													/>
												</div>
												<span class="files-list__name-text">Название файла</span
												>
											</div>
										</li>
									</ul>
								</div>
								<div class="files__body display-2" class:is-active={display}>
									<div class="files-table files-table--small-padding">
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input
															type="checkbox"
															id="checkbox1"
															bind:checked={flag}
														/>
														<label for="checkbox1" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button
														class="files-table__more"
														on:click={Util.dropdownToggle()}
													>
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox2" />
														<label for="checkbox2" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox3" />
														<label for="checkbox3" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox4" />
														<label for="checkbox4" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox5" />
														<label for="checkbox5" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox6" />
														<label for="checkbox6" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox7" />
														<label for="checkbox7" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox8" />
														<label for="checkbox8" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
																>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="files-table__row">
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="checkbox checkbox--border">
														<input type="checkbox" id="checkbox9" />
														<label for="checkbox9" />
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<div class="files-table__name">
														<div class="files-table__name-icon">
															<img
																src="./images/content/files-table-img1.png"
																alt=""
															/>
														</div>
														<span class="files-table__text"
															>Название файла.stl</span
														>
													</div>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">27 мб</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<span class="files-table__text">07.12.2021</span>
												</div>
											</div>
											<div class="files-table__column">
												<div class="files-table__cell">
													<button class="files-table__more">
														<svg
															width="24"
															height="24"
															viewBox="0 0 24 24"
															fill="none"
															xmlns="http://www.w3.org/2000/svg"
														>
															<circle cx="5" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="12" cy="12" r="2" fill="#5C5C5C" />
															<circle cx="19" cy="12" r="2" fill="#5C5C5C" />
														</svg>
													</button>
													<div class="files-table__dropdown">
														<div class="files-table__dropdown-list">
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Переименовать</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Удалить</a
																>
															</div>
															<div class="files-table__dropdown-item">
																<a href="#" class="files-table__dropdown-link"
																	>Скачать</a
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
			</div>
		</main>
	</div>
</div>
