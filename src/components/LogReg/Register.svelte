<script>
	import { useForm, Hint, HintGroup, validators, required, minLength, email } from "svelte-use-form";
	import { passwordMatch, containNumbers } from "../../customValidators";
	import { fade } from 'svelte/transition';
	
	const form = useForm();
	
	const requiredMessage = "This field is required";
</script>

<form use:form in:fade out:fade>
    <h1 class="display-6 text-center">Chongo</h1>
    <div style="text-align: center;"><img src="assets/img/logo.png" style="width: 100px;height: 100px;margin-right: auto;margin-left: auto;" alt="Logo of the chongo app family">
        <p id="Error" style="color: rgb(242,11,11);display: none;font-size: 10px;">Paragraph</p>
    </div>
    <div class="mb-3"><input placeholder="Email" type="email" name="email" use:validators={[required, email]} /></div>
    <HintGroup for="email">
        <Hint on="required">{requiredMessage}</Hint>
        <Hint on="email" hideWhenRequired>This must be a valid email</Hint>	
    </HintGroup>
    <div class="mb-3"><input placeholder="Name" type="text" name="name" use:validators={[required, minLength(5)]} /></div>
    <HintGroup for="name">
        <Hint on="required">{requiredMessage}</Hint>
        <Hint on="minLength" hideWhenRequired let:value>This field must have at least {value} characters.</Hint>
    </HintGroup>
    <div class="mb-3"><input placeholder="Password" type="password" name="password" use:validators={[required, minLength(5), containNumbers(2)]} /></div>
    <HintGroup for="password">
        <Hint on="required">{requiredMessage}</Hint>
        <Hint on="minLength" hideWhenRequired let:value>This field must have at least {value} characters.</Hint>	
        <Hint on="containNumbers" hideWhen="minLength" let:value>
            This field must contain at least {value} numbers.
        </Hint>	
    </HintGroup>
    <div class="mb-3"><input placeholder="Password Again" type="password" name="passwordConfirmation" use:validators={[required, passwordMatch]} /></div>
    <HintGroup for="passwordConfirmation">
        <Hint on="required">{requiredMessage}</Hint>
        <Hint on="passwordMatch" hideWhenRequired>Passwords do not match</Hint>	
    </HintGroup>
    <button disabled={!$form.valid} class="btn btn-primary d-block w-100" style="background: #60c659;margin-bottom:0px !important;">Register</button>
</form>
<style>
</style>