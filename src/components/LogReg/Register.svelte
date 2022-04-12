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
    import { passwordMatch, containNumbers } from "../../customValidators";
    import {xhr} from "../../Authenticate";
    import { fade } from "svelte/transition";
    const form = useForm();
    let errormsg = "";
    let password = "";
    let mail = "";
    let username = "";
    let name = "";
</script>

<form use:form class="LogReg" in:fade>
    <h1 class="display-6 text-center">Chongo</h1>
    <div style="text-align: center;">
        <img
            src="assets/img/logo.png"
            style="width: 100px;height: 100px;margin-right: auto;margin-left: auto;"
            alt="Logo of the chongo app family"
        />
        {#if errormsg != ""}
            <p
                id="Error"
                style="color: rgb(242,11,11);display: none;font-size: 10px;"
            >
                {errormsg}
            </p>
        {/if}
    </div>
    <div class="mb-3">
        <input
            placeholder="Email"
            class="form-control"
            type="email"
            bind:value="{mail}"
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
            placeholder="Name"
            class="form-control"
            type="text"
            name="name"
            bind:value="{name}"
            use:validators={[required, minLength(5)]}
        />
    </div>
    <HintGroup for="name">
        <Hint on="required">This is a mandatory field</Hint>
        <Hint on="minLength" hideWhenRequired let:value>This field must have at least {value} characters.</Hint>
    </HintGroup>
    <div class="mb-3">
        <input
            placeholder="Username"
            class="form-control"
            type="text"
            name="Username"
            bind:value="{username}"
            use:validators={[required, minLength(5)]}
        />
    </div>
    <HintGroup for="Username">
        <Hint on="required">This is a mandatory field</Hint>
    </HintGroup>
    <div class="mb-3">
        <input
            placeholder="Password"
            class="form-control"
            type="password"
            name="password"
            bind:value="{password}"
            use:validators={[required, minLength(5), containNumbers(2)]}
        />
    </div>
    <HintGroup for="password">
        <Hint on="required">This is a mandatory field</Hint>
        <Hint on="minLength" hideWhenRequired let:value
            >This field must have at least {value} characters.</Hint
        >
        <Hint on="containNumbers" hideWhen="minLength" let:value>
            This field must contain at least {value} numbers.
        </Hint>
    </HintGroup>

    <div class="mb-3">
        <input
            placeholder="Password Confirmation"
            class="form-control"
            type="password"
            name="passwordConfirmation"
            use:validators={[required, passwordMatch]}
        />
    </div>
    <HintGroup for="passwordConfirmation">
        <Hint on="required">This is a mandatory field</Hint>
        <Hint on="passwordMatch" hideWhenRequired>Passwords do not match</Hint>
    </HintGroup>
    <button
        type="button"
        disabled={!$form.valid}
        class="btn btn-primary d-block w-100"
        style="background: #60c659;margin-bottom:0px !important;"
        on:click="{()=>{
            let baseUrl = window.location.origin;
        let data = {
            username: username,
            name: name,
            email: mail,
            password: password,
        };
        xhr.open("POST", baseUrl + "/register", true);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.onreadystatechange = function () {
            if (this.readyState != 4) return;
            if (this.status == 201) {
                data = JSON.parse(this.responseText);
                window.sessionStorage.setItem("sessiontoken", data.token);
                window.location.href = baseUrl + "/home";
            } else {
                data = JSON.parse(this.responseText);
                errormsg = data.message;
            }
            // end of state change: it can be after some time (async)
        };
        //console.log(JSON.stringify(data));
        xhr.send(JSON.stringify(data));
        }}"
        >Register</button
    >
</form>

<style>
</style>
