<script>
	import Login from "../components/LogReg/Login.svelte";
	import Register from "../components/LogReg/Register.svelte";
	import {UserLogedin} from "../Authenticate"
	const baseUrl = window.location.origin;
	if(UserLogedin()){
		window.location.href = baseUrl + "/home";
	};
	let change = "Registration";
	let login = true;
	function handleClick(event) {
		login = !login;
		if (login == true){ 
			change = "Registration";
		}
		else {
			change = "Login";
		}
	}
</script>

<link rel="icon" type="image/png" sizes="500x500" href="assets/img/logo.png" />
<link rel="stylesheet" href="assets/css/styles.css" />
<link rel="stylesheet" href="assets/css/Login-Form-Dark.css" />
<section class="login-dark">
	{#if login}
		<Login />
	{:else}
		<Register />
	{/if}
	<button type="button" on:click={handleClick} class="{login ? 'btn btn-primary d-block changer':'btn btn-primary d-block changer1'}"
		>{change}</button>
</section>