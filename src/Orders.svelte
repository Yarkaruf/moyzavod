<script type="text/javascript">
	import Menu from "./Menu.svelte";
	import Nav from "./Nav.svelte";
	let tab = true;
</script>

<div class="wrapper">
	<Nav orders />
	<div class="general">
		<Menu />
		<main class="main">
			<div class="page page--offer">
				<div class="container">
					<div class="page__wrap">
						<div class="page__header">
							<h3 class="page-title">Мои заказы</h3>
						</div>
						<div class="page__body">
							<ul class="tabs-nav tabs-nav--offer">
								<li
									class="tabs-nav__item "
									class:tabs-nav__item--active={tab}
									on:click={() => {
										tab = true;
									}}
									data-tab-name="tab-1"
								>
									Активные заказы
								</li>
								<li
									class="tabs-nav__item"
									data-tab-name="tab-2"
									class:tabs-nav__item--active={!tab}
									on:click={() => {
										tab = false;
									}}
								>
									Выполненные заказы
								</li>
							</ul>
							<div class="tabs-content">
								<div class="tabs-content__item tab-1" class:is-active={tab}>
									<div class="substance">
										<div class="offer">
											<div class="offer__content">
												<div class="offer__left">
													<div class="offer-prev">
														<ul class="offer-prev__list">
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
														</ul>
													</div>
												</div>
												<div class="offer__right">
													<div class="offer-status">
														<span class="offer-status__text"
															>На рассмотрении</span
														>
													</div>
													<ul class="info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Номер заказа:</span
															>
															<span class="info-list__value">5</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Адрес доставки:</span
															>
															<span class="info-list__value"
																>Москва, Мира 1а, квартира 15</span
															>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Способ доставки:</span
															>
															<span class="info-list__value">Почта РФ</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Дата доставки:</span
															>
															<span class="info-list__value">21.11.2021</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Цена доставки:</span
															>
															<span class="info-list__value">562 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
													<div class="offer__buttons">
														<a href="#" class="button button--disabled"
															>Оплатить</a
														>
														<span class="offer__text"
															>Возможность оплаты появится после рассмотрения
															заказа исполнителем</span
														>
													</div>
												</div>
											</div>
										</div>
									</div>
									<div class="substance">
										<div class="offer">
											<div class="offer__content">
												<div class="offer__left">
													<div class="offer-prev">
														<ul class="offer-prev__list">
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
														</ul>
													</div>
												</div>
												<div class="offer__right">
													<div
														class="offer-status offer-status--given-production"
													>
														<span class="offer-status__text"
															>Отдан на производство</span
														>
													</div>
													<ul class="info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Номер заказа:</span
															>
															<span class="info-list__value">5</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Адрес доставки:</span
															>
															<span class="info-list__value"
																>Москва, Мира 1а, квартира 15</span
															>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Способ доставки:</span
															>
															<span class="info-list__value">Почта РФ</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Дата доставки:</span
															>
															<span class="info-list__value">21.11.2021</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Цена доставки:</span
															>
															<span class="info-list__value">562 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
												</div>
											</div>
										</div>
									</div>
									<div class="substance">
										<div class="offer">
											<div class="offer__content">
												<div class="offer__left">
													<div class="offer-prev">
														<ul class="offer-prev__list">
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
														</ul>
													</div>
												</div>
												<div class="offer__right">
													<div class="offer-status offer-status--manufactured">
														<span class="offer-status__text"
															>Изготавливается</span
														>
													</div>
													<ul class="info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Номер заказа:</span
															>
															<span class="info-list__value">5</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Адрес доставки:</span
															>
															<span class="info-list__value"
																>Москва, Мира 1а, квартира 15</span
															>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Способ доставки:</span
															>
															<span class="info-list__value">Почта РФ</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Дата доставки:</span
															>
															<span class="info-list__value">21.11.2021</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Цена доставки:</span
															>
															<span class="info-list__value">562 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
												</div>
											</div>
										</div>
									</div>
									<div class="substance">
										<div class="offer">
											<div class="offer__content">
												<div class="offer__left">
													<div class="offer-prev">
														<ul class="offer-prev__list">
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
														</ul>
													</div>
												</div>
												<div class="offer__right">
													<div
														class="offer-status offer-status--awaiting-payment"
													>
														<span class="offer-status__text"
															>Ожидает оплаты</span
														>
													</div>
													<ul class="info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Номер заказа:</span
															>
															<span class="info-list__value">5</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Адрес доставки:</span
															>
															<span class="info-list__value"
																>Москва, Мира 1а, квартира 15</span
															>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Способ доставки:</span
															>
															<span class="info-list__value">Почта РФ</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Дата доставки:</span
															>
															<span class="info-list__value">21.11.2021</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Цена доставки:</span
															>
															<span class="info-list__value">562 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
													<div class="offer__buttons">
														<a
															href="#"
															class="button button--no-uppercase-default"
															>Оплатить</a
														>
													</div>
												</div>
											</div>
										</div>
									</div>
								</div>
								<div class="tabs-content__item tab-2" class:is-active={!tab}>
									<div class="substance">
										<div class="offer">
											<div class="offer__content">
												<div class="offer__left">
													<div class="offer-prev">
														<ul class="offer-prev__list">
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
														</ul>
													</div>
												</div>
												<div class="offer__right">
													<div class="offer-status offer-status--delivered">
														<span class="offer-status__text">Доставлен</span>
													</div>
													<ul class="info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Номер заказа:</span
															>
															<span class="info-list__value">5</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Адрес доставки:</span
															>
															<span class="info-list__value"
																>Москва, Мира 1а, квартира 15</span
															>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Способ доставки:</span
															>
															<span class="info-list__value">Почта РФ</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Дата доставки:</span
															>
															<span class="info-list__value">21.11.2021</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Цена доставки:</span
															>
															<span class="info-list__value">562 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
												</div>
											</div>
										</div>
									</div>
									<div class="substance">
										<div class="offer">
											<div class="offer__content">
												<div class="offer__left">
													<div class="offer-prev">
														<ul class="offer-prev__list">
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
														</ul>
													</div>
												</div>
												<div class="offer__right">
													<div class="offer-status offer-status--canceled">
														<span class="offer-status__text">Отменен</span>
													</div>
													<ul class="info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Номер заказа:</span
															>
															<span class="info-list__value">5</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Адрес доставки:</span
															>
															<span class="info-list__value"
																>Москва, Мира 1а, квартира 15</span
															>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Способ доставки:</span
															>
															<span class="info-list__value">Почта РФ</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Дата доставки:</span
															>
															<span class="info-list__value">21.11.2021</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Цена доставки:</span
															>
															<span class="info-list__value">562 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
												</div>
											</div>
										</div>
									</div>
									<div class="substance">
										<div class="offer">
											<div class="offer__content">
												<div class="offer__left">
													<div class="offer-prev">
														<ul class="offer-prev__list">
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
															<li class="offer-prev__item">
																<a href="#" class="offer-prev__link">
																	<div class="offer-prev__pic">
																		<img
																			src="./images/content/offer-prev-img1.png"
																			alt=""
																		/>
																	</div>
																	<div class="offer-prev__info">
																		<span class="offer-prev__info-text"
																			>Filename.stl (5 шт)</span
																		>
																		<span class="offer-prev__info-count"
																			>1 шт: 2.456 ₽</span
																		>
																		<span class="offer-prev__info-price"
																			>Итого: 12.280 ₽</span
																		>
																	</div>
																</a>
															</li>
														</ul>
													</div>
												</div>
												<div class="offer__right">
													<div class="offer-status offer-status--canceled">
														<span class="offer-status__text">Отменен</span>
													</div>
													<ul class="info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Номер заказа:</span
															>
															<span class="info-list__value">5</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Адрес доставки:</span
															>
															<span class="info-list__value"
																>Москва, Мира 1а, квартира 15</span
															>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Способ доставки:</span
															>
															<span class="info-list__value">Почта РФ</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Дата доставки:</span
															>
															<span class="info-list__value">21.11.2021</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Цена доставки:</span
															>
															<span class="info-list__value">562 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
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
