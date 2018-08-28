function ProcessLogin(){
    username = $("#username").val();
    secret = $("#secret").val();

    promise = makeLogin(username, secret);

    promise.done(function(response){
        window.location.href = base_url;
    }).fail(function(response){
        $("#username").val("");
        $("#secret").val("");
        if(response.status >= 500){
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("silahkan mencoba sekali lagi");
        } else {
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("Username atau Password salah");
        }
    });
}

function ProcessLogout(){
    promise = makeLogout();

    promise.done(function(response){
        window.location.href = base_url;
    }).fail(function(response){
        location.reload();
    });
}

function ProcessRegister(){
    if($("#password").val() != $("#password-ver").val()){
        $("#request-alert-div").css("display", "block");
        $("#request-alert").html("kesalahan pada verifikasi password");
        return;
    } 

    fullname = $("#fullname").val();
    email = $("#email").val();
    password = $("#password").val();

    if(password.length <= 9){
        $("#request-alert-div").css("display", "block");
        $("#request-alert").html("Password kurang dari 10 karakter");
        return;
    }

    if(!validateEmail(email)){
        $("#request-alert-div").css("display", "block");
        $("#request-alert").html("Email tidak valid");
        return;
    }

    promise = registerUser(fullname, email, password);
    
    promise.done(function(response){
        window.location.href = base_url;
    }).fail(function(response){
        $("#username").val("");
        $("#secret").val("");
        if(response.status >= 500){
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("silahkan mencoba sekali lagi");
        } else {
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html(response.data.error_massage);
        }
    });
    
}

function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}