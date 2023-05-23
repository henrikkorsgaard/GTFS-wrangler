<script>

  import FileModal from './lib/FileModal.svelte';

</script>

<style>
    :global(body){
      height:100%;
      width:100%;
      margin:0px auto;
    }

    :global(#app){
      height:100%;
      width:100%;
      margin:0px auto;
    }
</style>

<FileModal></FileModal>

<!--
<div class="w-full fixed top-0 bottom-0 left-0 right-0 bg-zinc-50" ></div>
<div class="sidebar fixed top-0 bottom-0 w-64 inset-0 bg-gradient-to-b from-teal-950 via-teal-950 to-black">
<div class="text-slate-100 font-sans text-left text-xl m-2 p-2 pb-1 select-none hover:bg-teal-900" on:click={handleUploadClick} on:keydown={handleUploadClick}><Icon src="{CircleStack}" class="w-8 h-8 inline m-0 mb-1 mr-2" />Load GTFS.zip</div>
</div>-->

