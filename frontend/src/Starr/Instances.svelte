<script>
    import { Input, InputGroup, InputGroupText } from "sveltestrap"
    export let app
</script>

<h3>{app} Instance Info</h3>
<p>This page contains a list of instances that can be updated. You may provide api key and/or login credentials. Click the buttons to fetch information from the instance. Information such as queue, library, stats, etc.</p>
<InputGroup>
    <InputGroupText>This input doesn't do anything</InputGroupText>
    <Input />
</InputGroup>
