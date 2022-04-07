<script>
 	import Login from '../components/LogReg/Login.svelte';
	import Register from '../components/LogReg/Register.svelte';
	const requiredMessage = "This field is required";
	let change="Registration";
	let login=true;
    function handleClick(event) {
        login=!login;
		if(login==true)
			change="Registration";
		else
			change="Login";
    }
</script>
<div id="logincontent">
		<div id="Background">
		{#if login}
			<Login />
		{:else}
			<Register/>
		{/if}
		</div>
		<div id="Change"><button on:click={handleClick}>{change}</button></div>
        </div>
<style>
	#Background{
		height: 50%;
		width: 50%;
		margin-top: auto;
		background-color: rgba(0, 0, 0, 0.5);

	}
	#Change{
		height: 10%;
		width: 50%;
		margin-bottom: auto;
		background-color: rgba(0, 0, 0, 0.5);
	}
    #logincontent{
        display: flex;
        flex-direction:column;
        align-items: center;
        height: 100vh;
    }
</style>