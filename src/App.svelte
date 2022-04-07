<script>
<<<<<<< Updated upstream
	import Router from "./routes/index.svelte";
</script>

<main>
	<Router />
</main>

=======
	import { useForm, Hint, HintGroup, validators, required, minLength, email } from "svelte-use-form";
	import { passwordMatch, containNumbers } from "./customValidators";
	import { fade } from 'svelte/transition';
	
	const form = useForm();
	
	const requiredMessage = "This field is required";
	let change="Login";
	let login=true;
    function handleClick(event) {
        login=!login;
		if(login==true)
			change="Registration";
		else
			change="Login";

    }
</script>
	<main>
		<div id="Background">
		{#if login==false}
			 <div id="LogReg" out:fade>
			<form use:form >
				<h1>Registration</h1>
				<label for="email">Email</label>
				<input type="email" name="email" value="" use:validators={[required, email]} />
				<HintGroup for="email">
					<Hint on="required">This fild is mandatory</Hint>
					<Hint on="email" hideWhenRequired>This must be a valid email</Hint>	
				</HintGroup>

				<label for="name">Name</label>
				<input type="text" name="name" value="" use:validators={[required]}/>
				<HintGroup for="name">
					<Hint on="required">This fild is mandatory</Hint>
				</HintGroup>

				<label for="password">Password</label>
				<input type="password" name="password" use:validators={[required, minLength(5), containNumbers(2)]} />
				<HintGroup for="password">
					<Hint on="required">This fild is mandatory</Hint>
					<Hint on="minLength" hideWhenRequired let:value>This field must have at least {value} characters.</Hint>	
					<Hint on="containNumbers" hideWhen="minLength" let:value>
						This field must contain at least {value} numbers.
					</Hint>	
				</HintGroup>

				<label for="passwordConfirmation">Password Confirmation</label>
				<input type="password" name="passwordConfirmation" use:validators={[required, passwordMatch]} />
				<HintGroup for="passwordConfirmation">
					<Hint on="required">This fild is mandatory</Hint>
					<Hint on="passwordMatch" hideWhenRequired>Passwords do not match</Hint>	
				</HintGroup><br />

				<button on:click|preventDefault>Submit</button>
			</form>
			</div>
			
		{/if}
		{#if login == true}
			<div id="LogReg" out:fade>
			<form use:form >
				<h1>Login</h1>
				<label for="email">Email</label>
				<input type="email" name="email" use:validators={[required, email]} />
				<HintGroup for="email">
					<Hint on="required">This fild is mandatory</Hint>
					<Hint on="email" hideWhenRequired>Email is not valid</Hint>
				</HintGroup>
				<label for="passwordl">Password</label>
				<input type="password" name="passwordl" use:validators={[required]} />
				<HintGroup for="passwordl">
					<Hint on="required">This fild is mandatory</Hint>
				</HintGroup><br />
			
				<button on:click|preventDefault>Login</button>
			</form>
			</div>
		{/if}
		</div>
		<div id="Change"><button on:click={handleClick}>{change}</button></div>
		
	</main>
>>>>>>> Stashed changes
<style>
	main {
		text-align: center;
		padding: 0px;
		margin: 0 auto;
		height: 100vh;
		background-image: url("/assets/Nudli_Family.jpg");
		background-size: cover;
	}
</style>
