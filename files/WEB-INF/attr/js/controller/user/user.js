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