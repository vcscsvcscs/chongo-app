<script>
    import { fade } from "svelte/transition";
    let displayscrolltotop = false;
</script>
<svelte:window on:scroll="{displayscrolltotop = true}"></svelte:window>
{#if displayscrolltotop}
<a on:click="{()=>{setTimeout(() => {
    displayscrolltotop = false;
}, 100);}}" class="scroll-to-top" in:fade href="#page-top">
    <i class="fas fa-angle-up" />
</a>
{/if}
<style>
    .scroll-to-top {
        width: 3rem;
        height: 3rem;
        position: fixed;
        font-size: 2rem;
        top: 85vh;
        right: 5vw;
        color:azure;
        background-color: teal;
        border: none;
        border-radius: 1rem;
    }
</style>
