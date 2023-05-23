<script>
    let files;
    function fileInputChange(evt){
        console.log(evt)
        console.log(files)
    }

    function uploadFile(){
        if(files && files[0]){
            console.log(files)
        }
    }
</script>

<div class="w-1/2 h-auto bg-neutral-100 absolute m-auto left-0 right-0 border-stone-200 border-2 rounded p-10">
    
    <div
    class="w-full mb-2 inline-block text-neutral-700"
    >Upload GTFS.zip archive to get started. You can find examples <a href="https://transitfeeds.com/" target="new">here</a></div>

    <input class="relative m-0 block w-4/5 min-w-0 inline rounded border bg-neutral-50 border-solid border-neutral-300 px-3 py-[0.32rem] text-neutral-700 file:-mx-3 file:-my-[0.32rem] file:overflow-hidden  file:border-0 file:border-solid file:bg-neutral-50 file:px-3 file:py-[0.32rem] file:[border-inline-end-width:1px] file:[margin-inline-end:0.75rem] hover:file:text-black
    hover:file:bg-neutral-100" type="file" bind:files accept=".zip,json" on:change={fileInputChange}/>
    <button class="w-1/6 inline rounded border border-solid border-neutral-300 font-normal text-neutral-700 transition px-3 py-[0.32rem]" on:click={uploadFile}>Upload</button>

</div>