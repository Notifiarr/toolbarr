<script lang="ts">
  export let starrApp
  export let hidden = false

  import { Container, TabContent, TabPane } from "sveltestrap"
  import Instances from "./Instances/Index.svelte"
  import Database from "./Database/Index.svelte"
  import Source from "./Source/Index.svelte"
  import BGLogo from "../libs/BackgroundLogo.svelte"
</script>


<BGLogo {hidden} url="starr">
  <Container>
    <TabContent class="nav-justified">
      <TabPane tabId="Instances" tab="Instances" active>
        <Instances {starrApp} />
      </TabPane>
      <TabPane tabId="Database" tab="Database">
        <Database {starrApp} />
      </TabPane>
      <TabPane tabId="Source" tab="Source">
        <Source {starrApp} />
      </TabPane>
    </TabContent>
  </Container>
</BGLogo>