function ProcessRegisterCompany(){
    var companyType = $("#company-type").val();
    var companyName = $("#company-name").val();

    if(companyName=="" || companyType==""){
        SimpleSwal("Gagal melakukan pendaftaran perusahaan", "Pastikan nama dan jenis perusahaan benar", type_error,"tutup");
        return;
    }

    StartLoading();
    var swalConfirm = SwalConfimationProcess("Daftarkan Perusahaan", "Lanjutkan ?", type_question, "Daftarkan", "Batal");
    swalConfirm.then(function(){
        var promise = RegisterCompany(companyType, companyName);
        promise.done(function(response){
            FinishLoading();
            swal = SimpleSwal("Sukses melakukan pendaftaran perusahaan", "Sistem akan melakukan login ulang", type_success, "tutup");
            swal.then(function(){
                ProcessLogout();
            });
        }).fail(function(response){
            FinishLoading();
            var errMessage;
            if(response.status >= 500){
                errMessage = "silahkan mencoba sekali lagi"
            } else {
                errMessage = response.responseJSON.data
            }

            SimpleSwal("Gagal melakukan pendaftaran perusahaan", errMessage, type_error, "tutup");
        });
    });

    $("#company-name").val("");
}

function ProcessChangePassword(){
    var tokenOld = $("#token-old").val();
    var token = $("#token").val();
    var tokenVerification = $("#token-verification").val();

    if(tokenOld == ""){
        SimpleSwal("Ganti Password", "Masukkan Password lama anda", type_error, "Tutup");
        return 
    }

    if(token != tokenVerification){
        SimpleSwal("Ganti Password", "Password Baru dan Password Verifikasi salah", type_error, "Tutup");
        return 
    }

    var swal = SwalConfimationProcess("Ganti Password", "Lanjutkan ?", type_question, "Ganti Password", "Batal");
    swal.then(function(){
        StartLoading();
        promise = changePassword(tokenOld, token, tokenVerification);
        promise.done(function(response){
            FinishLoading();
            SimpleSwal("Ganti Password", "Berhasil melakukan penggantian password", type_success, "OK");
        }).fail(function(response){
            FinishLoading();
            if(response.responseJSON != undefined){
                SimpleSwal("Ganti Password", response.responseJSON.data.message, type_error, "Tutup");
            } else {
                SimpleSwal("Ganti Password", "silahkan mencoba sekali lagi", type_error, "Tutup");
            }
        });
    });


    $("#token-old").val("");
    $("#token").val("");
    $("#token-verification").val("");
}

function loadCompanyDetail(){
    promise = getCompanyDetail();

    StartLoading();
    promise.done(function(response){
        FinishLoading();
        response = response.data;    
        $("#company-name").html(response.data.company_name);    
        $("#company-create-time").html("Terdaftar Sejak :" + response.data.create_time);

        if(response.data.Status == 1){
            $("#company-status").html("Status : Aktif");
        } else {
            $("#company-status").html("Status : Tidak Aktif");
        }
        $("#company-name").val(response.data.company_type + response.data.company_name);

        $("#company-img").attr("src", base_url + response.data.logo_url);
    })
}

function initTableMember(){
    StartLoading();
    userTable = $('#userTable').DataTable();
    promise = getCompanyMember();
    promise.done(function(response){
        FinishLoading();
        users = response.data.data;
        var counter = 1;

        users.forEach(element => {
            var status = "Aktif";
            if(element.status!=1){
                status = "Tidak Aktif";
            }

            var typeUser = "Tidak Terdefinisi";
            if(element.type==1){
                typeUser = "Admin";
            } else if(element.type==2){
                typeUser = "Pengguna Biasa";
            }

            var imgUrl = `<img src="` + element.img_url + `" width="50px" height="50px" />`;
            
            userTable.row.add([
                counter,
                element.fullname,
                typeUser,
                status,
                imgUrl
            ]).draw(false);

            counter++;
        });
    });

}