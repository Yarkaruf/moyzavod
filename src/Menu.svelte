<script type="text/javascript">
	import { Link } from "svelte-routing";
	let opened = true;
</script>

<div>
	<div class="menu" class:menu--opened={opened}>
		<ul class="menu__list">
			<li
				class="menu__item"
				class:menu__item--active={window.location.pathname.startsWith(
					"/service"
				)}
			>
				<Link to="service" style="text-decoration: none;">
					<a href="#" class="menu__link">
						<div class="menu__icon">
							<svg
								width="20"
								height="20"
								viewBox="0 0 20 20"
								fill="none"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									d="M12.0909 1H19V7.90909H12.0909V1ZM12.0909 12.0909H19V19H12.0909V12.0909ZM1 1H7.90909V7.90909H1V1ZM1 12.0909H7.90909V19H1V12.0909Z"
									stroke="white"
									stroke-linecap="round"
									stroke-linejoin="round"
								/>
							</svg>
						</div>
						<span class="menu__text">Новый расчет</span>
					</a>
				</Link>
			</li>
			<li
				class="menu__item"
				class:menu__item--active={window.location.pathname.startsWith("/files")}
			>
				<Link to="files" style="text-decoration: none;">
					<a href="#" class="menu__link">
						<div class="menu__icon menu__icon--fill">
							<svg
								width="24"
								height="22"
								viewBox="0 0 24 22"
								fill="none"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									fill-rule="evenodd"
									clip-rule="evenodd"
									d="M16.8 21.5H2.7C1.2105 21.5 0 20.2895 0 18.8V7.25C0 6.008 1.008 5 2.25 5H3V14.693C3 16.7945 4.7055 18.5 6.807 18.5H19.5V18.8C19.5 20.2895 18.291 21.5 16.8 21.5ZM21.3 17H7.2C5.7105 17 4.5 15.7895 4.5 14.3V2.75C4.5 1.508 5.508 0.5 6.75 0.5H12.75C13.992 0.5 15 1.508 15 2.75V3.5H21.3C22.791 3.5 24 4.7105 24 6.2V14.3C24 15.7895 22.791 17 21.3 17Z"
									fill="white"
								/>
							</svg>
						</div>
						<span class="menu__text">Мои файлы</span>
					</a>
				</Link>
			</li>
			<li
				class="menu__item"
				class:menu__item--active={window.location.pathname.startsWith("/admin")}
			>
				<Link to="admin" style="text-decoration: none;">
					<a href="#" class="menu__link">
						<div class="menu__icon menu__icon--fill">
							<svg
								width="24"
								height="22"
								viewBox="0 0 24 22"
								fill="none"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									fill-rule="evenodd"
									clip-rule="evenodd"
									d="M16.8 21.5H2.7C1.2105 21.5 0 20.2895 0 18.8V7.25C0 6.008 1.008 5 2.25 5H3V14.693C3 16.7945 4.7055 18.5 6.807 18.5H19.5V18.8C19.5 20.2895 18.291 21.5 16.8 21.5ZM21.3 17H7.2C5.7105 17 4.5 15.7895 4.5 14.3V2.75C4.5 1.508 5.508 0.5 6.75 0.5H12.75C13.992 0.5 15 1.508 15 2.75V3.5H21.3C22.791 3.5 24 4.7105 24 6.2V14.3C24 15.7895 22.791 17 21.3 17Z"
									fill="white"
								/>
							</svg>
						</div>
						<span class="menu__text">Админка</span>
					</a>
				</Link>
			</li>
		</ul>
		<button
			class="menu-state"
			class:menu-state--active={opened}
			on:click={() => {
				opened = !opened;
			}}
		>
			<div class="menu-state__icon">
				<svg
					width="20"
					height="20"
					viewBox="0 0 20 20"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
				>
					<path
						fill-rule="evenodd"
						clip-rule="evenodd"
						d="M15 2.5H5C3.61929 2.5 2.5 3.61929 2.5 5V15C2.5 16.3807 3.61929 17.5 5 17.5H15C16.3807 17.5 17.5 16.3807 17.5 15V5C17.5 3.61929 16.3807 2.5 15 2.5ZM5 1.25C2.92893 1.25 1.25 2.92893 1.25 5V15C1.25 17.0711 2.92893 18.75 5 18.75H15C17.0711 18.75 18.75 17.0711 18.75 15V5C18.75 2.92893 17.0711 1.25 15 1.25H5Z"
						fill="#ADADAD"
					/>
					<path
						fill-rule="evenodd"
						clip-rule="evenodd"
						d="M12.2413 9.37469L6.25 9.37469C5.905 9.37469 5.625 9.65469 5.625 9.99969C5.625 10.3447 5.905 10.6247 6.25 10.6247L12.2413 10.6247L10.6094 12.2566C10.365 12.5003 10.365 12.8966 10.6094 13.1403C10.8531 13.3847 11.2494 13.3847 11.4931 13.1403L14.185 10.4491C14.1869 10.4466 14.1894 10.4441 14.1919 10.4416C14.3138 10.3197 14.375 10.1597 14.375 9.99969C14.375 9.83969 14.3138 9.67969 14.1919 9.55781C14.1894 9.55531 14.1869 9.55281 14.185 9.55031L11.4931 6.85906C11.2494 6.61469 10.8531 6.61469 10.6094 6.85906C10.365 7.10281 10.365 7.49906 10.6094 7.74281L12.2413 9.37469Z"
						fill="#ADADAD"
					/>
				</svg>
			</div>
			<span class="menu-state__text">Свернуть</span>
		</button>
	</div>
	<div class="menu-overlay" />
</div>
