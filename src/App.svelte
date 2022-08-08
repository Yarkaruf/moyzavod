<script>
	import { Router, Link, Route } from "svelte-routing";
	import Register from "./Register.svelte";
	import Login from "./Login.svelte";
	import Recovery from "./Recovery.svelte";
	import Service from "./Service.svelte";
	import Viewer from "./Viewer.svelte";
	import Configuration from "./Configuration.svelte";
	import Admin from "./Admin.svelte";
	import Account from "./Account.svelte";
	import AccountEdit from "./AccountEdit.svelte";
	import Orders from "./Orders.svelte";
	import Files from "./Files.svelte";
	import Folder from "./Folder.svelte";
	import Partner from "./Partner.svelte";
	import RegistOrder from "./RegistrationOrder.svelte";

	export let name;
	export let url = "";
	import axios from "axios";

	function checkAuth() {
		console.log("auth");
		axios
			.get("/api/user")
			.then(function (resp) {
				window.location.href = "/service";
			})
			.catch(function (err) {
				console.log(err);
				window.location.href = "/login";
			});
	}
	if (window.location.pathname == "/") {
		checkAuth();
	}
</script>

<main>
	<Router {url}>
		<div>
			<Route path="register" component={Register} />
			<Route path="files" component={Files} />
			<Route path="orders" component={Orders} />
			<Route path="login" component={Login} />
			<Route path="recovery" component={Recovery} />
			<Route path="service" component={Service} />
			<Route path="viewer" component={Viewer} />
			<Route path="configuration" component={Configuration} />
			<Route path="admin" component={Admin} />
			<Route path="account" component={Account} />
			<Route path="account-edit" component={AccountEdit} />
			<Route path="files-folder" component={Folder} />
			<Route path="partner" component={Partner} />
			<Route path="regist-order" component={RegistOrder} />
		</div>
	</Router>
</main>

<style>
</style>
