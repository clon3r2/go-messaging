let url = "http://localhost:8080/signup/"

let btn = document.getElementById("signup-submit")
btn.disabled = true;


let phoneInput = document.getElementById("phone")
let usernameInput = document.getElementById("username")
let passwordInput = document.getElementById("password")
phoneInput.oninput = ButtonEnableChecker
usernameInput.oninput = ButtonEnableChecker
passwordInput.oninput = ButtonEnableChecker

function ButtonEnableChecker() {
    btn.disabled = !(phoneInput.value !== "" && usernameInput.value !== "" && passwordInput.value !== "")
}

function ClearForm(){
    phoneInput.value = ""
    usernameInput.value = ""
    passwordInput.value = ""
}

async function Signup() {
    console.log("usernameInput.value ==> ", usernameInput.value)
    console.log("passwordInput.value ==> ", passwordInput.value)
    console.log("phoneInput.value ==> ", phoneInput.value)
    try {
        const response = await fetch(
            url,
            {
                method: "POST",
                headers: {
                    "Content-type": "application/json",
                },
                body: JSON.stringify({
                    username: usernameInput.value,
                    password: passwordInput.value,
                    phone: phoneInput.value
                })
            }
        );
        console.log("before response")
        console.log("response==> ", response)

        ClearForm()
    } catch (e) {
        console.log("err ==> ", e)
        alert(e)
        ClearForm()
    }


}
