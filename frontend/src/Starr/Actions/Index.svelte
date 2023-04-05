<script>
  export let starrApp
  export let hidden = false // Avoid querying backends if hidden.

  import Applogo from "../../libs/Applogo.svelte"
  import { conf } from "../../libs/config.js"
  import T, { _ } from "../../libs/Translate.svelte"
  import Action from "./action.svelte"
  import Tabs, { startTab } from "./tabs.svelte"
  import {
    Accordion,
    AccordionItem,
    Badge,
    Col,
    FormGroup,
    Input,
    InputGroup,
    InputGroupText,
    Row,
  } from "sveltestrap"

  let showTitle = true
  let selected = $conf.Instances[starrApp] ? $conf.Instances[starrApp][0] : undefined
  let instance = selected
  // This trick is to avoid reloading the instance data when expanding the header.
  $: if (selected != instance && selected != undefined) instance = selected

  let width
  let tab = startTab
  $: small = width < 1200
</script>

<svelte:window bind:innerWidth={width}/>

<Accordion>
  <AccordionItem active on:toggle={()=>showTitle=!showTitle}>

    <!-- This is the page title. It's collapsible. -->
    <span slot="header" style="width:95%;">
      <Applogo style="float:right" size="25px" app={starrApp}/>
      <h4 class="d-inline-block accordian-header">{@html $_("instances."+tab.link)}</h4>
      {#if instance}<Badge color="primary">{$_(instance.Name)}</Badge>{/if}
    </span>
    <p><T id="instances.ToolsSelector" {starrApp}/></p>

     <!-- Instance selector menu. -->
     <FormGroup>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Instance")}</InputGroupText>
        <Input invalid={!instance} type="select" bind:value={selected}>
        {#if $conf.Instances[starrApp] != null}
          {#each $conf.Instances[starrApp] as instance}
            <option value={instance}>{instance.Name}: {instance.URL}</option>
          {/each}
          {#if $conf.Instances[starrApp].length == 0}
            <option disabled>- {$_("instances.noInstancesConfigured")} -</option>
          {/if}
        {:else}
          <option disabled>- {$_("instances.noInstancesConfigured")} -</option>
        {/if}
        </Input>
      </InputGroup>
    </FormGroup>

    <!-- Display the nav links in the accordion when the screen is small. -->
    {#if small} <Tabs on:tab={(e) => {tab = e.detail}} fill pills {starrApp}/> {/if}
  </AccordionItem>
</Accordion>

<Row>
  <!-- Display the nav links in the side bar when the screen is not small. -->
  {#if !small}
  <Col xl="2" style="margin-top:8px;">
    <Tabs on:tab={(e) => {tab = e.detail}} vertical pills {starrApp}/>
  </Col>
  {/if}
  <!-- Display the selected tool, pass in selected instance. -->
  <Col xl="10">
    {#if !hidden} <Action {instance} {starrApp} {showTitle} {tab}/> {/if}
  </Col>
</Row>

<style>
  .accordian-header {
    margin-bottom: 0 !important;
  }
</style>