function loadlistProduct(totalRequest = 8){
    promise = getProductForMP(totalRequest);

    promise.done(function(response){
        if(response.data.length == 0){
            return;
        }

        for(var i=response.data.length-1;i>=0;i--){
            var priceInWeek = "NEGO"
            if(response.data[i].price_to_sell!=""){
                priceInWeek = response.data[i].price_to_sell +`/Minggu`;
            } 

            var template_choice = ``+
                `<div class="col-lg-3">` + 
                    `<div class="features-icons-item mx-auto">` + 
                        `<div class="features-icons-icon d-flex">` +
                            `<img src="` + base_url + response.data[i].img_url + `" style="min-width:10rem;">` +
                        `</div>` + 
                        `<h5 style="padding-top:1rem">` + response.data[i].product_name + `</h3>` + 
                        `<p class="desc" style="text-align:center;">`+ priceInWeek +
                        `<a class="btn btn-home btn-sm col-lg-12" onClick="javascript:goToDetailSewa('`+ response.data[i].product_name +`')" >SEWA</a>` +
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
        $("#product-id").val(response.data.product_id);        
        $("#product-img").attr("src", base_url + response.data.img_url);

        if(response.data.price_rent_daily!=""){
            $("#product-price-to-rent-daily").html(response.data.price_rent_daily + '/hari');            
        } else{
            $("#product-price-to-rent-daily").css("display", "none");
        }

        if(response.data.price_rent_weekly!=""){
            $("#product-price-to-rent-weekly").html(response.data.price_rent_weekly + '/minggu');        
        } else{
            $("#product-price-to-rent-weekly").css("display", "none");
        }

        if(response.data.price_rent_monthly!=""){
            $("#product-price-to-rent-monthly").html(response.data.price_rent_monthly + '/bulan');    
        } else{
            $("#product-price-to-rent-monthly").css("display", "none");
        }

    })
}

function goToDetailSewa(productName){
    var url = base_url + '/trek/' + productName;
    window.location.href = url;
}