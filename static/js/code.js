text = document.getElementById("textinput")
button = document.getElementById("ask")


async function SendPrompt(){
    try{
        const response = await fetch("http://localhost:9090/ask", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(text.value)
        });

        const result = await response.json();
        console.log("Success", result)

    }catch(error){
        console.error("Error sending data:", error)
    }
}


button.addEventListener('click', SendPrompt);