<script>
    let files;
    let buttonDisabled = true;
    function fileInputChange(evt){
        if(files && files[0]){
            buttonDisabled = false;
        }
    }

    async function uploadFile(){
        
        if(files && files[0]){

            // actually, I'd rather do a websocket connection here because I cannot get the progress in unpacking the file etc. 
            var formData = new FormData()
            formData.append("file", files[0])
            console.log(files[0])

            //There is something with the flow here!
            let socket = new WebSocket("ws://localhost:3000/ws/gtfs")

            socket.onopen  = function(evt) {
                console.log("websocket connection opened!")
                let file = files[0];
                let reader = new FileReader();
            
                reader.onload = function(evt){
                    let buffer = evt.target.result
                    
                    socket.binaryType = "blob"
                    socket.send(buffer)
                }

                reader.onerror = function(err){
                    console.log(err)
                }
            
                reader.readAsArrayBuffer(file);
            }

            socket.onmessage = function(evt){
                console.log(evt)
            }
           
            socket.onerror = function(err){
                console.error(err)
            }

            socket.onclose = function(){
                console.error("Websocket connections closeds")
            }


            
            /*
            name: "GTFS.zip"
​
            size: 54812635
​
            type: "application/zip"
            */
            /*
            const res = await fetch('http://localhost:3000/gtfs', {
		    	method: 'POST',
                headers: {},
			    body:formData
            })*/
            
            //console.log(res)
        }
    }   
</script>

<div class="w-1/2 h-auto bg-neutral-100 absolute m-auto left-0 right-0 border-stone-200 border-2 rounded p-10">
    
    <div
    class="w-full mb-2 inline-block text-neutral-700"
    >Upload GTFS.zip archive to get started. You can find examples <a href="https://transitfeeds.com/" target="new">here</a></div>

    <input class="relative m-0 block w-4/5 min-w-0 inline rounded border bg-neutral-50 border-solid border-neutral-300 px-3 py-[0.32rem] text-neutral-700 file:-mx-3 file:-my-[0.32rem] file:overflow-hidden  file:border-0 file:border-solid file:bg-neutral-50 file:px-3 file:py-[0.32rem] file:[border-inline-end-width:1px] file:[margin-inline-end:0.75rem] hover:file:text-black
    hover:file:bg-neutral-100" type="file" bind:files accept=".zip,.md" on:change={fileInputChange}/>
    <button disabled='{buttonDisabled}' class="w-1/6 inline rounded border border-solid border-neutral-300 font-normal text-neutral-700 transition px-3 py-[0.32rem] disabled:opacity-25" on:click={uploadFile} >Upload</button>

</div>