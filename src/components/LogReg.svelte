<script>
	import { useForm, Hint, HintGroup, validators, required, minLength, email } from "svelte-use-form";
	import { passwordMatch, containNumbers } from "../customValidators";
	import { fade } from 'svelte/transition';
	
	const form = useForm();
	
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
		{#if login==false}

			<form use:form class="LogReg" in:fade out:fade>
				<h1>
					Registration
				</h1>
				<label for="email">Email</label>
				<input type="email" name="email" use:validators={[required, email]} />
				<HintGroup for="email">
					<Hint on="required">{requiredMessage}</Hint>
					<Hint on="email" hideWhenRequired>This must be a valid email</Hint>	
				</HintGroup>

				<label for="name">Name</label>
				<input type="text" name="name"  />

				<label for="password">Password</label>
				<input type="password" name="password" use:validators={[required, minLength(5), containNumbers(2)]} />
				<HintGroup for="password">
					<Hint on="required">{requiredMessage}</Hint>
					<Hint on="minLength" hideWhenRequired let:value>This field must have at least {value} characters.</Hint>	
					<Hint on="containNumbers" hideWhen="minLength" let:value>
						This field must contain at least {value} numbers.
					</Hint>	
				</HintGroup>

				<label for="passwordConfirmation">Password Confirmation</label>
				<input type="password" name="passwordConfirmation" use:validators={[required, passwordMatch]} />
				<HintGroup for="passwordConfirmation">
					<Hint on="required">{requiredMessage}</Hint>
					<Hint on="passwordMatch" hideWhenRequired>Passwords do not match</Hint>	
				</HintGroup><br />

				<button disabled={!$form.valid} on:click|preventDefault>
					Submit
				</button>
			</form>
			
		{/if}
		{#if login == true}
			<form use:form class="LogReg"in:fade out:fade>
				<h1>Login</h1>
				<label for="email">Email</label>
				<input type="email" name="email" use:validators={[required, email]} />
				<HintGroup for="email">
				<Hint on="required">This is a mandatory field</Hint>
				<Hint on="email" hideWhenRequired>Email is not valid</Hint>
				</HintGroup>
				<label for="password">Password</label>
				<input type="password" name="password" use:validators={[required]} />
				<HintGroup for="password">
					<Hint on="required">This is a mandatory field</Hint>
				</HintGroup><br />
			
				<button disabled={!$form.valid}>Login</button>
			</form>
			
		{/if}
		</div>
		<div id="Change"><button on:click={handleClick}>{change}</button></div>
        </div>
<style>
	h1 {
		color: #0357f1;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
		margin:0px !important;
	}
	.LogReg{
		margin: auto;
		width: 100%;

	}
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
	form{
		color: aliceblue;
	}
    #logincontent{
        display: flex;
        flex-direction:column;
        align-items: center;
        height: 100vh;
    }
</style>