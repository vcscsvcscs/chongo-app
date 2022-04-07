<script>
	import { useForm, Hint, HintGroup, validators, required, minLength, email } from "svelte-use-form";
	import { fade } from 'svelte/transition';
	
	const form = useForm();
</script>

<form use:form class="LogReg" in:fade out:fade>
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
form{
		color: aliceblue;
	}
</style>