function ProcessSaveNewUser(){
    var swal = SwalConfimationProcess("Tambah Pengguna", "Tambah pengguna Baru ?", type_question, "Lanjutkan", "Batal");
    swal.then(function(){
        ProcessRequestUserSave();
    });
}

function ProcessUpdateUser(){
    var swal = SwalConfimationProcess("Konfigurasi Pengguna", "Ubah data pengguna ?", type_question, "Lanjutkan", "Batal");
    swal.then(function(){
        ProcessRequestUserSave(2);
    });
}

/*
    1 = save new
    2 = update
*/
function ProcessRequestUserSave(command = 1){
    if(command ==1 && ($("#password").val() != $("#password-ver").val())){
        var swal = SimpleSwal("Opps", "kesalahan pada verifikasi password", type_error, "tutup");
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
    
    StartLoading();
    promise.done(function(response){
        FinishLoading();
        var swal = SimpleSwal("Sukses", "sukses menyimpan data pengguna", type_success, "tutup");
        swal.then(function(){
            GoToIndex("admin/user");
        });
    }).fail(function(response){
        FinishLoading();
        if(response.status==400){
            var swal = SimpleSwal("Opps", "mohon isi seluruh data", type_error, "tutup");
        } else {
            var swal = SimpleSwal("Opps", "silahkan mencoba sekali lagi", type_error, "tutup");
        }
    });
}

function ProcessUploadImgUser(){
    var url = new URL(window.location.href);
    
    var userID = url.searchParams.get("user-id");
    var blobFile = $("#user-img-new")[0].files[0];
    
    var swal = SwalConfimationProcess("Upload Gambar", "Gambar akan di upload. Lanjutkan ?", type_question, "Lanjutkan", "Batal");
    swal.then(function(){
        StartLoading();
        promise = sendUpdateImgUser(userID, blobFile);
        promise.done(function(response){
            FinishLoading();
            var swal = SimpleSwal("Sukses", "sukses menyimpan gambar", type_success, "tutup");
            swal.then(function(){
                location.reload();
            });
        })
        .fail(function(response){
            FinishLoading();
            SimpleSwal("Opps", "silahkan mencoba sekali lagi", type_error, "tutup");
        });
    });    
}