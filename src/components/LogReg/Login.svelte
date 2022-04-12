<script>
    import {
        useForm,
        Hint,
        HintGroup,
        validators,
        required,
        minLength,
        email,
    } from "svelte-use-form";
    import { fade } from "svelte/transition";
    import {xhr} from "../../Authenticate";
    const baseUrl = window.location.origin;
    const form = useForm();
    let errormsg = "";
    let password = "";
    let mail = ""; 
</script>

<form use:form in:fade>
    <h1 class="display-6 text-center">Chongo</h1>
    <div style="text-align: center;">
        <img
            src="assets/img/logo.png"
            style="width: 100px;height: 100px;margin-right: auto;margin-left: auto;"
            alt="Logo of the chongo app family"
        />
        {#if errormsg != ""}
            <p id="Error" style="color: rgb(242,11,11);font-size: 10px;">
                {errormsg}
            </p>
        {/if}
    </div>
    <div class="mb-3">
        <input
            placeholder="Email"
            bind:value={mail}
            class="form-control"
            type="email"
            name="email"
            use:validators={[required, email]}
        />
    </div>
    <HintGroup for="email">
        <Hint on="required">This is a mandatory field</Hint>
        <Hint on="email" hideWhenRequired>Email is not valid</Hint>
    </HintGroup>
    <div class="mb-3">
        <input
            bind:value={password}
            placeholder="Password"
            class="form-control"
            type="password"
            name="password"
            use:validators={[required]}
        />
    </div>
    <HintGroup for="password">
        <Hint on="required">This is a mandatory field</Hint>
    </HintGroup>
    <button
        type="button"
        on:click="{()=>{
            let data = { email: mail, password: password };
            xhr.open("POST", baseUrl + "/login", true);
            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.onreadystatechange = function () {
                if (this.readyState != 4) return;
                if (this.status == 200) {
                    data = JSON.parse(this.responseText);
                    window.sessionStorage.setItem("sessiontoken", data.token);
                    window.location.href = baseUrl + "/home";
                } else {
                    data = JSON.parse(this.responseText);
                    errormsg = data.message;
                }
                // end of state change: it can be after some time (async)
            };
            xhr.send(JSON.stringify(data));
        }}"
        disabled={!$form.valid}
        class="btn btn-primary d-block w-100"
        style="background: #60c659; margin-bottom:0px !important;">Login</button
    >
</form>

<style>
</style>
