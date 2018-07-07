function loadlistProduct(){
    promise = getProductForMP();

    promise.done(function(response){
        if(response.data.length == 0){
            return;
        }

        for(var i=response.data.length-1;i>=0;i--){
            var template_choice = ``+
                `<div class="col-lg-3">` + 
                    `<div class="features-icons-item mx-auto">` + 
                        `<div class="features-icons-icon d-flex">` +
                            `<img src="` + base_url + response.data[i].img_url + `" style="min-width:10rem;">` +
                        `</div>` + 
                        `<h5 style="padding-top:1rem">` + response.data[i].product_name + `</h3>` + 
                        `<p class="desc" style="text-align:center;">`+ response.data[i].price_to_sell +`/Minggu` +
                        `<a class="btn btn-home btn-sm col-lg-12" onClick="javascript:goToDetailSewa(`+ response.data[i].product_id +`)" >SEWA</a>` +
                    `</div>` +
                `</div>`;

            $("#parent-list-products").after($(template_choice));
        }
    });
}

function loadProductDetail(){
    promise = getProductDetail();

    promise.done(function(response){
        $("#product-name").html(response.data.product_name);
        $("#product-price-to-rent").html(response.data.price_to_sell + '/minggu');
        $("#product-price-to-buy").html('Harga beli di toko ' + response.data.price_to_buy);
        $("#product-img").attr("src", base_url + response.data.img_url);
    })
}

function goToDetailSewa(idProduct){
    var url = base_url + '/provider/trek/' + idProduct;
    window.location.href = url;
}