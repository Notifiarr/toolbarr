<script>
  export let form
  export let updating
  export let starrApp
  export let field
  export let name = ""

  import T, { _ } from "../../../libs/Translate.svelte"
  import Fa from "svelte-fa"
  import { faCircleInfo } from "@fortawesome/free-solid-svg-icons"
  import { Dropdown, DropdownItem, DropdownMenu, DropdownToggle, Tooltip } from "sveltestrap"

  $: title = name ? $_(name) : $_(`instances.${field}Title`)
  let dropdown

  function toggleAll(key, on) {
    form.forEach((_, i) => {
      form[i][key] = on
    })
  }
</script>

<Tooltip placement="top" target={dropdown}>
  <T id={`instances.${field}Desc`} {starrApp}/>
</Tooltip>
<Dropdown size="sm">
  <DropdownToggle tag="span" class="link">
    <span bind:this={dropdown}>
      {title.match(/\b(\w)/g).join('').substring(0,2)} <Fa primaryColor="darkCyan" icon={faCircleInfo}/>
    </span>
  </DropdownToggle>
  <DropdownMenu>
    <DropdownItem header>{title}</DropdownItem>
    <DropdownItem disabled={updating} on:click={() => toggleAll(field, false)}>
      {$_("instances.DisableAll")}
    </DropdownItem>
    <DropdownItem disabled={updating} on:click={() => toggleAll(field, true)}>
      {$_("instances.EnableAll")}
    </DropdownItem>
  </DropdownMenu>
</Dropdown>
