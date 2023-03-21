<script lang="ts">
  export let starrApp
  export let hidden = false
  import { Container, TabContent, TabPane } from "sveltestrap"
  import Instances from "./Instances.svelte"
  import Database from "./Database.svelte"
  import Source from "./Source.svelte"
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