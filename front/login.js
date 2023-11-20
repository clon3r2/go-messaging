let usernameInput = document.getElementById("username")
let passwordInput = document.getElementById("password")
let btn = document.getElementById("login-submit")
btn.disabled = true;
const url = "http://localhost:8080/login/"

function ButtonEnableChecker() {
    btn.disabled = !(usernameInput.value !== "" && passwordInput.value !== "")
}

usernameInput.oninput = ButtonEnableChecker
passwordInput.oninput = ButtonEnableChecker
function ClearForm(){
    usernameInput.value = ""
    passwordInput.value = ""
}
async function Login() {
    try {
        let response = await fetch(
            url,
            {
                method: "POST",
                headers: {"Content-type": "application/json"},
                body: JSON.stringify({
                        username: usernameInput.value,
                        password: passwordInput.value
                    }
                )
            }
        )
        ClearForm()
        console.log("success!")
        console.log(response.json())
    } catch (e) {
        console.log("ERR!")
        console.log(e)
        alert(e)
        ClearForm()
    }


}