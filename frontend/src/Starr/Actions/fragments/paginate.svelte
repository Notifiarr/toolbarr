<!-- Custom reusable pagination with Sveltestrap. -->

<script lang="ts">
  import {
    Pagination,
    PaginationItem,
    PaginationLink,
    DropdownItem,
    DropdownMenu,
    DropdownToggle,
    ButtonDropdown,
    Tooltip,
  } from "@sveltestrap/sveltestrap"
  import { createEventDispatcher } from "svelte"
  import T, { _ } from "/src/libs/Translate.svelte"

  export let page: number = 1
  export let pageSize: number = 25
  export let pages: number = 1
  export let records: number
  export let total: number
  export let updating: boolean = false
  
  const dispatch = createEventDispatcher()
  const pageSizes = [15,25,35,50,100,200,500]
  let endPage: number = 1
  let pagePicker: boolean = false
  let sizePicker: boolean = false

  setPage(null, page)

  function max(num: number, max: number) {
    if (num < max) return num
    return max
  }

  function setPage(e: any, newPage: number) {
    if (e) e.preventDefault()
    endPage = max(newPage<4 ? 6 : newPage+2 > pages ? pages : newPage+2, pages)
    if (page == newPage) return
    page = newPage
    dispatch("update", true)
  }

  function startPage(): number {
    return endPage-5 > 0 ? endPage-5 : 1
  }

  function setSize(e: MouseEvent) {
    e.preventDefault()
    const target = e.target as HTMLSelectElement;
    const newSize = parseInt(target.value)
    if (newSize == pageSize) return
    page = 1 // need to start the page over too.
    pageSize = newSize
    dispatch("update", true)
  }
  
  // If all the items on the page get deleted, go back 1 page and refresh.
  $: if (records == 0 && total > 0) {
    if (page > 1) page--
    dispatch("update", true)
  }
</script>

<div id="container">
  <Pagination listClassName="mb-0" size="sm">
    <PaginationItem disabled={updating||page<2}>
      <Tooltip target="first">{$_("paginate.FirstPage")}</Tooltip>
      <PaginationLink id="first" first on:click={(e)=>setPage(e,1)}/>
    </PaginationItem>
    <PaginationItem disabled={updating||page<2}>
      <Tooltip target="previous">{$_("paginate.PreviousPage")}</Tooltip>
      <PaginationLink id="previous" previous on:click={(e)=>setPage(e,page-1)}/>
    </PaginationItem>

    {#each Array.from(Array(endPage+1).keys()).slice(startPage()) as i}
    <PaginationItem disabled={updating} active={page==i}>
      {#if page == i}
      <PaginationLink class="button" on:click={(e)=>{e.preventDefault();pagePicker=!pagePicker}}>
        <Tooltip target="page"><T id="paginate.Pageof" {page} {pages}/></Tooltip>
        <ButtonDropdown bind:isOpen={pagePicker}>
          <DropdownToggle id="page" tag="span" caret>{page}</DropdownToggle>
          <DropdownMenu>
            <DropdownItem header><T id="paginate.Pageof" {page} {pages}/></DropdownItem>
            {#each Array.from(Array(pages+1).keys()).slice(1) as i}
              <DropdownItem value={i} on:click={(e)=>setPage(e,i)} active={page==i}>{i}</DropdownItem>
            {/each}
          </DropdownMenu>
        </ButtonDropdown>
      </PaginationLink>
      {:else}
      <PaginationLink on:click={(e)=>setPage(e,i)}>{i}</PaginationLink>
      {/if}
    </PaginationItem>
    {/each}

    <PaginationItem disabled={updating||page>=pages}>
      <Tooltip target="next">{$_("paginate.NextPage")}</Tooltip>
      <PaginationLink id="next" next on:click={(e)=>setPage(e,page+1)}/>
    </PaginationItem>
    <PaginationItem disabled={updating||page>=pages}>
      <Tooltip target="last"><T id="paginate.LastPage" {pages}/></Tooltip>
      <PaginationLink id="last" last on:click={(e)=>setPage(e,pages)}/>
    </PaginationItem>

    <PaginationItem disabled={updating}>
      <Tooltip target="items">{$_("paginate.ItemsPerPage")}</Tooltip>
      <PaginationLink class="button" on:click={(e)=>{e.preventDefault();sizePicker=!sizePicker}}>
        <ButtonDropdown bind:isOpen={sizePicker}>
          <DropdownToggle id="items" caret tag="span">{pageSize}</DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("paginate.ItemsPerPage")}</DropdownItem>
              {#if !pageSizes.includes(pageSize)}
              <DropdownItem value={pageSize} on:click={setSize} active>{pageSize}</DropdownItem>
              {/if}
              {#each pageSizes as i}
              <DropdownItem value={i} on:click={setSize} active={pageSize==i}>{i}</DropdownItem>
              {/each}
          </DropdownMenu>
        </ButtonDropdown>
      </PaginationLink>
    </PaginationItem>
  </Pagination>

  <i style="font-size:0.7em"><T id="paginate.RecordsOf" {records} {total}/></i>
</div>

<style>
  #container :global(.button) {
    padding-top: 2px;
  }

  #container :global(.mb-0) {
    margin-bottom: -6px !important;
  }
  
  #container {
    margin: 4px auto 0px auto;
    width: max-content;
    text-align: center;
    line-height: 1.5em;
  }
</style>
