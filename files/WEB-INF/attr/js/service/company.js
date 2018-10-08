function getCompanyDetail(){
    var url = product_url + '/company/get-detail';
    var promise = $.ajax({
        url: url,
        type: 'GET',
        xhrFields: {
            withCredentials: true
         },
    });
    
    return promise;
    
}

function RegisterCompany(companyType, companyName){
    var url = product_url + '/company/register-company';
    var data = {
        companyType:companyType,
        companyName:companyName,
    };
    
    var promise = $.ajax({
        url: url,
        type: 'POST',
        data: data,
        xhrFields: {
            withCredentials: true
        }
    });

    return promise;
}