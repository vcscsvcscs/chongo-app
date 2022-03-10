<script>
    import { Router, Route } from 'svelte-routing';
    import Home from '../components/Home.svelte';
    import LogReg from '../components/LogReg.svelte';
  
    export let url = '';
  </script>
  
  <Router {url}>
      <Route path="home" component={Home} />
      <Route path="/" component={LogReg} />
  </Router>