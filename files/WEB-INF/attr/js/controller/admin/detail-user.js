function ProcessSaveNewUser(){
    ProcessRequestUserSave();
}

function ProcessUpdateUser(){
    ProcessRequestUserSave(2);
}

/*
    1 = save new
    2 = update
*/
function ProcessRequestUserSave(command = 1){
    if(command ==1 && ($("#password").val() != $("#password-ver").val())){
        $("#request-alert-div").css("display", "block");
        $("#request-alert").html("kesalahan pada verifikasi password");
        return;
    }
    
    var fullname = $("#fullname").val();
    var username = $("#username").val();
    var password = $("#password").val();
    var status = $("#status").val();
    var type = $("#type").val();

    if(command == 1){
        promise = sendNewUser(fullname, username, password, status, type);
    }
    if(command == 2){
        var url = new URL(window.location.href);
        var userID = url.searchParams.get("user-id");
        promise = sendUpdateUser(userID, fullname, username, password, status, type);
    }
    
    promise.done(function(response){
        GoToIndex("admin/user");
    }).fail(function(response){
        if(response.status==400){
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("mohon isi seluruh data");
        } else {
            $("#request-alert-div").css("display", "block");
            $("#request-alert").html("silahkan mencoba sekali lagi");
        }
    });
}

function ProcessUploadImgUser(){
    var url = new URL(window.location.href);
    
    var userID = url.searchParams.get("user-id");
    var blobFile = $("#user-img-new")[0].files[0];
    
    promise = sendUpdateImgUser(userID, blobFile);
    promise.done(function(response){
        GoToIndex("admin/user");
    })
    .fail(function(response){
        $("#request-alert-div").css("display", "block");
        $("#request-alert").html("silahkan mencoba sekali lagi");
    });
}