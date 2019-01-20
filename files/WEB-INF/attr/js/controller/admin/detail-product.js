function ProcessSaveNew(){
    ProcessRequestProductSave();
}

function ProcessUpdate(){
    ProcessRequestProductSave(2);
}

/*
    1 = save new
    2 = update
*/
function ProcessRequestProductSave(command = 1){
    var productName = $("#product-name").val();
    var typeProduct = $("#type").val();
    var status = $("#status").val();
    var priceRentDaily = $("#price-rent-daily").val();
    var priceRentWeekly = $("#price-rent-weekly").val();
    var priceRentMonthly = $("#price-rent-monthly").val();
    var path = $("#path").val();

    if(productName=="" || typeProduct == "" || priceRentDaily == "" || priceRentWeekly == "" || priceRentMonthly == "" || path == ""){
        SimpleSwal("Gagal menyimpan product baru", "Pastikan lengkapi seluruh keterangan data", type_error, "tutup");
        return;
    }

    var titleCommand = "Daftarkan Produk"
    if(command==2){
        titleCommand = "Update Produk"
    }

    var swal = SwalConfimationProcess(titleCommand, "Lanjutkan ?", type_question, "Lanjutkan", "Batalkan");
    swal.then(function(){
        if(command == 1){
            promise = sendNewProduct(productName, typeProduct, priceRentDaily, priceRentWeekly, priceRentMonthly, path);
        }
        if(command == 2){
            var url = new URL(window.location.href);
            var productID = url.searchParams.get("product-id");
            promise = sendUpdateProduct(productID, productName, typeProduct, status, priceRentDaily, priceRentWeekly, priceRentMonthly, path);
        }
        
        StartLoading();
        promise.done(function(response){
            FinishLoading();
            var successSwal = SimpleSwalWithoutDesc("Berhasil menyimpan product baru", type_success, "Menuju ke list");
            successSwal.then(function(){
                GoToIndex('admin/product');
            });
            return;
        }).fail(function(response){
            FinishLoading();
            if(response.status==400){
                SimpleSwal("Gagal menyimpan product baru", "Pastikan lengkapi seluruh keterangan data", type_error, "tutup");
            } else {
                SimpleSwal("Gagal menyimpan product baru", "Silahkan mencoba sekali lagi", type_error, "tutup");
            }
        });
    });
}

function ProcessUploadImg(){
    var url = new URL(window.location.href);
    
    var productID = url.searchParams.get("product-id");
    var blobFile = $("#product-img-new")[0].files[0];
    
    var swal = SwalConfimationProcess("Upload Gambar", "Gambar akan di upload. Lanjutkan ?", type_question, "Lanjutkan", "Batal");
    swal.then(function(){
        StartLoading();
        promise = sendUpdateImgProduct(productID, blobFile);
        promise.done(function(response){
            FinishLoading();
            var swal = SimpleSwal("Sukses", "Sukses menyimpan gambar", type_success, "tutup");
            swal.then(function(){
                location.reload();
            });
        }).fail(function(response){
            FinishLoading();
            location.reload();
            // var swal = SimpleSwal("Opps", "Silahkan mencoba sekali lagi", type_error, "tutup");
            // swal.then(function(){
            //     location.reload();
            // });
        });
    });
}

var isValidDate = function(value, format) {
    format = format || false;
    // lets parse the date to the best of our knowledge
    if (format) {
        value = parseDate(value);
    }

    var timestamp = Date.parse(value);

    return isNaN(timestamp) == false;
}

var parseDate = function(value) {
    var m = value.match(/^(\d{1,2})(\/|-)?(\d{1,2})(\/|-)?(\d{4})$/);
    if (m)
        value = m[5] + '-' + ("00" + m[3]).slice(-2) + '-' + ("00" + m[1]).slice(-2);

    return value;
}