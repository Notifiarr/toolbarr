<script type="ts">
  import golift from "./assets/images/golift.png"
  import starr from "./assets/images/servarr.png"
  import notifiarr from "./assets/images/notifiarr.png"

  export let url = golift

  switch (url) {
    case "starr":     url = starr;     break
    case "notifiarr": url = notifiarr; break
    default:          url = golift
  }
</script>

<div id="background" style="background-image:url({url});"><slot></slot></div>

<style>
  #background  {
    height: 100%;
    background-position: 10px bottom;
    background-size: 100px;
    background-repeat: no-repeat; 
  }
</style>