<script>
	import Menu from "./Menu.svelte";
	import Nav from "./Nav.svelte";
	import Util from "./util.js";

	let delivery = true;
	let tabFlag;
	let cancel = false;

	//Yandex map API
	ymaps.ready(init);
	function init() {
		// Создание карты.
		var myMap = new ymaps.Map("map", {
			// Координаты центра карты.
			// Порядок по умолчанию: «широта, долгота».
			// Чтобы не определять координаты центра карты вручную,
			// воспользуйтесь инструментом Определение координат.
			center: [55.76, 37.64],
			// Уровень масштабирования. Допустимые значения:
			// от 0 (весь мир) до 19.
			zoom: 10,
		});
	}
	window.onload = Util.tab;
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
							<h3 class="page-title">Оформление заказа</h3>
						</div>
						<div class="page__body">
							<div class="substance">
								<div class="offer offer--registration">
									<div class="offer__content">
										<div class="offer__left">
											<span class="account-title offer__title">Ваш заказ</span>
											<div class="offer-prev">
												<ul class="offer-prev__list">
													<li class="offer-prev__item">
														<a class="offer-prev__link">
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
																<div class="counter offer-prev__counter">
																	<button
																		class="counter__button counter__button--left"
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
																		value="5"
																	/>
																	<button
																		class="counter__button counter__button--right"
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
															<button class="offer-prev__btn-exit">
																<svg
																	width="8"
																	height="8"
																	viewBox="0 0 8 8"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<path
																		d="M7.38521 6.83569C7.42157 6.87167 7.45044 6.91452 7.47015 6.96176C7.48985 7.00901 7.5 7.0597 7.5 7.1109C7.5 7.16209 7.48985 7.21278 7.47015 7.26003C7.45044 7.30727 7.42157 7.35012 7.38521 7.3861C7.31231 7.45904 7.21349 7.5 7.11045 7.5C7.00742 7.5 6.9086 7.45904 6.8357 7.3861L4 4.54573L1.1643 7.3861C1.0914 7.45904 0.992582 7.5 0.889548 7.5C0.786514 7.5 0.687693 7.45904 0.614792 7.3861C0.578428 7.35012 0.549556 7.30727 0.529852 7.26003C0.510147 7.21278 0.5 7.16209 0.5 7.1109C0.5 7.0597 0.510147 7.00901 0.529852 6.96176C0.549556 6.91452 0.578428 6.87167 0.614792 6.83569L3.45049 3.99531L0.614792 1.15494C0.546113 1.08112 0.508723 0.983471 0.510501 0.882577C0.512278 0.781684 0.553083 0.68542 0.624319 0.614067C0.695555 0.542713 0.79166 0.501841 0.892388 0.500061C0.993115 0.498281 1.0906 0.535731 1.1643 0.604523L4 3.44489L6.8357 0.604523C6.9094 0.535731 7.00688 0.498281 7.10761 0.500061C7.20834 0.501841 7.30444 0.542713 7.37568 0.614067C7.44692 0.68542 7.48772 0.781684 7.4895 0.882577C7.49128 0.983471 7.45389 1.08112 7.38521 1.15494L4.54951 3.99531L7.38521 6.83569Z"
																		fill="#333333"
																	/>
																</svg>
															</button>
														</a>
													</li>
													<li class="offer-prev__item">
														<a class="offer-prev__link">
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
																<div class="counter offer-prev__counter">
																	<button
																		class="counter__button counter__button--left"
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
																		value="5"
																	/>
																	<button
																		class="counter__button counter__button--right"
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
															<button class="offer-prev__btn-exit">
																<svg
																	width="8"
																	height="8"
																	viewBox="0 0 8 8"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<path
																		d="M7.38521 6.83569C7.42157 6.87167 7.45044 6.91452 7.47015 6.96176C7.48985 7.00901 7.5 7.0597 7.5 7.1109C7.5 7.16209 7.48985 7.21278 7.47015 7.26003C7.45044 7.30727 7.42157 7.35012 7.38521 7.3861C7.31231 7.45904 7.21349 7.5 7.11045 7.5C7.00742 7.5 6.9086 7.45904 6.8357 7.3861L4 4.54573L1.1643 7.3861C1.0914 7.45904 0.992582 7.5 0.889548 7.5C0.786514 7.5 0.687693 7.45904 0.614792 7.3861C0.578428 7.35012 0.549556 7.30727 0.529852 7.26003C0.510147 7.21278 0.5 7.16209 0.5 7.1109C0.5 7.0597 0.510147 7.00901 0.529852 6.96176C0.549556 6.91452 0.578428 6.87167 0.614792 6.83569L3.45049 3.99531L0.614792 1.15494C0.546113 1.08112 0.508723 0.983471 0.510501 0.882577C0.512278 0.781684 0.553083 0.68542 0.624319 0.614067C0.695555 0.542713 0.79166 0.501841 0.892388 0.500061C0.993115 0.498281 1.0906 0.535731 1.1643 0.604523L4 3.44489L6.8357 0.604523C6.9094 0.535731 7.00688 0.498281 7.10761 0.500061C7.20834 0.501841 7.30444 0.542713 7.37568 0.614067C7.44692 0.68542 7.48772 0.781684 7.4895 0.882577C7.49128 0.983471 7.45389 1.08112 7.38521 1.15494L4.54951 3.99531L7.38521 6.83569Z"
																		fill="#333333"
																	/>
																</svg>
															</button>
														</a>
													</li>
													<li class="offer-prev__item">
														<a class="offer-prev__link">
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
																<div class="counter offer-prev__counter">
																	<button
																		class="counter__button counter__button--left"
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
																		value="5"
																	/>
																	<button
																		class="counter__button counter__button--right"
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
															<button class="offer-prev__btn-exit">
																<svg
																	width="8"
																	height="8"
																	viewBox="0 0 8 8"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<path
																		d="M7.38521 6.83569C7.42157 6.87167 7.45044 6.91452 7.47015 6.96176C7.48985 7.00901 7.5 7.0597 7.5 7.1109C7.5 7.16209 7.48985 7.21278 7.47015 7.26003C7.45044 7.30727 7.42157 7.35012 7.38521 7.3861C7.31231 7.45904 7.21349 7.5 7.11045 7.5C7.00742 7.5 6.9086 7.45904 6.8357 7.3861L4 4.54573L1.1643 7.3861C1.0914 7.45904 0.992582 7.5 0.889548 7.5C0.786514 7.5 0.687693 7.45904 0.614792 7.3861C0.578428 7.35012 0.549556 7.30727 0.529852 7.26003C0.510147 7.21278 0.5 7.16209 0.5 7.1109C0.5 7.0597 0.510147 7.00901 0.529852 6.96176C0.549556 6.91452 0.578428 6.87167 0.614792 6.83569L3.45049 3.99531L0.614792 1.15494C0.546113 1.08112 0.508723 0.983471 0.510501 0.882577C0.512278 0.781684 0.553083 0.68542 0.624319 0.614067C0.695555 0.542713 0.79166 0.501841 0.892388 0.500061C0.993115 0.498281 1.0906 0.535731 1.1643 0.604523L4 3.44489L6.8357 0.604523C6.9094 0.535731 7.00688 0.498281 7.10761 0.500061C7.20834 0.501841 7.30444 0.542713 7.37568 0.614067C7.44692 0.68542 7.48772 0.781684 7.4895 0.882577C7.49128 0.983471 7.45389 1.08112 7.38521 1.15494L4.54951 3.99531L7.38521 6.83569Z"
																		fill="#333333"
																	/>
																</svg>
															</button>
														</a>
													</li>
													<li class="offer-prev__item">
														<a class="offer-prev__link">
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
																<div class="counter offer-prev__counter">
																	<button
																		class="counter__button counter__button--left"
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
																		value="5"
																	/>
																	<button
																		class="counter__button counter__button--right"
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
															<button class="offer-prev__btn-exit">
																<svg
																	width="8"
																	height="8"
																	viewBox="0 0 8 8"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<path
																		d="M7.38521 6.83569C7.42157 6.87167 7.45044 6.91452 7.47015 6.96176C7.48985 7.00901 7.5 7.0597 7.5 7.1109C7.5 7.16209 7.48985 7.21278 7.47015 7.26003C7.45044 7.30727 7.42157 7.35012 7.38521 7.3861C7.31231 7.45904 7.21349 7.5 7.11045 7.5C7.00742 7.5 6.9086 7.45904 6.8357 7.3861L4 4.54573L1.1643 7.3861C1.0914 7.45904 0.992582 7.5 0.889548 7.5C0.786514 7.5 0.687693 7.45904 0.614792 7.3861C0.578428 7.35012 0.549556 7.30727 0.529852 7.26003C0.510147 7.21278 0.5 7.16209 0.5 7.1109C0.5 7.0597 0.510147 7.00901 0.529852 6.96176C0.549556 6.91452 0.578428 6.87167 0.614792 6.83569L3.45049 3.99531L0.614792 1.15494C0.546113 1.08112 0.508723 0.983471 0.510501 0.882577C0.512278 0.781684 0.553083 0.68542 0.624319 0.614067C0.695555 0.542713 0.79166 0.501841 0.892388 0.500061C0.993115 0.498281 1.0906 0.535731 1.1643 0.604523L4 3.44489L6.8357 0.604523C6.9094 0.535731 7.00688 0.498281 7.10761 0.500061C7.20834 0.501841 7.30444 0.542713 7.37568 0.614067C7.44692 0.68542 7.48772 0.781684 7.4895 0.882577C7.49128 0.983471 7.45389 1.08112 7.38521 1.15494L4.54951 3.99531L7.38521 6.83569Z"
																		fill="#333333"
																	/>
																</svg>
															</button>
														</a>
													</li>
													<li class="offer-prev__item">
														<a class="offer-prev__link">
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
																<div class="counter offer-prev__counter">
																	<button
																		class="counter__button counter__button--left"
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
																		value="5"
																	/>
																	<button
																		class="counter__button counter__button--right"
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
															<button class="offer-prev__btn-exit">
																<svg
																	width="8"
																	height="8"
																	viewBox="0 0 8 8"
																	fill="none"
																	xmlns="http://www.w3.org/2000/svg"
																>
																	<path
																		d="M7.38521 6.83569C7.42157 6.87167 7.45044 6.91452 7.47015 6.96176C7.48985 7.00901 7.5 7.0597 7.5 7.1109C7.5 7.16209 7.48985 7.21278 7.47015 7.26003C7.45044 7.30727 7.42157 7.35012 7.38521 7.3861C7.31231 7.45904 7.21349 7.5 7.11045 7.5C7.00742 7.5 6.9086 7.45904 6.8357 7.3861L4 4.54573L1.1643 7.3861C1.0914 7.45904 0.992582 7.5 0.889548 7.5C0.786514 7.5 0.687693 7.45904 0.614792 7.3861C0.578428 7.35012 0.549556 7.30727 0.529852 7.26003C0.510147 7.21278 0.5 7.16209 0.5 7.1109C0.5 7.0597 0.510147 7.00901 0.529852 6.96176C0.549556 6.91452 0.578428 6.87167 0.614792 6.83569L3.45049 3.99531L0.614792 1.15494C0.546113 1.08112 0.508723 0.983471 0.510501 0.882577C0.512278 0.781684 0.553083 0.68542 0.624319 0.614067C0.695555 0.542713 0.79166 0.501841 0.892388 0.500061C0.993115 0.498281 1.0906 0.535731 1.1643 0.604523L4 3.44489L6.8357 0.604523C6.9094 0.535731 7.00688 0.498281 7.10761 0.500061C7.20834 0.501841 7.30444 0.542713 7.37568 0.614067C7.44692 0.68542 7.48772 0.781684 7.4895 0.882577C7.49128 0.983471 7.45389 1.08112 7.38521 1.15494L4.54951 3.99531L7.38521 6.83569Z"
																		fill="#333333"
																	/>
																</svg>
															</button>
														</a>
													</li>
												</ul>
											</div>
										</div>
										<div class="offer__right">
											<span class="account-title offer__title"
												>Информация о заказе</span
											>
											<div class="offer__right-top">
												<ul class="info-list">
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
												</ul>
												<a class="button-link offer__right-btn-link"
													>Изменить данные</a
												>
											</div>
											<div class="offer__right-bottom">
												<ul class="info-list">
													<li class="info-list__item">
														<span class="info-list__category"
															>Стоимость деталей</span
														>
														<span class="info-list__value">66.312 ₽</span>
													</li>
													<li class="info-list__item info-list__item--big">
														<span class="info-list__category"
															>Итоговая стоимость</span
														>
														<span class="info-list__value">66.874 ₽</span>
													</li>
												</ul>
												<div class="offer__buttons">
													<a class="button button--disabled">Оформить заказ</a>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
							<div class="substance" class:remove={!delivery}>
								<div class="city">
									<h3 class="page-title city__title">Город доставки</h3>
									<p class="page-desc city__desc">
										Выберите ваш город из списка ниже, либо введите его в строку
										поиска
									</p>
									<ul class="tabs-nav city__tabs-nav">
										<li
											class="tabs-nav__item tabs-nav__item--active"
											data-tab-name="tab-1"
											on:click={() => (tabFlag = false)}
											on:click={init}
										>
											Москва
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-2"
											on:click={() => (tabFlag = true)}
										>
											Санкт-Петербург
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-3"
											on:click={() => (tabFlag = true)}
										>
											Новосибирск
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-4"
											on:click={() => (tabFlag = true)}
										>
											Нижний новгород
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-5"
											on:click={() => (tabFlag = true)}
										>
											Казань
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-6"
											on:click={() => (tabFlag = true)}
										>
											Челябинск
										</li>
									</ul>
									<div class="search city__search">
										<div class="field">
											<input
												type="text"
												name=""
												id=""
												class="field__input search__input"
												placeholder="Введите название города"
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
							</div>
							<div class="substance">
								<div class="delivery">
									<h3 class="page-title delivery__title">Способ доставки</h3>
									<p class="page-desc delivery__desc">
										Доступные способы доставки могут отличаться в зависимости от
										выбранного вами города
									</p>
									<ul class="tabs-nav delivery__tabs-nav">
										<li
											class="tabs-nav__item tabs-nav__item--active"
											data-tab-name="tab-7"
											on:click={() => (delivery = true)}
										>
											Самовывоз
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-8"
											on:click={() => (delivery = false)}
										>
											Курьерская доставка до двери
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-9"
											on:click={() => (delivery = false)}
										>
											Почта РФ
										</li>
										<li
											class="tabs-nav__item"
											data-tab-name="tab-10"
											on:click={() => (delivery = false)}
										>
											ТК «Деловые линии»
										</li>
									</ul>
									<div class="tabs-content delivery__tabs-content">
										<div class="tabs-content__item tab-8">
											<ul class="info-list delivery__info-gorizontal">
												<li class="info-list__item">
													<span class="info-list__category"
														>Сроки доставки:
													</span>
													<span class="info-list__value">2-3 дня</span>
												</li>
												<li class="info-list__item">
													<span class="info-list__category">Стоимость от:</span>
													<span class="info-list__value">450 ₽</span>
												</li>
											</ul>
											<div class="delivery__content">
												<div class="delivery__left">
													<form action="" class="delivery-form" method="post">
														<div class="field">
															<input
																type="text"
																class="field__input delivery-form__input"
																placeholder="Город"
															/>
														</div>
														<div class="field">
															<input
																type="text"
																class="field__input delivery-form__input delivery-form__input--validate"
																placeholder="Улица, дом, офис"
															/>
														</div>
														<div class="field">
															<input
																type="number"
																class="field__input delivery-form__input delivery-form__input--validate"
																placeholder="Почтовый индекс"
															/>
														</div>
														<div class="field">
															<input
																type="text"
																class="field__input delivery-form__input delivery-form__input--validate"
																placeholder="Контактное лицо, принимающее посылку"
															/>
														</div>
														<div class="field">
															<input
																type="tel"
																class="field__input delivery-form__input delivery-form__input--validate"
																placeholder="+7 (913) 042-23-42"
															/>
														</div>
														<div class="field">
															<textarea
																class="field__input delivery-form__input delivery-form__textarea"
																placeholder="Комментарии к заказу"
															/>
														</div>
													</form>
												</div>
												<div class="delivery__right delivery__right--cancel">
													<div
														class="delivery__overlay"
														class:is-active={cancel}
													>
														<!-- <div class="delivery-cancel">
                                                                    <div class="delivery-cancel__top">
                                                                        span.delivery
                                                                    </div>
                                                                    <div class="delivery-cancel__bottom"></div>
                                                                </div> -->
														<div class="order-delete delivery-delete">
															<div class="order-delete__top">
																<span class="order-delete__title"
																	>Отменить заказ?</span
																>
																<span class="order-delete__text"
																	>После отмены заказа выбранный вами товар
																	удалится из корзины</span
																>
															</div>
															<div class="order-delete__bottom">
																<button class="order-delete__button">Да</button>
																<button
																	class="order-delete__button"
																	on:click={() => (cancel = false)}>Нет</button
																>
															</div>
														</div>
													</div>
													<div class="payment">
														<span class="payment__title">Способ оплаты</span>
														<div class="payment__body">
															<div class="payment__item">
																<input
																	type="radio"
																	name="payment"
																	id="payment1"
																/>
																<label for="payment1"
																	>Онлайн оплата картой</label
																>
															</div>
															<div class="payment__item">
																<input
																	type="radio"
																	name="payment"
																	id="payment2"
																/>
																<label for="payment2">Оплата по счету</label>
															</div>
															<div class="payment__item">
																<input
																	type="radio"
																	name="payment"
																	id="payment3"
																/>
																<label for="payment3">Оплата наличными</label>
															</div>
														</div>
													</div>
													<ul class="info-list delivery__info-list">
														<li class="info-list__item">
															<span class="info-list__category"
																>Количество деталей:</span
															>
															<span class="info-list__value">27 шт</span>
														</li>
														<li class="info-list__item">
															<span class="info-list__category"
																>Стоимость деталей</span
															>
															<span class="info-list__value">66.312 ₽</span>
														</li>
														<li class="info-list__item info-list__item--big">
															<span class="info-list__category"
																>Итоговая стоимость</span
															>
															<span class="info-list__value">66.874 ₽</span>
														</li>
													</ul>
													<div class="delivery__buttons">
														<a class="button button--big">Оформить заказ</a>

														<a
															class="button button--border button--big"
															on:click={() => (cancel = true)}>Отменить заказ</a
														>
													</div>
												</div>
											</div>
										</div>
										{#if tabFlag}
											<div class="tabs-content__item tab-7 is-active">
												<div
													class="delivery__content delivery__content--small-gap"
												>
													<div class="delivery__left">
														<div class="company">
															<span class="company__title"
																>СДЭК (Самовывоз)</span
															>
															<ul class="info-list company__info-list">
																<li class="info-list__item">
																	<span class="info-list__category"
																		>Сроки доставки:</span
																	>
																	<span class="info-list__value">1-2 дня</span>
																</li>
																<li class="info-list__item">
																	<span class="info-list__category"
																		>Стоимость:</span
																	>
																	<span class="info-list__value">215 ₽</span>
																</li>
															</ul>
															<p class="company__desc">
																Доставка в один из пунктов самовывоза компании
																СДЭК
															</p>

															<a
																href="#"
																class="button button--no-uppercase-default"
																>Выбрать пункт выдачи</a
															>

															<div class="company__address">
																<span class="company__address-text"
																	>Вы сможете забрать заказ по адресу:</span
																>
																<span class="company__address-value"
																	>пр. Наставников, 21, лит.А, пом.15-Н</span
																>
															</div>
														</div>
													</div>
													<div class="delivery__right delivery__right--cancel">
														<div
															class="delivery__overlay"
															class:is-active={cancel}
														>
															<!-- <div class="delivery-cancel">
                                                                    <div class="delivery-cancel__top">
                                                                        span.delivery
                                                                    </div>
                                                                    <div class="delivery-cancel__bottom"></div>
                                                                </div> -->
															<div class="order-delete delivery-delete">
																<div class="order-delete__top">
																	<span class="order-delete__title"
																		>Отменить заказ?</span
																	>
																	<span class="order-delete__text"
																		>После отмены заказа выбранный вами товар
																		удалится из корзины</span
																	>
																</div>
																<div class="order-delete__bottom">
																	<button class="order-delete__button"
																		>Да</button
																	>
																	<button
																		class="order-delete__button"
																		on:click={() => (cancel = false)}
																		>Нет</button
																	>
																</div>
															</div>
														</div>
														<div class="payment">
															<span class="payment__title">Способ оплаты</span>
															<div class="payment__body">
																<div class="payment__item">
																	<input
																		type="radio"
																		name="payment"
																		id="payment1"
																	/>
																	<label for="payment1"
																		>Онлайн оплата картой</label
																	>
																</div>
																<div class="payment__item">
																	<input
																		type="radio"
																		name="payment"
																		id="payment2"
																	/>
																	<label for="payment2">Оплата по счету</label>
																</div>
																<div class="payment__item">
																	<input
																		type="radio"
																		name="payment"
																		id="payment3"
																	/>
																	<label for="payment3">Оплата наличными</label>
																</div>
															</div>
														</div>
														<ul class="info-list delivery__info-list">
															<li class="info-list__item">
																<span class="info-list__category"
																	>Количество деталей:</span
																>
																<span class="info-list__value">27 шт</span>
															</li>
															<li class="info-list__item">
																<span class="info-list__category"
																	>Стоимость деталей</span
																>
																<span class="info-list__value">66.312 ₽</span>
															</li>
															<li class="info-list__item info-list__item--big">
																<span class="info-list__category"
																	>Итоговая стоимость</span
																>
																<span class="info-list__value">66.874 ₽</span>
															</li>
														</ul>
														<div class="delivery__buttons">
															<a class="button button--big">Оформить заказ</a>

															<a
																class="button button--border button--big"
																on:click={() => (cancel = true)}
																>Отменить заказ</a
															>
														</div>
													</div>
												</div>
											</div>
										{:else}
											<div class="tabs-content__item tab-7 is-active">
												<ul class="info-list delivery__info-list">
													<li class="info-list__item">
														<span class="info-list__category">Стоимость: </span>
														<span class="info-list__value">0 ₽</span>
													</li>
												</ul>
												<div class="delivery__content">
													<div class="delivery__left">
														<div class="delivery__map" id="map" />
														<div class="delivery__contact">
															<p class="delivery__contact-text">
																Телефон: 8 (800) 550-13-29<br />
																Режим работы: Пн-Пт 11.00-19.00
															</p>
														</div>
													</div>
													<div class="delivery__right delivery__right--cancel">
														<div
															class="delivery__overlay"
															class:is-active={cancel}
														>
															<!-- <div class="delivery-cancel">
                                                                    <div class="delivery-cancel__top">
                                                                        span.delivery
                                                                    </div>
                                                                    <div class="delivery-cancel__bottom"></div>
                                                                </div> -->
															<div class="order-delete delivery-delete">
																<div class="order-delete__top">
																	<span class="order-delete__title"
																		>Отменить заказ?</span
																	>
																	<span class="order-delete__text"
																		>После отмены заказа выбранный вами товар
																		удалится из корзины</span
																	>
																</div>
																<div class="order-delete__bottom">
																	<button class="order-delete__button"
																		>Да</button
																	>
																	<button
																		class="order-delete__button"
																		on:click={() => (cancel = false)}
																		>Нет</button
																	>
																</div>
															</div>
														</div>
														<div class="payment">
															<span class="payment__title">Способ оплаты</span>
															<div class="payment__body">
																<div class="payment__item">
																	<input
																		type="radio"
																		name="payment"
																		id="payment1"
																	/>
																	<label for="payment1"
																		>Онлайн оплата картой</label
																	>
																</div>
																<div class="payment__item">
																	<input
																		type="radio"
																		name="payment"
																		id="payment2"
																	/>
																	<label for="payment2">Оплата по счету</label>
																</div>
																<div class="payment__item">
																	<input
																		type="radio"
																		name="payment"
																		id="payment3"
																	/>
																	<label for="payment3">Оплата наличными</label>
																</div>
															</div>
														</div>
														<ul class="info-list delivery__info-list">
															<li class="info-list__item">
																<span class="info-list__category"
																	>Количество деталей:</span
																>
																<span class="info-list__value">27 шт</span>
															</li>
															<li class="info-list__item">
																<span class="info-list__category"
																	>Стоимость деталей</span
																>
																<span class="info-list__value">66.312 ₽</span>
															</li>
															<li class="info-list__item info-list__item--big">
																<span class="info-list__category"
																	>Итоговая стоимость</span
																>
																<span class="info-list__value">66.874 ₽</span>
															</li>
														</ul>
														<div class="delivery__buttons">
															<a class="button button--big">Оформить заказ</a>

															<a
																class="button button--border button--big"
																on:click={() => (cancel = true)}
																>Отменить заказ</a
															>
														</div>
													</div>
												</div>
											</div>
										{/if}
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
