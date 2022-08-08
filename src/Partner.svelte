<script>
	import Menu from "./Menu.svelte";
	import Nav from "./Nav.svelte";
</script>

<div class="wrapper">
	<Nav />
	<div class="general">
		<Menu />
		<main class="main">
			<div class="page">
				<div class="container">
					<div class="page__wrap">
						<div class="page__header">
							<h3 class="page-title">Доступные заказы для производства</h3>
						</div>
						<div class="page__body">
							<div class="production">
								<div class="production__header">
									<div class="production__filter">
										<div class="production__filter-left">
											<div class="select production__filter-select">
												<div class="select__header">
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
											<div class="select production__filter-select">
												<div class="select__header">
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
										</div>
										<a href="#" class="button-link button-link--red"
											>Сбросить фильтры</a
										>
									</div>
									<div class="production-components">
										<div class="production-components__item">
											<span class="production-components__text"
												>3D печать SLA</span
											>
											<button class="production-components__exit">
												<svg
													width="8"
													height="7"
													viewBox="0 0 8 7"
													fill="none"
													xmlns="http://www.w3.org/2000/svg"
												>
													<path
														d="M7.38521 6.33569C7.42157 6.37167 7.45044 6.41452 7.47015 6.46176C7.48985 6.50901 7.5 6.5597 7.5 6.6109C7.5 6.66209 7.48985 6.71278 7.47015 6.76003C7.45044 6.80727 7.42157 6.85012 7.38521 6.8861C7.31231 6.95904 7.21349 7 7.11045 7C7.00742 7 6.9086 6.95904 6.8357 6.8861L4 4.04573L1.1643 6.8861C1.0914 6.95904 0.992582 7 0.889548 7C0.786514 7 0.687693 6.95904 0.614792 6.8861C0.578428 6.85012 0.549556 6.80727 0.529852 6.76003C0.510147 6.71278 0.5 6.66209 0.5 6.6109C0.5 6.5597 0.510147 6.50901 0.529852 6.46176C0.549556 6.41452 0.578428 6.37167 0.614792 6.33569L3.45049 3.49531L0.614792 0.654943C0.546113 0.581116 0.508723 0.483471 0.510501 0.382577C0.512278 0.281684 0.553083 0.18542 0.624319 0.114067C0.695555 0.042713 0.79166 0.00184082 0.892388 6.06744e-05C0.993115 -0.00171947 1.0906 0.0357315 1.1643 0.104523L4 2.94489L6.8357 0.104523C6.9094 0.0357315 7.00688 -0.00171947 7.10761 6.06744e-05C7.20834 0.00184082 7.30444 0.042713 7.37568 0.114067C7.44692 0.18542 7.48772 0.281684 7.4895 0.382577C7.49128 0.483471 7.45389 0.581116 7.38521 0.654943L4.54951 3.49531L7.38521 6.33569Z"
														fill="#6A96F8"
													/>
												</svg>
											</button>
										</div>
										<div class="production-components__item">
											<span class="production-components__text"
												>Литье в силикон</span
											>
											<button class="production-components__exit">
												<svg
													width="8"
													height="7"
													viewBox="0 0 8 7"
													fill="none"
													xmlns="http://www.w3.org/2000/svg"
												>
													<path
														d="M7.38521 6.33569C7.42157 6.37167 7.45044 6.41452 7.47015 6.46176C7.48985 6.50901 7.5 6.5597 7.5 6.6109C7.5 6.66209 7.48985 6.71278 7.47015 6.76003C7.45044 6.80727 7.42157 6.85012 7.38521 6.8861C7.31231 6.95904 7.21349 7 7.11045 7C7.00742 7 6.9086 6.95904 6.8357 6.8861L4 4.04573L1.1643 6.8861C1.0914 6.95904 0.992582 7 0.889548 7C0.786514 7 0.687693 6.95904 0.614792 6.8861C0.578428 6.85012 0.549556 6.80727 0.529852 6.76003C0.510147 6.71278 0.5 6.66209 0.5 6.6109C0.5 6.5597 0.510147 6.50901 0.529852 6.46176C0.549556 6.41452 0.578428 6.37167 0.614792 6.33569L3.45049 3.49531L0.614792 0.654943C0.546113 0.581116 0.508723 0.483471 0.510501 0.382577C0.512278 0.281684 0.553083 0.18542 0.624319 0.114067C0.695555 0.042713 0.79166 0.00184082 0.892388 6.06744e-05C0.993115 -0.00171947 1.0906 0.0357315 1.1643 0.104523L4 2.94489L6.8357 0.104523C6.9094 0.0357315 7.00688 -0.00171947 7.10761 6.06744e-05C7.20834 0.00184082 7.30444 0.042713 7.37568 0.114067C7.44692 0.18542 7.48772 0.281684 7.4895 0.382577C7.49128 0.483471 7.45389 0.581116 7.38521 0.654943L4.54951 3.49531L7.38521 6.33569Z"
														fill="#6A96F8"
													/>
												</svg>
											</button>
										</div>
										<div class="production-components__item">
											<span class="production-components__text"
												>Обработка с ЧПУ</span
											>
											<button class="production-components__exit">
												<svg
													width="8"
													height="7"
													viewBox="0 0 8 7"
													fill="none"
													xmlns="http://www.w3.org/2000/svg"
												>
													<path
														d="M7.38521 6.33569C7.42157 6.37167 7.45044 6.41452 7.47015 6.46176C7.48985 6.50901 7.5 6.5597 7.5 6.6109C7.5 6.66209 7.48985 6.71278 7.47015 6.76003C7.45044 6.80727 7.42157 6.85012 7.38521 6.8861C7.31231 6.95904 7.21349 7 7.11045 7C7.00742 7 6.9086 6.95904 6.8357 6.8861L4 4.04573L1.1643 6.8861C1.0914 6.95904 0.992582 7 0.889548 7C0.786514 7 0.687693 6.95904 0.614792 6.8861C0.578428 6.85012 0.549556 6.80727 0.529852 6.76003C0.510147 6.71278 0.5 6.66209 0.5 6.6109C0.5 6.5597 0.510147 6.50901 0.529852 6.46176C0.549556 6.41452 0.578428 6.37167 0.614792 6.33569L3.45049 3.49531L0.614792 0.654943C0.546113 0.581116 0.508723 0.483471 0.510501 0.382577C0.512278 0.281684 0.553083 0.18542 0.624319 0.114067C0.695555 0.042713 0.79166 0.00184082 0.892388 6.06744e-05C0.993115 -0.00171947 1.0906 0.0357315 1.1643 0.104523L4 2.94489L6.8357 0.104523C6.9094 0.0357315 7.00688 -0.00171947 7.10761 6.06744e-05C7.20834 0.00184082 7.30444 0.042713 7.37568 0.114067C7.44692 0.18542 7.48772 0.281684 7.4895 0.382577C7.49128 0.483471 7.45389 0.581116 7.38521 0.654943L4.54951 3.49531L7.38521 6.33569Z"
														fill="#6A96F8"
													/>
												</svg>
											</button>
										</div>
										<div class="production-components__item">
											<span class="production-components__text"
												>Рзка листовых материалов</span
											>
											<button class="production-components__exit">
												<svg
													width="8"
													height="7"
													viewBox="0 0 8 7"
													fill="none"
													xmlns="http://www.w3.org/2000/svg"
												>
													<path
														d="M7.38521 6.33569C7.42157 6.37167 7.45044 6.41452 7.47015 6.46176C7.48985 6.50901 7.5 6.5597 7.5 6.6109C7.5 6.66209 7.48985 6.71278 7.47015 6.76003C7.45044 6.80727 7.42157 6.85012 7.38521 6.8861C7.31231 6.95904 7.21349 7 7.11045 7C7.00742 7 6.9086 6.95904 6.8357 6.8861L4 4.04573L1.1643 6.8861C1.0914 6.95904 0.992582 7 0.889548 7C0.786514 7 0.687693 6.95904 0.614792 6.8861C0.578428 6.85012 0.549556 6.80727 0.529852 6.76003C0.510147 6.71278 0.5 6.66209 0.5 6.6109C0.5 6.5597 0.510147 6.50901 0.529852 6.46176C0.549556 6.41452 0.578428 6.37167 0.614792 6.33569L3.45049 3.49531L0.614792 0.654943C0.546113 0.581116 0.508723 0.483471 0.510501 0.382577C0.512278 0.281684 0.553083 0.18542 0.624319 0.114067C0.695555 0.042713 0.79166 0.00184082 0.892388 6.06744e-05C0.993115 -0.00171947 1.0906 0.0357315 1.1643 0.104523L4 2.94489L6.8357 0.104523C6.9094 0.0357315 7.00688 -0.00171947 7.10761 6.06744e-05C7.20834 0.00184082 7.30444 0.042713 7.37568 0.114067C7.44692 0.18542 7.48772 0.281684 7.4895 0.382577C7.49128 0.483471 7.45389 0.581116 7.38521 0.654943L4.54951 3.49531L7.38521 6.33569Z"
														fill="#6A96F8"
													/>
												</svg>
											</button>
										</div>
									</div>
								</div>
								<div class="production__body">
									<div class="production-order">
										<div class="production-order__header">
											<div class="production-order__label">
												<svg
													width="15"
													height="14"
													viewBox="0 0 15 14"
													fill="none"
													xmlns="http://www.w3.org/2000/svg"
												>
													<path
														d="M14.6667 10.0683C14.6667 9.89359 14.525 9.75195 14.3503 9.75195C14.1756 9.75195 14.034 9.89359 14.034 10.0683V11.941C14.034 12.2171 13.8101 12.441 13.534 12.441H2.13272C1.85657 12.441 1.63272 12.2171 1.63272 11.941V10.0683C1.63272 9.89359 1.49108 9.75195 1.31636 9.75195C1.14164 9.75195 1 9.89359 1 10.0683V12.5737C1 12.8499 1.22386 13.0737 1.5 13.0737H14.1667C14.4428 13.0737 14.6667 12.8499 14.6667 12.5737V10.0683Z"
														fill="#5C5C5C"
														stroke="#5C5C5C"
														stroke-width="0.4"
													/>
													<path
														d="M6.66305 2.9815C6.97804 2.66651 7.51661 2.8896 7.51661 3.33505V9.62563C7.51661 9.80035 7.65825 9.94198 7.83296 9.94198C8.00768 9.94198 8.14932 9.80035 8.14932 9.62563V3.33505C8.14932 2.8896 8.68789 2.66651 9.00288 2.9815L10.6486 4.62717C10.7709 4.74947 10.9691 4.74947 11.0915 4.62717C11.2138 4.50487 11.2138 4.30657 11.0915 4.18427L8.18652 1.27933C7.99126 1.08407 7.67467 1.08407 7.47941 1.27933L4.57448 4.18427C4.45217 4.30657 4.45217 4.50487 4.57448 4.62717C4.69678 4.74947 4.89508 4.74947 5.01738 4.62717L6.66305 2.9815Z"
														fill="#5C5C5C"
														stroke="#5C5C5C"
														stroke-width="0.4"
													/>
												</svg>
												<span class="production-order__label-text">001ь001</span
												>
											</div>
											<div class="production-order__label">
												<span class="production-order__label-text"
													>3 Детали</span
												>
											</div>
										</div>
										<div class="production-order__body">
											<div class="production-table">
												<div
													class="production-table__row production-table__header"
												>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text">Чертежи</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Название и номер</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Материал</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Срок производства</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Количество</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Стоимость</span
															>
														</div>
													</div>
												</div>
												<div class="production-table__body">
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="production-order__footer">
											<div class="production-order__footer-top">
												<span class="production-order__footer-price"
													>Общая стоимость: 982.400 ₽</span
												>
											</div>
											<div class="production-order__footer-bottom">
												<div class="production-order__footer-comments">
													<span class="production-order__footer-category"
														>Комментарии к производству</span
													>
													<span class="production-order__footer-value"
														>Допуск: ISO 268 - класс точности 12 (стандарт), 1
														точка Шероховатость поверхности: стандарт (RA 3,2
														мкм), вся деталь</span
													>
												</div>

												<a href="#" class="button button--no-uppercase-default"
													>Подать заяавку на заказ</a
												>
											</div>
										</div>
									</div>
									<div class="production-order">
										<div class="production-order__header">
											<div class="production-order__label">
												<svg
													width="15"
													height="14"
													viewBox="0 0 15 14"
													fill="none"
													xmlns="http://www.w3.org/2000/svg"
												>
													<path
														d="M14.6667 10.0683C14.6667 9.89359 14.525 9.75195 14.3503 9.75195C14.1756 9.75195 14.034 9.89359 14.034 10.0683V11.941C14.034 12.2171 13.8101 12.441 13.534 12.441H2.13272C1.85657 12.441 1.63272 12.2171 1.63272 11.941V10.0683C1.63272 9.89359 1.49108 9.75195 1.31636 9.75195C1.14164 9.75195 1 9.89359 1 10.0683V12.5737C1 12.8499 1.22386 13.0737 1.5 13.0737H14.1667C14.4428 13.0737 14.6667 12.8499 14.6667 12.5737V10.0683Z"
														fill="#5C5C5C"
														stroke="#5C5C5C"
														stroke-width="0.4"
													/>
													<path
														d="M6.66305 2.9815C6.97804 2.66651 7.51661 2.8896 7.51661 3.33505V9.62563C7.51661 9.80035 7.65825 9.94198 7.83296 9.94198C8.00768 9.94198 8.14932 9.80035 8.14932 9.62563V3.33505C8.14932 2.8896 8.68789 2.66651 9.00288 2.9815L10.6486 4.62717C10.7709 4.74947 10.9691 4.74947 11.0915 4.62717C11.2138 4.50487 11.2138 4.30657 11.0915 4.18427L8.18652 1.27933C7.99126 1.08407 7.67467 1.08407 7.47941 1.27933L4.57448 4.18427C4.45217 4.30657 4.45217 4.50487 4.57448 4.62717C4.69678 4.74947 4.89508 4.74947 5.01738 4.62717L6.66305 2.9815Z"
														fill="#5C5C5C"
														stroke="#5C5C5C"
														stroke-width="0.4"
													/>
												</svg>
												<span class="production-order__label-text">001ь001</span
												>
											</div>
											<div class="production-order__label">
												<span class="production-order__label-text"
													>3 Детали</span
												>
											</div>
										</div>
										<div class="production-order__body">
											<div class="production-table">
												<div
													class="production-table__row production-table__header"
												>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text">Чертежи</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Название и номер</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Материал</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Срок производства</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Количество</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Стоимость</span
															>
														</div>
													</div>
												</div>
												<div class="production-table__body">
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="production-order__footer">
											<div class="production-order__footer-top">
												<span class="production-order__footer-price"
													>Общая стоимость: 982.400 ₽</span
												>
											</div>
											<div class="production-order__footer-bottom">
												<div class="production-order__footer-comments">
													<span class="production-order__footer-category"
														>Комментарии к производству</span
													>
													<span class="production-order__footer-value"
														>Допуск: ISO 268 - класс точности 12 (стандарт), 1
														точка Шероховатость поверхности: стандарт (RA 3,2
														мкм), вся деталь</span
													>
												</div>

												<a href="#" class="button button--no-uppercase-default"
													>Подать заяавку на заказ</a
												>
											</div>
										</div>
									</div>
									<div class="production-order">
										<div class="production-order__header">
											<div class="production-order__label">
												<svg
													width="15"
													height="14"
													viewBox="0 0 15 14"
													fill="none"
													xmlns="http://www.w3.org/2000/svg"
												>
													<path
														d="M14.6667 10.0683C14.6667 9.89359 14.525 9.75195 14.3503 9.75195C14.1756 9.75195 14.034 9.89359 14.034 10.0683V11.941C14.034 12.2171 13.8101 12.441 13.534 12.441H2.13272C1.85657 12.441 1.63272 12.2171 1.63272 11.941V10.0683C1.63272 9.89359 1.49108 9.75195 1.31636 9.75195C1.14164 9.75195 1 9.89359 1 10.0683V12.5737C1 12.8499 1.22386 13.0737 1.5 13.0737H14.1667C14.4428 13.0737 14.6667 12.8499 14.6667 12.5737V10.0683Z"
														fill="#5C5C5C"
														stroke="#5C5C5C"
														stroke-width="0.4"
													/>
													<path
														d="M6.66305 2.9815C6.97804 2.66651 7.51661 2.8896 7.51661 3.33505V9.62563C7.51661 9.80035 7.65825 9.94198 7.83296 9.94198C8.00768 9.94198 8.14932 9.80035 8.14932 9.62563V3.33505C8.14932 2.8896 8.68789 2.66651 9.00288 2.9815L10.6486 4.62717C10.7709 4.74947 10.9691 4.74947 11.0915 4.62717C11.2138 4.50487 11.2138 4.30657 11.0915 4.18427L8.18652 1.27933C7.99126 1.08407 7.67467 1.08407 7.47941 1.27933L4.57448 4.18427C4.45217 4.30657 4.45217 4.50487 4.57448 4.62717C4.69678 4.74947 4.89508 4.74947 5.01738 4.62717L6.66305 2.9815Z"
														fill="#5C5C5C"
														stroke="#5C5C5C"
														stroke-width="0.4"
													/>
												</svg>
												<span class="production-order__label-text">001ь001</span
												>
											</div>
											<div class="production-order__label">
												<span class="production-order__label-text"
													>3 Детали</span
												>
											</div>
										</div>
										<div class="production-order__body">
											<div class="production-table">
												<div
													class="production-table__row production-table__header"
												>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text">Чертежи</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Название и номер</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Материал</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Срок производства</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Количество</span
															>
														</div>
													</div>
													<div class="production-table__column">
														<div class="production-table__cell">
															<span class="production-table__text"
																>Стоимость</span
															>
														</div>
													</div>
												</div>
												<div class="production-table__body">
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
													<div class="production-table__row">
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="production-table__pic">
																	<img
																		src="./images/content/production-table-img1.png"
																		alt=""
																	/>
																</div>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>Н 9457<br />
																	кришка верхная</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>СНС<br />
																	6061 Aluminum</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>2-3 дня</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<span class="production-table__text"
																	>100 штук</span
																>
															</div>
														</div>
														<div class="production-table__column">
															<div class="production-table__cell">
																<div class="price production-table__price">
																	<span class="price__small">1 шт: 2.456 ₽</span
																	>
																	<span class="price__total"
																		>Итого: 245.600 ₽</span
																	>
																</div>
															</div>
														</div>
													</div>
												</div>
											</div>
										</div>
										<div class="production-order__footer">
											<div class="production-order__footer-top">
												<span class="production-order__footer-price"
													>Общая стоимость: 982.400 ₽</span
												>
											</div>
											<div class="production-order__footer-bottom">
												<div class="production-order__footer-comments">
													<span class="production-order__footer-category"
														>Комментарии к производству</span
													>
													<span class="production-order__footer-value"
														>Допуск: ISO 268 - класс точности 12 (стандарт), 1
														точка Шероховатость поверхности: стандарт (RA 3,2
														мкм), вся деталь</span
													>
												</div>

												<a href="#" class="button button--no-uppercase-default"
													>Подать заяавку на заказ</a
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
		</main>
	</div>
</div>
