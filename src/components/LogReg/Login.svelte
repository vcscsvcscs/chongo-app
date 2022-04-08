<script>
	import { useForm, Hint, HintGroup, validators, required, minLength, email } from "svelte-use-form";
	import { fade } from 'svelte/transition';
	
	const form = useForm();
</script>

<form use:form in:fade>
    <h1 class="display-6 text-center">Chongo</h1>
    <div style="text-align: center;"><img src="assets/img/logo.png" style="width: 100px;height: 100px;margin-right: auto;margin-left: auto;" alt="Logo of the chongo app family">
        <p id="Error" style="color: rgb(242,11,11);display: none;font-size: 10px;">Paragraph</p>
    </div>
    <div class="mb-3"><input placeholder="Email" class="form-control" type="email" name="email" use:validators={[required, email]} /></div>
    <HintGroup for="email">
    <Hint on="required">This is a mandatory field</Hint>
    <Hint on="email" hideWhenRequired>Email is not valid</Hint>
    </HintGroup>
    <div class="mb-3"><input placeholder="Password" class="form-control" type="password" name="password" use:validators={[required]} /></div>
    <HintGroup for="password">
        <Hint on="required">This is a mandatory field</Hint>
    </HintGroup><br />
    <div class="mb-3"><button disabled={!$form.valid} class="btn btn-primary d-block w-100" style="background: #60c659;">Login</button></div>
</form>
<style>
</style>