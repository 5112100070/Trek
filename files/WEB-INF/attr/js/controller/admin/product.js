function loadlistProduct(totalRequest = 8){
    promise = getProductForMP(totalRequest);

    StartLoading();
    promise.done(function(response){
        FinishLoading();
        if(response.data.length == 0){
            return;
        }

        for(var i=response.data.length-1;i>=0;i--){
            var priceInWeek = "-"
            if(response.data[i].price_to_sell!=""){
                priceInWeek = response.data[i].price_to_sell +`/Minggu`;
            } 

            var template_choice = ``+
                `<div class="col-lg-3" style="margin-bottom:2%;">` + 
                    `<div class="features-icons-item mx-auto">` + 
                        `<div class="features-icons-icon d-flex">` +
                            `<img src="` + base_url + response.data[i].img_url + `" style="width:100%; height:10rem;">` +
                        `</div>` + 
                        `<h5 style="padding-top:1rem">` + response.data[i].product_name + `</h3>` + 
                        `<!-- <p class="desc" style="text-align:center;">`+ priceInWeek + `-->` +
                        `<a class="btn btn-home btn-sm col-lg-12" onClick="javascript:GoToIndex('admin/product/edit?product-id=`+ response.data[i].product_id +`')" >UPDATE PRODUK</a>` +
                    `</div>` +
                `</div>`;

            $("#parent-list-products").after($(template_choice));
        }
    });
}

function loadProductDetailByID(productID){
    promise = getProductDetailByID(productID);

    StartLoading();
    promise.done(function(response){
        FinishLoading();
        $("#product-name").val(response.data.product_name);
        $("#status").val(response.data.status);
        $("#type").val(response.data.type);
        $("#path").val(response.data.path);
        
        $("#product-img").attr("src", base_url + response.data.img_url);

        $("#product-price-to-rent-daily").html(response.data.price_rent_daily);
        $("#product-price-to-rent-weekly").html(response.data.price_rent_weekly);
        $("#product-price-to-rent-monthly").html(response.data.price_rent_monthly);
    })
}